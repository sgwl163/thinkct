// Package user
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package user

import (
	"context"
	"hotgo/api/admin/user/usercreditslog"
	"hotgo/internal/model/input/userin"
	"hotgo/internal/service"
)

var (
	UserCreditsLog = cUserCreditsLog{}
)

type cUserCreditsLog struct{}

// List 查看资产明细列表
func (c *cUserCreditsLog) List(ctx context.Context, req *usercreditslog.ListReq) (res *usercreditslog.ListRes, err error) {
	list, totalCount, err := service.UserUserCreditsLog().List(ctx, &req.UserCreditsLogListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*userin.UserCreditsLogListModel{}
	}

	res = new(usercreditslog.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出资产明细列表
func (c *cUserCreditsLog) Export(ctx context.Context, req *usercreditslog.ExportReq) (res *usercreditslog.ExportRes, err error) {
	err = service.UserUserCreditsLog().Export(ctx, &req.UserCreditsLogListInp)
	return
}

// Edit 更新资产明细
func (c *cUserCreditsLog) Edit(ctx context.Context, req *usercreditslog.EditReq) (res *usercreditslog.EditRes, err error) {
	err = service.UserUserCreditsLog().Edit(ctx, &req.UserCreditsLogEditInp)
	return
}

// View 获取指定资产明细信息
func (c *cUserCreditsLog) View(ctx context.Context, req *usercreditslog.ViewReq) (res *usercreditslog.ViewRes, err error) {
	data, err := service.UserUserCreditsLog().View(ctx, &req.UserCreditsLogViewInp)
	if err != nil {
		return
	}

	res = new(usercreditslog.ViewRes)
	res.UserCreditsLogViewModel = data
	return
}

// Delete 删除资产明细
func (c *cUserCreditsLog) Delete(ctx context.Context, req *usercreditslog.DeleteReq) (res *usercreditslog.DeleteRes, err error) {
	err = service.UserUserCreditsLog().Delete(ctx, &req.UserCreditsLogDeleteInp)
	return
}

// Status 更新资产明细状态
func (c *cUserCreditsLog) Status(ctx context.Context, req *usercreditslog.StatusReq) (res *usercreditslog.StatusRes, err error) {
	err = service.UserUserCreditsLog().Status(ctx, &req.UserCreditsLogStatusInp)
	return
}