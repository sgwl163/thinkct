package consts

import (
	"hotgo/internal/library/dict"
	"hotgo/internal/model"
)

const (
	ShopOrderStatusPendingPayment = "pending_payment"
	ShopOrderStatusPaid           = "paid"
	ShopOrderStatusCompleted      = "completed"
	ShopOrderStatusCancelled      = "cancelled"
	ShopOrderStatusClosed         = "closed"
)

var ShopOrderStatusOptions = []*model.Option{
	dict.GenWarningOption(ShopOrderStatusPendingPayment, "待支付"),
	dict.GenInfoOption(ShopOrderStatusPaid, "已支付"),
	dict.GenSuccessOption(ShopOrderStatusCompleted, "已完成"),
	dict.GenDefaultOption(ShopOrderStatusCancelled, "已取消"),
	dict.GenDefaultOption(ShopOrderStatusClosed, "已关闭"),
}

func init() {
	dict.RegisterEnums("ShopOrderStatusOptions", "商城订单状态选项", ShopOrderStatusOptions)
}
