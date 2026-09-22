package shop

import (
	"context"

	v1 "hotgo/api/api/shop/v1"
	"hotgo/internal/service"
)

func (c *ControllerV1) CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (res *v1.CreateOrderRes, err error) {
	return service.ApiShop().CreateOrder(ctx, req)
}

func (c *ControllerV1) OrderList(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error) {
	return service.ApiShop().OrderList(ctx, req)
}

func (c *ControllerV1) OrderDetail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error) {
	return service.ApiShop().OrderDetail(ctx, req)
}
