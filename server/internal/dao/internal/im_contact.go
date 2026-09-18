// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ImContactDao is the data access object for the table hg_im_contact.
type ImContactDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ImContactColumns   // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ImContactColumns defines and stores column names for the table hg_im_contact.
type ImContactColumns struct {
	Id        string // ID
	UserId    string // 用户ID
	FriendId  string // 好友ID
	Remark    string // 好友备注名
	Status    string // 状态: 1=正常, 2=拉黑
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// imContactColumns holds the columns for the table hg_im_contact.
var imContactColumns = ImContactColumns{
	Id:        "id",
	UserId:    "user_id",
	FriendId:  "friend_id",
	Remark:    "remark",
	Status:    "status",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewImContactDao creates and returns a new DAO object for table data access.
func NewImContactDao(handlers ...gdb.ModelHandler) *ImContactDao {
	return &ImContactDao{
		group:    "default",
		table:    "hg_im_contact",
		columns:  imContactColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ImContactDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ImContactDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ImContactDao) Columns() ImContactColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ImContactDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ImContactDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ImContactDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
