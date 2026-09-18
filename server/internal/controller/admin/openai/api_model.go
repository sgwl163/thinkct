// Package openai
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package openai

import (
	"context"
	"hotgo/api/admin/openai/apimodel"
	"hotgo/internal/model/input/openaiin"
	"hotgo/internal/service"
)

var (
	ApiModel = cApiModel{}
)

type cApiModel struct{}

// List 查看AI模型列表列表
func (c *cApiModel) List(ctx context.Context, req *apimodel.ListReq) (res *apimodel.ListRes, err error) {
	list, totalCount, err := service.OpenaiApiModel().List(ctx, &req.ApiModelListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*openaiin.ApiModelListModel{}
	}

	res = new(apimodel.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出AI模型列表列表
func (c *cApiModel) Export(ctx context.Context, req *apimodel.ExportReq) (res *apimodel.ExportRes, err error) {
	err = service.OpenaiApiModel().Export(ctx, &req.ApiModelListInp)
	return
}

// Edit 更新AI模型列表
func (c *cApiModel) Edit(ctx context.Context, req *apimodel.EditReq) (res *apimodel.EditRes, err error) {
	err = service.OpenaiApiModel().Edit(ctx, &req.ApiModelEditInp)
	return
}

// MaxSort 获取AI模型列表最大排序
func (c *cApiModel) MaxSort(ctx context.Context, req *apimodel.MaxSortReq) (res *apimodel.MaxSortRes, err error) {
	data, err := service.OpenaiApiModel().MaxSort(ctx, &req.ApiModelMaxSortInp)
	if err != nil {
		return
	}

	res = new(apimodel.MaxSortRes)
	res.ApiModelMaxSortModel = data
	return
}

// View 获取指定AI模型列表信息
func (c *cApiModel) View(ctx context.Context, req *apimodel.ViewReq) (res *apimodel.ViewRes, err error) {
	data, err := service.OpenaiApiModel().View(ctx, &req.ApiModelViewInp)
	if err != nil {
		return
	}

	res = new(apimodel.ViewRes)
	res.ApiModelViewModel = data
	return
}

// Delete 删除AI模型列表
func (c *cApiModel) Delete(ctx context.Context, req *apimodel.DeleteReq) (res *apimodel.DeleteRes, err error) {
	err = service.OpenaiApiModel().Delete(ctx, &req.ApiModelDeleteInp)
	return
}

// Status 更新AI模型列表状态
func (c *cApiModel) Status(ctx context.Context, req *apimodel.StatusReq) (res *apimodel.StatusRes, err error) {
	err = service.OpenaiApiModel().Status(ctx, &req.ApiModelStatusInp)
	return
}