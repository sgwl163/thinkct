// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ForumTopicDao is the data access object for the table hg_forum_topic.
type ForumTopicDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ForumTopicColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ForumTopicColumns defines and stores column names for the table hg_forum_topic.
type ForumTopicColumns struct {
	Id              string // 主题ID
	CategoryId      string // 板块ID
	UserId          string // 发布用户ID
	Title           string // 主题标题
	Keywords        string // 主题关键词，英文逗号分隔
	Description     string // 主题摘要
	Content         string // 主题内容
	Images          string // 主题图片，英文逗号分隔
	Status          string // 状态:normal,pending,rejected,hidden,closed
	IsTop           string // 是否置顶:0否,1是
	IsEssence       string // 是否精华:0否,1是
	CommentCount    string // 评论数量
	LikeCount       string // 点赞数量
	FavoriteCount   string // 收藏数量
	ViewCount       string // 浏览数量
	LastReplyUserId string // 最后回复用户ID
	LastReplyAt     string // 最后回复时间
	UpdatedAt       string // 修改时间
	CreatedAt       string // 创建时间
	DeletedAt       string // 删除时间
}

// forumTopicColumns holds the columns for the table hg_forum_topic.
var forumTopicColumns = ForumTopicColumns{
	Id:              "id",
	CategoryId:      "category_id",
	UserId:          "user_id",
	Title:           "title",
	Keywords:        "keywords",
	Description:     "description",
	Content:         "content",
	Images:          "images",
	Status:          "status",
	IsTop:           "is_top",
	IsEssence:       "is_essence",
	CommentCount:    "comment_count",
	LikeCount:       "like_count",
	FavoriteCount:   "favorite_count",
	ViewCount:       "view_count",
	LastReplyUserId: "last_reply_user_id",
	LastReplyAt:     "last_reply_at",
	UpdatedAt:       "updated_at",
	CreatedAt:       "created_at",
	DeletedAt:       "deleted_at",
}

// NewForumTopicDao creates and returns a new DAO object for table data access.
func NewForumTopicDao(handlers ...gdb.ModelHandler) *ForumTopicDao {
	return &ForumTopicDao{
		group:    "default",
		table:    "hg_forum_topic",
		columns:  forumTopicColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ForumTopicDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ForumTopicDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ForumTopicDao) Columns() ForumTopicColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ForumTopicDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ForumTopicDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ForumTopicDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
