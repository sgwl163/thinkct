// Package web
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package web

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/webin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sWebLinks struct{}

func NewWebLinks() *sWebLinks {
	return &sWebLinks{}
}

func init() {
	service.RegisterWebLinks(NewWebLinks())
}

// Model 友情链接ORM模型
func (s *sWebLinks) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.WebLinks.Ctx(ctx), option...)
}

// List 获取友情链接列表
func (s *sWebLinks) List(ctx context.Context, in *webin.LinksListInp) (list []*webin.LinksListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(webin.LinksListModel{})

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.WebLinks.Columns().Id, in.Id)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.WebLinks.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取友情链接列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出友情链接
func (s *sWebLinks) Export(ctx context.Context, in *webin.LinksListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(webin.LinksExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出友情链接-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []webin.LinksExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增友情链接
func (s *sWebLinks) Edit(ctx context.Context, in *webin.LinksEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(webin.LinksUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改友情链接失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(webin.LinksInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增友情链接失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除友情链接
func (s *sWebLinks) Delete(ctx context.Context, in *webin.LinksDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除友情链接失败，请稍后重试！")
		return
	}
	return
}

// View 获取友情链接指定信息
func (s *sWebLinks) View(ctx context.Context, in *webin.LinksViewInp) (res *webin.LinksViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取友情链接信息，请稍后重试！")
		return
	}
	return
}

// Status 更新友情链接状态
func (s *sWebLinks) Status(ctx context.Context, in *webin.LinksStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.WebLinks.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新友情链接状态失败，请稍后重试！")
		return
	}
	return
}