// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WebFeedback is the golang structure of table hg_web_feedback for DAO operations like Where/Data.
type WebFeedback struct {
	g.Meta     `orm:"table:hg_web_feedback, do:true"`
	Id         any // ID
	Name       any // 名称
	Email      any // 邮箱
	Phone      any // 手机号
	Company    any // 公司名称
	Content    any // 富文本
	Answer     any // 备注
	Ip         any // IP
	Status     any // 状态:0=禁用,1=启用
	CreateTime any // 创建时间
	UpdateTime any // 修改时间
}
