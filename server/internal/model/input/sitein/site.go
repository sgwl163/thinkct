package sitein

import (
	"hotgo/internal/model/entity"
)

// IndexInp 获取首页数据入参
type IndexInp struct {
}

// IndexModel 获取首页数据出参
type IndexModel struct {
	SiteConfig map[string]interface{} `json:"siteConfig"`
	NavMenu    []*NavMenuItem         `json:"navMenu"`
}

// NavMenuItem 前台导航树节点，仅包含展示所需字段。
type NavMenuItem struct {
	Id       int64          `json:"id"`
	Name     string         `json:"name"`
	Url      string         `json:"url"`
	Image    string         `json:"image"`
	ImageAlt string         `json:"imageAlt"`
	Children []*NavMenuItem `json:"children"`
}

// HomeInp 获取首页聚合数据入参
type HomeInp struct {
	ProductsChannelIds string `json:"productsChannelIds"`
	ProductsLimit      int    `json:"productsLimit"`
	CasesChannelIds    string `json:"casesChannelIds"`
	CasesLimit         int    `json:"casesLimit"`
	NewsChannelIds     string `json:"newsChannelIds"`
	NewsLimit          int    `json:"newsLimit"`
}

// ProductChannelModel 带有产品列表的分类结构
type ProductChannelModel struct {
	*entity.WebChannel
	List []*entity.ShopProduct `json:"list"`
}

// ArticleChannelModel 带有文章列表的分类结构
type ArticleChannelModel struct {
	*entity.WebChannel
	List []*entity.WebArticle `json:"list"`
}

// CarouselItem 首页轮播展示数据，不暴露管理字段。
type CarouselItem struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CtaText     string `json:"ctaText"`
	CtaLink     string `json:"ctaLink"`
	Color       string `json:"color"`
	Image       string `json:"image"`
	Alt         string `json:"alt"`
	AspectRatio string `json:"aspectRatio"`
}

// HomeModel 获取首页聚合数据出参
type HomeModel struct {
	Carousels        []*CarouselItem        `json:"carousels"`
	ProductsChannels []*ProductChannelModel `json:"productsChannels"`
	CasesChannels    []*ArticleChannelModel `json:"casesChannels"`
	NewsChannels     []*ArticleChannelModel `json:"newsChannels"`
	Links            []*entity.WebLinks     `json:"links"`
}

// ChannelInp 获取前台栏目列表入参
type ChannelInp struct {
	Channel string
	Page    int
	Limit   int
}

// ChannelModel 获取前台栏目列表出参
type ChannelModel struct {
	Channel     *entity.WebChannel   `json:"channel"`
	Breadcrumb  []*entity.WebChannel `json:"breadcrumb"`
	List        interface{}          `json:"list"`
	Total       int                  `json:"total"`
	Page        int                  `json:"page"`
	RelatedList []*entity.WebChannel `json:"relatedList"`
	HotList     interface{}          `json:"hotList"`
}

// ContentInp 获取内容详情入参
type ContentInp struct {
	Id      uint
	Channel string
}

// ContentAuthor 内容作者公开信息
type ContentAuthor struct {
	Id            uint   `json:"id" dc:"作者ID"`
	Username      string `json:"username" dc:"作者用户名"`
	Nickname      string `json:"nickname" dc:"作者昵称"`
	Avatar        string `json:"avatar" dc:"作者头像"`
	Motto         string `json:"motto" dc:"作者签名"`
	IsAdmin       bool   `json:"isAdmin" dc:"是否管理员作者"`
	ArticleCount  int    `json:"articleCount" dc:"已发布文章数"`
	ViewCount     int    `json:"viewCount" dc:"已发布文章阅读数"`
	FollowerCount int    `json:"followerCount" dc:"粉丝数，当前未启用关注功能时为0"`
}

// ContentModel 获取内容详情出参
type ContentModel struct {
	Content     interface{}          `json:"content"`
	Author      *ContentAuthor       `json:"author"`
	Channel     *entity.WebChannel   `json:"channel"`
	Breadcrumb  []*entity.WebChannel `json:"breadcrumb"`
	RelatedList []*entity.WebChannel `json:"relatedList"`
	HotList     interface{}          `json:"hotList"`
	Prev        interface{}          `json:"prev"`
	Next        interface{}          `json:"next"`
}

// SearchInp 搜索入参
type SearchInp struct {
	Keywords string
	Type     string
	Page     int
	Limit    int
}

// SearchModel 搜索出参
type SearchModel struct {
	List  interface{} `json:"list"`
	Total int         `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}

// FeedbackInp 在线留言入参
type FeedbackInp struct {
	Name                 string `json:"name"`
	Phone                string `json:"phone"`
	Email                string `json:"email"`
	Company              string `json:"company"`
	MessageTopic         string `json:"message_topic"`
	Content              string `json:"content"`
	ExpectedResponseTime string `json:"expected_response_time"`
	Ip                   string `json:"ip"`
}

// FeedbackModel 在线留言出参
type FeedbackModel struct{}
