package v1

import "github.com/gogf/gf/v2/frame/g"

type CartAddReq struct {
	g.Meta `path:"/shop/cart/add" method:"post" tags:"商城购物车" summary:"加入购物车"`
	OrderProduct
}
type CartAddRes struct {
	Id uint64 `json:"id"`
}
type CartListReq struct {
	g.Meta `path:"/shop/cart/list" method:"get" tags:"商城购物车" summary:"购物车列表"`
}
type CartItem struct {
	CartId      uint64  `json:"cartId"`
	ProductId   uint64  `json:"productId"`
	Title       string  `json:"title"`
	Image       string  `json:"image"`
	UnitPrice   float64 `json:"unitPrice"`
	Quantity    uint    `json:"quantity"`
	TotalAmount float64 `json:"totalAmount"`
}
type CartListRes struct {
	Items       []*CartItem `json:"items"`
	Balance     uint        `json:"balance"`
	TotalAmount float64     `json:"totalAmount"`
}
type CartUpdateReq struct {
	g.Meta   `path:"/shop/cart/update" method:"post" tags:"商城购物车" summary:"修改购物车数量"`
	Id       uint64 `json:"id" v:"required|min:1"`
	Quantity uint   `json:"quantity" v:"required|min:1|max:999"`
}
type CartUpdateRes struct{}
type CartDeleteReq struct {
	g.Meta `path:"/shop/cart/delete" method:"post" tags:"商城购物车" summary:"删除购物车商品"`
	Id     uint64 `json:"id" v:"required|min:1"`
}
type CartDeleteRes struct{}
type CartCheckoutReq struct {
	g.Meta    `path:"/shop/cart/checkout" method:"post" tags:"商城购物车" summary:"余额结算购物车"`
	CartIds   []uint64 `json:"cartIds" v:"required|length:1,50"`
	RequestId string   `json:"requestId" v:"required|regex:^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-4[0-9a-fA-F]{3}-[89aAbB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$"`
	Remark    string   `json:"remark" v:"length:0,500"`
}
type CartCheckoutRes struct {
	Order *OrderDetail `json:"order"`
}
