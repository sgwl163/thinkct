package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type LoginReq struct {
	g.Meta   `path:"/user/login" method:"post" tags:"用户" summary:"用户登录"`
	Username string `json:"username" v:"required#请输入用户名" dc:"用户名"`
	Password string `json:"password" v:"required#请输入密码" dc:"密码"`
}

type LoginRes struct {
	Token    string `json:"token" dc:"登录令牌"`
	Expire   int64  `json:"expire" dc:"过期时间"`
	Id       int64  `json:"id" dc:"用户ID"`
	Username string `json:"username" dc:"用户名"`
	Nickname string `json:"nickname" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
	Mobile   string `json:"mobile" dc:"手机号"`
	Email    string `json:"email" dc:"邮箱"`
	Money    uint   `json:"money" dc:"余额"`
	Score    uint   `json:"score" dc:"积分"`
}

type InfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"用户" summary:"获取用户信息"`
}

type InfoRes struct {
	Id       int64  `json:"id" dc:"用户ID"`
	Username string `json:"username" dc:"用户名"`
	Nickname string `json:"nickname" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
	Mobile   string `json:"mobile" dc:"手机号"`
	Email    string `json:"email" dc:"邮箱"`
	Money    uint   `json:"money" dc:"余额"`
	Score    uint   `json:"score" dc:"积分"`
}

type RegisterReq struct {
	g.Meta   `path:"/user/register" method:"post" tags:"用户" summary:"用户注册"`
	Username string `json:"username" v:"required#请输入用户名" dc:"用户名"`
	Password string `json:"password" v:"required|length:6,20#请输入密码|密码长度需要在6~20位之间" dc:"密码"`
	Nickname string `json:"nickname" dc:"昵称"`
	Mobile   string `json:"mobile" dc:"手机号"`
	Email    string `json:"email" dc:"邮箱"`
}

type RegisterRes struct {
	Token    string `json:"token" dc:"登录令牌"`
	Expire   int64  `json:"expire" dc:"过期时间"`
	Id       int64  `json:"id" dc:"用户ID"`
	Username string `json:"username" dc:"用户名"`
	Nickname string `json:"nickname" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
	Mobile   string `json:"mobile" dc:"手机号"`
	Email    string `json:"email" dc:"邮箱"`
	Money    uint   `json:"money" dc:"余额"`
	Score    uint   `json:"score" dc:"积分"`
}

type LogoutReq struct {
	g.Meta `path:"/user/logout" method:"post" tags:"用户" summary:"退出登录"`
}

type LogoutRes struct{}

type UpdateInfoReq struct {
	g.Meta   `path:"/user/update" method:"post" tags:"用户" summary:"修改用户信息"`
	Nickname string      `json:"nickname" v:"required#昵称不能为空" dc:"昵称"`
	Avatar   string      `json:"avatar" v:"required#头像不能为空" dc:"头像"`
	Mobile   string      `json:"mobile" v:"phone#手机号格式不正确" dc:"手机号"`
	Email    string      `json:"email" v:"email#邮箱格式不正确" dc:"邮箱"`
	Gender   int         `json:"gender" dc:"性别:0=未知,1=男,2=女"`
	Birthday *gtime.Time `json:"birthday" dc:"生日"`
	Motto    string      `json:"motto" dc:"签名"`
}

type UpdateInfoRes struct {
	Id       int64       `json:"id" dc:"用户ID"`
	Username string      `json:"username" dc:"用户名"`
	Nickname string      `json:"nickname" dc:"昵称"`
	Avatar   string      `json:"avatar" dc:"头像"`
	Mobile   string      `json:"mobile" dc:"手机号"`
	Email    string      `json:"email" dc:"邮箱"`
	Gender   int         `json:"gender" dc:"性别:0=未知,1=男,2=女"`
	Birthday *gtime.Time `json:"birthday" dc:"生日"`
	Motto    string      `json:"motto" dc:"签名"`
	Money    uint        `json:"money" dc:"余额"`
	Score    uint        `json:"score" dc:"积分"`
}
