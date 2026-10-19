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

// List 查看权限规则列表
func (c *cUserRule) List(ctx context.Context, req *userrule.ListReq) (res *userrule.ListRes, err error) {
	list, totalCount, err := service.UserUserRule().List(ctx, &req.UserRuleListInp)
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

// Export 导出权限规则列表
func (c *cUserRule) Export(ctx context.Context, req *userrule.ExportReq) (res *userrule.ExportRes, err error) {
	err = service.UserUserRule().Export(ctx, &req.UserRuleListInp)
	return
}

// Edit 更新权限规则
func (c *cUserRule) Edit(ctx context.Context, req *userrule.EditReq) (res *userrule.EditRes, err error) {
	err = service.UserUserRule().Edit(ctx, &req.UserRuleEditInp)
	return
}

// View 获取指定权限规则信息
func (c *cUserRule) View(ctx context.Context, req *userrule.ViewReq) (res *userrule.ViewRes, err error) {
	data, err := service.UserUserRule().View(ctx, &req.UserRuleViewInp)
	if err != nil {
		return
	}

	res = new(userrule.ViewRes)
	res.UserRuleViewModel = data
	return
}

// Delete 删除权限规则
func (c *cUserRule) Delete(ctx context.Context, req *userrule.DeleteReq) (res *userrule.DeleteRes, err error) {
	err = service.UserUserRule().Delete(ctx, &req.UserRuleDeleteInp)
	return
}