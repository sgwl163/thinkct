// Package im
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package im

import (
	"context"
	"hotgo/api/admin/im/conversation"
	"hotgo/internal/model/input/imin"
	"hotgo/internal/service"
)

var (
	Conversation = cConversation{}
)

type cConversation struct{}

// List 查看会话列表管理列表
func (c *cConversation) List(ctx context.Context, req *conversation.ListReq) (res *conversation.ListRes, err error) {
	list, totalCount, err := service.ImConversation().List(ctx, &req.ConversationListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*imin.ConversationListModel{}
	}

	res = new(conversation.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出会话列表管理列表
func (c *cConversation) Export(ctx context.Context, req *conversation.ExportReq) (res *conversation.ExportRes, err error) {
	err = service.ImConversation().Export(ctx, &req.ConversationListInp)
	return
}

// Edit 更新会话列表管理
func (c *cConversation) Edit(ctx context.Context, req *conversation.EditReq) (res *conversation.EditRes, err error) {
	err = service.ImConversation().Edit(ctx, &req.ConversationEditInp)
	return
}

// View 获取指定会话列表管理信息
func (c *cConversation) View(ctx context.Context, req *conversation.ViewReq) (res *conversation.ViewRes, err error) {
	data, err := service.ImConversation().View(ctx, &req.ConversationViewInp)
	if err != nil {
		return
	}

	res = new(conversation.ViewRes)
	res.ConversationViewModel = data
	return
}

// Delete 删除会话列表管理
func (c *cConversation) Delete(ctx context.Context, req *conversation.DeleteReq) (res *conversation.DeleteRes, err error) {
	err = service.ImConversation().Delete(ctx, &req.ConversationDeleteInp)
	return
}