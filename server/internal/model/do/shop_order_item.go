// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ShopOrderItem is the golang structure of table hg_shop_order_item for DAO operations like Where/Data.
type ShopOrderItem struct {
	g.Meta      `orm:"table:hg_shop_order_item, do:true"`
	Id          any // 订单明细ID
	OrderId     any // 订单ID
	ProductId   any // 商品ID
	Title       any // 商品名称快照
	Image       any // 商品图片快照
	UnitPrice   any // 成交单价
	Quantity    any // 购买数量
	TotalAmount any // 商品小计
	CreateTime  any // 创建时间
	UpdateTime  any // 修改时间
}
