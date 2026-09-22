// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumComment is the golang structure of table hg_forum_comment for DAO operations like Where/Data.
type ForumComment struct {
	g.Meta      `orm:"table:hg_forum_comment, do:true"`
	Id          any         // 评论ID
	TopicId     any         // 主题ID
	UserId      any         // 评论用户ID
	ParentId    any         // 父评论ID
	ReplyUserId any         // 被回复用户ID
	Content     any         // 评论内容
	Status      any         // 状态:normal,pending,rejected,hidden
	LikeCount   any         // 点赞数量
	UpdatedAt   *gtime.Time // 修改时间
	CreatedAt   *gtime.Time // 创建时间
	DeletedAt   *gtime.Time // 删除时间
}
