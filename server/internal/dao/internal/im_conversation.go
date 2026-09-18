// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ImConversationDao is the data access object for the table hg_im_conversation.
type ImConversationDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  ImConversationColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// ImConversationColumns defines and stores column names for the table hg_im_conversation.
type ImConversationColumns struct {
	Id             string // 会话ID
	UserId         string // 归属用户ID
	ChatType       string // 会话类型: 1=单聊, 2=群聊, 3=系统通知
	TargetId       string // 目标ID(好友ID或群组ID)
	LastMsgId      string // 最后一条消息ID
	LastMsgContent string // 最后消息内容摘要(用于列表展示)
	LastMsgTime    string // 最后消息时间
	UnreadCount    string // 未读消息数
	IsTop          string // 是否置顶: 0=否, 1=是
	IsDisturb      string // 消息免打扰: 0=否, 1=是
	CreatedAt      string // 创建时间
	UpdatedAt      string // 更新时间
}

// imConversationColumns holds the columns for the table hg_im_conversation.
var imConversationColumns = ImConversationColumns{
	Id:             "id",
	UserId:         "user_id",
	ChatType:       "chat_type",
	TargetId:       "target_id",
	LastMsgId:      "last_msg_id",
	LastMsgContent: "last_msg_content",
	LastMsgTime:    "last_msg_time",
	UnreadCount:    "unread_count",
	IsTop:          "is_top",
	IsDisturb:      "is_disturb",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewImConversationDao creates and returns a new DAO object for table data access.
func NewImConversationDao(handlers ...gdb.ModelHandler) *ImConversationDao {
	return &ImConversationDao{
		group:    "default",
		table:    "hg_im_conversation",
		columns:  imConversationColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ImConversationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ImConversationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ImConversationDao) Columns() ImConversationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ImConversationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ImConversationDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ImConversationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
