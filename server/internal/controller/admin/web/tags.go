// Package web
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package web

import (
	"context"
	"hotgo/api/admin/web/tags"
	"hotgo/internal/model/input/webin"
	"hotgo/internal/service"
)

var (
	Tags = cTags{}
)

type cTags struct{}

// List 查看标签管理列表
func (c *cTags) List(ctx context.Context, req *tags.ListReq) (res *tags.ListRes, err error) {
	list, totalCount, err := service.WebTags().List(ctx, &req.TagsListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*webin.TagsListModel{}
	}

	res = new(tags.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出标签管理列表
func (c *cTags) Export(ctx context.Context, req *tags.ExportReq) (res *tags.ExportRes, err error) {
	err = service.WebTags().Export(ctx, &req.TagsListInp)
	return
}

// Edit 更新标签管理
func (c *cTags) Edit(ctx context.Context, req *tags.EditReq) (res *tags.EditRes, err error) {
	err = service.WebTags().Edit(ctx, &req.TagsEditInp)
	return
}

// View 获取指定标签管理信息
func (c *cTags) View(ctx context.Context, req *tags.ViewReq) (res *tags.ViewRes, err error) {
	data, err := service.WebTags().View(ctx, &req.TagsViewInp)
	if err != nil {
		return
	}

	res = new(tags.ViewRes)
	res.TagsViewModel = data
	return
}

// Delete 删除标签管理
func (c *cTags) Delete(ctx context.Context, req *tags.DeleteReq) (res *tags.DeleteRes, err error) {
	err = service.WebTags().Delete(ctx, &req.TagsDeleteInp)
	return
}

// Status 更新标签管理状态
func (c *cTags) Status(ctx context.Context, req *tags.StatusReq) (res *tags.StatusRes, err error) {
	err = service.WebTags().Status(ctx, &req.TagsStatusInp)
	return
}