// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppLogDao is the data access object for the table hg_app_log.
type AppLogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AppLogColumns      // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AppLogColumns defines and stores column names for the table hg_app_log.
type AppLogColumns struct {
	Id        string // ID
	UserId    string // 用户ID
	ProductId string // 产品ID
	Host      string // 域名
	Content   string // 内容
	Ip        string // IP
	Version   string // 版本号
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

// appLogColumns holds the columns for the table hg_app_log.
var appLogColumns = AppLogColumns{
	Id:        "id",
	UserId:    "user_id",
	ProductId: "product_id",
	Host:      "host",
	Content:   "content",
	Ip:        "ip",
	Version:   "version",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewAppLogDao creates and returns a new DAO object for table data access.
func NewAppLogDao(handlers ...gdb.ModelHandler) *AppLogDao {
	return &AppLogDao{
		group:    "default",
		table:    "hg_app_log",
		columns:  appLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppLogDao) Columns() AppLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
