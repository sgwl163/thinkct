// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table hg_user for DAO operations like Where/Data.
type User struct {
	g.Meta        `orm:"table:hg_user, do:true"`
	Id            any         // ID
	GroupId       any         // 分组ID
	Username      any         // 用户名
	Nickname      any         // 昵称
	Email         any         // 邮箱
	Mobile        any         // 手机
	Avatar        any         // 头像
	Gender        any         // 性别:0=未知,1=男,2=女
	Birthday      *gtime.Time // 生日
	Money         any         // 余额
	Score         any         // 积分
	LastLoginTime any         // 上次登录时间
	LastLoginIp   any         // 上次登录IP
	LoginFailure  any         // 登录失败次数
	JoinIp        any         // 加入IP
	JoinTime      any         // 加入时间
	Motto         any         // 签名
	Password      any         // 密码
	Salt          any         // 密码盐（废弃待删）
	Status        any         // 状态:enable=启用,disable=禁用
	UpdateTime    any         // 更新时间
	CreateTime    any         // 创建时间
}
