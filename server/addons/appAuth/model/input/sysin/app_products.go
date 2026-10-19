// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package sysin

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

// AppProductsUpdateFields 修改产品管理字段过滤
type AppProductsUpdateFields struct {
	Cover       string  `json:"cover"       dc:"封面"`
	Title       string  `json:"title"       dc:"标题"`
	Keywords    string  `json:"keywords"    dc:"关键词"`
	Description string  `json:"description" dc:"描述"`
	Name        string  `json:"name"        dc:"名称"`
	Alias       string  `json:"alias"       dc:"别名"`
	Content     string  `json:"content"     dc:"内容"`
	Price       float64 `json:"price"       dc:"价格"`
	Ysweb       string  `json:"ysweb"       dc:"演示站"`
	Unique      string  `json:"unique"      dc:"特点"`
	Weigh       int     `json:"weigh"       dc:"权重"`
	Status      int     `json:"status"      dc:"状态:0=禁用,1=启用"`
}

// AppProductsInsertFields 新增产品管理字段过滤
type AppProductsInsertFields struct {
	Cover       string  `json:"cover"       dc:"封面"`
	Title       string  `json:"title"       dc:"标题"`
	Keywords    string  `json:"keywords"    dc:"关键词"`
	Description string  `json:"description" dc:"描述"`
	Name        string  `json:"name"        dc:"名称"`
	Alias       string  `json:"alias"       dc:"别名"`
	Content     string  `json:"content"     dc:"内容"`
	Price       float64 `json:"price"       dc:"价格"`
	Ysweb       string  `json:"ysweb"       dc:"演示站"`
	Unique      string  `json:"unique"      dc:"特点"`
	Weigh       int     `json:"weigh"       dc:"权重"`
	Status      int     `json:"status"      dc:"状态:0=禁用,1=启用"`
}

// AppProductsEditInp 修改/新增产品管理
type AppProductsEditInp struct {
	entity.AppProducts
}

func (in *AppProductsEditInp) Filter(ctx context.Context) (err error) {
	// 验证封面
	if err := g.Validator().Rules("required").Data(in.Cover).Messages("封面不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证价格
	if err := g.Validator().Rules("regex:(^[0-9]{1,10}$)|(^[0-9]{1,10}[\\.]{1}[0-9]{1,2}$)").Data(in.Price).Messages("价格最多允许输入10位整数及2位小数").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证特点
	if err := g.Validator().Rules("required").Data(in.Unique).Messages("特点不能为空").Run(ctx); err != nil {
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

type AppProductsEditModel struct{}

// AppProductsDeleteInp 删除产品管理
type AppProductsDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *AppProductsDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type AppProductsDeleteModel struct{}

// AppProductsViewInp 获取指定产品管理信息
type AppProductsViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *AppProductsViewInp) Filter(ctx context.Context) (err error) {
	return
}

type AppProductsViewModel struct {
	entity.AppProducts
}

// AppProductsListInp 获取产品管理列表
type AppProductsListInp struct {
	form.PageReq
	Id        int           `json:"id"        dc:"ID"`
	Status    int           `json:"status"    dc:"状态:0=禁用,1=启用"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *AppProductsListInp) Filter(ctx context.Context) (err error) {
	return
}

type AppProductsListModel struct {
	Id          int         `json:"id"          dc:"ID"`
	Cover       string      `json:"cover"       dc:"封面"`
	Title       string      `json:"title"       dc:"标题"`
	Keywords    string      `json:"keywords"    dc:"关键词"`
	Description string      `json:"description" dc:"描述"`
	Name        string      `json:"name"        dc:"名称"`
	Alias       string      `json:"alias"       dc:"别名"`
	Price       float64     `json:"price"       dc:"价格"`
	Ysweb       string      `json:"ysweb"       dc:"演示站"`
	Unique      string      `json:"unique"      dc:"特点"`
	Weigh       int         `json:"weigh"       dc:"权重"`
	Status      int         `json:"status"      dc:"状态:0=禁用,1=启用"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
}

// AppProductsExportModel 导出产品管理
type AppProductsExportModel struct {
	Id          int         `json:"id"          dc:"ID"`
	Cover       string      `json:"cover"       dc:"封面"`
	Title       string      `json:"title"       dc:"标题"`
	Keywords    string      `json:"keywords"    dc:"关键词"`
	Description string      `json:"description" dc:"描述"`
	Name        string      `json:"name"        dc:"名称"`
	Alias       string      `json:"alias"       dc:"别名"`
	Price       float64     `json:"price"       dc:"价格"`
	Ysweb       string      `json:"ysweb"       dc:"演示站"`
	Unique      string      `json:"unique"      dc:"特点"`
	Weigh       int         `json:"weigh"       dc:"权重"`
	Status      int         `json:"status"      dc:"状态:0=禁用,1=启用"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"修改时间"`
}

// AppProductsStatusInp 更新产品管理状态
type AppProductsStatusInp struct {
	Id     int `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int `json:"status" dc:"状态"`
}

func (in *AppProductsStatusInp) Filter(ctx context.Context) (err error) {
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

type AppProductsStatusModel struct{}