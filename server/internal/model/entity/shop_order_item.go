// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ShopOrderItem is the golang structure for table shop_order_item.
type ShopOrderItem struct {
	Id          uint64  `json:"id"          orm:"id"           description:"订单明细ID"`
	OrderId     uint64  `json:"orderId"     orm:"order_id"     description:"订单ID"`
	ProductId   uint64  `json:"productId"   orm:"product_id"   description:"商品ID"`
	Title       string  `json:"title"       orm:"title"        description:"商品名称快照"`
	Image       string  `json:"image"       orm:"image"        description:"商品图片快照"`
	UnitPrice   float64 `json:"unitPrice"   orm:"unit_price"   description:"成交单价"`
	Quantity    uint    `json:"quantity"    orm:"quantity"     description:"购买数量"`
	TotalAmount float64 `json:"totalAmount" orm:"total_amount" description:"商品小计"`
	CreateTime  uint64  `json:"createTime"  orm:"create_time"  description:"创建时间"`
	UpdateTime  uint64  `json:"updateTime"  orm:"update_time"  description:"修改时间"`
}
