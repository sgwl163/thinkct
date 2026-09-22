// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserCreditsLog is the golang structure of table hg_user_credits_log for DAO operations like Where/Data.
type UserCreditsLog struct {
	g.Meta     `orm:"table:hg_user_credits_log, do:true"`
	Id         any         // 流水ID
	UserId     any         // 用户ID
	CreditType any         // 资产类型:money=余额,score=积分
	BeforeNum  any         // 变动前数量
	Num        any         // 变动数量
	AfterNum   any         // 变动后数量
	Remark     any         // 备注
	MapId      any         // 关联ID
	Status     any         // 状态
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
}
