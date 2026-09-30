// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserLogDao is the data access object for the table hg_user_log.
type UserLogDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserLogColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserLogColumns defines and stores column names for the table hg_user_log.
type UserLogColumns struct {
	Id         string // 日志ID
	UserId     string // 用户ID
	Username   string // 用户名
	LogType    string // 记录类型:login=登录,operation=操作
	Action     string // 操作标识
	ActionName string // 操作名称
	EventAt    string // 事件时间
	Ip         string // IP地址
	UserAgent  string // UA信息
	Message    string // 结果消息
	Extra      string // 扩展数据
	Status     string // 状态
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// userLogColumns holds the columns for the table hg_user_log.
var userLogColumns = UserLogColumns{
	Id:         "id",
	UserId:     "user_id",
	Username:   "username",
	LogType:    "log_type",
	Action:     "action",
	ActionName: "action_name",
	EventAt:    "event_at",
	Ip:         "ip",
	UserAgent:  "user_agent",
	Message:    "message",
	Extra:      "extra",
	Status:     "status",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewUserLogDao creates and returns a new DAO object for table data access.
func NewUserLogDao(handlers ...gdb.ModelHandler) *UserLogDao {
	return &UserLogDao{
		group:    "default",
		table:    "hg_user_log",
		columns:  userLogColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserLogDao) Columns() UserLogColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserLogDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
