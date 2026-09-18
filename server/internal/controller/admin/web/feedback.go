// Package web
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package web

import (
	"context"
	"hotgo/api/admin/web/feedback"
	"hotgo/internal/model/input/webin"
	"hotgo/internal/service"
)

var (
	Feedback = cFeedback{}
)

type cFeedback struct{}

// List 查看留言管理列表
func (c *cFeedback) List(ctx context.Context, req *feedback.ListReq) (res *feedback.ListRes, err error) {
	list, totalCount, err := service.WebFeedback().List(ctx, &req.FeedbackListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*webin.FeedbackListModel{}
	}

	res = new(feedback.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出留言管理列表
func (c *cFeedback) Export(ctx context.Context, req *feedback.ExportReq) (res *feedback.ExportRes, err error) {
	err = service.WebFeedback().Export(ctx, &req.FeedbackListInp)
	return
}

// Edit 更新留言管理
func (c *cFeedback) Edit(ctx context.Context, req *feedback.EditReq) (res *feedback.EditRes, err error) {
	err = service.WebFeedback().Edit(ctx, &req.FeedbackEditInp)
	return
}

// View 获取指定留言管理信息
func (c *cFeedback) View(ctx context.Context, req *feedback.ViewReq) (res *feedback.ViewRes, err error) {
	data, err := service.WebFeedback().View(ctx, &req.FeedbackViewInp)
	if err != nil {
		return
	}

	res = new(feedback.ViewRes)
	res.FeedbackViewModel = data
	return
}

// Delete 删除留言管理
func (c *cFeedback) Delete(ctx context.Context, req *feedback.DeleteReq) (res *feedback.DeleteRes, err error) {
	err = service.WebFeedback().Delete(ctx, &req.FeedbackDeleteInp)
	return
}

// Status 更新留言管理状态
func (c *cFeedback) Status(ctx context.Context, req *feedback.StatusReq) (res *feedback.StatusRes, err error) {
	err = service.WebFeedback().Status(ctx, &req.FeedbackStatusInp)
	return
}