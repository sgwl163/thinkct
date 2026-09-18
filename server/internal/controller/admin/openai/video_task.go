// Package openai
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package openai

import (
	"context"
	"hotgo/api/admin/openai/videotask"
	"hotgo/internal/model/input/openaiin"
	"hotgo/internal/service"
)

var (
	VideoTask = cVideoTask{}
)

type cVideoTask struct{}

// List 查看AI视频任务列表
func (c *cVideoTask) List(ctx context.Context, req *videotask.ListReq) (res *videotask.ListRes, err error) {
	list, totalCount, err := service.OpenaiVideoTask().List(ctx, &req.VideoTaskListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*openaiin.VideoTaskListModel{}
	}

	res = new(videotask.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出AI视频任务列表
func (c *cVideoTask) Export(ctx context.Context, req *videotask.ExportReq) (res *videotask.ExportRes, err error) {
	err = service.OpenaiVideoTask().Export(ctx, &req.VideoTaskListInp)
	return
}

// Edit 更新AI视频任务
func (c *cVideoTask) Edit(ctx context.Context, req *videotask.EditReq) (res *videotask.EditRes, err error) {
	err = service.OpenaiVideoTask().Edit(ctx, &req.VideoTaskEditInp)
	return
}

// View 获取指定AI视频任务信息
func (c *cVideoTask) View(ctx context.Context, req *videotask.ViewReq) (res *videotask.ViewRes, err error) {
	data, err := service.OpenaiVideoTask().View(ctx, &req.VideoTaskViewInp)
	if err != nil {
		return
	}

	res = new(videotask.ViewRes)
	res.VideoTaskViewModel = data
	return
}

// Delete 删除AI视频任务
func (c *cVideoTask) Delete(ctx context.Context, req *videotask.DeleteReq) (res *videotask.DeleteRes, err error) {
	err = service.OpenaiVideoTask().Delete(ctx, &req.VideoTaskDeleteInp)
	return
}

// Status 更新AI视频任务状态
func (c *cVideoTask) Status(ctx context.Context, req *videotask.StatusReq) (res *videotask.StatusRes, err error) {
	err = service.OpenaiVideoTask().Status(ctx, &req.VideoTaskStatusInp)
	return
}