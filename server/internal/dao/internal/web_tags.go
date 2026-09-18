// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WebTagsDao is the data access object for the table hg_web_tags.
type WebTagsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  WebTagsColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// WebTagsColumns defines and stores column names for the table hg_web_tags.
type WebTagsColumns struct {
	Id            string // ID
	Name          string // 名称
	Type          string // 标签类型:default=default,success=success,info=info,warning=warning,danger=danger
	DocumentCount string // 文档数
	Seotitle      string // SEO标题
	Keywords      string // 关键词
	Description   string // 描述
	InnerChain    string // 自动内链:0=关,1=开
	Weigh         string // 权重
	Status        string // 状态:0=禁用,1=启用
	UpdateTime    string // 修改时间
	CreateTime    string // 创建时间
}

// webTagsColumns holds the columns for the table hg_web_tags.
var webTagsColumns = WebTagsColumns{
	Id:            "id",
	Name:          "name",
	Type:          "type",
	DocumentCount: "document_count",
	Seotitle:      "seotitle",
	Keywords:      "keywords",
	Description:   "description",
	InnerChain:    "inner_chain",
	Weigh:         "weigh",
	Status:        "status",
	UpdateTime:    "update_time",
	CreateTime:    "create_time",
}

// NewWebTagsDao creates and returns a new DAO object for table data access.
func NewWebTagsDao(handlers ...gdb.ModelHandler) *WebTagsDao {
	return &WebTagsDao{
		group:    "default",
		table:    "hg_web_tags",
		columns:  webTagsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WebTagsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WebTagsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WebTagsDao) Columns() WebTagsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WebTagsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WebTagsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *WebTagsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
