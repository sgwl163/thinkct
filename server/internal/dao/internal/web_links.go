// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WebLinksDao is the data access object for the table hg_web_links.
type WebLinksDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  WebLinksColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// WebLinksColumns defines and stores column names for the table hg_web_links.
type WebLinksColumns struct {
	Id         string // ID
	UserId     string // 用户ID
	Title      string // 标题
	Link       string // 链接
	Target     string // 跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口
	Logo       string // LOGO
	Weigh      string // 权重
	Status     string // 状态:disable=禁用,enable=启用,pending_trial=待审
	Remark     string // 备注
	UpdateTime string // 修改时间
	CreateTime string // 创建时间
}

// webLinksColumns holds the columns for the table hg_web_links.
var webLinksColumns = WebLinksColumns{
	Id:         "id",
	UserId:     "user_id",
	Title:      "title",
	Link:       "link",
	Target:     "target",
	Logo:       "logo",
	Weigh:      "weigh",
	Status:     "status",
	Remark:     "remark",
	UpdateTime: "update_time",
	CreateTime: "create_time",
}

// NewWebLinksDao creates and returns a new DAO object for table data access.
func NewWebLinksDao(handlers ...gdb.ModelHandler) *WebLinksDao {
	return &WebLinksDao{
		group:    "default",
		table:    "hg_web_links",
		columns:  webLinksColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WebLinksDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WebLinksDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WebLinksDao) Columns() WebLinksColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WebLinksDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WebLinksDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *WebLinksDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
