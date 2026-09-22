// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OauthVersionDao is the data access object for the table hg_oauth_version.
type OauthVersionDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  OauthVersionColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// OauthVersionColumns defines and stores column names for the table hg_oauth_version.
type OauthVersionColumns struct {
	Id        string // ID
	ProductId string // 产品ID
	Cover     string // 封面
	Title     string // 标题
	Content   string // 详情
	Version   string // 版本号
	File      string // 文件
	Status    string // 状态:1=启用,0=禁用
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

// oauthVersionColumns holds the columns for the table hg_oauth_version.
var oauthVersionColumns = OauthVersionColumns{
	Id:        "id",
	ProductId: "product_id",
	Cover:     "cover",
	Title:     "title",
	Content:   "content",
	Version:   "version",
	File:      "file",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewOauthVersionDao creates and returns a new DAO object for table data access.
func NewOauthVersionDao(handlers ...gdb.ModelHandler) *OauthVersionDao {
	return &OauthVersionDao{
		group:    "default",
		table:    "hg_oauth_version",
		columns:  oauthVersionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OauthVersionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OauthVersionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OauthVersionDao) Columns() OauthVersionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OauthVersionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OauthVersionDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OauthVersionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
