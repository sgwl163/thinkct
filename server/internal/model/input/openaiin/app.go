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

// AppUpdateFields 修改AI应用字段过滤
type AppUpdateFields struct {
	CategoryId  int64  `json:"categoryId"  dc:"分类ID"`
	Title       string `json:"title"       dc:"应用名称"`
	Icon        string `json:"icon"        dc:"应用图标"`
	Description string `json:"description" dc:"应用描述"`
	Prompt      string `json:"prompt"      dc:"系统提示词"`
	UseCount    int    `json:"useCount"    dc:"使用次数"`
	Sort        int    `json:"sort"        dc:"排序"`
	Status      int    `json:"status"      dc:"状态(1:启用, 2:禁用)"`
}

// AppInsertFields 新增AI应用字段过滤
type AppInsertFields struct {
	CategoryId  int64  `json:"categoryId"  dc:"分类ID"`
	Title       string `json:"title"       dc:"应用名称"`
	Icon        string `json:"icon"        dc:"应用图标"`
	Description string `json:"description" dc:"应用描述"`
	Prompt      string `json:"prompt"      dc:"系统提示词"`
	UseCount    int    `json:"useCount"    dc:"使用次数"`
	Sort        int    `json:"sort"        dc:"排序"`
	Status      int    `json:"status"      dc:"状态(1:启用, 2:禁用)"`
}

// AppEditInp 修改/新增AI应用
type AppEditInp struct {
	entity.OpenaiApp
}

func (in *AppEditInp) Filter(ctx context.Context) (err error) {
	// 验证分类ID
	if err := g.Validator().Rules("required").Data(in.CategoryId).Messages("分类ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证应用名称
	if err := g.Validator().Rules("required").Data(in.Title).Messages("应用名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证系统提示词
	if err := g.Validator().Rules("required").Data(in.Prompt).Messages("系统提示词不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证使用次数
	if err := g.Validator().Rules("required").Data(in.UseCount).Messages("使用次数不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证排序
	if err := g.Validator().Rules("required").Data(in.Sort).Messages("排序不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态(1:启用, 2:禁用)
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态(1:启用, 2:禁用)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态(1:启用, 2:禁用)值不正确").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type AppEditModel struct{}

// AppDeleteInp 删除AI应用
type AppDeleteInp struct {
	Id interface{} `json:"id" v:"required#应用ID不能为空" dc:"应用ID"`
}

func (in *AppDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type AppDeleteModel struct{}

// AppViewInp 获取指定AI应用信息
type AppViewInp struct {
	Id int64 `json:"id" v:"required#应用ID不能为空" dc:"应用ID"`
}

func (in *AppViewInp) Filter(ctx context.Context) (err error) {
	return
}

type AppViewModel struct {
	entity.OpenaiApp
}

// AppListInp 获取AI应用列表
type AppListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"应用ID"`
	Status    int           `json:"status"    dc:"状态(1:启用, 2:禁用)"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *AppListInp) Filter(ctx context.Context) (err error) {
	return
}

type AppListModel struct {
	Id         int64       `json:"id"         dc:"应用ID"`
	CategoryId int64       `json:"categoryId" dc:"分类ID"`
	Title      string      `json:"title"      dc:"应用名称"`
	Icon       string      `json:"icon"       dc:"应用图标"`
	UseCount   int         `json:"useCount"   dc:"使用次数"`
	Sort       int         `json:"sort"       dc:"排序"`
	Status     int         `json:"status"     dc:"状态(1:启用, 2:禁用)"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

// AppExportModel 导出AI应用
type AppExportModel struct {
	Id         int64       `json:"id"         dc:"应用ID"`
	CategoryId int64       `json:"categoryId" dc:"分类ID"`
	Title      string      `json:"title"      dc:"应用名称"`
	Icon       string      `json:"icon"       dc:"应用图标"`
	UseCount   int         `json:"useCount"   dc:"使用次数"`
	Sort       int         `json:"sort"       dc:"排序"`
	Status     int         `json:"status"     dc:"状态(1:启用, 2:禁用)"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

// AppMaxSortInp 获取AI应用最大排序
type AppMaxSortInp struct{}

func (in *AppMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type AppMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// AppStatusInp 更新AI应用状态
type AppStatusInp struct {
	Id     int64 `json:"id" v:"required#应用ID不能为空" dc:"应用ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *AppStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("应用ID不能为空")
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

type AppStatusModel struct{}