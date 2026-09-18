// Package openai
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package openai

import (
	"context"
	"hotgo/api/admin/openai/openaiappcategory"
	"hotgo/internal/model/input/openaiin"
	"hotgo/internal/service"
)

var (
	OpenaiAppCategory = cOpenaiAppCategory{}
)

type cOpenaiAppCategory struct{}

// List 查看AI应用分类列表
func (c *cOpenaiAppCategory) List(ctx context.Context, req *openaiappcategory.ListReq) (res *openaiappcategory.ListRes, err error) {
	list, totalCount, err := service.OpenaiAppCategory().List(ctx, &req.OpenaiAppCategoryListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*openaiin.OpenaiAppCategoryListModel{}
	}

	res = new(openaiappcategory.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出AI应用分类列表
func (c *cOpenaiAppCategory) Export(ctx context.Context, req *openaiappcategory.ExportReq) (res *openaiappcategory.ExportRes, err error) {
	err = service.OpenaiAppCategory().Export(ctx, &req.OpenaiAppCategoryListInp)
	return
}

// Edit 更新AI应用分类
func (c *cOpenaiAppCategory) Edit(ctx context.Context, req *openaiappcategory.EditReq) (res *openaiappcategory.EditRes, err error) {
	err = service.OpenaiAppCategory().Edit(ctx, &req.OpenaiAppCategoryEditInp)
	return
}

// MaxSort 获取AI应用分类最大排序
func (c *cOpenaiAppCategory) MaxSort(ctx context.Context, req *openaiappcategory.MaxSortReq) (res *openaiappcategory.MaxSortRes, err error) {
	data, err := service.OpenaiAppCategory().MaxSort(ctx, &req.OpenaiAppCategoryMaxSortInp)
	if err != nil {
		return
	}

	res = new(openaiappcategory.MaxSortRes)
	res.OpenaiAppCategoryMaxSortModel = data
	return
}

// View 获取指定AI应用分类信息
func (c *cOpenaiAppCategory) View(ctx context.Context, req *openaiappcategory.ViewReq) (res *openaiappcategory.ViewRes, err error) {
	data, err := service.OpenaiAppCategory().View(ctx, &req.OpenaiAppCategoryViewInp)
	if err != nil {
		return
	}

	res = new(openaiappcategory.ViewRes)
	res.OpenaiAppCategoryViewModel = data
	return
}

// Delete 删除AI应用分类
func (c *cOpenaiAppCategory) Delete(ctx context.Context, req *openaiappcategory.DeleteReq) (res *openaiappcategory.DeleteRes, err error) {
	err = service.OpenaiAppCategory().Delete(ctx, &req.OpenaiAppCategoryDeleteInp)
	return
}

// Status 更新AI应用分类状态
func (c *cOpenaiAppCategory) Status(ctx context.Context, req *openaiappcategory.StatusReq) (res *openaiappcategory.StatusRes, err error) {
	err = service.OpenaiAppCategory().Status(ctx, &req.OpenaiAppCategoryStatusInp)
	return
}