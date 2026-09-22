// Package forum
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package forum

import (
	"context"
	"hotgo/api/admin/forum/topic"
	"hotgo/internal/model/input/forumin"
	"hotgo/internal/service"
)

var (
	ForumTopic = cForumTopic{}
)

type cForumTopic struct{}

// List 查看社区主题列表
func (c *cForumTopic) List(ctx context.Context, req *topic.ListReq) (res *topic.ListRes, err error) {
	list, totalCount, err := service.ForumTopic().List(ctx, &req.TopicListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*forumin.TopicListModel{}
	}

	res = new(topic.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出社区主题列表
func (c *cForumTopic) Export(ctx context.Context, req *topic.ExportReq) (res *topic.ExportRes, err error) {
	err = service.ForumTopic().Export(ctx, &req.TopicListInp)
	return
}

// Edit 更新社区主题
func (c *cForumTopic) Edit(ctx context.Context, req *topic.EditReq) (res *topic.EditRes, err error) {
	err = service.ForumTopic().Edit(ctx, &req.TopicEditInp)
	return
}

// View 获取指定社区主题信息
func (c *cForumTopic) View(ctx context.Context, req *topic.ViewReq) (res *topic.ViewRes, err error) {
	data, err := service.ForumTopic().View(ctx, &req.TopicViewInp)
	if err != nil {
		return
	}

	res = new(topic.ViewRes)
	res.TopicViewModel = data
	return
}

// Delete 删除社区主题
func (c *cForumTopic) Delete(ctx context.Context, req *topic.DeleteReq) (res *topic.DeleteRes, err error) {
	err = service.ForumTopic().Delete(ctx, &req.TopicDeleteInp)
	return
}
