// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ImMessageDao is the data access object for the table hg_im_message.
type ImMessageDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ImMessageColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ImMessageColumns defines and stores column names for the table hg_im_message.
type ImMessageColumns struct {
	Id        string // 消息ID
	ChatType  string // 聊天类型: 1=单聊, 2=群聊, 3=系统消息
	FromId    string // 发送者ID
	ToId      string // 接收者ID(目标用户ID或群组ID)
	MsgType   string // 消息类型: 1=文本, 2=图片, 3=文件, 4=语音, 5=视频, 6=表情包
	Content   string // 消息内容(文本内容或文件URL/JSON)
	Status    string // 状态: 1=正常, 2=已撤回, 3=已删除
	CreatedAt string // 发送时间
	UpdatedAt string // 更新时间(如撤回时间)
}

// imMessageColumns holds the columns for the table hg_im_message.
var imMessageColumns = ImMessageColumns{
	Id:        "id",
	ChatType:  "chat_type",
	FromId:    "from_id",
	ToId:      "to_id",
	MsgType:   "msg_type",
	Content:   "content",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewImMessageDao creates and returns a new DAO object for table data access.
func NewImMessageDao(handlers ...gdb.ModelHandler) *ImMessageDao {
	return &ImMessageDao{
		group:    "default",
		table:    "hg_im_message",
		columns:  imMessageColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ImMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ImMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ImMessageDao) Columns() ImMessageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ImMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ImMessageDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ImMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
