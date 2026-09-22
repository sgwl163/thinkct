// Package oauthin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package oauthin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DaobanUpdateFields 修改盗版管理字段过滤
type DaobanUpdateFields struct {
	ProductId int    `json:"productId" dc:"产品ID"`
	Host      string `json:"host"      dc:"域名"`
	Ip        string `json:"ip"        dc:"IP"`
	Version   string `json:"version"   dc:"版本号"`
	Code      string `json:"code"      dc:"执行代码"`
	Data      string `json:"data"      dc:"数据"`
	Views     int    `json:"views"     dc:"浏览量"`
}

// DaobanInsertFields 新增盗版管理字段过滤
type DaobanInsertFields struct {
	ProductId int    `json:"productId" dc:"产品ID"`
	Host      string `json:"host"      dc:"域名"`
	Ip        string `json:"ip"        dc:"IP"`
	Version   string `json:"version"   dc:"版本号"`
	Code      string `json:"code"      dc:"执行代码"`
	Data      string `json:"data"      dc:"数据"`
	Views     int    `json:"views"     dc:"浏览量"`
}

// DaobanEditInp 修改/新增盗版管理
type DaobanEditInp struct {
	entity.OauthDaoban
}

func (in *DaobanEditInp) Filter(ctx context.Context) (err error) {
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

type DaobanEditModel struct{}

// DaobanDeleteInp 删除盗版管理
type DaobanDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *DaobanDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type DaobanDeleteModel struct{}

// DaobanViewInp 获取指定盗版管理信息
type DaobanViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *DaobanViewInp) Filter(ctx context.Context) (err error) {
	return
}

type DaobanViewModel struct {
	entity.OauthDaoban
}

// DaobanListInp 获取盗版管理列表
type DaobanListInp struct {
	form.PageReq
	Id        int           `json:"id"        dc:"ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *DaobanListInp) Filter(ctx context.Context) (err error) {
	return
}

type DaobanListModel struct {
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

// DaobanExportModel 导出盗版管理
type DaobanExportModel struct {
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