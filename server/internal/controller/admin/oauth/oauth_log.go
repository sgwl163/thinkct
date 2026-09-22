// Package oauth
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package oauth

import (
	"context"
	"hotgo/api/admin/oauth/oauthlog"
	"hotgo/internal/model/input/oauthin"
	"hotgo/internal/service"
)

var (
	OauthLog = cOauthLog{}
)

type cOauthLog struct{}

// List 查看升级记录列表
func (c *cOauthLog) List(ctx context.Context, req *oauthlog.ListReq) (res *oauthlog.ListRes, err error) {
	list, totalCount, err := service.OauthOauthLog().List(ctx, &req.OauthLogListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*oauthin.OauthLogListModel{}
	}

	res = new(oauthlog.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出升级记录列表
func (c *cOauthLog) Export(ctx context.Context, req *oauthlog.ExportReq) (res *oauthlog.ExportRes, err error) {
	err = service.OauthOauthLog().Export(ctx, &req.OauthLogListInp)
	return
}

// Edit 更新升级记录
func (c *cOauthLog) Edit(ctx context.Context, req *oauthlog.EditReq) (res *oauthlog.EditRes, err error) {
	err = service.OauthOauthLog().Edit(ctx, &req.OauthLogEditInp)
	return
}

// View 获取指定升级记录信息
func (c *cOauthLog) View(ctx context.Context, req *oauthlog.ViewReq) (res *oauthlog.ViewRes, err error) {
	data, err := service.OauthOauthLog().View(ctx, &req.OauthLogViewInp)
	if err != nil {
		return
	}

	res = new(oauthlog.ViewRes)
	res.OauthLogViewModel = data
	return
}

// Delete 删除升级记录
func (c *cOauthLog) Delete(ctx context.Context, req *oauthlog.DeleteReq) (res *oauthlog.DeleteRes, err error) {
	err = service.OauthOauthLog().Delete(ctx, &req.OauthLogDeleteInp)
	return
}