// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OauthLog is the golang structure of table hg_oauth_log for DAO operations like Where/Data.
type OauthLog struct {
	g.Meta    `orm:"table:hg_oauth_log, do:true"`
	Id        any         // ID
	UserId    any         // 用户ID
	ProductId any         // 产品ID
	Host      any         // 域名
	Content   any         // 内容
	Ip        any         // IP
	Version   any         // 版本号
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
