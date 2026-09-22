// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumCategory is the golang structure of table hg_forum_category for DAO operations like Where/Data.
type ForumCategory struct {
	g.Meta      `orm:"table:hg_forum_category, do:true"`
	Id          any         // 板块ID
	Pid         any         // 上级板块ID
	Name        any         // 板块名称
	Description any         // 板块简介
	Image       any         // 板块图标
	Sort        any         // 排序
	Status      any         // 状态:1正常,2禁用
	TopicCount  any         // 主题数量
	UpdatedAt   *gtime.Time // 修改时间
	CreatedAt   *gtime.Time // 创建时间
}
