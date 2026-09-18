// Package openaiin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package openaiin

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DrawTaskUpdateFields 修改AI绘画任务字段过滤
type DrawTaskUpdateFields struct {
	UserId         int64  `json:"userId"         dc:"用户ID"`
	Prompt         string `json:"prompt"         dc:"绘画提示词"`
	ReferenceImage string `json:"referenceImage" dc:"参考图URL"`
	Model          string `json:"model"          dc:"使用的模型"`
	Version        string `json:"version"        dc:"模型版本"`
	Ratio          string `json:"ratio"          dc:"图片比例"`
	Raw            int    `json:"raw"            dc:"是否开启RAW模式"`
	ImageUrl       string `json:"imageUrl"       dc:"生成的图片URL"`
	Status         int    `json:"status"         dc:"状态: 0=进行中, 1=成功, 2=失败"`
	FailReason     string `json:"failReason"     dc:"失败原因"`
	NegativePrompt string `json:"negativePrompt" dc:"负面提示词"`
	Stylize        int    `json:"stylize"        dc:"风格化程度(0-1000)"`
	Chaos          int    `json:"chaos"          dc:"多样性(0-100)"`
	Seed           int    `json:"seed"           dc:"随机种子"`
	MjTaskId       string `json:"mjTaskId"       dc:"MJ远端任务ID"`
	Buttons        string `json:"buttons"        dc:"MJ操作按钮JSON"`
	Progress       int    `json:"progress"       dc:"进度(0-100)"`
}

// DrawTaskInsertFields 新增AI绘画任务字段过滤
type DrawTaskInsertFields struct {
	UserId         int64  `json:"userId"         dc:"用户ID"`
	Prompt         string `json:"prompt"         dc:"绘画提示词"`
	ReferenceImage string `json:"referenceImage" dc:"参考图URL"`
	Model          string `json:"model"          dc:"使用的模型"`
	Version        string `json:"version"        dc:"模型版本"`
	Ratio          string `json:"ratio"          dc:"图片比例"`
	Raw            int    `json:"raw"            dc:"是否开启RAW模式"`
	ImageUrl       string `json:"imageUrl"       dc:"生成的图片URL"`
	Status         int    `json:"status"         dc:"状态: 0=进行中, 1=成功, 2=失败"`
	FailReason     string `json:"failReason"     dc:"失败原因"`
	NegativePrompt string `json:"negativePrompt" dc:"负面提示词"`
	Stylize        int    `json:"stylize"        dc:"风格化程度(0-1000)"`
	Chaos          int    `json:"chaos"          dc:"多样性(0-100)"`
	Seed           int    `json:"seed"           dc:"随机种子"`
	MjTaskId       string `json:"mjTaskId"       dc:"MJ远端任务ID"`
	Buttons        string `json:"buttons"        dc:"MJ操作按钮JSON"`
	Progress       int    `json:"progress"       dc:"进度(0-100)"`
}

// DrawTaskEditInp 修改/新增AI绘画任务
type DrawTaskEditInp struct {
	entity.OpenaiDrawTask
}

func (in *DrawTaskEditInp) Filter(ctx context.Context) (err error) {
	// 验证用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证绘画提示词
	if err := g.Validator().Rules("required").Data(in.Prompt).Messages("绘画提示词不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证使用的模型
	if err := g.Validator().Rules("required").Data(in.Model).Messages("使用的模型不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证进度(0-100)
	if err := g.Validator().Rules("required").Data(in.Progress).Messages("进度(0-100)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type DrawTaskEditModel struct{}

// DrawTaskDeleteInp 删除AI绘画任务
type DrawTaskDeleteInp struct {
	Id interface{} `json:"id" v:"required#任务ID不能为空" dc:"任务ID"`
}

func (in *DrawTaskDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type DrawTaskDeleteModel struct{}

// DrawTaskViewInp 获取指定AI绘画任务信息
type DrawTaskViewInp struct {
	Id int64 `json:"id" v:"required#任务ID不能为空" dc:"任务ID"`
}

func (in *DrawTaskViewInp) Filter(ctx context.Context) (err error) {
	return
}

type DrawTaskViewModel struct {
	entity.OpenaiDrawTask
}

// DrawTaskListInp 获取AI绘画任务列表
type DrawTaskListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"任务ID"`
	Status    int           `json:"status"    dc:"状态: 0=进行中, 1=成功, 2=失败"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *DrawTaskListInp) Filter(ctx context.Context) (err error) {
	return
}

type DrawTaskListModel struct {
	Id        int64       `json:"id"        dc:"任务ID"`
	UserId    int64       `json:"userId"    dc:"用户ID"`
	Model     string      `json:"model"     dc:"使用的模型"`
	Version   string      `json:"version"   dc:"模型版本"`
	Ratio     string      `json:"ratio"     dc:"图片比例"`
	Raw       int         `json:"raw"       dc:"是否开启RAW模式"`
	Status    int         `json:"status"    dc:"状态: 0=进行中, 1=成功, 2=失败"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
	Stylize   int         `json:"stylize"   dc:"风格化程度(0-1000)"`
	Chaos     int         `json:"chaos"     dc:"多样性(0-100)"`
	Seed      int         `json:"seed"      dc:"随机种子"`
	MjTaskId  string      `json:"mjTaskId"  dc:"MJ远端任务ID"`
	Progress  int         `json:"progress"  dc:"进度(0-100)"`
}

// DrawTaskExportModel 导出AI绘画任务
type DrawTaskExportModel struct {
	Id        int64       `json:"id"        dc:"任务ID"`
	UserId    int64       `json:"userId"    dc:"用户ID"`
	Model     string      `json:"model"     dc:"使用的模型"`
	Version   string      `json:"version"   dc:"模型版本"`
	Ratio     string      `json:"ratio"     dc:"图片比例"`
	Raw       int         `json:"raw"       dc:"是否开启RAW模式"`
	Status    int         `json:"status"    dc:"状态: 0=进行中, 1=成功, 2=失败"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
	Stylize   int         `json:"stylize"   dc:"风格化程度(0-1000)"`
	Chaos     int         `json:"chaos"     dc:"多样性(0-100)"`
	Seed      int         `json:"seed"      dc:"随机种子"`
	MjTaskId  string      `json:"mjTaskId"  dc:"MJ远端任务ID"`
	Progress  int         `json:"progress"  dc:"进度(0-100)"`
}

// DrawTaskStatusInp 更新AI绘画任务状态
type DrawTaskStatusInp struct {
	Id     int64 `json:"id" v:"required#任务ID不能为空" dc:"任务ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *DrawTaskStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("任务ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type DrawTaskStatusModel struct{}