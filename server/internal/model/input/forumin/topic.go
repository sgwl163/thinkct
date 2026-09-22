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

// TopicUpdateFields 修改社区主题字段过滤
type TopicUpdateFields struct {
	CategoryId      int64       `json:"categoryId"      dc:"板块ID"`
	UserId          int64       `json:"userId"          dc:"发布用户ID"`
	Title           string      `json:"title"           dc:"主题标题"`
	Keywords        string      `json:"keywords"        dc:"主题关键词，英文逗号分隔"`
	Description     string      `json:"description"     dc:"主题摘要"`
	Content         string      `json:"content"         dc:"主题内容"`
	Images          string      `json:"images"          dc:"主题图片，英文逗号分隔"`
	Status          string      `json:"status"          dc:"状态:normal,pending,rejected,hidden,closed"`
	IsTop           int         `json:"isTop"           dc:"是否置顶:0否,1是"`
	IsEssence       int         `json:"isEssence"       dc:"是否精华:0否,1是"`
	CommentCount    int         `json:"commentCount"    dc:"评论数量"`
	LikeCount       int         `json:"likeCount"       dc:"点赞数量"`
	FavoriteCount   int         `json:"favoriteCount"   dc:"收藏数量"`
	ViewCount       int         `json:"viewCount"       dc:"浏览数量"`
	LastReplyUserId int64       `json:"lastReplyUserId" dc:"最后回复用户ID"`
	LastReplyAt     *gtime.Time `json:"lastReplyAt"     dc:"最后回复时间"`
}

// TopicInsertFields 新增社区主题字段过滤
type TopicInsertFields struct {
	CategoryId      int64       `json:"categoryId"      dc:"板块ID"`
	UserId          int64       `json:"userId"          dc:"发布用户ID"`
	Title           string      `json:"title"           dc:"主题标题"`
	Keywords        string      `json:"keywords"        dc:"主题关键词，英文逗号分隔"`
	Description     string      `json:"description"     dc:"主题摘要"`
	Content         string      `json:"content"         dc:"主题内容"`
	Images          string      `json:"images"          dc:"主题图片，英文逗号分隔"`
	Status          string      `json:"status"          dc:"状态:normal,pending,rejected,hidden,closed"`
	IsTop           int         `json:"isTop"           dc:"是否置顶:0否,1是"`
	IsEssence       int         `json:"isEssence"       dc:"是否精华:0否,1是"`
	CommentCount    int         `json:"commentCount"    dc:"评论数量"`
	LikeCount       int         `json:"likeCount"       dc:"点赞数量"`
	FavoriteCount   int         `json:"favoriteCount"   dc:"收藏数量"`
	ViewCount       int         `json:"viewCount"       dc:"浏览数量"`
	LastReplyUserId int64       `json:"lastReplyUserId" dc:"最后回复用户ID"`
	LastReplyAt     *gtime.Time `json:"lastReplyAt"     dc:"最后回复时间"`
}

// TopicEditInp 修改/新增社区主题
type TopicEditInp struct {
	entity.ForumTopic
}

