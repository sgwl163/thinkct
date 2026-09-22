// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumInteraction is the golang structure of table hg_forum_interaction for DAO operations like Where/Data.
type ForumInteraction struct {
	g.Meta     `orm:"table:hg_forum_interaction, do:true"`
	Id         any         // 互动ID
	TargetType any         // 互动对象类型:topic,comment
	TargetId   any         // 互动对象ID
	UserId     any         // 用户ID
	ActionType any         // 互动类型:like,favorite
	CreatedAt  *gtime.Time // 创建时间
}
