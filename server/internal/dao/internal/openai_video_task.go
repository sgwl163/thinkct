// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OpenaiVideoTaskDao is the data access object for the table hg_openai_video_task.
type OpenaiVideoTaskDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  OpenaiVideoTaskColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// OpenaiVideoTaskColumns defines and stores column names for the table hg_openai_video_task.
type OpenaiVideoTaskColumns struct {
	Id             string // 任务ID
	UserId         string // 用户ID
	Prompt         string // 视频提示词
	ReferenceImage string // 参考图/首尾帧URL
	LastFrame      string // 尾帧图片URL
	VoiceDubUrl    string // 自定义配音URL
	BgMusicUrl     string // 背景音乐URL
	GenerateAudio  string // 是否生成同步音频: 1=是 0=否
	Model          string // 使用的模型
	VideoUrl       string // 生成的视频URL
	Status         string // 状态: 0=进行中 1=成功, 2=失败
	FailReason     string // 失败原因
	RemoteTaskId   string // 远端任务ID
	CreatedAt      string // 创建时间
	UpdatedAt      string // 修改时间
	DeletedAt      string // 删除时间
}

// openaiVideoTaskColumns holds the columns for the table hg_openai_video_task.
var openaiVideoTaskColumns = OpenaiVideoTaskColumns{
	Id:             "id",
	UserId:         "user_id",
	Prompt:         "prompt",
	ReferenceImage: "reference_image",
	LastFrame:      "last_frame",
	VoiceDubUrl:    "voice_dub_url",
	BgMusicUrl:     "bg_music_url",
	GenerateAudio:  "generate_audio",
	Model:          "model",
	VideoUrl:       "video_url",
	Status:         "status",
	FailReason:     "fail_reason",
	RemoteTaskId:   "remote_task_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
}

// NewOpenaiVideoTaskDao creates and returns a new DAO object for table data access.
func NewOpenaiVideoTaskDao(handlers ...gdb.ModelHandler) *OpenaiVideoTaskDao {
	return &OpenaiVideoTaskDao{
		group:    "default",
		table:    "hg_openai_video_task",
		columns:  openaiVideoTaskColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OpenaiVideoTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OpenaiVideoTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OpenaiVideoTaskDao) Columns() OpenaiVideoTaskColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OpenaiVideoTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OpenaiVideoTaskDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OpenaiVideoTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
