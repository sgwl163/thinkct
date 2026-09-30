// Package appin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package appin

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

// AppAuthUpdateFields 修改授权管理字段过滤
type AppAuthUpdateFields struct {
	UserId     int         `json:"userId"     dc:"用户ID"`
	ProductId  int         `json:"productId"  dc:"产品ID"`
	Name       string      `json:"name"       dc:"名称"`
	Host       string      `json:"host"       dc:"域名"`
	Ip         string      `json:"ip"         dc:"IP"`
	Qq         string      `json:"qq"         dc:"QQ"`
	Phone      string      `json:"phone"      dc:"电话"`
	Syskey     string      `json:"syskey"     dc:"授权码"`
	VersionId  string      `json:"versionId"  dc:"版本"`
	AuthMode   string      `json:"authMode"   dc:"验证方式:0=单域名验证,1=IP双重验证"`
	AuthDomain string      `json:"authDomain" dc:"验证域名:0=当前单域名方式,1=顶级泛域名方式"`
	Views      int         `json:"views"      dc:"浏览量"`
	Status     int         `json:"status"     dc:"状态:1=启用,0=禁用"`
	ExpiredAt  *gtime.Time `json:"expiredAt"  dc:"到期时间"`
}

// AppAuthInsertFields 新增授权管理字段过滤
type AppAuthInsertFields struct {
	UserId     int         `json:"userId"     dc:"用户ID"`
	ProductId  int         `json:"productId"  dc:"产品ID"`
	Name       string      `json:"name"       dc:"名称"`
	Host       string      `json:"host"       dc:"域名"`
	Ip         string      `json:"ip"         dc:"IP"`
	Qq         string      `json:"qq"         dc:"QQ"`
	Phone      string      `json:"phone"      dc:"电话"`
	Syskey     string      `json:"syskey"     dc:"授权码"`
	VersionId  string      `json:"versionId"  dc:"版本"`
	AuthMode   string      `json:"authMode"   dc:"验证方式:0=单域名验证,1=IP双重验证"`
	AuthDomain string      `json:"authDomain" dc:"验证域名:0=当前单域名方式,1=顶级泛域名方式"`
	Views      int         `json:"views"      dc:"浏览量"`
	Status     int         `json:"status"     dc:"状态:1=启用,0=禁用"`
	ExpiredAt  *gtime.Time `json:"expiredAt"  dc:"到期时间"`
}

// AppAuthEditInp 修改/新增授权管理
type AppAuthEditInp struct {
	entity.AppAuth
}

func (in *AppAuthEditInp) Filter(ctx context.Context) (err error) {
	// 验证用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证产品ID
	if err := g.Validator().Rules("required").Data(in.ProductId).Messages("产品ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证QQ
	if err := g.Validator().Rules("qq").Data(in.Qq).Messages("QQ不是QQ号码").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证电话
	if err := g.Validator().Rules("required").Data(in.Phone).Messages("电话不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证版本
	if err := g.Validator().Rules("required").Data(in.VersionId).Messages("版本不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证验证方式:0=单域名验证,1=IP双重验证
	if err := g.Validator().Rules("required").Data(in.AuthMode).Messages("验证方式:0=单域名验证,1=IP双重验证不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证验证域名:0=当前单域名方式,1=顶级泛域名方式
	if err := g.Validator().Rules("required").Data(in.AuthDomain).Messages("验证域名:0=当前单域名方式,1=顶级泛域名方式不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态:1=启用,0=禁用
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态:1=启用,0=禁用不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态:1=启用,0=禁用值不正确").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type AppAuthEditModel struct{}

// AppAuthDeleteInp 删除授权管理
type AppAuthDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *AppAuthDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type AppAuthDeleteModel struct{}

// AppAuthViewInp 获取指定授权管理信息
type AppAuthViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *AppAuthViewInp) Filter(ctx context.Context) (err error) {
	return
}

type AppAuthViewModel struct {
	entity.AppAuth
}

// AppAuthListInp 获取授权管理列表
type AppAuthListInp struct {
	form.PageReq
	Id        int           `json:"id"        dc:"ID"`
	Status    int           `json:"status"    dc:"状态:1=启用,0=禁用"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *AppAuthListInp) Filter(ctx context.Context) (err error) {
	return
}

type AppAuthListModel struct {
	Id         int         `json:"id"         dc:"ID"`
	UserId     int         `json:"userId"     dc:"用户ID"`
	ProductId  int         `json:"productId"  dc:"产品ID"`
	Name       string      `json:"name"       dc:"名称"`
	Host       string      `json:"host"       dc:"域名"`
	Ip         string      `json:"ip"         dc:"IP"`
	Qq         string      `json:"qq"         dc:"QQ"`
	Phone      string      `json:"phone"      dc:"电话"`
	Syskey     string      `json:"syskey"     dc:"授权码"`
	VersionId  string      `json:"versionId"  dc:"版本"`
	AuthMode   string      `json:"authMode"   dc:"验证方式:0=单域名验证,1=IP双重验证"`
	AuthDomain string      `json:"authDomain" dc:"验证域名:0=当前单域名方式,1=顶级泛域名方式"`
	Views      int         `json:"views"      dc:"浏览量"`
	Status     int         `json:"status"     dc:"状态:1=启用,0=禁用"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"修改时间"`
	ExpiredAt  *gtime.Time `json:"expiredAt"  dc:"到期时间"`
}

// AppAuthExportModel 导出授权管理
type AppAuthExportModel struct {
	Id         int         `json:"id"         dc:"ID"`
	UserId     int         `json:"userId"     dc:"用户ID"`
	ProductId  int         `json:"productId"  dc:"产品ID"`
	Name       string      `json:"name"       dc:"名称"`
	Host       string      `json:"host"       dc:"域名"`
	Ip         string      `json:"ip"         dc:"IP"`
	Qq         string      `json:"qq"         dc:"QQ"`
	Phone      string      `json:"phone"      dc:"电话"`
	Syskey     string      `json:"syskey"     dc:"授权码"`
	VersionId  string      `json:"versionId"  dc:"版本"`
	AuthMode   string      `json:"authMode"   dc:"验证方式:0=单域名验证,1=IP双重验证"`
	AuthDomain string      `json:"authDomain" dc:"验证域名:0=当前单域名方式,1=顶级泛域名方式"`
	Views      int         `json:"views"      dc:"浏览量"`
	Status     int         `json:"status"     dc:"状态:1=启用,0=禁用"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"修改时间"`
	ExpiredAt  *gtime.Time `json:"expiredAt"  dc:"到期时间"`
}

// AppAuthStatusInp 更新授权管理状态
type AppAuthStatusInp struct {
	Id     int `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int `json:"status" dc:"状态"`
}

func (in *AppAuthStatusInp) Filter(ctx context.Context) (err error) {
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

type AppAuthStatusModel struct{}