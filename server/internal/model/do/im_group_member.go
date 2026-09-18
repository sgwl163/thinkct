// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ImGroupMember is the golang structure of table hg_im_group_member for DAO operations like Where/Data.
type ImGroupMember struct {
	g.Meta     `orm:"table:hg_im_group_member, do:true"`
	Id         any         // ID
	GroupId    any         // 群组ID
	UserId     any         // 用户ID
	Role       any         // 角色: 1=群主, 2=管理员, 3=普通成员
	RemarkName any         // 群内昵称
	IsMute     any         // 是否禁言: 0=否, 1=是
	CreatedAt  *gtime.Time // 加入时间
}
