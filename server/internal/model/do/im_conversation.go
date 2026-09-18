// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ImConversation is the golang structure of table hg_im_conversation for DAO operations like Where/Data.
type ImConversation struct {
	g.Meta         `orm:"table:hg_im_conversation, do:true"`
	Id             any         // 会话ID
	UserId         any         // 归属用户ID
	ChatType       any         // 会话类型: 1=单聊, 2=群聊, 3=系统通知
	TargetId       any         // 目标ID(好友ID或群组ID)
	LastMsgId      any         // 最后一条消息ID
	LastMsgContent any         // 最后消息内容摘要(用于列表展示)
	LastMsgTime    *gtime.Time // 最后消息时间
	UnreadCount    any         // 未读消息数
	IsTop          any         // 是否置顶: 0=否, 1=是
	IsDisturb      any         // 消息免打扰: 0=否, 1=是
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 更新时间
}
