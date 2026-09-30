// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppAuth is the golang structure of table hg_app_auth for DAO operations like Where/Data.
type AppAuth struct {
	g.Meta     `orm:"table:hg_app_auth, do:true"`
	Id         any         // ID
	UserId     any         // 用户ID
	ProductId  any         // 产品ID
	Name       any         // 名称
	Host       any         // 域名
	Ip         any         // IP
	Qq         any         // QQ
	Phone      any         // 电话
	Syskey     any         // 授权码
	VersionId  any         // 版本
	AuthMode   any         // 验证方式:0=单域名验证,1=IP双重验证
	AuthDomain any         // 验证域名:0=当前单域名方式,1=顶级泛域名方式
	Views      any         // 浏览量
	Status     any         // 状态:1=启用,0=禁用
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 修改时间
	ExpiredAt  *gtime.Time // 到期时间
}
