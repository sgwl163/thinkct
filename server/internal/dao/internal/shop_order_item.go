// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShopOrderItemDao is the data access object for the table hg_shop_order_item.
type ShopOrderItemDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  ShopOrderItemColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// ShopOrderItemColumns defines and stores column names for the table hg_shop_order_item.
type ShopOrderItemColumns struct {
	Id          string // 订单明细ID
	OrderId     string // 订单ID
	ProductId   string // 商品ID
	Title       string // 商品名称快照
	Image       string // 商品图片快照
	UnitPrice   string // 成交单价
	Quantity    string // 购买数量
	TotalAmount string // 商品小计
	CreateTime  string // 创建时间
	UpdateTime  string // 修改时间
}

// shopOrderItemColumns holds the columns for the table hg_shop_order_item.
var shopOrderItemColumns = ShopOrderItemColumns{
	Id:          "id",
	OrderId:     "order_id",
	ProductId:   "product_id",
	Title:       "title",
	Image:       "image",
	UnitPrice:   "unit_price",
	Quantity:    "quantity",
	TotalAmount: "total_amount",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
}

// NewShopOrderItemDao creates and returns a new DAO object for table data access.
func NewShopOrderItemDao(handlers ...gdb.ModelHandler) *ShopOrderItemDao {
	return &ShopOrderItemDao{
		group:    "default",
		table:    "hg_shop_order_item",
		columns:  shopOrderItemColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ShopOrderItemDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ShopOrderItemDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ShopOrderItemDao) Columns() ShopOrderItemColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ShopOrderItemDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ShopOrderItemDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ShopOrderItemDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