func (in *TopicEditInp) Filter(ctx context.Context) (err error) {
	// 验证板块ID
	if err := g.Validator().Rules("required").Data(in.CategoryId).Messages("板块ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证发布用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("发布用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证主题标题
	if err := g.Validator().Rules("required").Data(in.Title).Messages("主题标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证主题关键词，英文逗号分隔
	if err := g.Validator().Rules("required").Data(in.Keywords).Messages("主题关键词，英文逗号分隔不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证主题摘要
	if err := g.Validator().Rules("required").Data(in.Description).Messages("主题摘要不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证主题内容
	if err := g.Validator().Rules("required").Data(in.Content).Messages("主题内容不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证主题图片，英文逗号分隔
	if err := g.Validator().Rules("required").Data(in.Images).Messages("主题图片，英文逗号分隔不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态:normal,pending,rejected,hidden,closed
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态:normal,pending,rejected,hidden,closed不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证是否置顶:0否,1是
	if err := g.Validator().Rules("required").Data(in.IsTop).Messages("是否置顶:0否,1是不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证是否精华:0否,1是
	if err := g.Validator().Rules("required").Data(in.IsEssence).Messages("是否精华:0否,1是不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证评论数量
	if err := g.Validator().Rules("required").Data(in.CommentCount).Messages("评论数量不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证点赞数量
	if err := g.Validator().Rules("required").Data(in.LikeCount).Messages("点赞数量不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证收藏数量
	if err := g.Validator().Rules("required").Data(in.FavoriteCount).Messages("收藏数量不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证浏览数量
	if err := g.Validator().Rules("required").Data(in.ViewCount).Messages("浏览数量不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证最后回复用户ID
	if err := g.Validator().Rules("required").Data(in.LastReplyUserId).Messages("最后回复用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type topicEditModel struct{}

// TopicDeleteInp 删除社区主题
type TopicDeleteInp struct {
	Id interface{} `json:"id" v:"required#主题ID不能为空" dc:"主题ID"`
}

func (in *TopicDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type topicDeleteModel struct{}

// TopicViewInp 获取指定社区主题信息
type TopicViewInp struct {
	Id int64 `json:"id" v:"required#主题ID不能为空" dc:"主题ID"`
}

func (in *TopicViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TopicViewModel struct {
	entity.ForumTopic
}

// TopicListInp 获取社区主题列表
type TopicListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"主题ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *TopicListInp) Filter(ctx context.Context) (err error) {
	return
}

type TopicListModel struct {
	Id              int64       `json:"id"              dc:"主题ID"`
	CategoryId      int64       `json:"categoryId"      dc:"板块ID"`
	UserId          int64       `json:"userId"          dc:"发布用户ID"`
	Title           string      `json:"title"           dc:"主题标题"`
	Status          string      `json:"status"          dc:"状态:normal,pending,rejected,hidden,closed"`
	IsTop           int         `json:"isTop"           dc:"是否置顶:0否,1是"`
	IsEssence       int         `json:"isEssence"       dc:"是否精华:0否,1是"`
	CommentCount    int         `json:"commentCount"    dc:"评论数量"`
	LikeCount       int         `json:"likeCount"       dc:"点赞数量"`
	FavoriteCount   int         `json:"favoriteCount"   dc:"收藏数量"`
	ViewCount       int         `json:"viewCount"       dc:"浏览数量"`
	LastReplyUserId int64       `json:"lastReplyUserId" dc:"最后回复用户ID"`
	LastReplyAt     *gtime.Time `json:"lastReplyAt"     dc:"最后回复时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       dc:"修改时间"`
	CreatedAt       *gtime.Time `json:"createdAt"       dc:"创建时间"`
}

// TopicExportModel 导出社区主题
type TopicExportModel struct {
	Id              int64       `json:"id"              dc:"主题ID"`
	CategoryId      int64       `json:"categoryId"      dc:"板块ID"`
	UserId          int64       `json:"userId"          dc:"发布用户ID"`
	Title           string      `json:"title"           dc:"主题标题"`
	Status          string      `json:"status"          dc:"状态:normal,pending,rejected,hidden,closed"`
	IsTop           int         `json:"isTop"           dc:"是否置顶:0否,1是"`
	IsEssence       int         `json:"isEssence"       dc:"是否精华:0否,1是"`
	CommentCount    int         `json:"commentCount"    dc:"评论数量"`
	LikeCount       int         `json:"likeCount"       dc:"点赞数量"`
	FavoriteCount   int         `json:"favoriteCount"   dc:"收藏数量"`
	ViewCount       int         `json:"viewCount"       dc:"浏览数量"`
	LastReplyUserId int64       `json:"lastReplyUserId" dc:"最后回复用户ID"`
	LastReplyAt     *gtime.Time `json:"lastReplyAt"     dc:"最后回复时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       dc:"修改时间"`
	CreatedAt       *gtime.Time `json:"createdAt"       dc:"创建时间"`
}
