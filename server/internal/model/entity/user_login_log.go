// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLoginLog is the golang structure for table user_login_log.
type UserLoginLog struct {
	Id        int64       `json:"id"        orm:"id"         description:"日志ID"`
	UserId    int64       `json:"userId"    orm:"user_id"    description:"用户ID"`
	Username  string      `json:"username"  orm:"username"   description:"用户名"`
	LoginAt   *gtime.Time `json:"loginAt"   orm:"login_at"   description:"登录时间"`
	LoginIp   string      `json:"loginIp"   orm:"login_ip"   description:"登录IP"`
	UserAgent string      `json:"userAgent" orm:"user_agent" description:"UA信息"`
	ErrMsg    string      `json:"errMsg"    orm:"err_msg"    description:"错误提示"`
	Status    int         `json:"status"    orm:"status"     description:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
