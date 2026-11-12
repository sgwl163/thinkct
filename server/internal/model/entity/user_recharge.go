// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRecharge is the golang structure for table user_recharge.
type UserRecharge struct {
	Id        int64       `json:"id"        orm:"id"         description:"充值订单ID"`
	UserId    int64       `json:"userId"    orm:"user_id"    description:"前台用户ID"`
	OrderSn   string      `json:"orderSn"   orm:"order_sn"   description:"充值订单号"`
	Money     uint        `json:"money"     orm:"money"      description:"充值金额，整数元，与会员余额同单位"`
	PayType   string      `json:"payType"   orm:"pay_type"   description:"支付渠道"`
	Status    string      `json:"status"    orm:"status"     description:"pending=待支付,completed=已到账"`
	PayAt     *gtime.Time `json:"payAt"     orm:"pay_at"     description:"支付时间"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
