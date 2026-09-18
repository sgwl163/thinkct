// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OpenaiChatMessageDao is the data access object for the table hg_openai_chat_message.
type OpenaiChatMessageDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  OpenaiChatMessageColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// OpenaiChatMessageColumns defines and stores column names for the table hg_openai_chat_message.
type OpenaiChatMessageColumns struct {
	Id          string // ID
	UserId      string // 用户ID
	SessionId   string // 会话ID
	Role        string // 角色(user/assistant/system)
	Content     string // 消息内容
	Model       string // 使用的模型
	Feedback    string // 反馈: 0=无, 1=喜欢, 2=不喜欢
	ReqTokens   string // 请求token数
	ResTokens   string // 响应token数
	TotalTokens string // 总token数
	Status      string // 状态
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
	DeletedAt   string // 删除时间
}

// openaiChatMessageColumns holds the columns for the table hg_openai_chat_message.
var openaiChatMessageColumns = OpenaiChatMessageColumns{
	Id:          "id",
	UserId:      "user_id",
	SessionId:   "session_id",
	Role:        "role",
	Content:     "content",
	Model:       "model",
	Feedback:    "feedback",
	ReqTokens:   "req_tokens",
	ResTokens:   "res_tokens",
	TotalTokens: "total_tokens",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewOpenaiChatMessageDao creates and returns a new DAO object for table data access.
func NewOpenaiChatMessageDao(handlers ...gdb.ModelHandler) *OpenaiChatMessageDao {
	return &OpenaiChatMessageDao{
		group:    "default",
		table:    "hg_openai_chat_message",
		columns:  openaiChatMessageColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OpenaiChatMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OpenaiChatMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OpenaiChatMessageDao) Columns() OpenaiChatMessageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OpenaiChatMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OpenaiChatMessageDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OpenaiChatMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
