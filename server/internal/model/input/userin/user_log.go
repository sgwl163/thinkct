// Package userin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package userin

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

// UserLogUpdateFields 修改操作记录字段过滤
type UserLogUpdateFields struct {
	UserId     int64       `json:"userId"     dc:"用户ID"`
	Username   string      `json:"username"   dc:"用户名"`
	LogType    string      `json:"logType"    dc:"记录类型:login=登录,operation=操作"`
	Action     string      `json:"action"     dc:"操作标识"`
	ActionName string      `json:"actionName" dc:"操作名称"`
	EventAt    *gtime.Time `json:"eventAt"    dc:"事件时间"`
	Ip         string      `json:"ip"         dc:"IP地址"`
	UserAgent  string      `json:"userAgent"  dc:"UA信息"`
	Message    string      `json:"message"    dc:"结果消息"`
	Extra      string      `json:"extra"      dc:"扩展数据"`
	Status     int         `json:"status"     dc:"状态"`
}

// UserLogInsertFields 新增操作记录字段过滤
type UserLogInsertFields struct {
	UserId     int64       `json:"userId"     dc:"用户ID"`
	Username   string      `json:"username"   dc:"用户名"`
	LogType    string      `json:"logType"    dc:"记录类型:login=登录,operation=操作"`
	Action     string      `json:"action"     dc:"操作标识"`
	ActionName string      `json:"actionName" dc:"操作名称"`
	EventAt    *gtime.Time `json:"eventAt"    dc:"事件时间"`
	Ip         string      `json:"ip"         dc:"IP地址"`
	UserAgent  string      `json:"userAgent"  dc:"UA信息"`
	Message    string      `json:"message"    dc:"结果消息"`
	Extra      string      `json:"extra"      dc:"扩展数据"`
	Status     int         `json:"status"     dc:"状态"`
}

// UserLogEditInp 修改/新增操作记录
type UserLogEditInp struct {
	entity.UserLog
}

func (in *UserLogEditInp) Filter(ctx context.Context) (err error) {
	// 验证用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证用户名
	if err := g.Validator().Rules("required").Data(in.Username).Messages("用户名不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证记录类型:login=登录,operation=操作
	if err := g.Validator().Rules("required").Data(in.LogType).Messages("记录类型:login=登录,operation=操作不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证操作标识
	if err := g.Validator().Rules("required").Data(in.Action).Messages("操作标识不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证操作名称
	if err := g.Validator().Rules("required").Data(in.ActionName).Messages("操作名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证IP地址
	if err := g.Validator().Rules("required").Data(in.Ip).Messages("IP地址不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证UA信息
	if err := g.Validator().Rules("required").Data(in.UserAgent).Messages("UA信息不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证结果消息
	if err := g.Validator().Rules("required").Data(in.Message).Messages("结果消息不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态值不正确").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type UserLogEditModel struct{}

// UserLogDeleteInp 删除操作记录
type UserLogDeleteInp struct {
	Id interface{} `json:"id" v:"required#日志ID不能为空" dc:"日志ID"`
}

func (in *UserLogDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type UserLogDeleteModel struct{}

// UserLogViewInp 获取指定操作记录信息
type UserLogViewInp struct {
	Id int64 `json:"id" v:"required#日志ID不能为空" dc:"日志ID"`
}

func (in *UserLogViewInp) Filter(ctx context.Context) (err error) {
	return
}

type UserLogViewModel struct {
	entity.UserLog
}

// UserLogListInp 获取操作记录列表
type UserLogListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"日志ID"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *UserLogListInp) Filter(ctx context.Context) (err error) {
	return
}

type UserLogListModel struct {
	Id         int64       `json:"id"         dc:"日志ID"`
	UserId     int64       `json:"userId"     dc:"用户ID"`
	Username   string      `json:"username"   dc:"用户名"`
	LogType    string      `json:"logType"    dc:"记录类型:login=登录,operation=操作"`
	Action     string      `json:"action"     dc:"操作标识"`
	ActionName string      `json:"actionName" dc:"操作名称"`
	EventAt    *gtime.Time `json:"eventAt"    dc:"事件时间"`
	Ip         string      `json:"ip"         dc:"IP地址"`
	Status     int         `json:"status"     dc:"状态"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

// UserLogExportModel 导出操作记录
type UserLogExportModel struct {
	Id         int64       `json:"id"         dc:"日志ID"`
	UserId     int64       `json:"userId"     dc:"用户ID"`
	Username   string      `json:"username"   dc:"用户名"`
	LogType    string      `json:"logType"    dc:"记录类型:login=登录,operation=操作"`
	Action     string      `json:"action"     dc:"操作标识"`
	ActionName string      `json:"actionName" dc:"操作名称"`
	EventAt    *gtime.Time `json:"eventAt"    dc:"事件时间"`
	Ip         string      `json:"ip"         dc:"IP地址"`
	Status     int         `json:"status"     dc:"状态"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

// UserLogStatusInp 更新操作记录状态
type UserLogStatusInp struct {
	Id     int64 `json:"id" v:"required#日志ID不能为空" dc:"日志ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *UserLogStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("日志ID不能为空")
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

type UserLogStatusModel struct{}