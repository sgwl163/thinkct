// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WebFeedbackDao is the data access object for the table hg_web_feedback.
type WebFeedbackDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  WebFeedbackColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// WebFeedbackColumns defines and stores column names for the table hg_web_feedback.
type WebFeedbackColumns struct {
	Id         string // ID
	Name       string // 名称
	Email      string // 邮箱
	Phone      string // 手机号
	Company    string // 公司名称
	Content    string // 富文本
	Answer     string // 备注
	Ip         string // IP
	Status     string // 状态:0=禁用,1=启用
	CreateTime string // 创建时间
	UpdateTime string // 修改时间
}

// webFeedbackColumns holds the columns for the table hg_web_feedback.
var webFeedbackColumns = WebFeedbackColumns{
	Id:         "id",
	Name:       "name",
	Email:      "email",
	Phone:      "phone",
	Company:    "company",
	Content:    "content",
	Answer:     "answer",
	Ip:         "ip",
	Status:     "status",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewWebFeedbackDao creates and returns a new DAO object for table data access.
func NewWebFeedbackDao(handlers ...gdb.ModelHandler) *WebFeedbackDao {
	return &WebFeedbackDao{
		group:    "default",
		table:    "hg_web_feedback",
		columns:  webFeedbackColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WebFeedbackDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WebFeedbackDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WebFeedbackDao) Columns() WebFeedbackColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WebFeedbackDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WebFeedbackDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *WebFeedbackDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
