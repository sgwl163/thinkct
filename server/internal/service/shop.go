// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "hotgo/api/api/shop/v1"
)

type (
	IApiShop interface {
		CartAdd(ctx context.Context, req *v1.CartAddReq) (*v1.CartAddRes, error)
		CartList(ctx context.Context, req *v1.CartListReq) (*v1.CartListRes, error)
		CartUpdate(ctx context.Context, req *v1.CartUpdateReq) (*v1.CartUpdateRes, error)
		CartDelete(ctx context.Context, req *v1.CartDeleteReq) (*v1.CartDeleteRes, error)
		CartCheckout(ctx context.Context, req *v1.CartCheckoutReq) (*v1.CartCheckoutRes, error)
		PurchaseOrder(ctx context.Context, req *v1.PurchaseOrderReq) (res *v1.PurchaseOrderRes, err error)
		CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (res *v1.CreateOrderRes, err error)
		OrderList(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error)
		OrderDetail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error)
	}
)

var (
	localApiShop IApiShop
)

func ApiShop() IApiShop {
	if localApiShop == nil {
		panic("implement not found for interface IApiShop, forgot register?")
	}
	return localApiShop
}

func RegisterApiShop(i IApiShop) {
	localApiShop = i
}
