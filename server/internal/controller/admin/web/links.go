// Package web
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package web

import (
	"context"
	"hotgo/api/admin/web/links"
	"hotgo/internal/model/input/webin"
	"hotgo/internal/service"
)

var (
	Links = cLinks{}
)

type cLinks struct{}

// List 查看友情链接列表
func (c *cLinks) List(ctx context.Context, req *links.ListReq) (res *links.ListRes, err error) {
	list, totalCount, err := service.WebLinks().List(ctx, &req.LinksListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*webin.LinksListModel{}
	}

	res = new(links.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出友情链接列表
func (c *cLinks) Export(ctx context.Context, req *links.ExportReq) (res *links.ExportRes, err error) {
	err = service.WebLinks().Export(ctx, &req.LinksListInp)
	return
}

// Edit 更新友情链接
func (c *cLinks) Edit(ctx context.Context, req *links.EditReq) (res *links.EditRes, err error) {
	err = service.WebLinks().Edit(ctx, &req.LinksEditInp)
	return
}

// View 获取指定友情链接信息
func (c *cLinks) View(ctx context.Context, req *links.ViewReq) (res *links.ViewRes, err error) {
	data, err := service.WebLinks().View(ctx, &req.LinksViewInp)
	if err != nil {
		return
	}

	res = new(links.ViewRes)
	res.LinksViewModel = data
	return
}

// Delete 删除友情链接
func (c *cLinks) Delete(ctx context.Context, req *links.DeleteReq) (res *links.DeleteRes, err error) {
	err = service.WebLinks().Delete(ctx, &req.LinksDeleteInp)
	return
}

// Status 更新友情链接状态
func (c *cLinks) Status(ctx context.Context, req *links.StatusReq) (res *links.StatusRes, err error) {
	err = service.WebLinks().Status(ctx, &req.LinksStatusInp)
	return
}