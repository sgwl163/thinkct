// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OauthDaobanDao is the data access object for the table hg_oauth_daoban.
type OauthDaobanDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  OauthDaobanColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// OauthDaobanColumns defines and stores column names for the table hg_oauth_daoban.
type OauthDaobanColumns struct {
	Id        string // ID
	ProductId string // 产品ID
	Host      string // 域名
	Ip        string // IP
	Version   string // 版本号
	Code      string // 执行代码
	Data      string // 数据
	Views     string // 浏览量
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

// oauthDaobanColumns holds the columns for the table hg_oauth_daoban.
var oauthDaobanColumns = OauthDaobanColumns{
	Id:        "id",
	ProductId: "product_id",
	Host:      "host",
	Ip:        "ip",
	Version:   "version",
	Code:      "code",
	Data:      "data",
	Views:     "views",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewOauthDaobanDao creates and returns a new DAO object for table data access.
func NewOauthDaobanDao(handlers ...gdb.ModelHandler) *OauthDaobanDao {
	return &OauthDaobanDao{
		group:    "default",
		table:    "hg_oauth_daoban",
		columns:  oauthDaobanColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OauthDaobanDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OauthDaobanDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OauthDaobanDao) Columns() OauthDaobanColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OauthDaobanDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OauthDaobanDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OauthDaobanDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
