// Package web
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package web

import (
	"context"
	"hotgo/api/admin/web/article"
	"hotgo/internal/model/input/webin"
	"hotgo/internal/service"
)

var (
	Article = cArticle{}
)

type cArticle struct{}

// List 查看文章管理列表
func (c *cArticle) List(ctx context.Context, req *article.ListReq) (res *article.ListRes, err error) {
	list, totalCount, err := service.WebArticle().List(ctx, &req.ArticleListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*webin.ArticleListModel{}
	}

	res = new(article.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出文章管理列表
func (c *cArticle) Export(ctx context.Context, req *article.ExportReq) (res *article.ExportRes, err error) {
	err = service.WebArticle().Export(ctx, &req.ArticleListInp)
	return
}

// Edit 更新文章管理
func (c *cArticle) Edit(ctx context.Context, req *article.EditReq) (res *article.EditRes, err error) {
	err = service.WebArticle().Edit(ctx, &req.ArticleEditInp)
	return
}

// View 获取指定文章管理信息
func (c *cArticle) View(ctx context.Context, req *article.ViewReq) (res *article.ViewRes, err error) {
	data, err := service.WebArticle().View(ctx, &req.ArticleViewInp)
	if err != nil {
		return
	}

	res = new(article.ViewRes)
	res.ArticleViewModel = data
	return
}

// Delete 删除文章管理
func (c *cArticle) Delete(ctx context.Context, req *article.DeleteReq) (res *article.DeleteRes, err error) {
	err = service.WebArticle().Delete(ctx, &req.ArticleDeleteInp)
	return
}

// Status 更新文章管理状态
func (c *cArticle) Status(ctx context.Context, req *article.StatusReq) (res *article.StatusRes, err error) {
	err = service.WebArticle().Status(ctx, &req.ArticleStatusInp)
	return
}