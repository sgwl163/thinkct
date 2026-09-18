// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// WebArticle is the golang structure for table web_article.
type WebArticle struct {
	Id          uint   `json:"id"          orm:"id"           description:"ID"`
	UserId      int    `json:"userId"      orm:"user_id"      description:"用户ID"`
	AdminId     uint   `json:"adminId"     orm:"admin_id"     description:"管理员ID"`
	ChannelId   uint   `json:"channelId"   orm:"channel_id"   description:"频道"`
	ChannelIds  string `json:"channelIds"  orm:"channel_ids"  description:"副频道"`
	Diyname     string `json:"diyname"     orm:"diyname"      description:"自定义名称"`
	Title       string `json:"title"       orm:"title"        description:"标题"`
	TitleStyle  string `json:"titleStyle"  orm:"title_style"  description:"标题样式"`
	Flag        string `json:"flag"        orm:"flag"         description:"标志:top=置顶,hot=热门,recommend=推荐,new=最新"`
	Image       string `json:"image"       orm:"image"        description:"图片"`
	Images      string `json:"images"      orm:"images"       description:"封面图片"`
	Content     string `json:"content"     orm:"content"      description:"内容"`
	Seotitle    string `json:"seotitle"    orm:"seotitle"     description:"SEO标题"`
	Keywords    string `json:"keywords"    orm:"keywords"     description:"关键词"`
	Description string `json:"description" orm:"description"  description:"描述"`
	Tags        string `json:"tags"        orm:"tags"         description:"标签"`
	Url         string `json:"url"         orm:"url"          description:"外部链接"`
	Target      string `json:"target"      orm:"target"       description:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Views       int    `json:"views"       orm:"views"        description:"浏览量"`
	Comments    int    `json:"comments"    orm:"comments"     description:"评论量"`
	Likes       int    `json:"likes"       orm:"likes"        description:"点赞数"`
	Dislikes    int    `json:"dislikes"    orm:"dislikes"     description:"点踩数"`
	Weigh       int    `json:"weigh"       orm:"weigh"        description:"权重"`
	Status      string `json:"status"      orm:"status"       description:"状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线"`
	PublishTime uint64 `json:"publishTime" orm:"publish_time" description:"发布时间"`
	UpdateTime  uint64 `json:"updateTime"  orm:"update_time"  description:"修改时间"`
	CreateTime  uint64 `json:"createTime"  orm:"create_time"  description:"创建时间"`
}
