// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ShopOrder is the golang structure of table hg_shop_order for DAO operations like Where/Data.
type ShopOrder struct {
	g.Meta         `orm:"table:hg_shop_order, do:true"`
	Id             any // 订单ID
	OrderSn        any // 订单号
	UserId         any // 下单用户ID
	Status         any // 订单状态:pending_payment=待支付,paid=已支付,completed=已完成,cancelled=已取消,closed=已关闭
	TotalAmount    any // 商品总金额
	DiscountAmount any // 优惠金额
	FreightAmount  any // 运费金额
	PayableAmount  any // 应付金额
	PaidAmount     any // 实付金额
	Remark         any // 用户备注
	PaymentType    any // 支付方式
	PaymentNo      any // 第三方支付单号
	PaymentTime    any // 支付时间
	CompletedTime  any // 完成时间
	CancelledTime  any // 取消时间
	UpdateTime     any // 修改时间
	CreateTime     any // 创建时间
}
