// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package sys

import (
	"context"
	"hotgo/addons/appAuth/api/admin/applog"
	"hotgo/addons/appAuth/model/input/sysin"
	"hotgo/addons/appAuth/service"
)

var (
	AppLog = cAppLog{}
)

type cAppLog struct{}

// List 查看升级记录列表
func (c *cAppLog) List(ctx context.Context, req *applog.ListReq) (res *applog.ListRes, err error) {
	list, totalCount, err := service.SysAppLog().List(ctx, &req.AppLogListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.AppLogListModel{}
	}

	res = new(applog.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出升级记录列表
func (c *cAppLog) Export(ctx context.Context, req *applog.ExportReq) (res *applog.ExportRes, err error) {
	err = service.SysAppLog().Export(ctx, &req.AppLogListInp)
	return
}

// Edit 更新升级记录
func (c *cAppLog) Edit(ctx context.Context, req *applog.EditReq) (res *applog.EditRes, err error) {
	err = service.SysAppLog().Edit(ctx, &req.AppLogEditInp)
	return
}

// View 获取指定升级记录信息
func (c *cAppLog) View(ctx context.Context, req *applog.ViewReq) (res *applog.ViewRes, err error) {
	data, err := service.SysAppLog().View(ctx, &req.AppLogViewInp)
	if err != nil {
		return
	}

	res = new(applog.ViewRes)
	res.AppLogViewModel = data
	return
}

// Delete 删除升级记录
func (c *cAppLog) Delete(ctx context.Context, req *applog.DeleteReq) (res *applog.DeleteRes, err error) {
	err = service.SysAppLog().Delete(ctx, &req.AppLogDeleteInp)
	return
}