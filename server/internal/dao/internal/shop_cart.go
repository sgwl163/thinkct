// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ShopCartDao is the data access object for the table hg_shop_cart.
type ShopCartDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ShopCartColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ShopCartColumns defines and stores column names for the table hg_shop_cart.
type ShopCartColumns struct {
	Id        string // 购物车ID
	UserId    string // 用户ID
	ProductId string // 商品ID
	Quantity  string // 购买数量
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// shopCartColumns holds the columns for the table hg_shop_cart.
var shopCartColumns = ShopCartColumns{
	Id:        "id",
	UserId:    "user_id",
	ProductId: "product_id",
	Quantity:  "quantity",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewShopCartDao creates and returns a new DAO object for table data access.
func NewShopCartDao(handlers ...gdb.ModelHandler) *ShopCartDao {
	return &ShopCartDao{
		group:    "default",
		table:    "hg_shop_cart",
		columns:  shopCartColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ShopCartDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ShopCartDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ShopCartDao) Columns() ShopCartColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ShopCartDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ShopCartDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ShopCartDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
