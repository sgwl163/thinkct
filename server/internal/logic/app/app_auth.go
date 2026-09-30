// Package app
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package app

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/appin"
	"hotgo/internal/model/input/form"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sAppAppAuth struct{}

func NewAppAppAuth() *sAppAppAuth {
	return &sAppAppAuth{}
}

func init() {
	service.RegisterAppAppAuth(NewAppAppAuth())
}

// Model 授权管理ORM模型
func (s *sAppAppAuth) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.AppAuth.Ctx(ctx), option...)
}

// List 获取授权管理列表
func (s *sAppAppAuth) List(ctx context.Context, in *appin.AppAuthListInp) (list []*appin.AppAuthListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(appin.AppAuthListModel{})

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.AppAuth.Columns().Id, in.Id)
	}

	// 查询状态:1=启用,0=禁用
	if in.Status > 0 {
		mod = mod.Where(dao.AppAuth.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.AppAuth.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.AppAuth.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取授权管理列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出授权管理
func (s *sAppAppAuth) Export(ctx context.Context, in *appin.AppAuthListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(appin.AppAuthExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出授权管理-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []appin.AppAuthExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增授权管理
func (s *sAppAppAuth) Edit(ctx context.Context, in *appin.AppAuthEditInp) (err error) {
	// 验证'Host'唯一
	if err = hgorm.IsUnique(ctx, &dao.AppAuth, g.Map{dao.AppAuth.Columns().Host: in.Host}, "域名已存在", in.Id); err != nil {
		return
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(appin.AppAuthUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改授权管理失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(appin.AppAuthInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增授权管理失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除授权管理
func (s *sAppAppAuth) Delete(ctx context.Context, in *appin.AppAuthDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除授权管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取授权管理指定信息
func (s *sAppAppAuth) View(ctx context.Context, in *appin.AppAuthViewInp) (res *appin.AppAuthViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取授权管理信息，请稍后重试！")
		return
	}
	return
}

// Status 更新授权管理状态
func (s *sAppAppAuth) Status(ctx context.Context, in *appin.AppAuthStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.AppAuth.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新授权管理状态失败，请稍后重试！")
		return
	}
	return
}