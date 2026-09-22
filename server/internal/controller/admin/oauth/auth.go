// Package oauth
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package oauth

import (
	"context"
	"hotgo/api/admin/oauth/auth"
	"hotgo/internal/model/input/oauthin"
	"hotgo/internal/service"
)

var (
	Auth = cAuth{}
)

type cAuth struct{}

// List 查看授权管理列表
func (c *cAuth) List(ctx context.Context, req *auth.ListReq) (res *auth.ListRes, err error) {
	list, totalCount, err := service.OauthAuth().List(ctx, &req.AuthListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*oauthin.AuthListModel{}
	}

	res = new(auth.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出授权管理列表
func (c *cAuth) Export(ctx context.Context, req *auth.ExportReq) (res *auth.ExportRes, err error) {
	err = service.OauthAuth().Export(ctx, &req.AuthListInp)
	return
}

// Edit 更新授权管理
func (c *cAuth) Edit(ctx context.Context, req *auth.EditReq) (res *auth.EditRes, err error) {
	err = service.OauthAuth().Edit(ctx, &req.AuthEditInp)
	return
}

// View 获取指定授权管理信息
func (c *cAuth) View(ctx context.Context, req *auth.ViewReq) (res *auth.ViewRes, err error) {
	data, err := service.OauthAuth().View(ctx, &req.AuthViewInp)
	if err != nil {
		return
	}

	res = new(auth.ViewRes)
	res.AuthViewModel = data
	return
}

// Delete 删除授权管理
func (c *cAuth) Delete(ctx context.Context, req *auth.DeleteReq) (res *auth.DeleteRes, err error) {
	err = service.OauthAuth().Delete(ctx, &req.AuthDeleteInp)
	return
}

// Status 更新授权管理状态
func (c *cAuth) Status(ctx context.Context, req *auth.StatusReq) (res *auth.StatusRes, err error) {
	err = service.OauthAuth().Status(ctx, &req.AuthStatusInp)
	return
}