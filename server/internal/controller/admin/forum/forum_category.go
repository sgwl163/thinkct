// Package forum
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package forum

import (
	"context"
	"hotgo/api/admin/forum/forumcategory"
	"hotgo/internal/model/input/forumin"
	"hotgo/internal/service"
)

var (
	ForumCategory = cForumCategory{}
)

type cForumCategory struct{}

// List 查看社区板块列表
func (c *cForumCategory) List(ctx context.Context, req *forumcategory.ListReq) (res *forumcategory.ListRes, err error) {
	list, totalCount, err := service.ForumForumCategory().List(ctx, &req.ForumCategoryListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*forumin.ForumCategoryListModel{}
	}

	res = new(forumcategory.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出社区板块列表
func (c *cForumCategory) Export(ctx context.Context, req *forumcategory.ExportReq) (res *forumcategory.ExportRes, err error) {
	err = service.ForumForumCategory().Export(ctx, &req.ForumCategoryListInp)
	return
}

// Edit 更新社区板块
func (c *cForumCategory) Edit(ctx context.Context, req *forumcategory.EditReq) (res *forumcategory.EditRes, err error) {
	err = service.ForumForumCategory().Edit(ctx, &req.ForumCategoryEditInp)
	return
}

// MaxSort 获取社区板块最大排序
func (c *cForumCategory) MaxSort(ctx context.Context, req *forumcategory.MaxSortReq) (res *forumcategory.MaxSortRes, err error) {
	data, err := service.ForumForumCategory().MaxSort(ctx, &req.ForumCategoryMaxSortInp)
	if err != nil {
		return
	}

	res = new(forumcategory.MaxSortRes)
	res.ForumCategoryMaxSortModel = data
	return
}

// View 获取指定社区板块信息
func (c *cForumCategory) View(ctx context.Context, req *forumcategory.ViewReq) (res *forumcategory.ViewRes, err error) {
	data, err := service.ForumForumCategory().View(ctx, &req.ForumCategoryViewInp)
	if err != nil {
		return
	}

	res = new(forumcategory.ViewRes)
	res.ForumCategoryViewModel = data
	return
}

// Delete 删除社区板块
func (c *cForumCategory) Delete(ctx context.Context, req *forumcategory.DeleteReq) (res *forumcategory.DeleteRes, err error) {
	err = service.ForumForumCategory().Delete(ctx, &req.ForumCategoryDeleteInp)
	return
}

// Status 更新社区板块状态
func (c *cForumCategory) Status(ctx context.Context, req *forumcategory.StatusReq) (res *forumcategory.StatusRes, err error) {
	err = service.ForumForumCategory().Status(ctx, &req.ForumCategoryStatusInp)
	return
}