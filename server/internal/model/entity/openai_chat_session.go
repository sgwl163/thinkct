// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiChatSession is the golang structure for table openai_chat_session.
type OpenaiChatSession struct {
	Id          int64       `json:"id"          orm:"id"          description:"ID"`
	UserId      int64       `json:"userId"      orm:"user_id"     description:"用户ID"`
	FolderId    int64       `json:"folderId"    orm:"folder_id"   description:"文件夹ID"`
	Title       string      `json:"title"       orm:"title"       description:"会话标题"`
	IsTop       int         `json:"isTop"       orm:"is_top"      description:"是否置顶"`
	Prompt      string      `json:"prompt"      orm:"prompt"      description:"系统提示词"`
	Model       string      `json:"model"       orm:"model"       description:"使用的模型"`
	Temperature float64     `json:"temperature" orm:"temperature" description:"温度值"`
	Status      int         `json:"status"      orm:"status"      description:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"修改时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"  description:"删除时间"`
}
