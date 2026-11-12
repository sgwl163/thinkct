// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRecharge is the golang structure of table hg_user_recharge for DAO operations like Where/Data.
type UserRecharge struct {
	g.Meta    `orm:"table:hg_user_recharge, do:true"`
	Id        any         // 充值订单ID
	UserId    any         // 前台用户ID
	OrderSn   any         // 充值订单号
	Money     any         // 充值金额，整数元，与会员余额同单位
	PayType   any         // 支付渠道
	Status    any         // pending=待支付,completed=已到账
	PayAt     *gtime.Time // 支付时间
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
}
