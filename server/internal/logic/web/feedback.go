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

type sWebFeedback struct{}

func NewWebFeedback() *sWebFeedback {
	return &sWebFeedback{}
}

func init() {
	service.RegisterWebFeedback(NewWebFeedback())
}

// Model 留言管理ORM模型
func (s *sWebFeedback) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.WebFeedback.Ctx(ctx), option...)
}

// List 获取留言管理列表
func (s *sWebFeedback) List(ctx context.Context, in *webin.FeedbackListInp) (list []*webin.FeedbackListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(webin.FeedbackListModel{})

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.WebFeedback.Columns().Id, in.Id)
	}

	// 查询状态:0=禁用,1=启用
	if in.Status > 0 {
		mod = mod.Where(dao.WebFeedback.Columns().Status, in.Status)
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.WebFeedback.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取留言管理列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出留言管理
func (s *sWebFeedback) Export(ctx context.Context, in *webin.FeedbackListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(webin.FeedbackExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出留言管理-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []webin.FeedbackExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增留言管理
func (s *sWebFeedback) Edit(ctx context.Context, in *webin.FeedbackEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(webin.FeedbackUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改留言管理失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(webin.FeedbackInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增留言管理失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除留言管理
func (s *sWebFeedback) Delete(ctx context.Context, in *webin.FeedbackDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除留言管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取留言管理指定信息
func (s *sWebFeedback) View(ctx context.Context, in *webin.FeedbackViewInp) (res *webin.FeedbackViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取留言管理信息，请稍后重试！")
		return
	}
	return
}

// Status 更新留言管理状态
func (s *sWebFeedback) Status(ctx context.Context, in *webin.FeedbackStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.WebFeedback.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新留言管理状态失败，请稍后重试！")
		return
	}
	return
}