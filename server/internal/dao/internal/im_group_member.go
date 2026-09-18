// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ImGroupMemberDao is the data access object for the table hg_im_group_member.
type ImGroupMemberDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  ImGroupMemberColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// ImGroupMemberColumns defines and stores column names for the table hg_im_group_member.
type ImGroupMemberColumns struct {
	Id         string // ID
	GroupId    string // 群组ID
	UserId     string // 用户ID
	Role       string // 角色: 1=群主, 2=管理员, 3=普通成员
	RemarkName string // 群内昵称
	IsMute     string // 是否禁言: 0=否, 1=是
	CreatedAt  string // 加入时间
}

// imGroupMemberColumns holds the columns for the table hg_im_group_member.
var imGroupMemberColumns = ImGroupMemberColumns{
	Id:         "id",
	GroupId:    "group_id",
	UserId:     "user_id",
	Role:       "role",
	RemarkName: "remark_name",
	IsMute:     "is_mute",
	CreatedAt:  "created_at",
}

// NewImGroupMemberDao creates and returns a new DAO object for table data access.
func NewImGroupMemberDao(handlers ...gdb.ModelHandler) *ImGroupMemberDao {
	return &ImGroupMemberDao{
		group:    "default",
		table:    "hg_im_group_member",
		columns:  imGroupMemberColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ImGroupMemberDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ImGroupMemberDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ImGroupMemberDao) Columns() ImGroupMemberColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ImGroupMemberDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ImGroupMemberDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ImGroupMemberDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
