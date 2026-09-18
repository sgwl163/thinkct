package v1

import "github.com/gogf/gf/v2/frame/g"

type IndexReq struct {
	g.Meta `path:"/site/index" method:"get" tags:"前台首页" summary:"获取首页数据"`
}

type IndexRes struct {
	SiteConfig g.Map `json:"siteConfig" dc:"网站基础配置"`
}

type HomeReq struct {
	g.Meta             `path:"/site/home" method:"get" tags:"前台首页" summary:"获取首页聚合数据"`
	ProductsChannelIds string `json:"productsChannelIds" dc:"产品分类栏目ID组(英文逗号分隔)"`
	ProductsLimit      int    `json:"productsLimit" dc:"每个产品分类显示的产品数量"`
	CasesChannelIds    string `json:"casesChannelIds" dc:"案例栏目ID组(英文逗号分隔)"`
	CasesLimit         int    `json:"casesLimit" dc:"每个案例分类显示的案例数量"`
	NewsChannelIds     string `json:"newsChannelIds" dc:"新闻栏目ID组(英文逗号分隔)"`
	NewsLimit          int    `json:"newsLimit" dc:"每个新闻分类显示的新闻数量"`
}

type HomeRes struct {
	ProductsChannels interface{} `json:"productsChannels" dc:"产品中心分类"`
	CasesChannels    interface{} `json:"casesChannels" dc:"经典案例分类"`
	NewsChannels     interface{} `json:"newsChannels" dc:"新闻中心分类"`
	Links            interface{} `json:"links" dc:"友情链接列表"`
}

type ChannelReq struct {
	g.Meta  `path:"/site/channel" method:"get" tags:"前台栏目" summary:"获取栏目列表与单栏目详情内容"`
	Channel string `json:"channel" dc:"栏目ID或自定义名称(diyname)(可选，传时返回当前栏目下的数据)"`
	Page    int    `json:"page" dc:"当前页码(传栏目时生效)"`
	Limit   int    `json:"limit" dc:"每页数量(传栏目时生效)"`
}

type ChannelRes struct {
	Channel     interface{} `json:"channel" dc:"当前栏目信息"`
	Breadcrumb  interface{} `json:"breadcrumb" dc:"面包屑导航"`
	List        interface{} `json:"list" dc:"关联的文章或产品列表"`
	Total       int         `json:"total" dc:"关联数据的总记录数"`
	Page        int         `json:"page" dc:"当前页码"`
	RelatedList interface{} `json:"relatedList" dc:"相关分类(同级栏目)"`
	HotList     interface{} `json:"hotList" dc:"热门内容"`
}

type ContentReq struct {
	g.Meta  `path:"/site/content" method:"get" tags:"前台内容" summary:"获取文章或产品的详情内容"`
	Id      uint   `json:"id" v:"required#内容ID不能为空" dc:"内容ID"`
	Channel string `json:"channel" v:"required#栏目不能为空" dc:"所属栏目ID或自定义名称(diyname)"`
}

type ContentRes struct {
	Content     interface{} `json:"content" dc:"详细内容"`
	Channel     interface{} `json:"channel" dc:"所属栏目信息"`
	Breadcrumb  interface{} `json:"breadcrumb" dc:"面包屑导航"`
	RelatedList interface{} `json:"relatedList" dc:"相关分类(同级栏目)"`
	HotList     interface{} `json:"hotList" dc:"热门内容"`
	Prev        interface{} `json:"prev" dc:"上一条内容"`
	Next        interface{} `json:"next" dc:"下一条内容"`
}

type SearchReq struct {
	g.Meta   `path:"/site/search" method:"get" tags:"前台搜索" summary:"全站内容搜索"`
	Keywords string `json:"keywords" v:"required#请输入搜索关键词" dc:"搜索关键词"`
	Type     string `json:"type" dc:"搜索类型(products:产品, article:文章)，默认为products"`
	Page     int    `json:"page" dc:"当前页码，默认为1"`
	Limit    int    `json:"limit" dc:"每页数量，默认为16"`
}

type SearchRes struct {
	List  interface{} `json:"list" dc:"数据列表"`
	Total int         `json:"total" dc:"总记录数"`
	Page  int         `json:"page" dc:"当前页码"`
	Limit int         `json:"limit" dc:"每页数量"`
}

type FeedbackReq struct {
	g.Meta               `path:"/site/feedback" method:"post" tags:"在线留言" summary:"提交在线留言"`
	Name                 string `json:"name" v:"required#请输入您的姓名" dc:"姓名"`
	Phone                string `json:"phone" v:"required#请输入您的手机号" dc:"手机号"`
	Email                string `json:"email" dc:"邮箱"`
	Company              string `json:"company" dc:"公司名称"`
	MessageTopic         string `json:"message_topic" dc:"留言主题"`
	Content              string `json:"content" v:"required#请输入留言内容" dc:"留言内容"`
	ExpectedResponseTime string `json:"expected_response_time" dc:"期望回复时间"`
}

type FeedbackRes struct{}
