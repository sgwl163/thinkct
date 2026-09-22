// Package oauthin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package oauthin

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

// VersionUpdateFields 修改版本更新字段过滤
type VersionUpdateFields struct {
	ProductId int    `json:"productId" dc:"产品ID"`
	Cover     string `json:"cover"     dc:"封面"`
	Title     string `json:"title"     dc:"标题"`
	Content   string `json:"content"   dc:"详情"`
	Version   string `json:"version"   dc:"版本号"`
	File      string `json:"file"      dc:"文件"`
	Status    int    `json:"status"    dc:"状态:1=启用,0=禁用"`
}

// VersionInsertFields 新增版本更新字段过滤
type VersionInsertFields struct {
	ProductId int    `json:"productId" dc:"产品ID"`
	Cover     string `json:"cover"     dc:"封面"`
	Title     string `json:"title"     dc:"标题"`
	Content   string `json:"content"   dc:"详情"`
	Version   string `json:"version"   dc:"版本号"`
	File      string `json:"file"      dc:"文件"`
	Status    int    `json:"status"    dc:"状态:1=启用,0=禁用"`
}

// VersionEditInp 修改/新增版本更新
type VersionEditInp struct {
	entity.OauthVersion
}

func (in *VersionEditInp) Filter(ctx context.Context) (err error) {
	// 验证产品ID
	if err := g.Validator().Rules("required").Data(in.ProductId).Messages("产品ID不能为空").Run(ctx); err != nil {
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

type VersionEditModel struct{}

// VersionDeleteInp 删除版本更新
type VersionDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *VersionDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type VersionDeleteModel struct{}

// VersionViewInp 获取指定版本更新信息
type VersionViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *VersionViewInp) Filter(ctx context.Context) (err error) {
	return
}

type VersionViewModel struct {
	entity.OauthVersion
}

// VersionListInp 获取版本更新列表
type VersionListInp struct {
	form.PageReq
	Id        int           `json:"id"        dc:"ID"`
	Status    int           `json:"status"    dc:"状态:1=启用,0=禁用"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *VersionListInp) Filter(ctx context.Context) (err error) {
	return
}

type VersionListModel struct {
	Id        int         `json:"id"        dc:"ID"`
	ProductId int         `json:"productId" dc:"产品ID"`
	Cover     string      `json:"cover"     dc:"封面"`
	Title     string      `json:"title"     dc:"标题"`
	Version   string      `json:"version"   dc:"版本号"`
	File      string      `json:"file"      dc:"文件"`
	Status    int         `json:"status"    dc:"状态:1=启用,0=禁用"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// VersionExportModel 导出版本更新
type VersionExportModel struct {
	Id        int         `json:"id"        dc:"ID"`
	ProductId int         `json:"productId" dc:"产品ID"`
	Cover     string      `json:"cover"     dc:"封面"`
	Title     string      `json:"title"     dc:"标题"`
	Version   string      `json:"version"   dc:"版本号"`
	File      string      `json:"file"      dc:"文件"`
	Status    int         `json:"status"    dc:"状态:1=启用,0=禁用"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// VersionStatusInp 更新版本更新状态
type VersionStatusInp struct {
	Id     int `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int `json:"status" dc:"状态"`
}

func (in *VersionStatusInp) Filter(ctx context.Context) (err error) {
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

type VersionStatusModel struct{}