// Package im
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package im

import (
	"context"
	"hotgo/api/admin/im/message"
	"hotgo/internal/model/input/imin"
	"hotgo/internal/service"
)

var (
	Message = cMessage{}
)

type cMessage struct{}

// List 查看IM-聊天记录表列表
func (c *cMessage) List(ctx context.Context, req *message.ListReq) (res *message.ListRes, err error) {
	list, totalCount, err := service.ImMessage().List(ctx, &req.MessageListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*imin.MessageListModel{}
	}

	res = new(message.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出IM-聊天记录表列表
func (c *cMessage) Export(ctx context.Context, req *message.ExportReq) (res *message.ExportRes, err error) {
	err = service.ImMessage().Export(ctx, &req.MessageListInp)
	return
}

// Edit 更新IM-聊天记录表
func (c *cMessage) Edit(ctx context.Context, req *message.EditReq) (res *message.EditRes, err error) {
	err = service.ImMessage().Edit(ctx, &req.MessageEditInp)
	return
}

// View 获取指定IM-聊天记录表信息
func (c *cMessage) View(ctx context.Context, req *message.ViewReq) (res *message.ViewRes, err error) {
	data, err := service.ImMessage().View(ctx, &req.MessageViewInp)
	if err != nil {
		return
	}

	res = new(message.ViewRes)
	res.MessageViewModel = data
	return
}

// Delete 删除IM-聊天记录表
func (c *cMessage) Delete(ctx context.Context, req *message.DeleteReq) (res *message.DeleteRes, err error) {
	err = service.ImMessage().Delete(ctx, &req.MessageDeleteInp)
	return
}

// Status 更新IM-聊天记录表状态
func (c *cMessage) Status(ctx context.Context, req *message.StatusReq) (res *message.StatusRes, err error) {
	err = service.ImMessage().Status(ctx, &req.MessageStatusInp)
	return
}