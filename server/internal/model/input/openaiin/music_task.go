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

// MusicTaskUpdateFields 修改AI音乐任务字段过滤
type MusicTaskUpdateFields struct {
	UserId           int64   `json:"userId"           dc:"用户ID"`
	Prompt           string  `json:"prompt"           dc:"提示词(灵感或歌词)"`
	Type             int     `json:"type"             dc:"创作模式: 1=灵感模式, 2=定制模式"`
	MakeInstrumental int     `json:"makeInstrumental" dc:"纯音乐: 0=否 1=是"`
	Title            string  `json:"title"            dc:"歌名"`
	Tags             string  `json:"tags"             dc:"风格标签"`
	Model            string  `json:"model"            dc:"使用的模型(mv)"`
	AudioUrl         string  `json:"audioUrl"         dc:"生成的音频URL"`
	ImageUrl         string  `json:"imageUrl"         dc:"封面图URL"`
	VideoUrl         string  `json:"videoUrl"         dc:"生成的视频URL"`
	Duration         float64 `json:"duration"         dc:"时长"`
	Status           int     `json:"status"           dc:"状态: 0=进行中 1=成功, 2=失败"`
	FailReason       string  `json:"failReason"       dc:"失败原因"`
	RemoteTaskId     string  `json:"remoteTaskId"     dc:"远端任务ID"`
	RemoteClipId     string  `json:"remoteClipId"     dc:"远端歌曲ID"`
}

// MusicTaskInsertFields 新增AI音乐任务字段过滤
type MusicTaskInsertFields struct {
	UserId           int64   `json:"userId"           dc:"用户ID"`
	Prompt           string  `json:"prompt"           dc:"提示词(灵感或歌词)"`
	Type             int     `json:"type"             dc:"创作模式: 1=灵感模式, 2=定制模式"`
	MakeInstrumental int     `json:"makeInstrumental" dc:"纯音乐: 0=否 1=是"`
	Title            string  `json:"title"            dc:"歌名"`
	Tags             string  `json:"tags"             dc:"风格标签"`
	Model            string  `json:"model"            dc:"使用的模型(mv)"`
	AudioUrl         string  `json:"audioUrl"         dc:"生成的音频URL"`
	ImageUrl         string  `json:"imageUrl"         dc:"封面图URL"`
	VideoUrl         string  `json:"videoUrl"         dc:"生成的视频URL"`
	Duration         float64 `json:"duration"         dc:"时长"`
	Status           int     `json:"status"           dc:"状态: 0=进行中 1=成功, 2=失败"`
	FailReason       string  `json:"failReason"       dc:"失败原因"`
	RemoteTaskId     string  `json:"remoteTaskId"     dc:"远端任务ID"`
	RemoteClipId     string  `json:"remoteClipId"     dc:"远端歌曲ID"`
}

// MusicTaskEditInp 修改/新增AI音乐任务
type MusicTaskEditInp struct {
	entity.OpenaiMusicTask
}

func (in *MusicTaskEditInp) Filter(ctx context.Context) (err error) {
	// 验证用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证提示词(灵感或歌词)
	if err := g.Validator().Rules("required").Data(in.Prompt).Messages("提示词(灵感或歌词)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证使用的模型(mv)
	if err := g.Validator().Rules("required").Data(in.Model).Messages("使用的模型(mv)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type MusicTaskEditModel struct{}

// MusicTaskDeleteInp 删除AI音乐任务
type MusicTaskDeleteInp struct {
	Id interface{} `json:"id" v:"required#任务ID不能为空" dc:"任务ID"`
}

func (in *MusicTaskDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type MusicTaskDeleteModel struct{}

// MusicTaskViewInp 获取指定AI音乐任务信息
type MusicTaskViewInp struct {
	Id int64 `json:"id" v:"required#任务ID不能为空" dc:"任务ID"`
}

func (in *MusicTaskViewInp) Filter(ctx context.Context) (err error) {
	return
}

type MusicTaskViewModel struct {
	entity.OpenaiMusicTask
}

// MusicTaskListInp 获取AI音乐任务列表
type MusicTaskListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"任务ID"`
	Status    int           `json:"status"    dc:"状态: 0=进行中 1=成功, 2=失败"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *MusicTaskListInp) Filter(ctx context.Context) (err error) {
	return
}

type MusicTaskListModel struct {
	Id               int64       `json:"id"               dc:"任务ID"`
	UserId           int64       `json:"userId"           dc:"用户ID"`
	Type             int         `json:"type"             dc:"创作模式: 1=灵感模式, 2=定制模式"`
	MakeInstrumental int         `json:"makeInstrumental" dc:"纯音乐: 0=否 1=是"`
	Title            string      `json:"title"            dc:"歌名"`
	Tags             string      `json:"tags"             dc:"风格标签"`
	Model            string      `json:"model"            dc:"使用的模型(mv)"`
	Duration         float64     `json:"duration"         dc:"时长"`
	Status           int         `json:"status"           dc:"状态: 0=进行中 1=成功, 2=失败"`
	RemoteTaskId     string      `json:"remoteTaskId"     dc:"远端任务ID"`
	RemoteClipId     string      `json:"remoteClipId"     dc:"远端歌曲ID"`
	CreatedAt        *gtime.Time `json:"createdAt"        dc:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        dc:"修改时间"`
}

// MusicTaskExportModel 导出AI音乐任务
type MusicTaskExportModel struct {
	Id               int64       `json:"id"               dc:"任务ID"`
	UserId           int64       `json:"userId"           dc:"用户ID"`
	Type             int         `json:"type"             dc:"创作模式: 1=灵感模式, 2=定制模式"`
	MakeInstrumental int         `json:"makeInstrumental" dc:"纯音乐: 0=否 1=是"`
	Title            string      `json:"title"            dc:"歌名"`
	Tags             string      `json:"tags"             dc:"风格标签"`
	Model            string      `json:"model"            dc:"使用的模型(mv)"`
	Duration         float64     `json:"duration"         dc:"时长"`
	Status           int         `json:"status"           dc:"状态: 0=进行中 1=成功, 2=失败"`
	RemoteTaskId     string      `json:"remoteTaskId"     dc:"远端任务ID"`
	RemoteClipId     string      `json:"remoteClipId"     dc:"远端歌曲ID"`
	CreatedAt        *gtime.Time `json:"createdAt"        dc:"创建时间"`
	UpdatedAt        *gtime.Time `json:"updatedAt"        dc:"修改时间"`
}

// MusicTaskStatusInp 更新AI音乐任务状态
type MusicTaskStatusInp struct {
	Id     int64 `json:"id" v:"required#任务ID不能为空" dc:"任务ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *MusicTaskStatusInp) Filter(ctx context.Context) (err error) {
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

type MusicTaskStatusModel struct{}