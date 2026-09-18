// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ImConversation is the golang structure for table im_conversation.
type ImConversation struct {
	Id             int64       `json:"id"             orm:"id"               description:"会话ID"`
	UserId         int64       `json:"userId"         orm:"user_id"          description:"归属用户ID"`
	ChatType       int         `json:"chatType"       orm:"chat_type"        description:"会话类型: 1=单聊, 2=群聊, 3=系统通知"`
	TargetId       int64       `json:"targetId"       orm:"target_id"        description:"目标ID(好友ID或群组ID)"`
	LastMsgId      int64       `json:"lastMsgId"      orm:"last_msg_id"      description:"最后一条消息ID"`
	LastMsgContent string      `json:"lastMsgContent" orm:"last_msg_content" description:"最后消息内容摘要(用于列表展示)"`
	LastMsgTime    *gtime.Time `json:"lastMsgTime"    orm:"last_msg_time"    description:"最后消息时间"`
	UnreadCount    int         `json:"unreadCount"    orm:"unread_count"     description:"未读消息数"`
	IsTop          int         `json:"isTop"          orm:"is_top"           description:"是否置顶: 0=否, 1=是"`
	IsDisturb      int         `json:"isDisturb"      orm:"is_disturb"       description:"消息免打扰: 0=否, 1=是"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:"更新时间"`
}
