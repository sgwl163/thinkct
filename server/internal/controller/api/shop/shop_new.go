package shop

import (
	"hotgo/api/api/shop"
)

type ControllerV1 struct{}

func NewV1() shop.IShopV1 {
	return &ControllerV1{}
}
