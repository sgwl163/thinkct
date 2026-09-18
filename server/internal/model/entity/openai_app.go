// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiApp is the golang structure for table openai_app.
type OpenaiApp struct {
	Id          int64       `json:"id"          orm:"id"          description:"应用ID"`
	CategoryId  int64       `json:"categoryId"  orm:"category_id" description:"分类ID"`
	Title       string      `json:"title"       orm:"title"       description:"应用名称"`
	Icon        string      `json:"icon"        orm:"icon"        description:"应用图标"`
	Description string      `json:"description" orm:"description" description:"应用描述"`
	Prompt      string      `json:"prompt"      orm:"prompt"      description:"系统提示词"`
	UseCount    int         `json:"useCount"    orm:"use_count"   description:"使用次数"`
	Sort        int         `json:"sort"        orm:"sort"        description:"排序"`
	Status      int         `json:"status"      orm:"status"      description:"状态(1:启用, 2:禁用)"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`
}
