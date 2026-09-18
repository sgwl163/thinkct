// Package openai
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package openai

import (
	"context"
	"hotgo/api/admin/openai/musictask"
	"hotgo/internal/model/input/openaiin"
	"hotgo/internal/service"
)

var (
	MusicTask = cMusicTask{}
)

type cMusicTask struct{}

// List 查看AI音乐任务列表
func (c *cMusicTask) List(ctx context.Context, req *musictask.ListReq) (res *musictask.ListRes, err error) {
	list, totalCount, err := service.OpenaiMusicTask().List(ctx, &req.MusicTaskListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*openaiin.MusicTaskListModel{}
	}

	res = new(musictask.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出AI音乐任务列表
func (c *cMusicTask) Export(ctx context.Context, req *musictask.ExportReq) (res *musictask.ExportRes, err error) {
	err = service.OpenaiMusicTask().Export(ctx, &req.MusicTaskListInp)
	return
}

// Edit 更新AI音乐任务
func (c *cMusicTask) Edit(ctx context.Context, req *musictask.EditReq) (res *musictask.EditRes, err error) {
	err = service.OpenaiMusicTask().Edit(ctx, &req.MusicTaskEditInp)
	return
}

// View 获取指定AI音乐任务信息
func (c *cMusicTask) View(ctx context.Context, req *musictask.ViewReq) (res *musictask.ViewRes, err error) {
	data, err := service.OpenaiMusicTask().View(ctx, &req.MusicTaskViewInp)
	if err != nil {
		return
	}

	res = new(musictask.ViewRes)
	res.MusicTaskViewModel = data
	return
}

// Delete 删除AI音乐任务
func (c *cMusicTask) Delete(ctx context.Context, req *musictask.DeleteReq) (res *musictask.DeleteRes, err error) {
	err = service.OpenaiMusicTask().Delete(ctx, &req.MusicTaskDeleteInp)
	return
}

// Status 更新AI音乐任务状态
func (c *cMusicTask) Status(ctx context.Context, req *musictask.StatusReq) (res *musictask.StatusRes, err error) {
	err = service.OpenaiMusicTask().Status(ctx, &req.MusicTaskStatusInp)
	return
}