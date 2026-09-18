// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiChatFolder is the golang structure for table openai_chat_folder.
type OpenaiChatFolder struct {
	Id        int64       `json:"id"        orm:"id"         description:"ID"`
	UserId    int64       `json:"userId"    orm:"user_id"    description:"用户ID"`
	Name      string      `json:"name"      orm:"name"       description:"文件夹名称"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`
	Status    int         `json:"status"    orm:"status"     description:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"修改时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
