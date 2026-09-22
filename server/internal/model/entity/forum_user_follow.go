// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumUserFollow is the golang structure for table forum_user_follow.
type ForumUserFollow struct {
	Id           uint64      `json:"id"           orm:"id"             description:"关注ID"`
	UserId       uint64      `json:"userId"       orm:"user_id"        description:"关注用户ID"`
	FollowUserId uint64      `json:"followUserId" orm:"follow_user_id" description:"被关注用户ID"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"     description:"关注时间"`
}
