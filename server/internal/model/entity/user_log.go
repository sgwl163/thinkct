// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLog is the golang structure for table user_log.
type UserLog struct {
	Id         int64       `json:"id"         orm:"id"          description:"日志ID"`
	UserId     int64       `json:"userId"     orm:"user_id"     description:"用户ID"`
	Username   string      `json:"username"   orm:"username"    description:"用户名"`
	LogType    string      `json:"logType"    orm:"log_type"    description:"记录类型:login=登录,operation=操作"`
	Action     string      `json:"action"     orm:"action"      description:"操作标识"`
	ActionName string      `json:"actionName" orm:"action_name" description:"操作名称"`
	EventAt    *gtime.Time `json:"eventAt"    orm:"event_at"    description:"事件时间"`
	Ip         string      `json:"ip"         orm:"ip"          description:"IP地址"`
	UserAgent  string      `json:"userAgent"  orm:"user_agent"  description:"UA信息"`
	Message    string      `json:"message"    orm:"message"     description:"结果消息"`
	Extra      string      `json:"extra"      orm:"extra"       description:"扩展数据"`
	Status     int         `json:"status"     orm:"status"      description:"状态"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"`
}
