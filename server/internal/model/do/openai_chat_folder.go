// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiChatFolder is the golang structure of table hg_openai_chat_folder for DAO operations like Where/Data.
type OpenaiChatFolder struct {
	g.Meta    `orm:"table:hg_openai_chat_folder, do:true"`
	Id        any         // ID
	UserId    any         // 用户ID
	Name      any         // 文件夹名称
	Sort      any         // 排序
	Status    any         // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
	DeletedAt *gtime.Time // 删除时间
}
