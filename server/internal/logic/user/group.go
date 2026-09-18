// Package user
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package user

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/userin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sUsergroup struct{}

func NewUsergroup() *sUsergroup {
	return &sUsergroup{}
}

func init() {
	service.RegisterUsergroup(NewUsergroup())
}

// Model 分组管理ORM模型
func (s *sUsergroup) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.UserGroup.Ctx(ctx), option...)
}

// List 获取分组管理列表
func (s *sUsergroup) List(ctx context.Context, in *userin.GroupListInp) (list []*userin.GroupListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(userin.GroupListModel{})

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.UserGroup.Columns().Id, in.Id)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.UserGroup.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取分组管理列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出分组管理
func (s *sUsergroup) Export(ctx context.Context, in *userin.GroupListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(userin.GroupExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出分组管理-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []userin.GroupExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增分组管理
func (s *sUsergroup) Edit(ctx context.Context, in *userin.GroupEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(userin.GroupUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改分组管理失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(userin.GroupInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
				err = gerror.Wrap(err, "新增分组管理失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除分组管理
func (s *sUsergroup) Delete(ctx context.Context, in *userin.GroupDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除分组管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取分组管理指定信息
func (s *sUsergroup) View(ctx context.Context, in *userin.GroupViewInp) (res *userin.GroupViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取分组管理信息，请稍后重试！")
		return
	}
	return
}

// Status 更新分组管理状态
func (s *sUsergroup) Status(ctx context.Context, in *userin.GroupStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.UserGroup.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新分组管理状态失败，请稍后重试！")
		return
	}
	return
}
