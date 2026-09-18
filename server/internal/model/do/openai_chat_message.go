// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiChatMessage is the golang structure of table hg_openai_chat_message for DAO operations like Where/Data.
type OpenaiChatMessage struct {
	g.Meta      `orm:"table:hg_openai_chat_message, do:true"`
	Id          any         // ID
	UserId      any         // 用户ID
	SessionId   any         // 会话ID
	Role        any         // 角色(user/assistant/system)
	Content     any         // 消息内容
	Model       any         // 使用的模型
	Feedback    any         // 反馈: 0=无, 1=喜欢, 2=不喜欢
	ReqTokens   any         // 请求token数
	ResTokens   any         // 响应token数
	TotalTokens any         // 总token数
	Status      any         // 状态
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
	DeletedAt   *gtime.Time // 删除时间
}
