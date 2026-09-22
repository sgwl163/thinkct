// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ForumInteractionDao is the data access object for the table hg_forum_interaction.
type ForumInteractionDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  ForumInteractionColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// ForumInteractionColumns defines and stores column names for the table hg_forum_interaction.
type ForumInteractionColumns struct {
	Id         string // 互动ID
	TargetType string // 互动对象类型:topic,comment
	TargetId   string // 互动对象ID
	UserId     string // 用户ID
	ActionType string // 互动类型:like,favorite
	CreatedAt  string // 创建时间
}

// forumInteractionColumns holds the columns for the table hg_forum_interaction.
var forumInteractionColumns = ForumInteractionColumns{
	Id:         "id",
	TargetType: "target_type",
	TargetId:   "target_id",
	UserId:     "user_id",
	ActionType: "action_type",
	CreatedAt:  "created_at",
}

// NewForumInteractionDao creates and returns a new DAO object for table data access.
func NewForumInteractionDao(handlers ...gdb.ModelHandler) *ForumInteractionDao {
	return &ForumInteractionDao{
		group:    "default",
		table:    "hg_forum_interaction",
		columns:  forumInteractionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ForumInteractionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ForumInteractionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ForumInteractionDao) Columns() ForumInteractionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ForumInteractionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ForumInteractionDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ForumInteractionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
