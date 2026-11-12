// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ShopCart is the golang structure for table shop_cart.
type ShopCart struct {
	Id        uint64 `json:"id"        orm:"id"         description:"购物车ID"`
	UserId    uint64 `json:"userId"    orm:"user_id"    description:"用户ID"`
	ProductId uint64 `json:"productId" orm:"product_id" description:"商品ID"`
	Quantity  uint   `json:"quantity"  orm:"quantity"   description:"购买数量"`
	CreatedAt uint64 `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt uint64 `json:"updatedAt" orm:"updated_at" description:"更新时间"`
}
