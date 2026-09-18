// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiAppCategory is the golang structure of table hg_openai_app_category for DAO operations like Where/Data.
type OpenaiAppCategory struct {
	g.Meta    `orm:"table:hg_openai_app_category, do:true"`
	Id        any         // 分类ID
	Name      any         // 分类名称
	Sort      any         // 排序
	Status    any         // 状态(1:启用, 2:禁用)
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
