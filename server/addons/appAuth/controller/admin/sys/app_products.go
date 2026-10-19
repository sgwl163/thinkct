// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package sys

import (
	"context"
	"hotgo/addons/appAuth/api/admin/appproducts"
	"hotgo/addons/appAuth/model/input/sysin"
	"hotgo/addons/appAuth/service"
)

var (
	AppProducts = cAppProducts{}
)

type cAppProducts struct{}

// List 查看产品管理列表
func (c *cAppProducts) List(ctx context.Context, req *appproducts.ListReq) (res *appproducts.ListRes, err error) {
	list, totalCount, err := service.SysAppProducts().List(ctx, &req.AppProductsListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.AppProductsListModel{}
	}

	res = new(appproducts.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出产品管理列表
func (c *cAppProducts) Export(ctx context.Context, req *appproducts.ExportReq) (res *appproducts.ExportRes, err error) {
	err = service.SysAppProducts().Export(ctx, &req.AppProductsListInp)
	return
}

// Edit 更新产品管理
func (c *cAppProducts) Edit(ctx context.Context, req *appproducts.EditReq) (res *appproducts.EditRes, err error) {
	err = service.SysAppProducts().Edit(ctx, &req.AppProductsEditInp)
	return
}

// View 获取指定产品管理信息
func (c *cAppProducts) View(ctx context.Context, req *appproducts.ViewReq) (res *appproducts.ViewRes, err error) {
	data, err := service.SysAppProducts().View(ctx, &req.AppProductsViewInp)
	if err != nil {
		return
	}

	res = new(appproducts.ViewRes)
	res.AppProductsViewModel = data
	return
}

// Delete 删除产品管理
func (c *cAppProducts) Delete(ctx context.Context, req *appproducts.DeleteReq) (res *appproducts.DeleteRes, err error) {
	err = service.SysAppProducts().Delete(ctx, &req.AppProductsDeleteInp)
	return
}

// Status 更新产品管理状态
func (c *cAppProducts) Status(ctx context.Context, req *appproducts.StatusReq) (res *appproducts.StatusRes, err error) {
	err = service.SysAppProducts().Status(ctx, &req.AppProductsStatusInp)
	return
}