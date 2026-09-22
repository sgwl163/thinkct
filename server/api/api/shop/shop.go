package shop

import (
	"context"

	v1 "hotgo/api/api/shop/v1"
)

type IShopV1 interface {
	CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (res *v1.CreateOrderRes, err error)
	OrderList(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error)
	OrderDetail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error)
}
