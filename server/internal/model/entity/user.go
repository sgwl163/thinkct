// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id            uint        `json:"id"            orm:"id"              description:"ID"`
	GroupId       uint        `json:"groupId"       orm:"group_id"        description:"分组ID"`
	Username      string      `json:"username"      orm:"username"        description:"用户名"`
	Nickname      string      `json:"nickname"      orm:"nickname"        description:"昵称"`
	Email         string      `json:"email"         orm:"email"           description:"邮箱"`
	Mobile        string      `json:"mobile"        orm:"mobile"          description:"手机"`
	Avatar        string      `json:"avatar"        orm:"avatar"          description:"头像"`
	Gender        uint        `json:"gender"        orm:"gender"          description:"性别:0=未知,1=男,2=女"`
	Birthday      *gtime.Time `json:"birthday"      orm:"birthday"        description:"生日"`
	Money         uint        `json:"money"         orm:"money"           description:"余额"`
	Score         uint        `json:"score"         orm:"score"           description:"积分"`
	LastLoginTime uint64      `json:"lastLoginTime" orm:"last_login_time" description:"上次登录时间"`
	LastLoginIp   string      `json:"lastLoginIp"   orm:"last_login_ip"   description:"上次登录IP"`
	LoginFailure  uint        `json:"loginFailure"  orm:"login_failure"   description:"登录失败次数"`
	JoinIp        string      `json:"joinIp"        orm:"join_ip"         description:"加入IP"`
	JoinTime      uint64      `json:"joinTime"      orm:"join_time"       description:"加入时间"`
	Motto         string      `json:"motto"         orm:"motto"           description:"签名"`
	Password      string      `json:"password"      orm:"password"        description:"密码"`
	Salt          string      `json:"salt"          orm:"salt"            description:"密码盐（废弃待删）"`
	Status        string      `json:"status"        orm:"status"          description:"状态:enable=启用,disable=禁用"`
	UpdateTime    uint64      `json:"updateTime"    orm:"update_time"     description:"更新时间"`
	CreateTime    uint64      `json:"createTime"    orm:"create_time"     description:"创建时间"`
}
