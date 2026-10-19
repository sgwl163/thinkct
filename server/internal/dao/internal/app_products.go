// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppProductsDao is the data access object for the table hg_app_products.
type AppProductsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AppProductsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AppProductsColumns defines and stores column names for the table hg_app_products.
type AppProductsColumns struct {
	Id          string // ID
	Cover       string // 封面
	Title       string // 标题
	Keywords    string // 关键词
	Description string // 描述
	Name        string // 名称
	Alias       string // 别名
	Content     string // 内容
	Price       string // 价格
	Ysweb       string // 演示站
	Unique      string // 特点
	Weigh       string // 权重
	Status      string // 状态:0=禁用,1=启用
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
}

// appProductsColumns holds the columns for the table hg_app_products.
var appProductsColumns = AppProductsColumns{
	Id:          "id",
	Cover:       "cover",
	Title:       "title",
	Keywords:    "keywords",
	Description: "description",
	Name:        "name",
	Alias:       "alias",
	Content:     "content",
	Price:       "price",
	Ysweb:       "ysweb",
	Unique:      "unique",
	Weigh:       "weigh",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewAppProductsDao creates and returns a new DAO object for table data access.
func NewAppProductsDao(handlers ...gdb.ModelHandler) *AppProductsDao {
	return &AppProductsDao{
		group:    "default",
		table:    "hg_app_products",
		columns:  appProductsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppProductsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppProductsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppProductsDao) Columns() AppProductsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppProductsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppProductsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppProductsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
