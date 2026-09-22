// Package forum
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package forum

import (
	"context"
	"hotgo/api/admin/forum/category"
	"hotgo/internal/model/input/forumin"
	"hotgo/internal/service"
)

var (
	Category = cCategory{}
)

type cCategory struct{}

// List 查看社区板块列表
func (c *cCategory) List(ctx context.Context, req *category.ListReq) (res *category.ListRes, err error) {
	list, totalCount, err := service.ForumCategory().List(ctx, &req.CategoryListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*forumin.CategoryListModel{}
	}

	res = new(category.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出社区板块列表
func (c *cCategory) Export(ctx context.Context, req *category.ExportReq) (res *category.ExportRes, err error) {
	err = service.ForumCategory().Export(ctx, &req.CategoryListInp)
	return
}

// Edit 更新社区板块
func (c *cCategory) Edit(ctx context.Context, req *category.EditReq) (res *category.EditRes, err error) {
	err = service.ForumCategory().Edit(ctx, &req.CategoryEditInp)
	return
}

// MaxSort 获取社区板块最大排序
func (c *cCategory) MaxSort(ctx context.Context, req *category.MaxSortReq) (res *category.MaxSortRes, err error) {
	data, err := service.ForumCategory().MaxSort(ctx, &req.CategoryMaxSortInp)
	if err != nil {
		return
	}

	res = new(category.MaxSortRes)
	res.CategoryMaxSortModel = data
	return
}

// View 获取指定社区板块信息
func (c *cCategory) View(ctx context.Context, req *category.ViewReq) (res *category.ViewRes, err error) {
	data, err := service.ForumCategory().View(ctx, &req.CategoryViewInp)
	if err != nil {
		return
	}

	res = new(category.ViewRes)
	res.CategoryViewModel = data
	return
}

// Delete 删除社区板块
func (c *cCategory) Delete(ctx context.Context, req *category.DeleteReq) (res *category.DeleteRes, err error) {
	err = service.ForumCategory().Delete(ctx, &req.CategoryDeleteInp)
	return
}

// Status 更新社区板块状态
func (c *cCategory) Status(ctx context.Context, req *category.StatusReq) (res *category.StatusRes, err error) {
	err = service.ForumCategory().Status(ctx, &req.CategoryStatusInp)
	return
}
