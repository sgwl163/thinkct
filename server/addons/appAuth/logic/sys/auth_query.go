package sys

import (
	"context"
	"fmt"
	"sort"
	"strings"

	ac "hotgo/addons/appAuth/consts"
	am "hotgo/addons/appAuth/model"
	"hotgo/addons/appAuth/model/input/sysin"
	"hotgo/addons/appAuth/service"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/dict"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/net/idna"
	"golang.org/x/net/publicsuffix"
)

// authQueryRecord 只读取公开查询判定所需字段，绝不读取私有授权码、用户或联系方式。
type authQueryRecord struct {
	Id         uint
	ProductId  int
	Name       string
	Host       string
	AuthDomain string
	AuthMode   string
	Status     uint
	CreatedAt  *gtime.Time
	ExpiredAt  *gtime.Time
}

type authQueryProduct struct {
	Id     uint
	Name   string
	Title  string
	Status uint
}

// Query 仅查询域名的授权登记状态，不作为软件客户端的授权码或IP验证接口。
// 精确域名优先于泛域名，更具体的泛域名优先；禁用或到期的精确记录不会回退到泛域名。
func (s *sSysAppAuth) Query(ctx context.Context, in *sysin.AuthQueryInp) (res *sysin.AuthQueryModel, err error) {
	conf, err := service.SysConfig().GetBasic(ctx)
	if err != nil {
		return nil, err
	}
	// 只有正常得到查询结果才计数，未登记、到期和禁用同样属于完成的查询。
	defer func() {
		if err != nil {
			return
		}
		res.AuthQueryCount, err = service.SysConfig().IncrementAuthQueryCount(ctx)
		if err != nil {
			res = nil
		}
	}()
	res = &sysin.AuthQueryModel{Host: in.Host, Status: ac.AppAuthQueryUnauthorized, Message: conf.UnauthorizedMessage}
	exact, wildcard := authQueryHosts(in.Host)
	cols := dao.AppAuth.Columns()
	// 域名不区分大小写，兼容后台旧记录的首尾空格和一个末尾根域名点。
	hostColumn := fmt.Sprintf("LOWER(TRIM(%s))", cols.Host)
	mod := dao.AppAuth.Ctx(ctx)
	mod = mod.Where(mod.Builder().WhereIn(hostColumn, exact).WhereOr(
		mod.Builder().WhereIn(hostColumn, wildcard).Where(cols.AuthDomain, ac.AppAuthDomainWildcard),
	))
	if in.ProductId > 0 {
		mod = mod.Where(cols.ProductId, in.ProductId)
	}
	var rows []*authQueryRecord
	if err = mod.Fields(authQueryRecord{}).OrderDesc(cols.Id).Scan(&rows); err != nil {
		return nil, gerror.Wrap(err, "查询域名授权失败，请稍后重试")
	}
	row := selectAuthQueryRecord(in.Host, rows)
	if row == nil {
		return res, nil
	}

	// 产品被禁用或删除时，所属授权不再作为有效授权公开展示。
	var product *authQueryProduct
	productCols := dao.AppProducts.Columns()
	if err = dao.AppProducts.Ctx(ctx).Fields(authQueryProduct{}).Where(productCols.Id, row.ProductId).Scan(&product); err != nil {
		return nil, gerror.Wrap(err, "查询授权产品失败，请稍后重试")
	}
	if row.Status != uint(consts.StatusEnabled) || product == nil || product.Status != uint(consts.StatusEnabled) {
		res.Status, res.Message = ac.AppAuthQueryDisabled, conf.DisabledMessage
		return res, nil
	}
	if row.ExpiredAt != nil && !row.ExpiredAt.IsZero() && !row.ExpiredAt.After(gtime.Now()) {
		res.Status, res.Message = ac.AppAuthQueryExpired, conf.ExpiredMessage
		return res, nil
	}

	productName := product.Name
	if productName == "" {
		productName = product.Title
	}
	certificate := &sysin.AuthCertificate{
		Domain: in.Host, Company: row.Name, ProductName: productName,
		Type: dict.GetOptionLabel(ac.AppAuthDomainOptions, row.AuthDomain),
		// 稳定的公开登记参考号，不是授权表中的syskey，不能用于软件激活。
		License: fmt.Sprintf("AUTH-%08d", row.Id), ExpireDate: "长期有效",
	}
	if row.CreatedAt != nil && !row.CreatedAt.IsZero() {
		certificate.IssueDate = row.CreatedAt.Format("Y-m-d")
	}
	if row.ExpiredAt != nil && !row.ExpiredAt.IsZero() {
		certificate.ExpireDate = row.ExpiredAt.String()
	}
	res.Authorized, res.Status, res.Message = true, ac.AppAuthQueryAuthorized, "该域名已获得产品授权"
	res.Certificate = certificate
	return res, nil
}

// authQueryHosts 仅生成精确值候选，不使用LIKE，防止域名边界绕过或读取无关授权记录。
// 泛域名只查至可注册域名，不会将com、co.uk等公共后缀作为授权范围。
func authQueryHosts(host string) (exact, wildcard []string) {
	variants := func(value string) []string {
		values := []string{value, value + "."}
		if unicode, err := idna.Lookup.ToUnicode(value); err == nil && unicode != value {
			values = append(values, unicode, unicode+".")
		}
		return values
	}
	exact = variants(host)
	root, _ := publicsuffix.EffectiveTLDPlusOne(host)
	for base := host; base != ""; {
		for _, candidate := range variants(base) {
			wildcard = append(wildcard, candidate, "*."+candidate)
		}
		if base == root {
			break
		}
		_, rest, ok := strings.Cut(base, ".")
		if !ok {
			break
		}
		base = rest
	}
	return
}

func selectAuthQueryRecord(host string, rows []*authQueryRecord) *authQueryRecord {
	type match struct {
		row   *authQueryRecord
		base  string
		exact bool
	}
	var matches []match
	for _, row := range rows {
		if row == nil || (row.AuthMode != ac.AppAuthModeDomain && row.AuthMode != ac.AppAuthModeDomainIP) {
			continue
		}
		stored := strings.ToLower(strings.TrimSpace(row.Host))
		isWildcard := strings.HasPrefix(stored, "*.")
		base, err := am.NormalizeAuthHost(strings.TrimPrefix(stored, "*."))
		if err != nil {
			continue
		}
		exact := !isWildcard && base == host
		switch row.AuthDomain {
		case ac.AppAuthDomainSingle:
			if !exact {
				continue
			}
		case ac.AppAuthDomainWildcard:
			if base != host && !strings.HasSuffix(host, "."+base) {
				continue
			}
		default:
			continue
		}
		matches = append(matches, match{row: row, base: base, exact: exact})
	}
	sort.SliceStable(matches, func(i, j int) bool {
		if matches[i].exact != matches[j].exact {
			return matches[i].exact
		}
		if len(matches[i].base) != len(matches[j].base) {
			return len(matches[i].base) > len(matches[j].base)
		}
		return matches[i].row.Id > matches[j].row.Id
	})
	if len(matches) == 0 {
		return nil
	}
	return matches[0].row
}
