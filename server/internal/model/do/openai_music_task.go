// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiMusicTask is the golang structure of table hg_openai_music_task for DAO operations like Where/Data.
type OpenaiMusicTask struct {
	g.Meta           `orm:"table:hg_openai_music_task, do:true"`
	Id               any         // 任务ID
	UserId           any         // 用户ID
	Prompt           any         // 提示词(灵感或歌词)
	Type             any         // 创作模式: 1=灵感模式, 2=定制模式
	MakeInstrumental any         // 纯音乐: 0=否 1=是
	Title            any         // 歌名
	Tags             any         // 风格标签
	Model            any         // 使用的模型(mv)
	AudioUrl         any         // 生成的音频URL
	ImageUrl         any         // 封面图URL
	VideoUrl         any         // 生成的视频URL
	Duration         any         // 时长
	Status           any         // 状态: 0=进行中 1=成功, 2=失败
	FailReason       any         // 失败原因
	RemoteTaskId     any         // 远端任务ID
	RemoteClipId     any         // 远端歌曲ID
	CreatedAt        *gtime.Time // 创建时间
	UpdatedAt        *gtime.Time // 修改时间
	DeletedAt        *gtime.Time // 删除时间
}
