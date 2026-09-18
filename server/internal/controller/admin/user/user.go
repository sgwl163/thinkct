// Package user
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package user

import (
	"context"
	"hotgo/api/admin/user/user"
	"hotgo/internal/model/input/userin"
	"hotgo/internal/service"
)

var (
	User = cUser{}
)

type cUser struct{}

// List 查看会员管理列表
func (c *cUser) List(ctx context.Context, req *user.ListReq) (res *user.ListRes, err error) {
	list, totalCount, err := service.UserUser().List(ctx, &req.UserListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*userin.UserListModel{}
	}

	res = new(user.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出会员管理列表
func (c *cUser) Export(ctx context.Context, req *user.ExportReq) (res *user.ExportRes, err error) {
	err = service.UserUser().Export(ctx, &req.UserListInp)
	return
}

// Edit 更新会员管理
func (c *cUser) Edit(ctx context.Context, req *user.EditReq) (res *user.EditRes, err error) {
	err = service.UserUser().Edit(ctx, &req.UserEditInp)
	return
}

// View 获取指定会员管理信息
func (c *cUser) View(ctx context.Context, req *user.ViewReq) (res *user.ViewRes, err error) {
	data, err := service.UserUser().View(ctx, &req.UserViewInp)
	if err != nil {
		return
	}

	res = new(user.ViewRes)
	res.UserViewModel = data
	return
}

// Delete 删除会员管理
func (c *cUser) Delete(ctx context.Context, req *user.DeleteReq) (res *user.DeleteRes, err error) {
	err = service.UserUser().Delete(ctx, &req.UserDeleteInp)
	return
}

// Status 更新会员管理状态
func (c *cUser) Status(ctx context.Context, req *user.StatusReq) (res *user.StatusRes, err error) {
	err = service.UserUser().Status(ctx, &req.UserStatusInp)
	return
}

// Switch 更新会员管理开关状态
func (c *cUser) Switch(ctx context.Context, req *user.SwitchReq) (res *user.SwitchRes, err error) {
	err = service.UserUser().Switch(ctx, &req.UserSwitchInp)
	return
}