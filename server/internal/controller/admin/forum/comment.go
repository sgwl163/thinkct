// Package forum
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package forum

import (
	"context"
	"hotgo/api/admin/forum/comment"
	"hotgo/internal/model/input/forumin"
	"hotgo/internal/service"
)

var (
	Comment = cComment{}
)

type cComment struct{}

// List 查看评论管理列表
func (c *cComment) List(ctx context.Context, req *comment.ListReq) (res *comment.ListRes, err error) {
	list, totalCount, err := service.ForumComment().List(ctx, &req.CommentListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*forumin.CommentListModel{}
	}

	res = new(comment.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出评论管理列表
func (c *cComment) Export(ctx context.Context, req *comment.ExportReq) (res *comment.ExportRes, err error) {
	err = service.ForumComment().Export(ctx, &req.CommentListInp)
	return
}

// Edit 更新评论管理
func (c *cComment) Edit(ctx context.Context, req *comment.EditReq) (res *comment.EditRes, err error) {
	err = service.ForumComment().Edit(ctx, &req.CommentEditInp)
	return
}

// View 获取指定评论管理信息
func (c *cComment) View(ctx context.Context, req *comment.ViewReq) (res *comment.ViewRes, err error) {
	data, err := service.ForumComment().View(ctx, &req.CommentViewInp)
	if err != nil {
		return
	}

	res = new(comment.ViewRes)
	res.CommentViewModel = data
	return
}

// Delete 删除评论管理
func (c *cComment) Delete(ctx context.Context, req *comment.DeleteReq) (res *comment.DeleteRes, err error) {
	err = service.ForumComment().Delete(ctx, &req.CommentDeleteInp)
	return
}