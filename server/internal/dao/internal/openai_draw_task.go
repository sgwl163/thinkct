// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OpenaiDrawTaskDao is the data access object for the table hg_openai_draw_task.
type OpenaiDrawTaskDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  OpenaiDrawTaskColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// OpenaiDrawTaskColumns defines and stores column names for the table hg_openai_draw_task.
type OpenaiDrawTaskColumns struct {
	Id             string // 任务ID
	UserId         string // 用户ID
	Prompt         string // 绘画提示词
	ReferenceImage string // 参考图URL
	Model          string // 使用的模型
	Version        string // 模型版本
	Ratio          string // 图片比例
	Raw            string // 是否开启RAW模式
	ImageUrl       string // 生成的图片URL
	Status         string // 状态: 0=进行中, 1=成功, 2=失败
	FailReason     string // 失败原因
	CreatedAt      string // 创建时间
	UpdatedAt      string // 修改时间
	DeletedAt      string // 删除时间
	NegativePrompt string // 负面提示词
	Stylize        string // 风格化程度(0-1000)
	Chaos          string // 多样性(0-100)
	Seed           string // 随机种子
	MjTaskId       string // MJ远端任务ID
	Buttons        string // MJ操作按钮JSON
	Progress       string // 进度(0-100)
}

// openaiDrawTaskColumns holds the columns for the table hg_openai_draw_task.
var openaiDrawTaskColumns = OpenaiDrawTaskColumns{
	Id:             "id",
	UserId:         "user_id",
	Prompt:         "prompt",
	ReferenceImage: "reference_image",
	Model:          "model",
	Version:        "version",
	Ratio:          "ratio",
	Raw:            "raw",
	ImageUrl:       "image_url",
	Status:         "status",
	FailReason:     "fail_reason",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
	NegativePrompt: "negative_prompt",
	Stylize:        "stylize",
	Chaos:          "chaos",
	Seed:           "seed",
	MjTaskId:       "mj_task_id",
	Buttons:        "buttons",
	Progress:       "progress",
}

// NewOpenaiDrawTaskDao creates and returns a new DAO object for table data access.
func NewOpenaiDrawTaskDao(handlers ...gdb.ModelHandler) *OpenaiDrawTaskDao {
	return &OpenaiDrawTaskDao{
		group:    "default",
		table:    "hg_openai_draw_task",
		columns:  openaiDrawTaskColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OpenaiDrawTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OpenaiDrawTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OpenaiDrawTaskDao) Columns() OpenaiDrawTaskColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OpenaiDrawTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OpenaiDrawTaskDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OpenaiDrawTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
