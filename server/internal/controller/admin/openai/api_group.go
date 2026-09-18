// Package openai
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package openai

import (
	"context"
	"hotgo/api/admin/openai/apigroup"
	"hotgo/internal/model/input/openaiin"
	"hotgo/internal/service"
)

var (
	ApiGroup = cApiGroup{}
)

type cApiGroup struct{}

// List 查看AI接口分组列表
func (c *cApiGroup) List(ctx context.Context, req *apigroup.ListReq) (res *apigroup.ListRes, err error) {
	list, totalCount, err := service.OpenaiApiGroup().List(ctx, &req.ApiGroupListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*openaiin.ApiGroupListModel{}
	}

	res = new(apigroup.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出AI接口分组列表
func (c *cApiGroup) Export(ctx context.Context, req *apigroup.ExportReq) (res *apigroup.ExportRes, err error) {
	err = service.OpenaiApiGroup().Export(ctx, &req.ApiGroupListInp)
	return
}

// Edit 更新AI接口分组
func (c *cApiGroup) Edit(ctx context.Context, req *apigroup.EditReq) (res *apigroup.EditRes, err error) {
	err = service.OpenaiApiGroup().Edit(ctx, &req.ApiGroupEditInp)
	return
}

// MaxSort 获取AI接口分组最大排序
func (c *cApiGroup) MaxSort(ctx context.Context, req *apigroup.MaxSortReq) (res *apigroup.MaxSortRes, err error) {
	data, err := service.OpenaiApiGroup().MaxSort(ctx, &req.ApiGroupMaxSortInp)
	if err != nil {
		return
	}

	res = new(apigroup.MaxSortRes)
	res.ApiGroupMaxSortModel = data
	return
}

// View 获取指定AI接口分组信息
func (c *cApiGroup) View(ctx context.Context, req *apigroup.ViewReq) (res *apigroup.ViewRes, err error) {
	data, err := service.OpenaiApiGroup().View(ctx, &req.ApiGroupViewInp)
	if err != nil {
		return
	}

	res = new(apigroup.ViewRes)
	res.ApiGroupViewModel = data
	return
}

// Delete 删除AI接口分组
func (c *cApiGroup) Delete(ctx context.Context, req *apigroup.DeleteReq) (res *apigroup.DeleteRes, err error) {
	err = service.OpenaiApiGroup().Delete(ctx, &req.ApiGroupDeleteInp)
	return
}

// Status 更新AI接口分组状态
func (c *cApiGroup) Status(ctx context.Context, req *apigroup.StatusReq) (res *apigroup.StatusRes, err error) {
	err = service.OpenaiApiGroup().Status(ctx, &req.ApiGroupStatusInp)
	return
}