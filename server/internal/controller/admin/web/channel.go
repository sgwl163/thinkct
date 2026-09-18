// Package web
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package web

import (
	"context"
	"hotgo/api/admin/web/channel"
	"hotgo/internal/model/input/webin"
	"hotgo/internal/service"
)

var (
	Channel = cChannel{}
)

type cChannel struct{}

// List 查看栏目管理列表
func (c *cChannel) List(ctx context.Context, req *channel.ListReq) (res *channel.ListRes, err error) {
	list, totalCount, err := service.WebChannel().List(ctx, &req.ChannelListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*webin.ChannelListModel{}
	}

	res = new(channel.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出栏目管理列表
func (c *cChannel) Export(ctx context.Context, req *channel.ExportReq) (res *channel.ExportRes, err error) {
	err = service.WebChannel().Export(ctx, &req.ChannelListInp)
	return
}

// Edit 更新栏目管理
func (c *cChannel) Edit(ctx context.Context, req *channel.EditReq) (res *channel.EditRes, err error) {
	err = service.WebChannel().Edit(ctx, &req.ChannelEditInp)
	return
}

// View 获取指定栏目管理信息
func (c *cChannel) View(ctx context.Context, req *channel.ViewReq) (res *channel.ViewRes, err error) {
	data, err := service.WebChannel().View(ctx, &req.ChannelViewInp)
	if err != nil {
		return
	}

	res = new(channel.ViewRes)
	res.ChannelViewModel = data
	return
}

// Delete 删除栏目管理
func (c *cChannel) Delete(ctx context.Context, req *channel.DeleteReq) (res *channel.DeleteRes, err error) {
	err = service.WebChannel().Delete(ctx, &req.ChannelDeleteInp)
	return
}

// Status 更新栏目管理状态
func (c *cChannel) Status(ctx context.Context, req *channel.StatusReq) (res *channel.StatusRes, err error) {
	err = service.WebChannel().Status(ctx, &req.ChannelStatusInp)
	return
}

// TreeOption 获取栏目管理关系树选项
func (c *cChannel) TreeOption(ctx context.Context, req *channel.TreeOptionReq) (res *channel.TreeOptionRes, err error) {
	data, err := service.WebChannel().TreeOption(ctx, req.Model)
	if err != nil {
		return nil, err
	}

	if len(data) > 0 {
		res = (*channel.TreeOptionRes)(&data)
	} else {
		temp := make(channel.TreeOptionRes, 0)
		res = &temp
	}
	return
}