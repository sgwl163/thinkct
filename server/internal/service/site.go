// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/model/input/sitein"
)

type (
	IApiSite interface {
		// Index 获取首页数据
		Index(ctx context.Context, in sitein.IndexInp) (res *sitein.IndexModel, err error)
		// Home 获取首页各个栏目的聚合数据
		Home(ctx context.Context, in sitein.HomeInp) (res *sitein.HomeModel, err error)
		// Channel 获取前台栏目列表与详情
		Channel(ctx context.Context, in sitein.ChannelInp) (res *sitein.ChannelModel, err error)
		// Content 获取内容详情
		Content(ctx context.Context, in sitein.ContentInp) (res *sitein.ContentModel, err error)
		// Search 全站内容搜索
		Search(ctx context.Context, in sitein.SearchInp) (res *sitein.SearchModel, err error)
		// Feedback 提交在线留言
		Feedback(ctx context.Context, in sitein.FeedbackInp) (res *sitein.FeedbackModel, err error)
	}
)

var (
	localApiSite IApiSite
)

func ApiSite() IApiSite {
	if localApiSite == nil {
		panic("implement not found for interface IApiSite, forgot register?")
	}
	return localApiSite
}

func RegisterApiSite(i IApiSite) {
	localApiSite = i
}
