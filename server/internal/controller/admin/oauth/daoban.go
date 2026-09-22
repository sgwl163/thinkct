// Package oauth
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package oauth

import (
	"context"
	"hotgo/api/admin/oauth/daoban"
	"hotgo/internal/model/input/oauthin"
	"hotgo/internal/service"
)

var (
	Daoban = cDaoban{}
)

type cDaoban struct{}

// List 查看盗版管理列表
func (c *cDaoban) List(ctx context.Context, req *daoban.ListReq) (res *daoban.ListRes, err error) {
	list, totalCount, err := service.OauthDaoban().List(ctx, &req.DaobanListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*oauthin.DaobanListModel{}
	}

	res = new(daoban.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出盗版管理列表
func (c *cDaoban) Export(ctx context.Context, req *daoban.ExportReq) (res *daoban.ExportRes, err error) {
	err = service.OauthDaoban().Export(ctx, &req.DaobanListInp)
	return
}

// Edit 更新盗版管理
func (c *cDaoban) Edit(ctx context.Context, req *daoban.EditReq) (res *daoban.EditRes, err error) {
	err = service.OauthDaoban().Edit(ctx, &req.DaobanEditInp)
	return
}

// View 获取指定盗版管理信息
func (c *cDaoban) View(ctx context.Context, req *daoban.ViewReq) (res *daoban.ViewRes, err error) {
	data, err := service.OauthDaoban().View(ctx, &req.DaobanViewInp)
	if err != nil {
		return
	}

	res = new(daoban.ViewRes)
	res.DaobanViewModel = data
	return
}

// Delete 删除盗版管理
func (c *cDaoban) Delete(ctx context.Context, req *daoban.DeleteReq) (res *daoban.DeleteRes, err error) {
	err = service.OauthDaoban().Delete(ctx, &req.DaobanDeleteInp)
	return
}