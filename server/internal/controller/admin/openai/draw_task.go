// Package openai
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package openai

import (
	"context"
	"hotgo/api/admin/openai/drawtask"
	"hotgo/internal/model/input/openaiin"
	"hotgo/internal/service"
)

var (
	DrawTask = cDrawTask{}
)

type cDrawTask struct{}

// List 查看AI绘画任务列表
func (c *cDrawTask) List(ctx context.Context, req *drawtask.ListReq) (res *drawtask.ListRes, err error) {
	list, totalCount, err := service.OpenaiDrawTask().List(ctx, &req.DrawTaskListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*openaiin.DrawTaskListModel{}
	}

	res = new(drawtask.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出AI绘画任务列表
func (c *cDrawTask) Export(ctx context.Context, req *drawtask.ExportReq) (res *drawtask.ExportRes, err error) {
	err = service.OpenaiDrawTask().Export(ctx, &req.DrawTaskListInp)
	return
}

// Edit 更新AI绘画任务
func (c *cDrawTask) Edit(ctx context.Context, req *drawtask.EditReq) (res *drawtask.EditRes, err error) {
	err = service.OpenaiDrawTask().Edit(ctx, &req.DrawTaskEditInp)
	return
}

// View 获取指定AI绘画任务信息
func (c *cDrawTask) View(ctx context.Context, req *drawtask.ViewReq) (res *drawtask.ViewRes, err error) {
	data, err := service.OpenaiDrawTask().View(ctx, &req.DrawTaskViewInp)
	if err != nil {
		return
	}

	res = new(drawtask.ViewRes)
	res.DrawTaskViewModel = data
	return
}

// Delete 删除AI绘画任务
func (c *cDrawTask) Delete(ctx context.Context, req *drawtask.DeleteReq) (res *drawtask.DeleteRes, err error) {
	err = service.OpenaiDrawTask().Delete(ctx, &req.DrawTaskDeleteInp)
	return
}

// Status 更新AI绘画任务状态
func (c *cDrawTask) Status(ctx context.Context, req *drawtask.StatusReq) (res *drawtask.StatusRes, err error) {
	err = service.OpenaiDrawTask().Status(ctx, &req.DrawTaskStatusInp)
	return
}