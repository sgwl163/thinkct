// Package user
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package user

import (
	"context"
	"hotgo/api/admin/user/userrule"
	"hotgo/internal/model/input/userin"
	"hotgo/internal/service"
)

var (
	UserRule = cUserRule{}
)

type cUserRule struct{}

// List 查看规则管理列表
func (c *cUserRule) List(ctx context.Context, req *userrule.ListReq) (res *userrule.ListRes, err error) {
	list, totalCount, err := service.UserRule().List(ctx, &req.UserRuleListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*userin.UserRuleListModel{}
	}

	res = new(userrule.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出规则管理列表
func (c *cUserRule) Export(ctx context.Context, req *userrule.ExportReq) (res *userrule.ExportRes, err error) {
	err = service.UserRule().Export(ctx, &req.UserRuleListInp)
	return
}

// Edit 更新规则管理
func (c *cUserRule) Edit(ctx context.Context, req *userrule.EditReq) (res *userrule.EditRes, err error) {
	err = service.UserRule().Edit(ctx, &req.UserRuleEditInp)
	return
}

// View 获取指定规则管理信息
func (c *cUserRule) View(ctx context.Context, req *userrule.ViewReq) (res *userrule.ViewRes, err error) {
	data, err := service.UserRule().View(ctx, &req.UserRuleViewInp)
	if err != nil {
		return
	}

	res = new(userrule.ViewRes)
	res.UserRuleViewModel = data
	return
}

// Delete 删除规则管理
func (c *cUserRule) Delete(ctx context.Context, req *userrule.DeleteReq) (res *userrule.DeleteRes, err error) {
	err = service.UserRule().Delete(ctx, &req.UserRuleDeleteInp)
	return
}

// Status 更新规则管理状态
func (c *cUserRule) Status(ctx context.Context, req *userrule.StatusReq) (res *userrule.StatusRes, err error) {
	err = service.UserRule().Status(ctx, &req.UserRuleStatusInp)
	return
}