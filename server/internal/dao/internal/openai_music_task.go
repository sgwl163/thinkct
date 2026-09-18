// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// OpenaiMusicTaskDao is the data access object for the table hg_openai_music_task.
type OpenaiMusicTaskDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  OpenaiMusicTaskColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// OpenaiMusicTaskColumns defines and stores column names for the table hg_openai_music_task.
type OpenaiMusicTaskColumns struct {
	Id               string // 任务ID
	UserId           string // 用户ID
	Prompt           string // 提示词(灵感或歌词)
	Type             string // 创作模式: 1=灵感模式, 2=定制模式
	MakeInstrumental string // 纯音乐: 0=否 1=是
	Title            string // 歌名
	Tags             string // 风格标签
	Model            string // 使用的模型(mv)
	AudioUrl         string // 生成的音频URL
	ImageUrl         string // 封面图URL
	VideoUrl         string // 生成的视频URL
	Duration         string // 时长
	Status           string // 状态: 0=进行中 1=成功, 2=失败
	FailReason       string // 失败原因
	RemoteTaskId     string // 远端任务ID
	RemoteClipId     string // 远端歌曲ID
	CreatedAt        string // 创建时间
	UpdatedAt        string // 修改时间
	DeletedAt        string // 删除时间
}

// openaiMusicTaskColumns holds the columns for the table hg_openai_music_task.
var openaiMusicTaskColumns = OpenaiMusicTaskColumns{
	Id:               "id",
	UserId:           "user_id",
	Prompt:           "prompt",
	Type:             "type",
	MakeInstrumental: "make_instrumental",
	Title:            "title",
	Tags:             "tags",
	Model:            "model",
	AudioUrl:         "audio_url",
	ImageUrl:         "image_url",
	VideoUrl:         "video_url",
	Duration:         "duration",
	Status:           "status",
	FailReason:       "fail_reason",
	RemoteTaskId:     "remote_task_id",
	RemoteClipId:     "remote_clip_id",
	CreatedAt:        "created_at",
	UpdatedAt:        "updated_at",
	DeletedAt:        "deleted_at",
}

// NewOpenaiMusicTaskDao creates and returns a new DAO object for table data access.
func NewOpenaiMusicTaskDao(handlers ...gdb.ModelHandler) *OpenaiMusicTaskDao {
	return &OpenaiMusicTaskDao{
		group:    "default",
		table:    "hg_openai_music_task",
		columns:  openaiMusicTaskColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *OpenaiMusicTaskDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *OpenaiMusicTaskDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *OpenaiMusicTaskDao) Columns() OpenaiMusicTaskColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *OpenaiMusicTaskDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *OpenaiMusicTaskDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *OpenaiMusicTaskDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
