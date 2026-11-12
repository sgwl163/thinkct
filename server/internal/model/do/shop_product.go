// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ShopProduct is the golang structure of table hg_shop_product for DAO operations like Where/Data.
type ShopProduct struct {
	g.Meta      `orm:"table:hg_shop_product, do:true"`
	Id          any // ID
	UserId      any // 用户ID
	AdminId     any // 管理员ID
	ChannelId   any // 频道
	ChannelIds  any // 副频道
	Diyname     any // 自定义名称
	Title       any // 标题
	TitleStyle  any // 标题样式
	Flag        any // 标志:top=置顶,hot=热门,recommend=推荐,new=最新
	Image       any // 图片
	Images      any // 封面图片
	Content     any // 内容
	Seotitle    any // SEO标题
	Keywords    any // 关键词
	Description any // 描述
	Tags        any // 标签
	Url         any // 外部链接
	Target      any // 跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口
	Views       any // 浏览量
	Comments    any // 评论量
	Likes       any // 点赞数
	Dislikes    any // 点踩数
	Weigh       any // 权重
	Price       any // 价格
	Status      any // 状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线
	PublishTime any // 发布时间
	UpdateTime  any // 修改时间
	CreateTime  any // 创建时间
	Name        any // 型号名称
}
