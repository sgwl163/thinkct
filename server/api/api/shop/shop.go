package shop

import (
	"context"

	v1 "hotgo/api/api/shop/v1"
)

type IShopV1 interface {
	PurchaseOrder(ctx context.Context, req *v1.PurchaseOrderReq) (res *v1.PurchaseOrderRes, err error)
	CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (res *v1.CreateOrderRes, err error)
	OrderList(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error)
	OrderDetail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error)
	CartAdd(ctx context.Context, req *v1.CartAddReq) (res *v1.CartAddRes, err error)
	CartList(ctx context.Context, req *v1.CartListReq) (res *v1.CartListRes, err error)
	CartUpdate(ctx context.Context, req *v1.CartUpdateReq) (res *v1.CartUpdateRes, err error)
	CartDelete(ctx context.Context, req *v1.CartDeleteReq) (res *v1.CartDeleteRes, err error)
	CartCheckout(ctx context.Context, req *v1.CartCheckoutReq) (res *v1.CartCheckoutRes, err error)
}
