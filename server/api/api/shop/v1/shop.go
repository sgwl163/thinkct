package v1

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

type OrderProduct struct {
	ProductId uint64 `json:"productId" v:"required|min:1#商品ID不能为空|商品ID不正确" dc:"商品ID"`
	Quantity  uint   `json:"quantity" v:"required|min:1|max:999#购买数量不能为空|购买数量最少为1|购买数量最多为999" dc:"购买数量"`
}

type CreateOrderReq struct {
	g.Meta   `path:"/shop/order/create" method:"post" tags:"商城订单" summary:"创建商城订单"`
	Products []OrderProduct `json:"products" v:"required|length:1,50#商品不能为空|商品数量最多50种" dc:"商品列表"`
	Remark   string         `json:"remark" v:"length:0,500#订单备注不能超过500个字符" dc:"订单备注"`
}

type CreateOrderRes struct {
	Order *OrderDetail `json:"order" dc:"订单详情"`
}

type OrderListReq struct {
	g.Meta `path:"/shop/order/list" method:"get" tags:"商城订单" summary:"获取我的商城订单"`
	form.PageReq
	Status string `json:"status" dc:"订单状态，可选"`
}

type OrderListRes struct {
	form.PageRes
	List []*OrderDetail `json:"list" dc:"订单列表"`
}

type OrderDetailReq struct {
	g.Meta  `path:"/shop/order/detail" method:"get" tags:"商城订单" summary:"获取商城订单详情"`
	OrderSn string `json:"orderSn" v:"required#订单号不能为空" dc:"订单号"`
}

type OrderDetailRes struct {
	Order *OrderDetail `json:"order" dc:"订单详情"`
}

type OrderDetail struct {
	entity.ShopOrder
	Items []*entity.ShopOrderItem `json:"items" dc:"订单商品明细"`
}
