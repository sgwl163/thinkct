// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumReport is the golang structure for table forum_report.
type ForumReport struct {
	Id          uint64      `json:"id"          orm:"id"          description:"举报ID"`
	TargetType  string      `json:"targetType"  orm:"target_type" description:"举报对象类型:topic,comment"`
	TargetId    uint64      `json:"targetId"    orm:"target_id"   description:"举报对象ID"`
	UserId      uint64      `json:"userId"      orm:"user_id"     description:"举报用户ID"`
	Reason      string      `json:"reason"      orm:"reason"      description:"举报原因"`
	Description string      `json:"description" orm:"description" description:"补充说明"`
	Status      string      `json:"status"      orm:"status"      description:"处理状态:pending,processed,rejected"`
	HandledBy   uint64      `json:"handledBy"   orm:"handled_by"  description:"处理管理员ID"`
	HandledAt   *gtime.Time `json:"handledAt"   orm:"handled_at"  description:"处理时间"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`
}
