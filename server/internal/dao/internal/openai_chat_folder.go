// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OpenaiChatFolderDao is the data access object for the table hg_openai_chat_folder.
type OpenaiChatFolderDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  OpenaiChatFolderColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// OpenaiChatFolderColumns defines and stores column names for the table hg_openai_chat_folder.
type OpenaiChatFolderColumns struct {
	Id        string // ID
	UserId    string // 用户ID
	Name      string // 文件夹名称
	Sort      string // 排序
	Status    string // 状态
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
	DeletedAt string // 删除时间
}

// openaiChatFolderColumns holds the columns for the table hg_openai_chat_folder.
var openaiChatFolderColumns = OpenaiChatFolderColumns{
	Id:        "id",
	UserId:    "user_id",
	Name:      "name",
	Sort:      "sort",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewOpenaiChatFolderDao creates and returns a new DAO object for table data access.
func NewOpenaiChatFolderDao(handlers ...gdb.ModelHandler) *OpenaiChatFolderDao {
	return &OpenaiChatFolderDao{
		group:    "default",
		table:    "hg_openai_chat_folder",
		columns:  openaiChatFolderColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OpenaiChatFolderDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OpenaiChatFolderDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OpenaiChatFolderDao) Columns() OpenaiChatFolderColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OpenaiChatFolderDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OpenaiChatFolderDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OpenaiChatFolderDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
