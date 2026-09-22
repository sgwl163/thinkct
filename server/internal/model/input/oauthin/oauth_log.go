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
)

// OauthLogUpdateFields 修改升级记录字段过滤
type OauthLogUpdateFields struct {
	UserId    int    `json:"userId"    dc:"用户ID"`
	ProductId int    `json:"productId" dc:"产品ID"`
	Host      string `json:"host"      dc:"域名"`
	Content   string `json:"content"   dc:"内容"`
	Ip        string `json:"ip"        dc:"IP"`
	Version   string `json:"version"   dc:"版本号"`
}

// OauthLogInsertFields 新增升级记录字段过滤
type OauthLogInsertFields struct {
	UserId    int    `json:"userId"    dc:"用户ID"`
	ProductId int    `json:"productId" dc:"产品ID"`
	Host      string `json:"host"      dc:"域名"`
	Content   string `json:"content"   dc:"内容"`
	Ip        string `json:"ip"        dc:"IP"`
	Version   string `json:"version"   dc:"版本号"`
}

// OauthLogEditInp 修改/新增升级记录
type OauthLogEditInp struct {
	entity.OauthLog
}

func (in *OauthLogEditInp) Filter(ctx context.Context) (err error) {
	// 验证用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证产品ID
	if err := g.Validator().Rules("required").Data(in.ProductId).Messages("产品ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证域名
	if err := g.Validator().Rules("required").Data(in.Host).Messages("域名不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证内容
	if err := g.Validator().Rules("required").Data(in.Content).Messages("内容不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证IP
	if err := g.Validator().Rules("required").Data(in.Ip).Messages("IP不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证版本号
	if err := g.Validator().Rules("required").Data(in.Version).Messages("版本号不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type OauthLogEditModel struct{}

// OauthLogDeleteInp 删除升级记录
type OauthLogDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *OauthLogDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type OauthLogDeleteModel struct{}

// OauthLogViewInp 获取指定升级记录信息
type OauthLogViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *OauthLogViewInp) Filter(ctx context.Context) (err error) {
	return
}

type OauthLogViewModel struct {
	entity.OauthLog
}

// OauthLogListInp 获取升级记录列表
type OauthLogListInp struct {
	form.PageReq
	Id        int     `json:"id"        dc:"ID"`
	CreatedAt []int64 `json:"createdAt" dc:"创建时间"`
}

func (in *OauthLogListInp) Filter(ctx context.Context) (err error) {
	return
}

type OauthLogListModel struct {
	Id        int    `json:"id"        dc:"ID"`
	UserId    int    `json:"userId"    dc:"用户ID"`
	ProductId int    `json:"productId" dc:"产品ID"`
	Host      string `json:"host"      dc:"域名"`
	Ip        string `json:"ip"        dc:"IP"`
	Version   string `json:"version"   dc:"版本号"`
	CreatedAt int64  `json:"createdAt" dc:"创建时间"`
	UpdatedAt int64  `json:"updatedAt" dc:"修改时间"`
}

// OauthLogExportModel 导出升级记录
type OauthLogExportModel struct {
	Id        int    `json:"id"        dc:"ID"`
	UserId    int    `json:"userId"    dc:"用户ID"`
	ProductId int    `json:"productId" dc:"产品ID"`
	Host      string `json:"host"      dc:"域名"`
	Ip        string `json:"ip"        dc:"IP"`
	Version   string `json:"version"   dc:"版本号"`
	CreatedAt int64  `json:"createdAt" dc:"创建时间"`
	UpdatedAt int64  `json:"updatedAt" dc:"修改时间"`
}