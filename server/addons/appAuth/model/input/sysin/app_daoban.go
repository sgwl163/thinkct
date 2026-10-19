// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package sysin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppDaobanUpdateFields 修改盗版管理字段过滤
type AppDaobanUpdateFields struct {
	ProductId int    `json:"productId" dc:"产品ID"`
	Host      string `json:"host"      dc:"域名"`
	Ip        string `json:"ip"        dc:"IP"`
	Version   string `json:"version"   dc:"版本号"`
	Code      string `json:"code"      dc:"执行代码"`
	Data      string `json:"data"      dc:"数据"`
	Views     int    `json:"views"     dc:"浏览量"`
}

// AppDaobanInsertFields 新增盗版管理字段过滤
type AppDaobanInsertFields struct {
	ProductId int    `json:"productId" dc:"产品ID"`
	Host      string `json:"host"      dc:"域名"`
	Ip        string `json:"ip"        dc:"IP"`
	Version   string `json:"version"   dc:"版本号"`
	Code      string `json:"code"      dc:"执行代码"`
	Data      string `json:"data"      dc:"数据"`
	Views     int    `json:"views"     dc:"浏览量"`
}

// AppDaobanEditInp 修改/新增盗版管理
type AppDaobanEditInp struct {
	entity.AppDaoban
}

func (in *AppDaobanEditInp) Filter(ctx context.Context) (err error) {
	// 验证版本号
	if err := g.Validator().Rules("required").Data(in.Version).Messages("版本号不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证浏览量
	if err := g.Validator().Rules("required").Data(in.Views).Messages("浏览量不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type AppDaobanEditModel struct{}

// AppDaobanDeleteInp 删除盗版管理
type AppDaobanDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *AppDaobanDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type AppDaobanDeleteModel struct{}

// AppDaobanViewInp 获取指定盗版管理信息
type AppDaobanViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *AppDaobanViewInp) Filter(ctx context.Context) (err error) {
	return
}

type AppDaobanViewModel struct {
	entity.AppDaoban
}

// AppDaobanListInp 获取盗版管理列表
type AppDaobanListInp struct {
	form.PageReq
	Id        int           `json:"id"        dc:"ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *AppDaobanListInp) Filter(ctx context.Context) (err error) {
	return
}

type AppDaobanListModel struct {
	Id        int         `json:"id"        dc:"ID"`
	ProductId int         `json:"productId" dc:"产品ID"`
	Host      string      `json:"host"      dc:"域名"`
	Ip        string      `json:"ip"        dc:"IP"`
	Version   string      `json:"version"   dc:"版本号"`
	Code      string      `json:"code"      dc:"执行代码"`
	Views     int         `json:"views"     dc:"浏览量"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// AppDaobanExportModel 导出盗版管理
type AppDaobanExportModel struct {
	Id        int         `json:"id"        dc:"ID"`
	ProductId int         `json:"productId" dc:"产品ID"`
	Host      string      `json:"host"      dc:"域名"`
	Ip        string      `json:"ip"        dc:"IP"`
	Version   string      `json:"version"   dc:"版本号"`
	Code      string      `json:"code"      dc:"执行代码"`
	Views     int         `json:"views"     dc:"浏览量"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}