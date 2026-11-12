package shop

import (
	"context"

	v1 "hotgo/api/api/shop/v1"
	"hotgo/internal/service"
)

func (c *ControllerV1) CartAdd(ctx context.Context, req *v1.CartAddReq) (*v1.CartAddRes, error) {
	return service.ApiShop().CartAdd(ctx, req)
}
func (c *ControllerV1) CartList(ctx context.Context, req *v1.CartListReq) (*v1.CartListRes, error) {
	return service.ApiShop().CartList(ctx, req)
}
func (c *ControllerV1) CartUpdate(ctx context.Context, req *v1.CartUpdateReq) (*v1.CartUpdateRes, error) {
	return service.ApiShop().CartUpdate(ctx, req)
}
func (c *ControllerV1) CartDelete(ctx context.Context, req *v1.CartDeleteReq) (*v1.CartDeleteRes, error) {
	return service.ApiShop().CartDelete(ctx, req)
}
func (c *ControllerV1) CartCheckout(ctx context.Context, req *v1.CartCheckoutReq) (*v1.CartCheckoutRes, error) {
	return service.ApiShop().CartCheckout(ctx, req)
}

func (c *ControllerV1) PurchaseOrder(ctx context.Context, req *v1.PurchaseOrderReq) (res *v1.PurchaseOrderRes, err error) {
	return service.ApiShop().PurchaseOrder(ctx, req)
}

func (c *ControllerV1) CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (res *v1.CreateOrderRes, err error) {
	return service.ApiShop().CreateOrder(ctx, req)
}

func (c *ControllerV1) OrderList(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error) {
	return service.ApiShop().OrderList(ctx, req)
}

func (c *ControllerV1) OrderDetail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error) {
	return service.ApiShop().OrderDetail(ctx, req)
}
