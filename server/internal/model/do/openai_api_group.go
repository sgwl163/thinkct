// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiApiGroup is the golang structure of table hg_openai_api_group for DAO operations like Where/Data.
type OpenaiApiGroup struct {
	g.Meta    `orm:"table:hg_openai_api_group, do:true"`
	Id        any         // ID
	Name      any         // 组名称(如OpenAI、文心一言)
	BaseUrl   any         // 接口地址
	ApiKey    any         // API Key
	Status    any         // 状态
	Sort      any         // 排序
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
	DeletedAt *gtime.Time // 删除时间
}
