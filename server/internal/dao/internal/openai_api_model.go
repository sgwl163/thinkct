// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OpenaiApiModelDao is the data access object for the table hg_openai_api_model.
type OpenaiApiModelDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  OpenaiApiModelColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// OpenaiApiModelColumns defines and stores column names for the table hg_openai_api_model.
type OpenaiApiModelColumns struct {
	Id        string // ID
	GroupId   string // 分组ID
	Name      string // 模型别名(如GPT-4)
	Model     string // 模型标识(如gpt-4-turbo)
	IsVision  string // 是否支持识图
	IsFile    string // 是否支持上传文件
	IsSearch  string // 是否支持联网搜索
	Status    string // 状态
	Sort      string // 排序
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
	DeletedAt string // 删除时间
}

// openaiApiModelColumns holds the columns for the table hg_openai_api_model.
var openaiApiModelColumns = OpenaiApiModelColumns{
	Id:        "id",
	GroupId:   "group_id",
	Name:      "name",
	Model:     "model",
	IsVision:  "is_vision",
	IsFile:    "is_file",
	IsSearch:  "is_search",
	Status:    "status",
	Sort:      "sort",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewOpenaiApiModelDao creates and returns a new DAO object for table data access.
func NewOpenaiApiModelDao(handlers ...gdb.ModelHandler) *OpenaiApiModelDao {
	return &OpenaiApiModelDao{
		group:    "default",
		table:    "hg_openai_api_model",
		columns:  openaiApiModelColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OpenaiApiModelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OpenaiApiModelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OpenaiApiModelDao) Columns() OpenaiApiModelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OpenaiApiModelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OpenaiApiModelDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OpenaiApiModelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
