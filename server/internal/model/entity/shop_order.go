// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ShopOrder is the golang structure for table shop_order.
type ShopOrder struct {
	Id             uint64  `json:"id"             orm:"id"              description:"订单ID"`
	OrderSn        string  `json:"orderSn"        orm:"order_sn"        description:"订单号"`
	UserId         uint64  `json:"userId"         orm:"user_id"         description:"下单用户ID"`
	Status         string  `json:"status"         orm:"status"          description:"订单状态:pending_payment=待支付,paid=已支付,completed=已完成,cancelled=已取消,closed=已关闭"`
	TotalAmount    float64 `json:"totalAmount"    orm:"total_amount"    description:"商品总金额"`
	DiscountAmount float64 `json:"discountAmount" orm:"discount_amount" description:"优惠金额"`
	FreightAmount  float64 `json:"freightAmount"  orm:"freight_amount"  description:"运费金额"`
	PayableAmount  float64 `json:"payableAmount"  orm:"payable_amount"  description:"应付金额"`
	PaidAmount     float64 `json:"paidAmount"     orm:"paid_amount"     description:"实付金额"`
	Remark         string  `json:"remark"         orm:"remark"          description:"用户备注"`
	PaymentType    string  `json:"paymentType"    orm:"payment_type"    description:"支付方式"`
	PaymentNo      string  `json:"paymentNo"      orm:"payment_no"      description:"第三方支付单号"`
	PaymentTime    uint64  `json:"paymentTime"    orm:"payment_time"    description:"支付时间"`
	CompletedTime  uint64  `json:"completedTime"  orm:"completed_time"  description:"完成时间"`
	CancelledTime  uint64  `json:"cancelledTime"  orm:"cancelled_time"  description:"取消时间"`
	UpdateTime     uint64  `json:"updateTime"     orm:"update_time"     description:"修改时间"`
	CreateTime     uint64  `json:"createTime"     orm:"create_time"     description:"创建时间"`
}
