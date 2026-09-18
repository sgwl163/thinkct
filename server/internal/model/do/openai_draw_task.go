// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiDrawTask is the golang structure of table hg_openai_draw_task for DAO operations like Where/Data.
type OpenaiDrawTask struct {
	g.Meta         `orm:"table:hg_openai_draw_task, do:true"`
	Id             any         // 任务ID
	UserId         any         // 用户ID
	Prompt         any         // 绘画提示词
	ReferenceImage any         // 参考图URL
	Model          any         // 使用的模型
	Version        any         // 模型版本
	Ratio          any         // 图片比例
	Raw            any         // 是否开启RAW模式
	ImageUrl       any         // 生成的图片URL
	Status         any         // 状态: 0=进行中, 1=成功, 2=失败
	FailReason     any         // 失败原因
	CreatedAt      *gtime.Time // 创建时间
	UpdatedAt      *gtime.Time // 修改时间
	DeletedAt      *gtime.Time // 删除时间
	NegativePrompt any         // 负面提示词
	Stylize        any         // 风格化程度(0-1000)
	Chaos          any         // 多样性(0-100)
	Seed           any         // 随机种子
	MjTaskId       any         // MJ远端任务ID
	Buttons        any         // MJ操作按钮JSON
	Progress       any         // 进度(0-100)
}
