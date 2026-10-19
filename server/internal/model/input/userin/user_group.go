// Package userin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package userin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserGroupUpdateFields 修改用户分组字段过滤
type UserGroupUpdateFields struct {
	Name   string `json:"name"   dc:"组名"`
	Rules  string `json:"rules"  dc:"权限节点"`
	Status string `json:"status" dc:"状态:0=禁用,1=启用"`
}

// UserGroupInsertFields 新增用户分组字段过滤
type UserGroupInsertFields struct {
	Name   string `json:"name"   dc:"组名"`
	Rules  string `json:"rules"  dc:"权限节点"`
	Status string `json:"status" dc:"状态:0=禁用,1=启用"`
}

// UserGroupEditInp 修改/新增用户分组
type UserGroupEditInp struct {
	entity.UserGroup
}

func (in *UserGroupEditInp) Filter(ctx context.Context) (err error) {
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

type UserGroupEditModel struct{}

// UserGroupDeleteInp 删除用户分组
type UserGroupDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *UserGroupDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type UserGroupDeleteModel struct{}

// UserGroupViewInp 获取指定用户分组信息
type UserGroupViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *UserGroupViewInp) Filter(ctx context.Context) (err error) {
	return
}

type UserGroupViewModel struct {
	entity.UserGroup
}

// UserGroupListInp 获取用户分组列表
type UserGroupListInp struct {
	form.PageReq
	Id        int           `json:"id"        dc:"ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *UserGroupListInp) Filter(ctx context.Context) (err error) {
	return
}

type UserGroupListModel struct {
	Id        int         `json:"id"        dc:"ID"`
	Name      string      `json:"name"      dc:"组名"`
	Status    string      `json:"status"    dc:"状态:0=禁用,1=启用"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}

// UserGroupExportModel 导出用户分组
type UserGroupExportModel struct {
	Id        int         `json:"id"        dc:"ID"`
	Name      string      `json:"name"      dc:"组名"`
	Status    string      `json:"status"    dc:"状态:0=禁用,1=启用"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
}