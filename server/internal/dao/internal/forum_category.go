// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ForumCategoryDao is the data access object for the table hg_forum_category.
type ForumCategoryDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  ForumCategoryColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// ForumCategoryColumns defines and stores column names for the table hg_forum_category.
type ForumCategoryColumns struct {
	Id          string // 板块ID
	Pid         string // 上级板块ID
	Name        string // 板块名称
	Description string // 板块简介
	Image       string // 板块图标
	Sort        string // 排序
	Status      string // 状态:1正常,2禁用
	TopicCount  string // 主题数量
	UpdatedAt   string // 修改时间
	CreatedAt   string // 创建时间
}

// forumCategoryColumns holds the columns for the table hg_forum_category.
var forumCategoryColumns = ForumCategoryColumns{
	Id:          "id",
	Pid:         "pid",
	Name:        "name",
	Description: "description",
	Image:       "image",
	Sort:        "sort",
	Status:      "status",
	TopicCount:  "topic_count",
	UpdatedAt:   "updated_at",
	CreatedAt:   "created_at",
}

// NewForumCategoryDao creates and returns a new DAO object for table data access.
func NewForumCategoryDao(handlers ...gdb.ModelHandler) *ForumCategoryDao {
	return &ForumCategoryDao{
		group:    "default",
		table:    "hg_forum_category",
		columns:  forumCategoryColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ForumCategoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ForumCategoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ForumCategoryDao) Columns() ForumCategoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ForumCategoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ForumCategoryDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ForumCategoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
