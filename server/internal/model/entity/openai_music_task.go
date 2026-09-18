// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiMusicTask is the golang structure for table openai_music_task.
type OpenaiMusicTask struct {
	Id               int64       `json:"id"               orm:"id"                description:"任务ID"`
	UserId           int64       `json:"userId"           orm:"user_id"           description:"用户ID"`
	Prompt           string      `json:"prompt"           orm:"prompt"            description:"提示词(灵感或歌词)"`
	Type             int         `json:"type"             orm:"type"              description:"创作模式: 1=灵感模式, 2=定制模式"`
	MakeInstrumental int         `json:"makeInstrumental" orm:"make_instrumental" description:"纯音乐: 0=否 1=是"`
	Title            string      `json:"title"            orm:"title"             description:"歌名"`
	Tags             string      `json:"tags"             orm:"tags"              description:"风格标签"`
	Model            string      `json:"model"            orm:"model"             description:"使用的模型(mv)"`
	AudioUrl         string      `json:"audioUrl"         orm:"audio_url"         description:"生成的音频URL"`
	ImageUrl         string      `json:"imageUrl"         orm:"image_url"         description:"封面图URL"`
	VideoUrl         string      `json:"videoUrl"         orm:"video_url"         description:"生成的视频URL"`
	Duration         float64     `json:"duration"         orm:"duration"          description:"时长"`
	Status           int         `json:"status"           orm:"status"            description:"状态: 0=进行中 1=成功, 2=失败"`
	FailReason       string      `json:"failReason"       orm:"fail_reason"       description:"失败原因"`
	RemoteTaskId     string      `json:"remoteTaskId"     orm:"remote_task_id"    description:"远端任务ID"`
	RemoteClipId     string      `json:"remoteClipId"     orm:"remote_clip_id"    description:"远端歌曲ID"`
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"        description:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"        description:"修改时间"`
	DeletedAt        *gtime.Time `json:"deletedAt"        orm:"deleted_at"        description:"删除时间"`
}
