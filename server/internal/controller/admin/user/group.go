// Package user
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package user

import (
	"context"
	group "hotgo/api/admin/user/group"
	"hotgo/internal/model/input/userin"
	"hotgo/internal/service"
)

var (
	Group = cGroup{}
)

type cGroup struct{}

// List 查看分组管理列表
func (c *cGroup) List(ctx context.Context, req *group.ListReq) (res *group.ListRes, err error) {
	list, totalCount, err := service.Usergroup().List(ctx, &req.GroupListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*userin.GroupListModel{}
	}

	res = new(group.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出分组管理列表
func (c *cGroup) Export(ctx context.Context, req *group.ExportReq) (res *group.ExportRes, err error) {
	err = service.Usergroup().Export(ctx, &req.GroupListInp)
	return
}

// Edit 更新分组管理
func (c *cGroup) Edit(ctx context.Context, req *group.EditReq) (res *group.EditRes, err error) {
	err = service.Usergroup().Edit(ctx, &req.GroupEditInp)
	return
}

// View 获取指定分组管理信息
func (c *cGroup) View(ctx context.Context, req *group.ViewReq) (res *group.ViewRes, err error) {
	data, err := service.Usergroup().View(ctx, &req.GroupViewInp)
	if err != nil {
		return
	}

	res = new(group.ViewRes)
	res.GroupViewModel = data
	return
}

// Delete 删除分组管理
func (c *cGroup) Delete(ctx context.Context, req *group.DeleteReq) (res *group.DeleteRes, err error) {
	err = service.Usergroup().Delete(ctx, &req.GroupDeleteInp)
	return
}

// Status 更新分组管理状态
func (c *cGroup) Status(ctx context.Context, req *group.StatusReq) (res *group.StatusRes, err error) {
	err = service.Usergroup().Status(ctx, &req.GroupStatusInp)
	return
}
