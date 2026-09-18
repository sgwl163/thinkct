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

// VideoTaskUpdateFields 修改AI视频任务字段过滤
type VideoTaskUpdateFields struct {
	UserId         int64  `json:"userId"         dc:"用户ID"`
	Prompt         string `json:"prompt"         dc:"视频提示词"`
	ReferenceImage string `json:"referenceImage" dc:"参考图/首帧URL"`
	LastFrame      string `json:"lastFrame"      dc:"尾帧图片URL"`
	VoiceDubUrl    string `json:"voiceDubUrl"    dc:"自定义配音URL（音频）"`
	BgMusicUrl     string `json:"bgMusicUrl"     dc:"自定义背景音乐URL（音频）"`
	GenerateAudio  int    `json:"generateAudio"  dc:"是否生成同步音频: 1=是 0=否"`
	Model          string `json:"model"          dc:"使用的模型"`
	VideoUrl       string `json:"videoUrl"       dc:"生成的视频URL"`
	Status         int    `json:"status"         dc:"状态: 0=进行中 1=成功, 2=失败"`
	FailReason     string `json:"failReason"     dc:"失败原因"`
	RemoteTaskId   string `json:"remoteTaskId"   dc:"远端任务ID"`
}

// VideoTaskInsertFields 新增AI视频任务字段过滤
type VideoTaskInsertFields struct {
	UserId         int64  `json:"userId"         dc:"用户ID"`
	Prompt         string `json:"prompt"         dc:"视频提示词"`
	ReferenceImage string `json:"referenceImage" dc:"参考图/首帧URL"`
	LastFrame      string `json:"lastFrame"      dc:"尾帧图片URL"`
	VoiceDubUrl    string `json:"voiceDubUrl"    dc:"自定义配音URL（音频）"`
	BgMusicUrl     string `json:"bgMusicUrl"     dc:"自定义背景音乐URL（音频）"`
	GenerateAudio  int    `json:"generateAudio"  dc:"是否生成同步音频: 1=是 0=否"`
	Model          string `json:"model"          dc:"使用的模型"`
	VideoUrl       string `json:"videoUrl"       dc:"生成的视频URL"`
	Status         int    `json:"status"         dc:"状态: 0=进行中 1=成功, 2=失败"`
	FailReason     string `json:"failReason"     dc:"失败原因"`
	RemoteTaskId   string `json:"remoteTaskId"   dc:"远端任务ID"`
}

// VideoTaskEditInp 修改/新增AI视频任务
type VideoTaskEditInp struct {
	entity.OpenaiVideoTask
}

func (in *VideoTaskEditInp) Filter(ctx context.Context) (err error) {
	// 验证用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证视频提示词
	if err := g.Validator().Rules("required").Data(in.Prompt).Messages("视频提示词不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证使用的模型
	if err := g.Validator().Rules("required").Data(in.Model).Messages("使用的模型不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type VideoTaskEditModel struct{}

// VideoTaskDeleteInp 删除AI视频任务
type VideoTaskDeleteInp struct {
	Id interface{} `json:"id" v:"required#任务ID不能为空" dc:"任务ID"`
}

func (in *VideoTaskDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type VideoTaskDeleteModel struct{}

// VideoTaskViewInp 获取指定AI视频任务信息
type VideoTaskViewInp struct {
	Id int64 `json:"id" v:"required#任务ID不能为空" dc:"任务ID"`
}

func (in *VideoTaskViewInp) Filter(ctx context.Context) (err error) {
	return
}

type VideoTaskViewModel struct {
	entity.OpenaiVideoTask
}

// VideoTaskListInp 获取AI视频任务列表
type VideoTaskListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"任务ID"`
	Status    int           `json:"status"    dc:"状态: 0=进行中 1=成功, 2=失败"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *VideoTaskListInp) Filter(ctx context.Context) (err error) {
	return
}

type VideoTaskListModel struct {
	Id           int64       `json:"id"           dc:"任务ID"`
	UserId       int64       `json:"userId"       dc:"用户ID"`
	Model        string      `json:"model"        dc:"使用的模型"`
	Status       int         `json:"status"       dc:"状态: 0=进行中 1=成功, 2=失败"`
	RemoteTaskId string      `json:"remoteTaskId" dc:"远端任务ID"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"修改时间"`
}

// VideoTaskExportModel 导出AI视频任务
type VideoTaskExportModel struct {
	Id           int64       `json:"id"           dc:"任务ID"`
	UserId       int64       `json:"userId"       dc:"用户ID"`
	Model        string      `json:"model"        dc:"使用的模型"`
	Status       int         `json:"status"       dc:"状态: 0=进行中 1=成功, 2=失败"`
	RemoteTaskId string      `json:"remoteTaskId" dc:"远端任务ID"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"修改时间"`
}

// VideoTaskStatusInp 更新AI视频任务状态
type VideoTaskStatusInp struct {
	Id     int64 `json:"id" v:"required#任务ID不能为空" dc:"任务ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *VideoTaskStatusInp) Filter(ctx context.Context) (err error) {
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

type VideoTaskStatusModel struct{}