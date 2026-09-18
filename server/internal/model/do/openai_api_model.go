// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiApiModel is the golang structure of table hg_openai_api_model for DAO operations like Where/Data.
type OpenaiApiModel struct {
	g.Meta    `orm:"table:hg_openai_api_model, do:true"`
	Id        any         // ID
	GroupId   any         // 分组ID
	Name      any         // 模型别名(如GPT-4)
	Model     any         // 模型标识(如gpt-4-turbo)
	IsVision  any         // 是否支持识图
	IsFile    any         // 是否支持上传文件
	IsSearch  any         // 是否支持联网搜索
	Status    any         // 状态
	Sort      any         // 排序
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
	DeletedAt *gtime.Time // 删除时间
}
