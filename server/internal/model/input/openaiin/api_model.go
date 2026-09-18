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

// ApiModelUpdateFields 修改AI模型列表字段过滤
type ApiModelUpdateFields struct {
	GroupId  int64  `json:"groupId"  dc:"分组ID"`
	Name     string `json:"name"     dc:"模型别名(如GPT-4)"`
	Model    string `json:"model"    dc:"模型标识(如gpt-4-turbo)"`
	IsVision int    `json:"isVision" dc:"是否支持识图"`
	IsFile   int    `json:"isFile"   dc:"是否支持上传文件"`
	IsSearch int    `json:"isSearch" dc:"是否支持联网搜索"`
	Status   int    `json:"status"   dc:"状态"`
	Sort     int    `json:"sort"     dc:"排序"`
}

// ApiModelInsertFields 新增AI模型列表字段过滤
type ApiModelInsertFields struct {
	GroupId  int64  `json:"groupId"  dc:"分组ID"`
	Name     string `json:"name"     dc:"模型别名(如GPT-4)"`
	Model    string `json:"model"    dc:"模型标识(如gpt-4-turbo)"`
	IsVision int    `json:"isVision" dc:"是否支持识图"`
	IsFile   int    `json:"isFile"   dc:"是否支持上传文件"`
	IsSearch int    `json:"isSearch" dc:"是否支持联网搜索"`
	Status   int    `json:"status"   dc:"状态"`
	Sort     int    `json:"sort"     dc:"排序"`
}

// ApiModelEditInp 修改/新增AI模型列表
type ApiModelEditInp struct {
	entity.OpenaiApiModel
}

func (in *ApiModelEditInp) Filter(ctx context.Context) (err error) {
	// 验证分组ID
	if err := g.Validator().Rules("required").Data(in.GroupId).Messages("分组ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证模型别名(如GPT-4)
	if err := g.Validator().Rules("required").Data(in.Name).Messages("模型别名(如GPT-4)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证模型标识(如gpt-4-turbo)
	if err := g.Validator().Rules("required").Data(in.Model).Messages("模型标识(如gpt-4-turbo)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type ApiModelEditModel struct{}

// ApiModelDeleteInp 删除AI模型列表
type ApiModelDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *ApiModelDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ApiModelDeleteModel struct{}

// ApiModelViewInp 获取指定AI模型列表信息
type ApiModelViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *ApiModelViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ApiModelViewModel struct {
	entity.OpenaiApiModel
}

// ApiModelListInp 获取AI模型列表列表
type ApiModelListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"ID"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *ApiModelListInp) Filter(ctx context.Context) (err error) {
	return
}

type ApiModelListModel struct {
	Id        int64       `json:"id"        dc:"ID"`
	GroupId   int64       `json:"groupId"   dc:"分组ID"`
	Name      string      `json:"name"      dc:"模型别名(如GPT-4)"`
	Model     string      `json:"model"     dc:"模型标识(如gpt-4-turbo)"`
	IsVision  int         `json:"isVision"  dc:"是否支持识图"`
	IsFile    int         `json:"isFile"    dc:"是否支持上传文件"`
	IsSearch  int         `json:"isSearch"  dc:"是否支持联网搜索"`
	Status    int         `json:"status"    dc:"状态"`
	Sort      int         `json:"sort"      dc:"排序"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// ApiModelExportModel 导出AI模型列表
type ApiModelExportModel struct {
	Id        int64       `json:"id"        dc:"ID"`
	GroupId   int64       `json:"groupId"   dc:"分组ID"`
	Name      string      `json:"name"      dc:"模型别名(如GPT-4)"`
	Model     string      `json:"model"     dc:"模型标识(如gpt-4-turbo)"`
	IsVision  int         `json:"isVision"  dc:"是否支持识图"`
	IsFile    int         `json:"isFile"    dc:"是否支持上传文件"`
	IsSearch  int         `json:"isSearch"  dc:"是否支持联网搜索"`
	Status    int         `json:"status"    dc:"状态"`
	Sort      int         `json:"sort"      dc:"排序"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// ApiModelMaxSortInp 获取AI模型列表最大排序
type ApiModelMaxSortInp struct{}

func (in *ApiModelMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type ApiModelMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// ApiModelStatusInp 更新AI模型列表状态
type ApiModelStatusInp struct {
	Id     int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *ApiModelStatusInp) Filter(ctx context.Context) (err error) {
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

type ApiModelStatusModel struct{}