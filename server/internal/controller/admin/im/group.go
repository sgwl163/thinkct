// Package im
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package im

import (
	"context"
	"hotgo/api/admin/im/group"
	"hotgo/internal/model/input/imin"
	"hotgo/internal/service"
)

var (
	Group = cGroup{}
)

type cGroup struct{}

// List 查看IM-群组表列表
func (c *cGroup) List(ctx context.Context, req *group.ListReq) (res *group.ListRes, err error) {
	list, totalCount, err := service.ImGroup().List(ctx, &req.GroupListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*imin.GroupListModel{}
	}

	res = new(group.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出IM-群组表列表
func (c *cGroup) Export(ctx context.Context, req *group.ExportReq) (res *group.ExportRes, err error) {
	err = service.ImGroup().Export(ctx, &req.GroupListInp)
	return
}

// Edit 更新IM-群组表
func (c *cGroup) Edit(ctx context.Context, req *group.EditReq) (res *group.EditRes, err error) {
	err = service.ImGroup().Edit(ctx, &req.GroupEditInp)
	return
}

// View 获取指定IM-群组表信息
func (c *cGroup) View(ctx context.Context, req *group.ViewReq) (res *group.ViewRes, err error) {
	data, err := service.ImGroup().View(ctx, &req.GroupViewInp)
	if err != nil {
		return
	}

	res = new(group.ViewRes)
	res.GroupViewModel = data
	return
}

// Delete 删除IM-群组表
func (c *cGroup) Delete(ctx context.Context, req *group.DeleteReq) (res *group.DeleteRes, err error) {
	err = service.ImGroup().Delete(ctx, &req.GroupDeleteInp)
	return
}

// Status 更新IM-群组表状态
func (c *cGroup) Status(ctx context.Context, req *group.StatusReq) (res *group.StatusRes, err error) {
	err = service.ImGroup().Status(ctx, &req.GroupStatusInp)
	return
}