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

type sWebArticle struct{}

func NewWebArticle() *sWebArticle {
	return &sWebArticle{}
}

func init() {
	service.RegisterWebArticle(NewWebArticle())
}

// Model 文章管理ORM模型
func (s *sWebArticle) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.WebArticle.Ctx(ctx), option...)
}

// List 获取文章管理列表
func (s *sWebArticle) List(ctx context.Context, in *webin.ArticleListInp) (list []*webin.ArticleListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(webin.ArticleListModel{})

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.WebArticle.Columns().Id, in.Id)
	}
	
	if in.ChannelId > 0 {
		mod = mod.Where(dao.WebArticle.Columns().ChannelId, in.ChannelId)
	}

	if in.ChannelIds != "" {
		mod = mod.WhereLike(dao.WebArticle.Columns().ChannelIds, "%"+in.ChannelIds+"%")
	}

	if in.Title != "" {
		mod = mod.WhereLike(dao.WebArticle.Columns().Title, "%"+in.Title+"%")
	}

	if in.Flag != "" {
		mod = mod.WhereLike(dao.WebArticle.Columns().Flag, "%"+in.Flag+"%")
	}

	if in.Status != "" {
		mod = mod.Where(dao.WebArticle.Columns().Status, in.Status)
	}
	
	if len(in.CreateTime) == 2 {
		mod = mod.WhereBetween(dao.WebArticle.Columns().CreateTime, in.CreateTime[0], in.CreateTime[1])
	}
	if len(in.UpdateTime) == 2 {
		mod = mod.WhereBetween(dao.WebArticle.Columns().UpdateTime, in.UpdateTime[0], in.UpdateTime[1])
	}
	if len(in.PublishTime) == 2 {
		mod = mod.WhereBetween(dao.WebArticle.Columns().PublishTime, in.PublishTime[0], in.PublishTime[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.WebArticle.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取文章管理列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出文章管理
func (s *sWebArticle) Export(ctx context.Context, in *webin.ArticleListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(webin.ArticleExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出文章管理-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []webin.ArticleExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增文章管理
func (s *sWebArticle) Edit(ctx context.Context, in *webin.ArticleEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(webin.ArticleUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改文章管理失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(webin.ArticleInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增文章管理失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除文章管理
func (s *sWebArticle) Delete(ctx context.Context, in *webin.ArticleDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除文章管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取文章管理指定信息
func (s *sWebArticle) View(ctx context.Context, in *webin.ArticleViewInp) (res *webin.ArticleViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取文章管理信息，请稍后重试！")
		return
	}
	return
}

// Status 更新文章管理状态
func (s *sWebArticle) Status(ctx context.Context, in *webin.ArticleStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.WebArticle.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新文章管理状态失败，请稍后重试！")
		return
	}
	return
}