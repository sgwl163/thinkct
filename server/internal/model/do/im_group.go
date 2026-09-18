// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ImGroup is the golang structure of table hg_im_group for DAO operations like Where/Data.
type ImGroup struct {
	g.Meta    `orm:"table:hg_im_group, do:true"`
	Id        any         // 群组ID
	Name      any         // 群名称
	Avatar    any         // 群头像
	OwnerId   any         // 群主ID
	Notice    any         // 群公告
	Status    any         // 状态: 1=正常, 2=已解散
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
