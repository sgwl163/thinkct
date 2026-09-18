package service

import (
	"context"
	"hotgo/internal/model/input/sitein"
)

type (
	IApiSite interface {
		Index(ctx context.Context, in sitein.IndexInp) (res *sitein.IndexModel, err error)
		Home(ctx context.Context, in sitein.HomeInp) (res *sitein.HomeModel, err error)
		Channel(ctx context.Context, in sitein.ChannelInp) (res *sitein.ChannelModel, err error)
		Content(ctx context.Context, in sitein.ContentInp) (res *sitein.ContentModel, err error)
		Search(ctx context.Context, in sitein.SearchInp) (res *sitein.SearchModel, err error)
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
