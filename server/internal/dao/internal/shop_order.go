// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShopOrderDao is the data access object for the table hg_shop_order.
type ShopOrderDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ShopOrderColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ShopOrderColumns defines and stores column names for the table hg_shop_order.
type ShopOrderColumns struct {
	Id             string // 订单ID
	OrderSn        string // 订单号
	UserId         string // 下单用户ID
	Status         string // 订单状态:pending_payment=待支付,paid=已支付,completed=已完成,cancelled=已取消,closed=已关闭
	TotalAmount    string // 商品总金额
	DiscountAmount string // 优惠金额
	FreightAmount  string // 运费金额
	PayableAmount  string // 应付金额
	PaidAmount     string // 实付金额
	Remark         string // 用户备注
	PaymentType    string // 支付方式
	PaymentNo      string // 第三方支付单号
	PaymentTime    string // 支付时间
	CompletedTime  string // 完成时间
	CancelledTime  string // 取消时间
	UpdateTime     string // 修改时间
	CreateTime     string // 创建时间
}

// shopOrderColumns holds the columns for the table hg_shop_order.
var shopOrderColumns = ShopOrderColumns{
	Id:             "id",
	OrderSn:        "order_sn",
	UserId:         "user_id",
	Status:         "status",
	TotalAmount:    "total_amount",
	DiscountAmount: "discount_amount",
	FreightAmount:  "freight_amount",
	PayableAmount:  "payable_amount",
	PaidAmount:     "paid_amount",
	Remark:         "remark",
	PaymentType:    "payment_type",
	PaymentNo:      "payment_no",
	PaymentTime:    "payment_time",
	CompletedTime:  "completed_time",
	CancelledTime:  "cancelled_time",
	UpdateTime:     "update_time",
	CreateTime:     "create_time",
}

// NewShopOrderDao creates and returns a new DAO object for table data access.
func NewShopOrderDao(handlers ...gdb.ModelHandler) *ShopOrderDao {
	return &ShopOrderDao{
		group:    "default",
		table:    "hg_shop_order",
		columns:  shopOrderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ShopOrderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ShopOrderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ShopOrderDao) Columns() ShopOrderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ShopOrderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ShopOrderDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ShopOrderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
