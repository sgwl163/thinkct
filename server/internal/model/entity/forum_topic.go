// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumTopic is the golang structure for table forum_topic.
type ForumTopic struct {
	Id              uint64      `json:"id"              orm:"id"                 description:"主题ID"`
	CategoryId      uint64      `json:"categoryId"      orm:"category_id"        description:"板块ID"`
	UserId          uint64      `json:"userId"          orm:"user_id"            description:"发布用户ID"`
	Title           string      `json:"title"           orm:"title"              description:"主题标题"`
	Keywords        string      `json:"keywords"        orm:"keywords"           description:"主题关键词，英文逗号分隔"`
	Description     string      `json:"description"     orm:"description"        description:"主题摘要"`
	Content         string      `json:"content"         orm:"content"            description:"主题内容"`
	Images          string      `json:"images"          orm:"images"             description:"主题图片，英文逗号分隔"`
	Status          string      `json:"status"          orm:"status"             description:"状态:normal,pending,rejected,hidden,closed"`
	IsTop           int         `json:"isTop"           orm:"is_top"             description:"是否置顶:0否,1是"`
	IsEssence       int         `json:"isEssence"       orm:"is_essence"         description:"是否精华:0否,1是"`
	CommentCount    uint        `json:"commentCount"    orm:"comment_count"      description:"评论数量"`
	LikeCount       uint        `json:"likeCount"       orm:"like_count"         description:"点赞数量"`
	FavoriteCount   uint        `json:"favoriteCount"   orm:"favorite_count"     description:"收藏数量"`
	ViewCount       uint        `json:"viewCount"       orm:"view_count"         description:"浏览数量"`
	LastReplyUserId uint64      `json:"lastReplyUserId" orm:"last_reply_user_id" description:"最后回复用户ID"`
	LastReplyAt     *gtime.Time `json:"lastReplyAt"     orm:"last_reply_at"      description:"最后回复时间"`
	UpdatedAt       *gtime.Time `json:"updatedAt"       orm:"updated_at"         description:"修改时间"`
	CreatedAt       *gtime.Time `json:"createdAt"       orm:"created_at"         description:"创建时间"`
	DeletedAt       *gtime.Time `json:"deletedAt"       orm:"deleted_at"         description:"删除时间"`
}
