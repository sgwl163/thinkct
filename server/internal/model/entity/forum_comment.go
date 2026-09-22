// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumComment is the golang structure for table forum_comment.
type ForumComment struct {
	Id          uint64      `json:"id"          orm:"id"            description:"评论ID"`
	TopicId     uint64      `json:"topicId"     orm:"topic_id"      description:"主题ID"`
	UserId      uint64      `json:"userId"      orm:"user_id"       description:"评论用户ID"`
	ParentId    uint64      `json:"parentId"    orm:"parent_id"     description:"父评论ID"`
	ReplyUserId uint64      `json:"replyUserId" orm:"reply_user_id" description:"被回复用户ID"`
	Content     string      `json:"content"     orm:"content"       description:"评论内容"`
	Status      string      `json:"status"      orm:"status"        description:"状态:normal,pending,rejected,hidden"`
	LikeCount   uint        `json:"likeCount"   orm:"like_count"    description:"点赞数量"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"    description:"修改时间"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"    description:"创建时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"    description:"删除时间"`
}
