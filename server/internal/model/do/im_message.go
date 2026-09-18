// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ImMessage is the golang structure of table hg_im_message for DAO operations like Where/Data.
type ImMessage struct {
	g.Meta    `orm:"table:hg_im_message, do:true"`
	Id        any         // 消息ID
	ChatType  any         // 聊天类型: 1=单聊, 2=群聊, 3=系统消息
	FromId    any         // 发送者ID
	ToId      any         // 接收者ID(目标用户ID或群组ID)
	MsgType   any         // 消息类型: 1=文本, 2=图片, 3=文件, 4=语音, 5=视频, 6=表情包
	Content   any         // 消息内容(文本内容或文件URL/JSON)
	Status    any         // 状态: 1=正常, 2=已撤回, 3=已删除
	CreatedAt *gtime.Time // 发送时间
	UpdatedAt *gtime.Time // 更新时间(如撤回时间)
}
