// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumUserFollow is the golang structure of table hg_forum_user_follow for DAO operations like Where/Data.
type ForumUserFollow struct {
	g.Meta       `orm:"table:hg_forum_user_follow, do:true"`
	Id           any         // 关注ID
	UserId       any         // 关注用户ID
	FollowUserId any         // 被关注用户ID
	CreatedAt    *gtime.Time // 关注时间
}
