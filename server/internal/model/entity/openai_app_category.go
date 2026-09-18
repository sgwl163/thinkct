// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiAppCategory is the golang structure for table openai_app_category.
type OpenaiAppCategory struct {
	Id        int64       `json:"id"        orm:"id"         description:"分类ID"`
	Name      string      `json:"name"      orm:"name"       description:"分类名称"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`
	Status    int         `json:"status"    orm:"status"     description:"状态(1:启用, 2:禁用)"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
