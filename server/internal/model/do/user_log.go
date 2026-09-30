// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLog is the golang structure of table hg_user_log for DAO operations like Where/Data.
type UserLog struct {
	g.Meta     `orm:"table:hg_user_log, do:true"`
	Id         any         // 日志ID
	UserId     any         // 用户ID
	Username   any         // 用户名
	LogType    any         // 记录类型:login=登录,operation=操作
	Action     any         // 操作标识
	ActionName any         // 操作名称
	EventAt    *gtime.Time // 事件时间
	Ip         any         // IP地址
	UserAgent  any         // UA信息
	Message    any         // 结果消息
	Extra      any         // 扩展数据
	Status     any         // 状态
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
