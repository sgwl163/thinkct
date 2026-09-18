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

// ApiGroupUpdateFields 修改AI接口分组字段过滤
type ApiGroupUpdateFields struct {
	Name    string `json:"name"    dc:"组名称(如OpenAI、文心一言)"`
	BaseUrl string `json:"baseUrl" dc:"接口地址"`
	ApiKey  string `json:"apiKey"  dc:"API Key"`
	Status  int    `json:"status"  dc:"状态"`
	Sort    int    `json:"sort"    dc:"排序"`
}

// ApiGroupInsertFields 新增AI接口分组字段过滤
type ApiGroupInsertFields struct {
	Name    string `json:"name"    dc:"组名称(如OpenAI、文心一言)"`
	BaseUrl string `json:"baseUrl" dc:"接口地址"`
	ApiKey  string `json:"apiKey"  dc:"API Key"`
	Status  int    `json:"status"  dc:"状态"`
	Sort    int    `json:"sort"    dc:"排序"`
}

// ApiGroupEditInp 修改/新增AI接口分组
type ApiGroupEditInp struct {
	entity.OpenaiApiGroup
}

func (in *ApiGroupEditInp) Filter(ctx context.Context) (err error) {
	// 验证组名称(如OpenAI、文心一言)
	if err := g.Validator().Rules("required").Data(in.Name).Messages("组名称(如OpenAI、文心一言)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type ApiGroupEditModel struct{}

// ApiGroupDeleteInp 删除AI接口分组
type ApiGroupDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *ApiGroupDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ApiGroupDeleteModel struct{}

// ApiGroupViewInp 获取指定AI接口分组信息
type ApiGroupViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *ApiGroupViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ApiGroupViewModel struct {
	entity.OpenaiApiGroup
}

// ApiGroupListInp 获取AI接口分组列表
type ApiGroupListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"ID"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *ApiGroupListInp) Filter(ctx context.Context) (err error) {
	return
}

type ApiGroupListModel struct {
	Id        int64       `json:"id"        dc:"ID"`
	Name      string      `json:"name"      dc:"组名称(如OpenAI、文心一言)"`
	BaseUrl   string      `json:"baseUrl"   dc:"接口地址"`
	ApiKey    string      `json:"apiKey"    dc:"API Key"`
	Status    int         `json:"status"    dc:"状态"`
	Sort      int         `json:"sort"      dc:"排序"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// ApiGroupExportModel 导出AI接口分组
type ApiGroupExportModel struct {
	Id        int64       `json:"id"        dc:"ID"`
	Name      string      `json:"name"      dc:"组名称(如OpenAI、文心一言)"`
	BaseUrl   string      `json:"baseUrl"   dc:"接口地址"`
	ApiKey    string      `json:"apiKey"    dc:"API Key"`
	Status    int         `json:"status"    dc:"状态"`
	Sort      int         `json:"sort"      dc:"排序"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// ApiGroupMaxSortInp 获取AI接口分组最大排序
type ApiGroupMaxSortInp struct{}

func (in *ApiGroupMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type ApiGroupMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// ApiGroupStatusInp 更新AI接口分组状态
type ApiGroupStatusInp struct {
	Id     int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *ApiGroupStatusInp) Filter(ctx context.Context) (err error) {
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

type ApiGroupStatusModel struct{}