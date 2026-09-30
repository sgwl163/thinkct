// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppAuthDao is the data access object for the table hg_app_auth.
type AppAuthDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AppAuthColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AppAuthColumns defines and stores column names for the table hg_app_auth.
type AppAuthColumns struct {
	Id         string // ID
	UserId     string // 用户ID
	ProductId  string // 产品ID
	Name       string // 名称
	Host       string // 域名
	Ip         string // IP
	Qq         string // QQ
	Phone      string // 电话
	Syskey     string // 授权码
	VersionId  string // 版本
	AuthMode   string // 验证方式:0=单域名验证,1=IP双重验证
	AuthDomain string // 验证域名:0=当前单域名方式,1=顶级泛域名方式
	Views      string // 浏览量
	Status     string // 状态:1=启用,0=禁用
	CreatedAt  string // 创建时间
	UpdatedAt  string // 修改时间
	ExpiredAt  string // 到期时间
}

// appAuthColumns holds the columns for the table hg_app_auth.
var appAuthColumns = AppAuthColumns{
	Id:         "id",
	UserId:     "user_id",
	ProductId:  "product_id",
	Name:       "name",
	Host:       "host",
	Ip:         "ip",
	Qq:         "qq",
	Phone:      "phone",
	Syskey:     "syskey",
	VersionId:  "version_id",
	AuthMode:   "auth_mode",
	AuthDomain: "auth_domain",
	Views:      "views",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	ExpiredAt:  "expired_at",
}

// NewAppAuthDao creates and returns a new DAO object for table data access.
func NewAppAuthDao(handlers ...gdb.ModelHandler) *AppAuthDao {
	return &AppAuthDao{
		group:    "default",
		table:    "hg_app_auth",
		columns:  appAuthColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppAuthDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppAuthDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppAuthDao) Columns() AppAuthColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppAuthDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppAuthDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppAuthDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
