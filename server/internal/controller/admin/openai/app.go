// Package openai
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package openai

import (
	"context"
	"hotgo/api/admin/openai/app"
	"hotgo/internal/model/input/openaiin"
	"hotgo/internal/service"
)

var (
	App = cApp{}
)

type cApp struct{}

// List 查看AI应用列表
func (c *cApp) List(ctx context.Context, req *app.ListReq) (res *app.ListRes, err error) {
	list, totalCount, err := service.OpenaiApp().List(ctx, &req.AppListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*openaiin.AppListModel{}
	}

	res = new(app.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出AI应用列表
func (c *cApp) Export(ctx context.Context, req *app.ExportReq) (res *app.ExportRes, err error) {
	err = service.OpenaiApp().Export(ctx, &req.AppListInp)
	return
}

// Edit 更新AI应用
func (c *cApp) Edit(ctx context.Context, req *app.EditReq) (res *app.EditRes, err error) {
	err = service.OpenaiApp().Edit(ctx, &req.AppEditInp)
	return
}

// MaxSort 获取AI应用最大排序
func (c *cApp) MaxSort(ctx context.Context, req *app.MaxSortReq) (res *app.MaxSortRes, err error) {
	data, err := service.OpenaiApp().MaxSort(ctx, &req.AppMaxSortInp)
	if err != nil {
		return
	}

	res = new(app.MaxSortRes)
	res.AppMaxSortModel = data
	return
}

// View 获取指定AI应用信息
func (c *cApp) View(ctx context.Context, req *app.ViewReq) (res *app.ViewRes, err error) {
	data, err := service.OpenaiApp().View(ctx, &req.AppViewInp)
	if err != nil {
		return
	}

	res = new(app.ViewRes)
	res.AppViewModel = data
	return
}

// Delete 删除AI应用
func (c *cApp) Delete(ctx context.Context, req *app.DeleteReq) (res *app.DeleteRes, err error) {
	err = service.OpenaiApp().Delete(ctx, &req.AppDeleteInp)
	return
}

// Status 更新AI应用状态
func (c *cApp) Status(ctx context.Context, req *app.StatusReq) (res *app.StatusRes, err error) {
	err = service.OpenaiApp().Status(ctx, &req.AppStatusInp)
	return
}