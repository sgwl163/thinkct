// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ImContact is the golang structure of table hg_im_contact for DAO operations like Where/Data.
type ImContact struct {
	g.Meta    `orm:"table:hg_im_contact, do:true"`
	Id        any         // ID
	UserId    any         // 用户ID
	FriendId  any         // 好友ID
	Remark    any         // 好友备注名
	Status    any         // 状态: 1=正常, 2=拉黑
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
