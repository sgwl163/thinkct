// Package app
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package app

import (
	"context"
	"hotgo/api/admin/app/appauth"
	"hotgo/internal/model/input/appin"
	"hotgo/internal/service"
)

var (
	AppAuth = cAppAuth{}
)

type cAppAuth struct{}

// List 查看授权管理列表
func (c *cAppAuth) List(ctx context.Context, req *appauth.ListReq) (res *appauth.ListRes, err error) {
	list, totalCount, err := service.AppAppAuth().List(ctx, &req.AppAuthListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*appin.AppAuthListModel{}
	}

	res = new(appauth.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出授权管理列表
func (c *cAppAuth) Export(ctx context.Context, req *appauth.ExportReq) (res *appauth.ExportRes, err error) {
	err = service.AppAppAuth().Export(ctx, &req.AppAuthListInp)
	return
}

// Edit 更新授权管理
func (c *cAppAuth) Edit(ctx context.Context, req *appauth.EditReq) (res *appauth.EditRes, err error) {
	err = service.AppAppAuth().Edit(ctx, &req.AppAuthEditInp)
	return
}

// View 获取指定授权管理信息
func (c *cAppAuth) View(ctx context.Context, req *appauth.ViewReq) (res *appauth.ViewRes, err error) {
	data, err := service.AppAppAuth().View(ctx, &req.AppAuthViewInp)
	if err != nil {
		return
	}

	res = new(appauth.ViewRes)
	res.AppAuthViewModel = data
	return
}

// Delete 删除授权管理
func (c *cAppAuth) Delete(ctx context.Context, req *appauth.DeleteReq) (res *appauth.DeleteRes, err error) {
	err = service.AppAppAuth().Delete(ctx, &req.AppAuthDeleteInp)
	return
}

// Status 更新授权管理状态
func (c *cAppAuth) Status(ctx context.Context, req *appauth.StatusReq) (res *appauth.StatusRes, err error) {
	err = service.AppAppAuth().Status(ctx, &req.AppAuthStatusInp)
	return
}