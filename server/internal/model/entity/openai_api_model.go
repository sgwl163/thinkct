// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiApiModel is the golang structure for table openai_api_model.
type OpenaiApiModel struct {
	Id        int64       `json:"id"        orm:"id"         description:"ID"`
	GroupId   int64       `json:"groupId"   orm:"group_id"   description:"分组ID"`
	Name      string      `json:"name"      orm:"name"       description:"模型别名(如GPT-4)"`
	Model     string      `json:"model"     orm:"model"      description:"模型标识(如gpt-4-turbo)"`
	IsVision  int         `json:"isVision"  orm:"is_vision"  description:"是否支持识图"`
	IsFile    int         `json:"isFile"    orm:"is_file"    description:"是否支持上传文件"`
	IsSearch  int         `json:"isSearch"  orm:"is_search"  description:"是否支持联网搜索"`
	Status    int         `json:"status"    orm:"status"     description:"状态"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"修改时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
