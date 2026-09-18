// Package webin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package webin

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// LinksUpdateFields 修改友情链接字段过滤
type LinksUpdateFields struct {
	UserId     int    `json:"userId"     dc:"用户ID"`
	Title      string `json:"title"      dc:"标题"`
	Link       string `json:"link"       dc:"链接"`
	Target     string `json:"target"     dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Logo       string `json:"logo"       dc:"LOGO"`
	Weigh      int    `json:"weigh"      dc:"权重"`
	Status     string `json:"status"     dc:"状态:disable=禁用,enable=启用,pending_trial=待审"`
	Remark     string `json:"remark"     dc:"备注"`
	UpdateTime int64  `json:"updateTime" dc:"修改时间"`
	CreateTime int64  `json:"createTime" dc:"创建时间"`
}

// LinksInsertFields 新增友情链接字段过滤
type LinksInsertFields struct {
	UserId     int    `json:"userId"     dc:"用户ID"`
	Title      string `json:"title"      dc:"标题"`
	Link       string `json:"link"       dc:"链接"`
	Target     string `json:"target"     dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Logo       string `json:"logo"       dc:"LOGO"`
	Weigh      int    `json:"weigh"      dc:"权重"`
	Status     string `json:"status"     dc:"状态:disable=禁用,enable=启用,pending_trial=待审"`
	Remark     string `json:"remark"     dc:"备注"`
	UpdateTime int64  `json:"updateTime" dc:"修改时间"`
	CreateTime int64  `json:"createTime" dc:"创建时间"`
}

// LinksEditInp 修改/新增友情链接
type LinksEditInp struct {
	entity.WebLinks
}

func (in *LinksEditInp) Filter(ctx context.Context) (err error) {
	// 验证用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证标题
	if err := g.Validator().Rules("required").Data(in.Title).Messages("标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证链接
	if err := g.Validator().Rules("required").Data(in.Link).Messages("链接不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口
	if err := g.Validator().Rules("required").Data(in.Target).Messages("跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证LOGO
	if err := g.Validator().Rules("required").Data(in.Logo).Messages("LOGO不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态:disable=禁用,enable=启用,pending_trial=待审
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态:disable=禁用,enable=启用,pending_trial=待审不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type LinksEditModel struct{}

// LinksDeleteInp 删除友情链接
type LinksDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *LinksDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type LinksDeleteModel struct{}

// LinksViewInp 获取指定友情链接信息
type LinksViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *LinksViewInp) Filter(ctx context.Context) (err error) {
	return
}

type LinksViewModel struct {
	entity.WebLinks
}

// LinksListInp 获取友情链接列表
type LinksListInp struct {
	form.PageReq
	Id int `json:"id" dc:"ID"`
}

func (in *LinksListInp) Filter(ctx context.Context) (err error) {
	return
}

type LinksListModel struct {
	Id         int    `json:"id"         dc:"ID"`
	UserId     int    `json:"userId"     dc:"用户ID"`
	Title      string `json:"title"      dc:"标题"`
	Link       string `json:"link"       dc:"链接"`
	Target     string `json:"target"     dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Logo       string `json:"logo"       dc:"LOGO"`
	Weigh      int    `json:"weigh"      dc:"权重"`
	Status     string `json:"status"     dc:"状态:disable=禁用,enable=启用,pending_trial=待审"`
	Remark     string `json:"remark"     dc:"备注"`
	UpdateTime int64  `json:"updateTime" dc:"修改时间"`
	CreateTime int64  `json:"createTime" dc:"创建时间"`
}

// LinksExportModel 导出友情链接
type LinksExportModel struct {
	Id         int    `json:"id"         dc:"ID"`
	UserId     int    `json:"userId"     dc:"用户ID"`
	Title      string `json:"title"      dc:"标题"`
	Link       string `json:"link"       dc:"链接"`
	Target     string `json:"target"     dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Logo       string `json:"logo"       dc:"LOGO"`
	Weigh      int    `json:"weigh"      dc:"权重"`
	Status     string `json:"status"     dc:"状态:disable=禁用,enable=启用,pending_trial=待审"`
	Remark     string `json:"remark"     dc:"备注"`
	UpdateTime int64  `json:"updateTime" dc:"修改时间"`
	CreateTime int64  `json:"createTime" dc:"创建时间"`
}

// LinksStatusInp 更新友情链接状态
type LinksStatusInp struct {
	Id     int `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int `json:"status" dc:"状态"`
}

func (in *LinksStatusInp) Filter(ctx context.Context) (err error) {
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

type LinksStatusModel struct{}