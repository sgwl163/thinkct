// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ImContact is the golang structure for table im_contact.
type ImContact struct {
	Id        int64       `json:"id"        orm:"id"         description:"ID"`
	UserId    int64       `json:"userId"    orm:"user_id"    description:"用户ID"`
	FriendId  int64       `json:"friendId"  orm:"friend_id"  description:"好友ID"`
	Remark    string      `json:"remark"    orm:"remark"     description:"好友备注名"`
	Status    int         `json:"status"    orm:"status"     description:"状态: 1=正常, 2=拉黑"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
