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
)

// GroupUpdateFields 修改分组管理字段过滤
type GroupUpdateFields struct {
	Name       string `json:"name"       dc:"组名"`
	Rules      string `json:"rules"      dc:"权限节点"`
	Status     string `json:"status"     dc:"状态:0=禁用,1=启用"`
	UpdateTime int64  `json:"updateTime" dc:"更新时间"`
	CreateTime int64  `json:"createTime" dc:"创建时间"`
}

// GroupInsertFields 新增分组管理字段过滤
type GroupInsertFields struct {
	Name       string `json:"name"       dc:"组名"`
	Rules      string `json:"rules"      dc:"权限节点"`
	Status     string `json:"status"     dc:"状态:0=禁用,1=启用"`
	UpdateTime int64  `json:"updateTime" dc:"更新时间"`
	CreateTime int64  `json:"createTime" dc:"创建时间"`
}

// GroupEditInp 修改/新增分组管理
type GroupEditInp struct {
	entity.UserGroup
}

func (in *GroupEditInp) Filter(ctx context.Context) (err error) {
	// 验证组名
	if err := g.Validator().Rules("required").Data(in.Name).Messages("组名不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态:0=禁用,1=启用
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态:0=禁用,1=启用不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type GroupEditModel struct{}

// GroupDeleteInp 删除分组管理
type GroupDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *GroupDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type GroupDeleteModel struct{}

// GroupViewInp 获取指定分组管理信息
type GroupViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *GroupViewInp) Filter(ctx context.Context) (err error) {
	return
}

type GroupViewModel struct {
	entity.UserGroup
}

// GroupListInp 获取分组管理列表
type GroupListInp struct {
	form.PageReq
	Id int `json:"id" dc:"ID"`
}

func (in *GroupListInp) Filter(ctx context.Context) (err error) {
	return
}

type GroupListModel struct {
	Id         int    `json:"id"         dc:"ID"`
	Name       string `json:"name"       dc:"组名"`
	Status     string `json:"status"     dc:"状态:0=禁用,1=启用"`
	UpdateTime int64  `json:"updateTime" dc:"更新时间"`
	CreateTime int64  `json:"createTime" dc:"创建时间"`
}

// GroupExportModel 导出分组管理
type GroupExportModel struct {
	Id         int    `json:"id"         dc:"ID"`
	Name       string `json:"name"       dc:"组名"`
	Status     string `json:"status"     dc:"状态:0=禁用,1=启用"`
	UpdateTime int64  `json:"updateTime" dc:"更新时间"`
	CreateTime int64  `json:"createTime" dc:"创建时间"`
}

// GroupStatusInp 更新分组管理状态
type GroupStatusInp struct {
	Id     int `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int `json:"status" dc:"状态"`
}

func (in *GroupStatusInp) Filter(ctx context.Context) (err error) {
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

type GroupStatusModel struct{}