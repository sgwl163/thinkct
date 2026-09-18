// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ImMessage is the golang structure for table im_message.
type ImMessage struct {
	Id        int64       `json:"id"        orm:"id"         description:"消息ID"`
	ChatType  int         `json:"chatType"  orm:"chat_type"  description:"聊天类型: 1=单聊, 2=群聊, 3=系统消息"`
	FromId    int64       `json:"fromId"    orm:"from_id"    description:"发送者ID"`
	ToId      int64       `json:"toId"      orm:"to_id"      description:"接收者ID(目标用户ID或群组ID)"`
	MsgType   int         `json:"msgType"   orm:"msg_type"   description:"消息类型: 1=文本, 2=图片, 3=文件, 4=语音, 5=视频, 6=表情包"`
	Content   string      `json:"content"   orm:"content"    description:"消息内容(文本内容或文件URL/JSON)"`
	Status    int         `json:"status"    orm:"status"     description:"状态: 1=正常, 2=已撤回, 3=已删除"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"发送时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间(如撤回时间)"`
}
