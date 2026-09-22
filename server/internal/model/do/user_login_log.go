// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserLoginLog is the golang structure of table hg_user_login_log for DAO operations like Where/Data.
type UserLoginLog struct {
	g.Meta    `orm:"table:hg_user_login_log, do:true"`
	Id        any         // 日志ID
	UserId    any         // 用户ID
	Username  any         // 用户名
	LoginAt   *gtime.Time // 登录时间
	LoginIp   any         // 登录IP
	UserAgent any         // UA信息
	ErrMsg    any         // 错误提示
	Status    any         // 状态
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
