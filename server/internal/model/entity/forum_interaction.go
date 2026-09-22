// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumInteraction is the golang structure for table forum_interaction.
type ForumInteraction struct {
	Id         uint64      `json:"id"         orm:"id"          description:"互动ID"`
	TargetType string      `json:"targetType" orm:"target_type" description:"互动对象类型:topic,comment"`
	TargetId   uint64      `json:"targetId"   orm:"target_id"   description:"互动对象ID"`
	UserId     uint64      `json:"userId"     orm:"user_id"     description:"用户ID"`
	ActionType string      `json:"actionType" orm:"action_type" description:"互动类型:like,favorite"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
}
