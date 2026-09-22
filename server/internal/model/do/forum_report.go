// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ForumReport is the golang structure of table hg_forum_report for DAO operations like Where/Data.
type ForumReport struct {
	g.Meta      `orm:"table:hg_forum_report, do:true"`
	Id          any         // 举报ID
	TargetType  any         // 举报对象类型:topic,comment
	TargetId    any         // 举报对象ID
	UserId      any         // 举报用户ID
	Reason      any         // 举报原因
	Description any         // 补充说明
	Status      any         // 处理状态:pending,processed,rejected
	HandledBy   any         // 处理管理员ID
	HandledAt   *gtime.Time // 处理时间
	CreatedAt   *gtime.Time // 创建时间
}
