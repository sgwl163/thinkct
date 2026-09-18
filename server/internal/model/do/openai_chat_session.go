// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiChatSession is the golang structure of table hg_openai_chat_session for DAO operations like Where/Data.
type OpenaiChatSession struct {
	g.Meta      `orm:"table:hg_openai_chat_session, do:true"`
	Id          any         // ID
	UserId      any         // 用户ID
	FolderId    any         // 文件夹ID
	Title       any         // 会话标题
	IsTop       any         // 是否置顶
	Prompt      any         // 系统提示词
	Model       any         // 使用的模型
	Temperature any         // 温度值
	Status      any         // 状态
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
	DeletedAt   *gtime.Time // 删除时间
}
