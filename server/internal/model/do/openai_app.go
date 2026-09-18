// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiApp is the golang structure of table hg_openai_app for DAO operations like Where/Data.
type OpenaiApp struct {
	g.Meta      `orm:"table:hg_openai_app, do:true"`
	Id          any         // 应用ID
	CategoryId  any         // 分类ID
	Title       any         // 应用名称
	Icon        any         // 应用图标
	Description any         // 应用描述
	Prompt      any         // 系统提示词
	UseCount    any         // 使用次数
	Sort        any         // 排序
	Status      any         // 状态(1:启用, 2:禁用)
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
