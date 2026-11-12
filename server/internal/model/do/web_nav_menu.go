// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WebNavMenu is the golang structure of table hg_web_nav_menu for DAO operations like Where/Data.
type WebNavMenu struct {
	g.Meta    `orm:"table:hg_web_nav_menu, do:true"`
	Id        any         // 导航ID
	Pid       any         // 上级ID
	Level     any         // 关系树级别
	Tree      any         // 关系树
	Name      any         // 导航名称
	Url       any         // 链接地址
	Image     any         // 图片
	ImageAlt  any         // 图片替代文本
	Sort      any         // 排序
	Status    any         // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
