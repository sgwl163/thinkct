// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package sys

import (
	"context"
	"hotgo/addons/appAuth/api/admin/appversion"
	"hotgo/addons/appAuth/model/input/sysin"
	"hotgo/addons/appAuth/service"
)

var (
	AppVersion = cAppVersion{}
)

type cAppVersion struct{}

// List 查看版本更新列表
func (c *cAppVersion) List(ctx context.Context, req *appversion.ListReq) (res *appversion.ListRes, err error) {
	list, totalCount, err := service.SysAppVersion().List(ctx, &req.AppVersionListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*sysin.AppVersionListModel{}
	}

	res = new(appversion.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出版本更新列表
func (c *cAppVersion) Export(ctx context.Context, req *appversion.ExportReq) (res *appversion.ExportRes, err error) {
	err = service.SysAppVersion().Export(ctx, &req.AppVersionListInp)
	return
}

// Edit 更新版本更新
func (c *cAppVersion) Edit(ctx context.Context, req *appversion.EditReq) (res *appversion.EditRes, err error) {
	err = service.SysAppVersion().Edit(ctx, &req.AppVersionEditInp)
	return
}

// View 获取指定版本更新信息
func (c *cAppVersion) View(ctx context.Context, req *appversion.ViewReq) (res *appversion.ViewRes, err error) {
	data, err := service.SysAppVersion().View(ctx, &req.AppVersionViewInp)
	if err != nil {
		return
	}

	res = new(appversion.ViewRes)
	res.AppVersionViewModel = data
	return
}

// Delete 删除版本更新
func (c *cAppVersion) Delete(ctx context.Context, req *appversion.DeleteReq) (res *appversion.DeleteRes, err error) {
	err = service.SysAppVersion().Delete(ctx, &req.AppVersionDeleteInp)
	return
}

// Status 更新版本更新状态
func (c *cAppVersion) Status(ctx context.Context, req *appversion.StatusReq) (res *appversion.StatusRes, err error) {
	err = service.SysAppVersion().Status(ctx, &req.AppVersionStatusInp)
	return
}