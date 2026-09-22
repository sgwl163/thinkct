// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ForumReportDao is the data access object for the table hg_forum_report.
type ForumReportDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ForumReportColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ForumReportColumns defines and stores column names for the table hg_forum_report.
type ForumReportColumns struct {
	Id          string // 举报ID
	TargetType  string // 举报对象类型:topic,comment
	TargetId    string // 举报对象ID
	UserId      string // 举报用户ID
	Reason      string // 举报原因
	Description string // 补充说明
	Status      string // 处理状态:pending,processed,rejected
	HandledBy   string // 处理管理员ID
	HandledAt   string // 处理时间
	CreatedAt   string // 创建时间
}

// forumReportColumns holds the columns for the table hg_forum_report.
var forumReportColumns = ForumReportColumns{
	Id:          "id",
	TargetType:  "target_type",
	TargetId:    "target_id",
	UserId:      "user_id",
	Reason:      "reason",
	Description: "description",
	Status:      "status",
	HandledBy:   "handled_by",
	HandledAt:   "handled_at",
	CreatedAt:   "created_at",
}

// NewForumReportDao creates and returns a new DAO object for table data access.
func NewForumReportDao(handlers ...gdb.ModelHandler) *ForumReportDao {
	return &ForumReportDao{
		group:    "default",
		table:    "hg_forum_report",
		columns:  forumReportColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ForumReportDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ForumReportDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ForumReportDao) Columns() ForumReportColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ForumReportDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ForumReportDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ForumReportDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
