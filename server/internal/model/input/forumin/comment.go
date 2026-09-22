// Package forumin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package forumin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CommentUpdateFields 修改评论管理字段过滤
type CommentUpdateFields struct {
	TopicId     int64  `json:"topicId"     dc:"主题ID"`
	UserId      int64  `json:"userId"      dc:"评论用户ID"`
	ParentId    int64  `json:"parentId"    dc:"父评论ID"`
	ReplyUserId int64  `json:"replyUserId" dc:"被回复用户ID"`
	Content     string `json:"content"     dc:"评论内容"`
	Status      string `json:"status"      dc:"状态:normal,pending,rejected,hidden"`
	LikeCount   int    `json:"likeCount"   dc:"点赞数量"`
}

// CommentInsertFields 新增评论管理字段过滤
type CommentInsertFields struct {
	TopicId     int64  `json:"topicId"     dc:"主题ID"`
	UserId      int64  `json:"userId"      dc:"评论用户ID"`
	ParentId    int64  `json:"parentId"    dc:"父评论ID"`
	ReplyUserId int64  `json:"replyUserId" dc:"被回复用户ID"`
	Content     string `json:"content"     dc:"评论内容"`
	Status      string `json:"status"      dc:"状态:normal,pending,rejected,hidden"`
	LikeCount   int    `json:"likeCount"   dc:"点赞数量"`
}

// CommentEditInp 修改/新增评论管理
type CommentEditInp struct {
	entity.ForumComment
}

func (in *CommentEditInp) Filter(ctx context.Context) (err error) {
	// 验证主题ID
	if err := g.Validator().Rules("required").Data(in.TopicId).Messages("主题ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证评论用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("评论用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证父评论ID
	if err := g.Validator().Rules("required").Data(in.ParentId).Messages("父评论ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证被回复用户ID
	if err := g.Validator().Rules("required").Data(in.ReplyUserId).Messages("被回复用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证评论内容
	if err := g.Validator().Rules("required").Data(in.Content).Messages("评论内容不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态:normal,pending,rejected,hidden
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态:normal,pending,rejected,hidden不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证点赞数量
	if err := g.Validator().Rules("required").Data(in.LikeCount).Messages("点赞数量不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type CommentEditModel struct{}

// CommentDeleteInp 删除评论管理
type CommentDeleteInp struct {
	Id interface{} `json:"id" v:"required#评论ID不能为空" dc:"评论ID"`
}

func (in *CommentDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CommentDeleteModel struct{}

// CommentViewInp 获取指定评论管理信息
type CommentViewInp struct {
	Id int64 `json:"id" v:"required#评论ID不能为空" dc:"评论ID"`
}

func (in *CommentViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CommentViewModel struct {
	entity.ForumComment
}

// CommentListInp 获取评论管理列表
type CommentListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"评论ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *CommentListInp) Filter(ctx context.Context) (err error) {
	return
}

type CommentListModel struct {
	Id          int64       `json:"id"          dc:"评论ID"`
	TopicId     int64       `json:"topicId"     dc:"主题ID"`
	UserId      int64       `json:"userId"      dc:"评论用户ID"`
	ParentId    int64       `json:"parentId"    dc:"父评论ID"`
	ReplyUserId int64       `json:"replyUserId" dc:"被回复用户ID"`
	Status      string      `json:"status"      dc:"状态:normal,pending,rejected,hidden"`
	LikeCount   int         `json:"likeCount"   dc:"点赞数量"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
}

// CommentExportModel 导出评论管理
type CommentExportModel struct {
	Id          int64       `json:"id"          dc:"评论ID"`
	TopicId     int64       `json:"topicId"     dc:"主题ID"`
	UserId      int64       `json:"userId"      dc:"评论用户ID"`
	ParentId    int64       `json:"parentId"    dc:"父评论ID"`
	ReplyUserId int64       `json:"replyUserId" dc:"被回复用户ID"`
	Status      string      `json:"status"      dc:"状态:normal,pending,rejected,hidden"`
	LikeCount   int         `json:"likeCount"   dc:"点赞数量"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
}