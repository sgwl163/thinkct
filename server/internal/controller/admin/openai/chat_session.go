// Package openai
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package openai

import (
	"context"
	"hotgo/api/admin/openai/chatsession"
	"hotgo/internal/model/input/openaiin"
	"hotgo/internal/service"
)

var (
	ChatSession = cChatSession{}
)

type cChatSession struct{}

// List 查看聊天会话列表
func (c *cChatSession) List(ctx context.Context, req *chatsession.ListReq) (res *chatsession.ListRes, err error) {
	list, totalCount, err := service.OpenaiChatSession().List(ctx, &req.ChatSessionListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*openaiin.ChatSessionListModel{}
	}

	res = new(chatsession.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出聊天会话列表
func (c *cChatSession) Export(ctx context.Context, req *chatsession.ExportReq) (res *chatsession.ExportRes, err error) {
	err = service.OpenaiChatSession().Export(ctx, &req.ChatSessionListInp)
	return
}

// Edit 更新聊天会话
func (c *cChatSession) Edit(ctx context.Context, req *chatsession.EditReq) (res *chatsession.EditRes, err error) {
	err = service.OpenaiChatSession().Edit(ctx, &req.ChatSessionEditInp)
	return
}

// View 获取指定聊天会话信息
func (c *cChatSession) View(ctx context.Context, req *chatsession.ViewReq) (res *chatsession.ViewRes, err error) {
	data, err := service.OpenaiChatSession().View(ctx, &req.ChatSessionViewInp)
	if err != nil {
		return
	}

	res = new(chatsession.ViewRes)
	res.ChatSessionViewModel = data
	return
}

// Delete 删除聊天会话
func (c *cChatSession) Delete(ctx context.Context, req *chatsession.DeleteReq) (res *chatsession.DeleteRes, err error) {
	err = service.OpenaiChatSession().Delete(ctx, &req.ChatSessionDeleteInp)
	return
}

// Status 更新聊天会话状态
func (c *cChatSession) Status(ctx context.Context, req *chatsession.StatusReq) (res *chatsession.StatusRes, err error) {
	err = service.OpenaiChatSession().Status(ctx, &req.ChatSessionStatusInp)
	return
}