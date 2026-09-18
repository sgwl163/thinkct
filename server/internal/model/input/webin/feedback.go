// Package webin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package webin

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// FeedbackUpdateFields 修改留言管理字段过滤
type FeedbackUpdateFields struct {
	Name       string `json:"name"       dc:"名称"`
	Email      string `json:"email"      dc:"邮箱"`
	Phone      string `json:"phone"      dc:"手机号"`
	Company    string `json:"company"    dc:"公司名称"`
	Content    string `json:"content"    dc:"富文本"`
	Answer     string `json:"answer"     dc:"备注"`
	Ip         string `json:"ip"         dc:"IP"`
	Status     int    `json:"status"     dc:"状态:0=禁用,1=启用"`
	CreateTime int64  `json:"createTime" dc:"创建时间"`
	UpdateTime int64  `json:"updateTime" dc:"修改时间"`
}

// FeedbackInsertFields 新增留言管理字段过滤
type FeedbackInsertFields struct {
	Name       string `json:"name"       dc:"名称"`
	Email      string `json:"email"      dc:"邮箱"`
	Phone      string `json:"phone"      dc:"手机号"`
	Company    string `json:"company"    dc:"公司名称"`
	Content    string `json:"content"    dc:"富文本"`
	Answer     string `json:"answer"     dc:"备注"`
	Ip         string `json:"ip"         dc:"IP"`
	Status     int    `json:"status"     dc:"状态:0=禁用,1=启用"`
	CreateTime int64  `json:"createTime" dc:"创建时间"`
	UpdateTime int64  `json:"updateTime" dc:"修改时间"`
}

// FeedbackEditInp 修改/新增留言管理
type FeedbackEditInp struct {
	entity.WebFeedback
}

func (in *FeedbackEditInp) Filter(ctx context.Context) (err error) {
	// 验证邮箱
	if err := g.Validator().Rules("email").Data(in.Email).Messages("邮箱不是邮箱地址").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证备注
	if err := g.Validator().Rules("required").Data(in.Answer).Messages("备注不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态:0=禁用,1=启用
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态:0=禁用,1=启用不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态:0=禁用,1=启用值不正确").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type FeedbackEditModel struct{}

// FeedbackDeleteInp 删除留言管理
type FeedbackDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *FeedbackDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type FeedbackDeleteModel struct{}

// FeedbackViewInp 获取指定留言管理信息
type FeedbackViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *FeedbackViewInp) Filter(ctx context.Context) (err error) {
	return
}

type FeedbackViewModel struct {
	entity.WebFeedback
}

// FeedbackListInp 获取留言管理列表
type FeedbackListInp struct {
	form.PageReq
	Id     int `json:"id"     dc:"ID"`
	Status int `json:"status" dc:"状态:0=禁用,1=启用"`
}

func (in *FeedbackListInp) Filter(ctx context.Context) (err error) {
	return
}

type FeedbackListModel struct {
	Id         int    `json:"id"         dc:"ID"`
	Name       string `json:"name"       dc:"名称"`
	Email      string `json:"email"      dc:"邮箱"`
	Phone      string `json:"phone"      dc:"手机号"`
	Company    string `json:"company"    dc:"公司名称"`
	Answer     string `json:"answer"     dc:"备注"`
	Ip         string `json:"ip"         dc:"IP"`
	Status     int    `json:"status"     dc:"状态:0=禁用,1=启用"`
	CreateTime int64  `json:"createTime" dc:"创建时间"`
	UpdateTime int64  `json:"updateTime" dc:"修改时间"`
}

// FeedbackExportModel 导出留言管理
type FeedbackExportModel struct {
	Id         int    `json:"id"         dc:"ID"`
	Name       string `json:"name"       dc:"名称"`
	Email      string `json:"email"      dc:"邮箱"`
	Phone      string `json:"phone"      dc:"手机号"`
	Company    string `json:"company"    dc:"公司名称"`
	Answer     string `json:"answer"     dc:"备注"`
	Ip         string `json:"ip"         dc:"IP"`
	Status     int    `json:"status"     dc:"状态:0=禁用,1=启用"`
	CreateTime int64  `json:"createTime" dc:"创建时间"`
	UpdateTime int64  `json:"updateTime" dc:"修改时间"`
}

// FeedbackStatusInp 更新留言管理状态
type FeedbackStatusInp struct {
	Id     int `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int `json:"status" dc:"状态"`
}

func (in *FeedbackStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
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

type FeedbackStatusModel struct{}