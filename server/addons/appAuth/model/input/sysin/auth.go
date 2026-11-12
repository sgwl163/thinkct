package sysin

import (
	"context"
	"hotgo/addons/appAuth/model"

	"github.com/gogf/gf/v2/errors/gerror"
)

// AuthQueryInp 公开查询域名授权登记状态，不用于客户端授权码或IP双重验证。
type AuthQueryInp struct {
	Host      string `json:"host" dc:"待查询的裸域名"`
	ProductId int    `json:"productId" v:"integer|min:0#产品ID必须为非负整数#产品ID不能为负数" dc:"可选产品ID，省略或零时查询该域名的授权登记"`
}

func (in *AuthQueryInp) Filter(ctx context.Context) (err error) {
	if in.ProductId < 0 {
		return gerror.New("产品ID不能为负数")
	}
	in.Host, err = model.NormalizeAuthHost(in.Host)
	return
}

type AuthQueryModel struct {
	Host           string           `json:"host" dc:"查询域名"`
	Authorized     bool             `json:"authorized" dc:"是否存在有效授权登记"`
	Status         string           `json:"status" dc:"authorized/unauthorized/expired/disabled"`
	Message        string           `json:"message" dc:"授权状态提示"`
	Certificate    *AuthCertificate `json:"certificate" dc:"有效授权证书展示信息，非授权码"`
	AuthQueryCount int64            `json:"authQueryCount" dc:"包含本次查询的累计授权查询次数"`
}

// AuthCertificate 公开展示白名单，不包含私有授权码、用户ID、电话、QQ或服务器IP。
type AuthCertificate struct {
	Domain      string `json:"domain" dc:"查询域名"`
	Company     string `json:"company" dc:"授权登记名称"`
	ProductName string `json:"productName" dc:"授权产品名称"`
	Type        string `json:"type" dc:"授权域名范围"`
	License     string `json:"license" dc:"公开证书参考编号，不是授权码"`
	IssueDate   string `json:"issueDate" dc:"授权登记日期"`
	ExpireDate  string `json:"expireDate" dc:"授权到期日期或长期有效"`
}

// AuthConfigInp 获取公开产品授权配置，不接受后台配置分组或插件名称。
type AuthConfigInp struct{}

func (in *AuthConfigInp) Filter(ctx context.Context) (err error) {
	return
}

// AuthConfigModel 独立的公开白名单，避免未来后台新增私有配置时被自动暴露。
type AuthConfigModel struct {
	UnauthorizedMessage string `json:"unauthorizedMessage"`
	ExpiredMessage      string `json:"expiredMessage"`
	DisabledMessage     string `json:"disabledMessage"`
	PurchaseUrl         string `json:"purchaseUrl"`
	ServiceName         string `json:"serviceName"`
	ServiceQQ           string `json:"serviceQQ"`
	ServicePhone        string `json:"servicePhone"`
	SupportUrl          string `json:"supportUrl"`
	AuthQueryCount      int64  `json:"authQueryCount" dc:"累计授权查询次数，可手动调整，完成授权查询后自动递增"`
}
