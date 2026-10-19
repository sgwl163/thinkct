// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package sys

import (
	"context"
	"fmt"
	"hotgo/addons/appAuth/model/input/sysin"
	"hotgo/addons/appAuth/service"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysAppProducts struct{}

func NewSysAppProducts() *sSysAppProducts {
	return &sSysAppProducts{}
}

func init() {
	service.RegisterSysAppProducts(NewSysAppProducts())
}

// Model 产品管理ORM模型
func (s *sSysAppProducts) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.AppProducts.Ctx(ctx), option...)
}

// List 获取产品管理列表
func (s *sSysAppProducts) List(ctx context.Context, in *sysin.AppProductsListInp) (list []*sysin.AppProductsListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.AppProductsListModel{})

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.AppProducts.Columns().Id, in.Id)
	}

	// 查询状态:0=禁用,1=启用
	if in.Status > 0 {
		mod = mod.Where(dao.AppProducts.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.AppProducts.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.AppProducts.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取产品管理列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出产品管理
func (s *sSysAppProducts) Export(ctx context.Context, in *sysin.AppProductsListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.AppProductsExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出产品管理-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.AppProductsExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增产品管理
func (s *sSysAppProducts) Edit(ctx context.Context, in *sysin.AppProductsEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.AppProductsUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改产品管理失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.AppProductsInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增产品管理失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除产品管理
func (s *sSysAppProducts) Delete(ctx context.Context, in *sysin.AppProductsDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除产品管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取产品管理指定信息
func (s *sSysAppProducts) View(ctx context.Context, in *sysin.AppProductsViewInp) (res *sysin.AppProductsViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取产品管理信息，请稍后重试！")
		return
	}
	return
}

// Status 更新产品管理状态
func (s *sSysAppProducts) Status(ctx context.Context, in *sysin.AppProductsStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.AppProducts.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新产品管理状态失败，请稍后重试！")
		return
	}
	return
}