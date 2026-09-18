// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WebLinks is the golang structure of table hg_web_links for DAO operations like Where/Data.
type WebLinks struct {
	g.Meta     `orm:"table:hg_web_links, do:true"`
	Id         any // ID
	UserId     any // 用户ID
	Title      any // 标题
	Link       any // 链接
	Target     any // 跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口
	Logo       any // LOGO
	Weigh      any // 权重
	Status     any // 状态:disable=禁用,enable=启用,pending_trial=待审
	Remark     any // 备注
	UpdateTime any // 修改时间
	CreateTime any // 创建时间
}
