// Package user
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package user

import (
	"context"
	"hotgo/api/admin/user/userlog"
	"hotgo/internal/model/input/userin"
	"hotgo/internal/service"
)

var (
	UserLog = cUserLog{}
)

type cUserLog struct{}

// List 查看操作记录列表
func (c *cUserLog) List(ctx context.Context, req *userlog.ListReq) (res *userlog.ListRes, err error) {
	list, totalCount, err := service.UserUserLog().List(ctx, &req.UserLogListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*userin.UserLogListModel{}
	}

	res = new(userlog.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出操作记录列表
func (c *cUserLog) Export(ctx context.Context, req *userlog.ExportReq) (res *userlog.ExportRes, err error) {
	err = service.UserUserLog().Export(ctx, &req.UserLogListInp)
	return
}

// Edit 更新操作记录
func (c *cUserLog) Edit(ctx context.Context, req *userlog.EditReq) (res *userlog.EditRes, err error) {
	err = service.UserUserLog().Edit(ctx, &req.UserLogEditInp)
	return
}

// View 获取指定操作记录信息
func (c *cUserLog) View(ctx context.Context, req *userlog.ViewReq) (res *userlog.ViewRes, err error) {
	data, err := service.UserUserLog().View(ctx, &req.UserLogViewInp)
	if err != nil {
		return
	}

	res = new(userlog.ViewRes)
	res.UserLogViewModel = data
	return
}

// Delete 删除操作记录
func (c *cUserLog) Delete(ctx context.Context, req *userlog.DeleteReq) (res *userlog.DeleteRes, err error) {
	err = service.UserUserLog().Delete(ctx, &req.UserLogDeleteInp)
	return
}

// Status 更新操作记录状态
func (c *cUserLog) Status(ctx context.Context, req *userlog.StatusReq) (res *userlog.StatusRes, err error) {
	err = service.UserUserLog().Status(ctx, &req.UserLogStatusInp)
	return
}