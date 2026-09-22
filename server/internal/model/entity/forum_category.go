// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumCategory is the golang structure for table forum_category.
type ForumCategory struct {
	Id          uint64      `json:"id"          orm:"id"          description:"板块ID"`
	Pid         uint64      `json:"pid"         orm:"pid"         description:"上级板块ID"`
	Name        string      `json:"name"        orm:"name"        description:"板块名称"`
	Description string      `json:"description" orm:"description" description:"板块简介"`
	Image       string      `json:"image"       orm:"image"       description:"板块图标"`
	Sort        int         `json:"sort"        orm:"sort"        description:"排序"`
	Status      int         `json:"status"      orm:"status"      description:"状态:1正常,2禁用"`
	TopicCount  uint        `json:"topicCount"  orm:"topic_count" description:"主题数量"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"修改时间"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`
}
