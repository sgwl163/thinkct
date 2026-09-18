// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiVideoTask is the golang structure for table openai_video_task.
type OpenaiVideoTask struct {
	Id             int64       `json:"id"             orm:"id"              description:"任务ID"`
	UserId         int64       `json:"userId"         orm:"user_id"         description:"用户ID"`
	Prompt         string      `json:"prompt"         orm:"prompt"          description:"视频提示词"`
	ReferenceImage string      `json:"referenceImage" orm:"reference_image" description:"参考图/首尾帧URL"`
	LastFrame      string      `json:"lastFrame"      orm:"last_frame"      description:"尾帧图片URL"`
	VoiceDubUrl    string      `json:"voiceDubUrl"    orm:"voice_dub_url"   description:"自定义配音URL"`
	BgMusicUrl     string      `json:"bgMusicUrl"     orm:"bg_music_url"    description:"背景音乐URL"`
	GenerateAudio  int         `json:"generateAudio"  orm:"generate_audio"  description:"是否生成同步音频: 1=是 0=否"`
	Model          string      `json:"model"          orm:"model"           description:"使用的模型"`
	VideoUrl       string      `json:"videoUrl"       orm:"video_url"       description:"生成的视频URL"`
	Status         int         `json:"status"         orm:"status"          description:"状态: 0=进行中 1=成功, 2=失败"`
	FailReason     string      `json:"failReason"     orm:"fail_reason"     description:"失败原因"`
	RemoteTaskId   string      `json:"remoteTaskId"   orm:"remote_task_id"  description:"远端任务ID"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"修改时间"`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:"删除时间"`
}
