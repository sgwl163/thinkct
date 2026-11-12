// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ShopCart is the golang structure of table hg_shop_cart for DAO operations like Where/Data.
type ShopCart struct {
	g.Meta    `orm:"table:hg_shop_cart, do:true"`
	Id        any // 购物车ID
	UserId    any // 用户ID
	ProductId any // 商品ID
	Quantity  any // 购买数量
	CreatedAt any // 创建时间
	UpdatedAt any // 更新时间
}
