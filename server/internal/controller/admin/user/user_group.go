// Package user
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package user

import (
	"context"
	"hotgo/api/admin/user/usergroup"
	"hotgo/internal/model/input/userin"
	"hotgo/internal/service"
)

var (
	UserGroup = cUserGroup{}
)

type cUserGroup struct{}

// List 查看用户分组列表
func (c *cUserGroup) List(ctx context.Context, req *usergroup.ListReq) (res *usergroup.ListRes, err error) {
	list, totalCount, err := service.UserUserGroup().List(ctx, &req.UserGroupListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*userin.UserGroupListModel{}
	}

	res = new(usergroup.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出用户分组列表
func (c *cUserGroup) Export(ctx context.Context, req *usergroup.ExportReq) (res *usergroup.ExportRes, err error) {
	err = service.UserUserGroup().Export(ctx, &req.UserGroupListInp)
	return
}

// Edit 更新用户分组
func (c *cUserGroup) Edit(ctx context.Context, req *usergroup.EditReq) (res *usergroup.EditRes, err error) {
	err = service.UserUserGroup().Edit(ctx, &req.UserGroupEditInp)
	return
}

// View 获取指定用户分组信息
func (c *cUserGroup) View(ctx context.Context, req *usergroup.ViewReq) (res *usergroup.ViewRes, err error) {
	data, err := service.UserUserGroup().View(ctx, &req.UserGroupViewInp)
	if err != nil {
		return
	}

	res = new(usergroup.ViewRes)
	res.UserGroupViewModel = data
	return
}

// Delete 删除用户分组
func (c *cUserGroup) Delete(ctx context.Context, req *usergroup.DeleteReq) (res *usergroup.DeleteRes, err error) {
	err = service.UserUserGroup().Delete(ctx, &req.UserGroupDeleteInp)
	return
}