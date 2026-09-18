// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OpenaiChatSessionDao is the data access object for the table hg_openai_chat_session.
type OpenaiChatSessionDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  OpenaiChatSessionColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// OpenaiChatSessionColumns defines and stores column names for the table hg_openai_chat_session.
type OpenaiChatSessionColumns struct {
	Id          string // ID
	UserId      string // 用户ID
	FolderId    string // 文件夹ID
	Title       string // 会话标题
	IsTop       string // 是否置顶
	Prompt      string // 系统提示词
	Model       string // 使用的模型
	Temperature string // 温度值
	Status      string // 状态
	CreatedAt   string // 创建时间
	UpdatedAt   string // 修改时间
	DeletedAt   string // 删除时间
}

// openaiChatSessionColumns holds the columns for the table hg_openai_chat_session.
var openaiChatSessionColumns = OpenaiChatSessionColumns{
	Id:          "id",
	UserId:      "user_id",
	FolderId:    "folder_id",
	Title:       "title",
	IsTop:       "is_top",
	Prompt:      "prompt",
	Model:       "model",
	Temperature: "temperature",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewOpenaiChatSessionDao creates and returns a new DAO object for table data access.
func NewOpenaiChatSessionDao(handlers ...gdb.ModelHandler) *OpenaiChatSessionDao {
	return &OpenaiChatSessionDao{
		group:    "default",
		table:    "hg_openai_chat_session",
		columns:  openaiChatSessionColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OpenaiChatSessionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OpenaiChatSessionDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OpenaiChatSessionDao) Columns() OpenaiChatSessionColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OpenaiChatSessionDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OpenaiChatSessionDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OpenaiChatSessionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
