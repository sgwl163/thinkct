package site

import (
	"context"

	v1 "hotgo/api/api/site/v1"
	"hotgo/internal/model/input/sitein"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Site = cSite{}
)

type cSite struct{}

func NewV1() *cSite {
	return &cSite{}
}

// Index 获取首页数据
func (c *cSite) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	data, err := service.ApiSite().Index(ctx, sitein.IndexInp{})
	if err != nil {
		return nil, err
	}

	res = new(v1.IndexRes)
	res.SiteConfig = data.SiteConfig
	return res, nil
}

// Home 获取首页聚合数据
func (c *cSite) Home(ctx context.Context, req *v1.HomeReq) (res *v1.HomeRes, err error) {
	data, err := service.ApiSite().Home(ctx, sitein.HomeInp{
		ProductsChannelIds: req.ProductsChannelIds,
		ProductsLimit:      req.ProductsLimit,
		CasesChannelIds:    req.CasesChannelIds,
		CasesLimit:         req.CasesLimit,
		NewsChannelIds:     req.NewsChannelIds,
		NewsLimit:          req.NewsLimit,
	})
	if err != nil {
		return nil, err
	}

	res = new(v1.HomeRes)
	res.ProductsChannels = data.ProductsChannels
	res.CasesChannels = data.CasesChannels
	res.NewsChannels = data.NewsChannels
	res.Links = data.Links
	return res, nil
}

// Channel 获取栏目列表与单栏目详情内容
func (c *cSite) Channel(ctx context.Context, req *v1.ChannelReq) (res *v1.ChannelRes, err error) {
	data, err := service.ApiSite().Channel(ctx, sitein.ChannelInp{
		Channel: req.Channel,
		Page:    req.Page,
		Limit:   req.Limit,
	})
	if err != nil {
		return nil, err
	}
	res = new(v1.ChannelRes)
	res.Channel = data.Channel
	res.Breadcrumb = data.Breadcrumb
	res.List = data.List
	res.Total = data.Total
	res.Page = data.Page
	res.RelatedList = data.RelatedList
	res.HotList = data.HotList
	return res, nil
}

// Content 获取内容详情
func (c *cSite) Content(ctx context.Context, req *v1.ContentReq) (res *v1.ContentRes, err error) {
	data, err := service.ApiSite().Content(ctx, sitein.ContentInp{
		Id:      req.Id,
		Channel: req.Channel,
	})
	if err != nil {
		return nil, err
	}
	res = new(v1.ContentRes)
	res.Content = data.Content
	res.Author = data.Author
	res.Channel = data.Channel
	res.Breadcrumb = data.Breadcrumb
	res.RelatedList = data.RelatedList
	res.HotList = data.HotList
	res.Prev = data.Prev
	res.Next = data.Next
	return res, nil
}

// Search 全站内容搜索
func (c *cSite) Search(ctx context.Context, req *v1.SearchReq) (res *v1.SearchRes, err error) {
	data, err := service.ApiSite().Search(ctx, sitein.SearchInp{
		Keywords: req.Keywords,
		Type:     req.Type,
		Page:     req.Page,
		Limit:    req.Limit,
	})
	if err != nil {
		return nil, err
	}
	res = new(v1.SearchRes)
	res.List = data.List
	res.Total = data.Total
	res.Page = data.Page
	res.Limit = data.Limit
	return res, nil
}

// Feedback 提交在线留言
func (c *cSite) Feedback(ctx context.Context, req *v1.FeedbackReq) (res *v1.FeedbackRes, err error) {
	r := g.RequestFromCtx(ctx)
	clientIp := ""
	if r != nil {
		clientIp = r.GetClientIp()
	}

	_, err = service.ApiSite().Feedback(ctx, sitein.FeedbackInp{
		Name:                 req.Name,
		Phone:                req.Phone,
		Email:                req.Email,
		Company:              req.Company,
		MessageTopic:         req.MessageTopic,
		Content:              req.Content,
		ExpectedResponseTime: req.ExpectedResponseTime,
		Ip:                   clientIp,
	})
	if err != nil {
		return nil, err
	}
	return &v1.FeedbackRes{}, nil
}
