// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiChatMessage is the golang structure for table openai_chat_message.
type OpenaiChatMessage struct {
	Id          int64       `json:"id"          orm:"id"           description:"ID"`
	UserId      int64       `json:"userId"      orm:"user_id"      description:"用户ID"`
	SessionId   int64       `json:"sessionId"   orm:"session_id"   description:"会话ID"`
	Role        string      `json:"role"        orm:"role"         description:"角色(user/assistant/system)"`
	Content     string      `json:"content"     orm:"content"      description:"消息内容"`
	Model       string      `json:"model"       orm:"model"        description:"使用的模型"`
	Feedback    int         `json:"feedback"    orm:"feedback"     description:"反馈: 0=无, 1=喜欢, 2=不喜欢"`
	ReqTokens   int         `json:"reqTokens"   orm:"req_tokens"   description:"请求token数"`
	ResTokens   int         `json:"resTokens"   orm:"res_tokens"   description:"响应token数"`
	TotalTokens int         `json:"totalTokens" orm:"total_tokens" description:"总token数"`
	Status      int         `json:"status"      orm:"status"       description:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"修改时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`
}
