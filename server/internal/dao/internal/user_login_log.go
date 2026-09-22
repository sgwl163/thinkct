// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserLoginLogDao is the data access object for the table hg_user_login_log.
type UserLoginLogDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  UserLoginLogColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// UserLoginLogColumns defines and stores column names for the table hg_user_login_log.
type UserLoginLogColumns struct {
	Id        string // 日志ID
	UserId    string // 用户ID
	Username  string // 用户名
	LoginAt   string // 登录时间
	LoginIp   string // 登录IP
	UserAgent string // UA信息
	ErrMsg    string // 错误提示
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// userLoginLogColumns holds the columns for the table hg_user_login_log.
var userLoginLogColumns = UserLoginLogColumns{
	Id:        "id",
	UserId:    "user_id",
	Username:  "username",
	LoginAt:   "login_at",
	LoginIp:   "login_ip",
	UserAgent: "user_agent",
	ErrMsg:    "err_msg",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewUserLoginLogDao creates and returns a new DAO object for table data access.
func NewUserLoginLogDao(handlers ...gdb.ModelHandler) *UserLoginLogDao {
	return &UserLoginLogDao{
		group:    "default",
		table:    "hg_user_login_log",
		columns:  userLoginLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserLoginLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserLoginLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserLoginLogDao) Columns() UserLoginLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserLoginLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserLoginLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserLoginLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
