// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiApiGroup is the golang structure for table openai_api_group.
type OpenaiApiGroup struct {
	Id        int64       `json:"id"        orm:"id"         description:"ID"`
	Name      string      `json:"name"      orm:"name"       description:"组名称(如OpenAI、文心一言)"`
	BaseUrl   string      `json:"baseUrl"   orm:"base_url"   description:"接口地址"`
	ApiKey    string      `json:"apiKey"    orm:"api_key"    description:"API Key"`
	Status    int         `json:"status"    orm:"status"     description:"状态"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"修改时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
