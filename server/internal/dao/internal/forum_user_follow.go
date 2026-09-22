// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ForumUserFollowDao is the data access object for the table hg_forum_user_follow.
type ForumUserFollowDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  ForumUserFollowColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// ForumUserFollowColumns defines and stores column names for the table hg_forum_user_follow.
type ForumUserFollowColumns struct {
	Id           string // 关注ID
	UserId       string // 关注用户ID
	FollowUserId string // 被关注用户ID
	CreatedAt    string // 关注时间
}

// forumUserFollowColumns holds the columns for the table hg_forum_user_follow.
var forumUserFollowColumns = ForumUserFollowColumns{
	Id:           "id",
	UserId:       "user_id",
	FollowUserId: "follow_user_id",
	CreatedAt:    "created_at",
}

// NewForumUserFollowDao creates and returns a new DAO object for table data access.
func NewForumUserFollowDao(handlers ...gdb.ModelHandler) *ForumUserFollowDao {
	return &ForumUserFollowDao{
		group:    "default",
		table:    "hg_forum_user_follow",
		columns:  forumUserFollowColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ForumUserFollowDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ForumUserFollowDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ForumUserFollowDao) Columns() ForumUserFollowColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ForumUserFollowDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ForumUserFollowDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ForumUserFollowDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
