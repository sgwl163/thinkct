// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserGroup is the golang structure of table hg_user_group for DAO operations like Where/Data.
type UserGroup struct {
	g.Meta    `orm:"table:hg_user_group, do:true"`
	Id        any         // ID
	Name      any         // 组名
	Rules     any         // 权限节点
	Status    any         // 状态:0=禁用,1=启用
	UpdatedAt *gtime.Time // 更新时间
	CreatedAt *gtime.Time // 创建时间
}
