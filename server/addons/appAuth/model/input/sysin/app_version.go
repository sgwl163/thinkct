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

// AppVersionUpdateFields 修改版本更新字段过滤
type AppVersionUpdateFields struct {
	ProductId int    `json:"productId" dc:"产品ID"`
	Cover     string `json:"cover"     dc:"封面"`
	Title     string `json:"title"     dc:"标题"`
	Content   string `json:"content"   dc:"详情"`
	Version   string `json:"version"   dc:"版本号"`
	File      string `json:"file"      dc:"文件"`
	Status    int    `json:"status"    dc:"状态:1=启用,0=禁用"`
}

// AppVersionInsertFields 新增版本更新字段过滤
type AppVersionInsertFields struct {
	ProductId int    `json:"productId" dc:"产品ID"`
	Cover     string `json:"cover"     dc:"封面"`
	Title     string `json:"title"     dc:"标题"`
	Content   string `json:"content"   dc:"详情"`
	Version   string `json:"version"   dc:"版本号"`
	File      string `json:"file"      dc:"文件"`
	Status    int    `json:"status"    dc:"状态:1=启用,0=禁用"`
}

// AppVersionEditInp 修改/新增版本更新
type AppVersionEditInp struct {
	entity.AppVersion
}

func (in *AppVersionEditInp) Filter(ctx context.Context) (err error) {
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

type AppVersionEditModel struct{}

// AppVersionDeleteInp 删除版本更新
type AppVersionDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *AppVersionDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type AppVersionDeleteModel struct{}

// AppVersionViewInp 获取指定版本更新信息
type AppVersionViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *AppVersionViewInp) Filter(ctx context.Context) (err error) {
	return
}

type AppVersionViewModel struct {
	entity.AppVersion
}

// AppVersionListInp 获取版本更新列表
type AppVersionListInp struct {
	form.PageReq
	Id        int           `json:"id"        dc:"ID"`
	Status    int           `json:"status"    dc:"状态:1=启用,0=禁用"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *AppVersionListInp) Filter(ctx context.Context) (err error) {
	return
}

type AppVersionListModel struct {
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

// AppVersionExportModel 导出版本更新
type AppVersionExportModel struct {
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

// AppVersionStatusInp 更新版本更新状态
type AppVersionStatusInp struct {
	Id     int `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int `json:"status" dc:"状态"`
}

func (in *AppVersionStatusInp) Filter(ctx context.Context) (err error) {
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

type AppVersionStatusModel struct{}