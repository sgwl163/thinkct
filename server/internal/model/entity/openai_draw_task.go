// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OpenaiDrawTask is the golang structure for table openai_draw_task.
type OpenaiDrawTask struct {
	Id             int64       `json:"id"             orm:"id"              description:"任务ID"`
	UserId         int64       `json:"userId"         orm:"user_id"         description:"用户ID"`
	Prompt         string      `json:"prompt"         orm:"prompt"          description:"绘画提示词"`
	ReferenceImage string      `json:"referenceImage" orm:"reference_image" description:"参考图URL"`
	Model          string      `json:"model"          orm:"model"           description:"使用的模型"`
	Version        string      `json:"version"        orm:"version"         description:"模型版本"`
	Ratio          string      `json:"ratio"          orm:"ratio"           description:"图片比例"`
	Raw            int         `json:"raw"            orm:"raw"             description:"是否开启RAW模式"`
	ImageUrl       string      `json:"imageUrl"       orm:"image_url"       description:"生成的图片URL"`
	Status         int         `json:"status"         orm:"status"          description:"状态: 0=进行中, 1=成功, 2=失败"`
	FailReason     string      `json:"failReason"     orm:"fail_reason"     description:"失败原因"`
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:"创建时间"`
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:"修改时间"`
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:"删除时间"`
	NegativePrompt string      `json:"negativePrompt" orm:"negative_prompt" description:"负面提示词"`
	Stylize        int         `json:"stylize"        orm:"stylize"         description:"风格化程度(0-1000)"`
	Chaos          int         `json:"chaos"          orm:"chaos"           description:"多样性(0-100)"`
	Seed           int         `json:"seed"           orm:"seed"            description:"随机种子"`
	MjTaskId       string      `json:"mjTaskId"       orm:"mj_task_id"      description:"MJ远端任务ID"`
	Buttons        string      `json:"buttons"        orm:"buttons"         description:"MJ操作按钮JSON"`
	Progress       int         `json:"progress"       orm:"progress"        description:"进度(0-100)"`
}
