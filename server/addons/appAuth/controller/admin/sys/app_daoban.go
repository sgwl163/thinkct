// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package sys

import (
	"context"
	"hotgo/addons/appAuth/api/admin/appdaoban"
	"hotgo/addons/appAuth/model/input/sysin"
	"hotgo/addons/appAuth/service"
)

var (
	AppDaoban = cAppDaoban{}
)

type cAppDaoban struct{}

// List 查看盗版管理列表
func (c *cAppDaoban) List(ctx context.Context, req *appdaoban.ListReq) (res *appdaoban.ListRes, err error) {
	list, totalCount, err := service.SysAppDaoban().List(ctx, &req.AppDaobanListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.AppDaobanListModel{}
	}

	res = new(appdaoban.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出盗版管理列表
func (c *cAppDaoban) Export(ctx context.Context, req *appdaoban.ExportReq) (res *appdaoban.ExportRes, err error) {
	err = service.SysAppDaoban().Export(ctx, &req.AppDaobanListInp)
	return
}

// Edit 更新盗版管理
func (c *cAppDaoban) Edit(ctx context.Context, req *appdaoban.EditReq) (res *appdaoban.EditRes, err error) {
	err = service.SysAppDaoban().Edit(ctx, &req.AppDaobanEditInp)
	return
}

// View 获取指定盗版管理信息
func (c *cAppDaoban) View(ctx context.Context, req *appdaoban.ViewReq) (res *appdaoban.ViewRes, err error) {
	data, err := service.SysAppDaoban().View(ctx, &req.AppDaobanViewInp)
	if err != nil {
		return
	}

	res = new(appdaoban.ViewRes)
	res.AppDaobanViewModel = data
	return
}

// Delete 删除盗版管理
func (c *cAppDaoban) Delete(ctx context.Context, req *appdaoban.DeleteReq) (res *appdaoban.DeleteRes, err error) {
	err = service.SysAppDaoban().Delete(ctx, &req.AppDaobanDeleteInp)
	return
}