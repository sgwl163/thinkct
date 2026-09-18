// Package imin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package imin

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

// GroupUpdateFields 修改IM-群组表字段过滤
type GroupUpdateFields struct {
	Name    string `json:"name"    dc:"群名称"`
	Avatar  string `json:"avatar"  dc:"群头像"`
	OwnerId int64  `json:"ownerId" dc:"群主ID"`
	Notice  string `json:"notice"  dc:"群公告"`
	Status  int    `json:"status"  dc:"状态: 1=正常, 2=已解散"`
}

// GroupInsertFields 新增IM-群组表字段过滤
type GroupInsertFields struct {
	Name    string `json:"name"    dc:"群名称"`
	Avatar  string `json:"avatar"  dc:"群头像"`
	OwnerId int64  `json:"ownerId" dc:"群主ID"`
	Notice  string `json:"notice"  dc:"群公告"`
	Status  int    `json:"status"  dc:"状态: 1=正常, 2=已解散"`
}

// GroupEditInp 修改/新增IM-群组表
type GroupEditInp struct {
	entity.ImGroup
}

func (in *GroupEditInp) Filter(ctx context.Context) (err error) {
	// 验证群名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("群名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证群主ID
	if err := g.Validator().Rules("required").Data(in.OwnerId).Messages("群主ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态: 1=正常, 2=已解散
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态: 1=正常, 2=已解散不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态: 1=正常, 2=已解散值不正确").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type GroupEditModel struct{}

// GroupDeleteInp 删除IM-群组表
type GroupDeleteInp struct {
	Id interface{} `json:"id" v:"required#群组ID不能为空" dc:"群组ID"`
}

func (in *GroupDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type GroupDeleteModel struct{}

// GroupViewInp 获取指定IM-群组表信息
type GroupViewInp struct {
	Id int64 `json:"id" v:"required#群组ID不能为空" dc:"群组ID"`
}

func (in *GroupViewInp) Filter(ctx context.Context) (err error) {
	return
}

type GroupViewModel struct {
	entity.ImGroup
}

// GroupListInp 获取IM-群组表列表
type GroupListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"群组ID"`
	Status    int           `json:"status"    dc:"状态: 1=正常, 2=已解散"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *GroupListInp) Filter(ctx context.Context) (err error) {
	return
}

type GroupListModel struct {
	Id        int64       `json:"id"        dc:"群组ID"`
	Name      string      `json:"name"      dc:"群名称"`
	Avatar    string      `json:"avatar"    dc:"群头像"`
	OwnerId   int64       `json:"ownerId"   dc:"群主ID"`
	Status    int         `json:"status"    dc:"状态: 1=正常, 2=已解散"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// GroupExportModel 导出IM-群组表
type GroupExportModel struct {
	Id        int64       `json:"id"        dc:"群组ID"`
	Name      string      `json:"name"      dc:"群名称"`
	Avatar    string      `json:"avatar"    dc:"群头像"`
	OwnerId   int64       `json:"ownerId"   dc:"群主ID"`
	Status    int         `json:"status"    dc:"状态: 1=正常, 2=已解散"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// GroupStatusInp 更新IM-群组表状态
type GroupStatusInp struct {
	Id     int64 `json:"id" v:"required#群组ID不能为空" dc:"群组ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *GroupStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("群组ID不能为空")
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