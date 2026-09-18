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

// OpenaiAppCategoryUpdateFields 修改AI应用分类字段过滤
type OpenaiAppCategoryUpdateFields struct {
	Name   string `json:"name"   dc:"分类名称"`
	Sort   int    `json:"sort"   dc:"排序"`
	Status int    `json:"status" dc:"状态(1:启用, 2:禁用)"`
}

// OpenaiAppCategoryInsertFields 新增AI应用分类字段过滤
type OpenaiAppCategoryInsertFields struct {
	Name   string `json:"name"   dc:"分类名称"`
	Sort   int    `json:"sort"   dc:"排序"`
	Status int    `json:"status" dc:"状态(1:启用, 2:禁用)"`
}

// OpenaiAppCategoryEditInp 修改/新增AI应用分类
type OpenaiAppCategoryEditInp struct {
	entity.OpenaiAppCategory
}

func (in *OpenaiAppCategoryEditInp) Filter(ctx context.Context) (err error) {
	// 验证分类名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("分类名称不能为空").Run(ctx); err != nil {
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

type OpenaiAppCategoryEditModel struct{}

// OpenaiAppCategoryDeleteInp 删除AI应用分类
type OpenaiAppCategoryDeleteInp struct {
	Id interface{} `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *OpenaiAppCategoryDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type OpenaiAppCategoryDeleteModel struct{}

// OpenaiAppCategoryViewInp 获取指定AI应用分类信息
type OpenaiAppCategoryViewInp struct {
	Id int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
}

func (in *OpenaiAppCategoryViewInp) Filter(ctx context.Context) (err error) {
	return
}

type OpenaiAppCategoryViewModel struct {
	entity.OpenaiAppCategory
}

// OpenaiAppCategoryListInp 获取AI应用分类列表
type OpenaiAppCategoryListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"分类ID"`
	Status    int           `json:"status"    dc:"状态(1:启用, 2:禁用)"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *OpenaiAppCategoryListInp) Filter(ctx context.Context) (err error) {
	return
}

type OpenaiAppCategoryListModel struct {
	Id        int64       `json:"id"        dc:"分类ID"`
	Name      string      `json:"name"      dc:"分类名称"`
	Sort      int         `json:"sort"      dc:"排序"`
	Status    int         `json:"status"    dc:"状态(1:启用, 2:禁用)"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// OpenaiAppCategoryExportModel 导出AI应用分类
type OpenaiAppCategoryExportModel struct {
	Id        int64       `json:"id"        dc:"分类ID"`
	Name      string      `json:"name"      dc:"分类名称"`
	Sort      int         `json:"sort"      dc:"排序"`
	Status    int         `json:"status"    dc:"状态(1:启用, 2:禁用)"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间"`
}

// OpenaiAppCategoryMaxSortInp 获取AI应用分类最大排序
type OpenaiAppCategoryMaxSortInp struct{}

func (in *OpenaiAppCategoryMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type OpenaiAppCategoryMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// OpenaiAppCategoryStatusInp 更新AI应用分类状态
type OpenaiAppCategoryStatusInp struct {
	Id     int64 `json:"id" v:"required#分类ID不能为空" dc:"分类ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *OpenaiAppCategoryStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("分类ID不能为空")
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

type OpenaiAppCategoryStatusModel struct{}