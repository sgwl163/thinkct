// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ForumCommentDao is the data access object for the table hg_forum_comment.
type ForumCommentDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  ForumCommentColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// ForumCommentColumns defines and stores column names for the table hg_forum_comment.
type ForumCommentColumns struct {
	Id          string // 评论ID
	TopicId     string // 主题ID
	UserId      string // 评论用户ID
	ParentId    string // 父评论ID
	ReplyUserId string // 被回复用户ID
	Content     string // 评论内容
	Status      string // 状态:normal,pending,rejected,hidden
	LikeCount   string // 点赞数量
	UpdatedAt   string // 修改时间
	CreatedAt   string // 创建时间
	DeletedAt   string // 删除时间
}

// forumCommentColumns holds the columns for the table hg_forum_comment.
var forumCommentColumns = ForumCommentColumns{
	Id:          "id",
	TopicId:     "topic_id",
	UserId:      "user_id",
	ParentId:    "parent_id",
	ReplyUserId: "reply_user_id",
	Content:     "content",
	Status:      "status",
	LikeCount:   "like_count",
	UpdatedAt:   "updated_at",
	CreatedAt:   "created_at",
	DeletedAt:   "deleted_at",
}

// NewForumCommentDao creates and returns a new DAO object for table data access.
func NewForumCommentDao(handlers ...gdb.ModelHandler) *ForumCommentDao {
	return &ForumCommentDao{
		group:    "default",
		table:    "hg_forum_comment",
		columns:  forumCommentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ForumCommentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ForumCommentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ForumCommentDao) Columns() ForumCommentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ForumCommentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ForumCommentDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ForumCommentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
