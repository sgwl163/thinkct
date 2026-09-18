// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiVideoTask is the golang structure of table hg_openai_video_task for DAO operations like Where/Data.
type OpenaiVideoTask struct {
	g.Meta         `orm:"table:hg_openai_video_task, do:true"`
	Id             any         // 任务ID
	UserId         any         // 用户ID
	Prompt         any         // 视频提示词
	ReferenceImage any         // 参考图/首尾帧URL
	LastFrame      any         // 尾帧图片URL
	VoiceDubUrl    any         // 自定义配音URL
	BgMusicUrl     any         // 背景音乐URL
	GenerateAudio  any         // 是否生成同步音频: 1=是 0=否
	Model          any         // 使用的模型
	VideoUrl       any         // 生成的视频URL
	Status         any         // 状态: 0=进行中 1=成功, 2=失败
	FailReason     any         // 失败原因
	RemoteTaskId   any         // 远端任务ID
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 修改时间
	DeletedAt      *gtime.Time // 删除时间
}
