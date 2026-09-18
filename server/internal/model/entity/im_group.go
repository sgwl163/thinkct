// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ImGroup is the golang structure for table im_group.
type ImGroup struct {
	Id        int64       `json:"id"        orm:"id"         description:"群组ID"`
	Name      string      `json:"name"      orm:"name"       description:"群名称"`
	Avatar    string      `json:"avatar"    orm:"avatar"     description:"群头像"`
	OwnerId   int64       `json:"ownerId"   orm:"owner_id"   description:"群主ID"`
	Notice    string      `json:"notice"    orm:"notice"     description:"群公告"`
	Status    int         `json:"status"    orm:"status"     description:"状态: 1=正常, 2=已解散"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
