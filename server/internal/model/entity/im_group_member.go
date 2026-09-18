// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ImGroupMember is the golang structure for table im_group_member.
type ImGroupMember struct {
	Id         int64       `json:"id"         orm:"id"          description:"ID"`
	GroupId    int64       `json:"groupId"    orm:"group_id"    description:"群组ID"`
	UserId     int64       `json:"userId"     orm:"user_id"     description:"用户ID"`
	Role       int         `json:"role"       orm:"role"        description:"角色: 1=群主, 2=管理员, 3=普通成员"`
	RemarkName string      `json:"remarkName" orm:"remark_name" description:"群内昵称"`
	IsMute     int         `json:"isMute"     orm:"is_mute"     description:"是否禁言: 0=否, 1=是"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"加入时间"`
}
