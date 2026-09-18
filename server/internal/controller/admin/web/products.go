// Package web
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package web

import (
	"context"
	"hotgo/api/admin/web/products"
	"hotgo/internal/model/input/webin"
	"hotgo/internal/service"
)

var (
	Products = cProducts{}
)

type cProducts struct{}

// List 查看产品管理列表
func (c *cProducts) List(ctx context.Context, req *products.ListReq) (res *products.ListRes, err error) {
	list, totalCount, err := service.WebProducts().List(ctx, &req.ProductsListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*webin.ProductsListModel{}
	}

	res = new(products.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出产品管理列表
func (c *cProducts) Export(ctx context.Context, req *products.ExportReq) (res *products.ExportRes, err error) {
	err = service.WebProducts().Export(ctx, &req.ProductsListInp)
	return
}

// Edit 更新产品管理
func (c *cProducts) Edit(ctx context.Context, req *products.EditReq) (res *products.EditRes, err error) {
	err = service.WebProducts().Edit(ctx, &req.ProductsEditInp)
	return
}

// View 获取指定产品管理信息
func (c *cProducts) View(ctx context.Context, req *products.ViewReq) (res *products.ViewRes, err error) {
	data, err := service.WebProducts().View(ctx, &req.ProductsViewInp)
	if err != nil {
		return
	}

	res = new(products.ViewRes)
	res.ProductsViewModel = data
	return
}

// Delete 删除产品管理
func (c *cProducts) Delete(ctx context.Context, req *products.DeleteReq) (res *products.DeleteRes, err error) {
	err = service.WebProducts().Delete(ctx, &req.ProductsDeleteInp)
	return
}

// Status 更新产品管理状态
func (c *cProducts) Status(ctx context.Context, req *products.StatusReq) (res *products.StatusRes, err error) {
	err = service.WebProducts().Status(ctx, &req.ProductsStatusInp)
	return
}