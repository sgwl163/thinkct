// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumTopic is the golang structure of table hg_forum_topic for DAO operations like Where/Data.
type ForumTopic struct {
	g.Meta          `orm:"table:hg_forum_topic, do:true"`
	Id              any         // 主题ID
	CategoryId      any         // 板块ID
	UserId          any         // 发布用户ID
	Title           any         // 主题标题
	Keywords        any         // 主题关键词，英文逗号分隔
	Description     any         // 主题摘要
	Content         any         // 主题内容
	Images          any         // 主题图片，英文逗号分隔
	Status          any         // 状态:normal,pending,rejected,hidden,closed
	IsTop           any         // 是否置顶:0否,1是
	IsEssence       any         // 是否精华:0否,1是
	CommentCount    any         // 评论数量
	LikeCount       any         // 点赞数量
	FavoriteCount   any         // 收藏数量
	ViewCount       any         // 浏览数量
	LastReplyUserId any         // 最后回复用户ID
	LastReplyAt     *gtime.Time // 最后回复时间
	UpdatedAt       *gtime.Time // 修改时间
	CreatedAt       *gtime.Time // 创建时间
	DeletedAt       *gtime.Time // 删除时间
}
