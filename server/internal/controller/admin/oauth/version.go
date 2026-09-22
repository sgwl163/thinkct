// Package oauth
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package oauth

import (
	"context"
	"hotgo/api/admin/oauth/version"
	"hotgo/internal/model/input/oauthin"
	"hotgo/internal/service"
)

var (
	Version = cVersion{}
)

type cVersion struct{}

// List 查看版本更新列表
func (c *cVersion) List(ctx context.Context, req *version.ListReq) (res *version.ListRes, err error) {
	list, totalCount, err := service.OauthVersion().List(ctx, &req.VersionListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*oauthin.VersionListModel{}
	}

	res = new(version.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出版本更新列表
func (c *cVersion) Export(ctx context.Context, req *version.ExportReq) (res *version.ExportRes, err error) {
	err = service.OauthVersion().Export(ctx, &req.VersionListInp)
	return
}

// Edit 更新版本更新
func (c *cVersion) Edit(ctx context.Context, req *version.EditReq) (res *version.EditRes, err error) {
	err = service.OauthVersion().Edit(ctx, &req.VersionEditInp)
	return
}

// View 获取指定版本更新信息
func (c *cVersion) View(ctx context.Context, req *version.ViewReq) (res *version.ViewRes, err error) {
	data, err := service.OauthVersion().View(ctx, &req.VersionViewInp)
	if err != nil {
		return
	}

	res = new(version.ViewRes)
	res.VersionViewModel = data
	return
}

// Delete 删除版本更新
func (c *cVersion) Delete(ctx context.Context, req *version.DeleteReq) (res *version.DeleteRes, err error) {
	err = service.OauthVersion().Delete(ctx, &req.VersionDeleteInp)
	return
}

// Status 更新版本更新状态
func (c *cVersion) Status(ctx context.Context, req *version.StatusReq) (res *version.StatusRes, err error) {
	err = service.OauthVersion().Status(ctx, &req.VersionStatusInp)
	return
}