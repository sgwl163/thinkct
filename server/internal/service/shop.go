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
