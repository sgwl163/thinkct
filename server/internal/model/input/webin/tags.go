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

// TagsUpdateFields 修改标签管理字段过滤
type TagsUpdateFields struct {
	Name          string `json:"name"          dc:"名称"`
	Type          string `json:"type"          dc:"标签类型:default=default,success=success,info=info,warning=warning,danger=danger"`
	DocumentCount int    `json:"documentCount" dc:"文档数"`
	Seotitle      string `json:"seotitle"      dc:"SEO标题"`
	Keywords      string `json:"keywords"      dc:"关键词"`
	Description   string `json:"description"   dc:"描述"`
	InnerChain    int    `json:"innerChain"    dc:"自动内链:0=关,1=开"`
	Weigh         int    `json:"weigh"         dc:"权重"`
	Status        int    `json:"status"        dc:"状态:0=禁用,1=启用"`
	UpdateTime    int64  `json:"updateTime"    dc:"修改时间"`
	CreateTime    int64  `json:"createTime"    dc:"创建时间"`
}

// TagsInsertFields 新增标签管理字段过滤
type TagsInsertFields struct {
	Name          string `json:"name"          dc:"名称"`
	Type          string `json:"type"          dc:"标签类型:default=default,success=success,info=info,warning=warning,danger=danger"`
	DocumentCount int    `json:"documentCount" dc:"文档数"`
	Seotitle      string `json:"seotitle"      dc:"SEO标题"`
	Keywords      string `json:"keywords"      dc:"关键词"`
	Description   string `json:"description"   dc:"描述"`
	InnerChain    int    `json:"innerChain"    dc:"自动内链:0=关,1=开"`
	Weigh         int    `json:"weigh"         dc:"权重"`
	Status        int    `json:"status"        dc:"状态:0=禁用,1=启用"`
	UpdateTime    int64  `json:"updateTime"    dc:"修改时间"`
	CreateTime    int64  `json:"createTime"    dc:"创建时间"`
}

// TagsEditInp 修改/新增标签管理
type TagsEditInp struct {
	entity.WebTags
}

func (in *TagsEditInp) Filter(ctx context.Context) (err error) {
	// 验证名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证标签类型:default=default,success=success,info=info,warning=warning,danger=danger
	if err := g.Validator().Rules("required").Data(in.Type).Messages("标签类型:default=default,success=success,info=info,warning=warning,danger=danger不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证文章数量
	// if err := g.Validator().Rules("required").Data(in.DocumentCount).Messages("文章数量不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证SEO标题
	// if err := g.Validator().Rules("required").Data(in.Seotitle).Messages("SEO标题不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证关键词
	// if err := g.Validator().Rules("required").Data(in.Keywords).Messages("关键词不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证描述
	// if err := g.Validator().Rules("required").Data(in.Description).Messages("描述不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证自动内链:0=关,1=开
	// if err := g.Validator().Rules("required").Data(in.InnerChain).Messages("自动内链:0=关,1=开不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证状态:0=禁用,1=启用
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态:0=禁用,1=启用不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态:0=禁用,1=启用值不正确").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type TagsEditModel struct{}

// TagsDeleteInp 删除标签管理
type TagsDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *TagsDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type TagsDeleteModel struct{}

// TagsViewInp 获取指定标签管理信息
type TagsViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *TagsViewInp) Filter(ctx context.Context) (err error) {
	return
}

type TagsViewModel struct {
	entity.WebTags
}

// TagsListInp 获取标签管理列表
type TagsListInp struct {
	form.PageReq
	Id     int `json:"id"     dc:"ID"`
	Status int `json:"status" dc:"状态:0=禁用,1=启用"`
}

func (in *TagsListInp) Filter(ctx context.Context) (err error) {
	return
}

type TagsListModel struct {
	Id            int    `json:"id"            dc:"ID"`
	Name          string `json:"name"          dc:"名称"`
	Type          string `json:"type"          dc:"标签类型:default=default,success=success,info=info,warning=warning,danger=danger"`
	DocumentCount int    `json:"documentCount" dc:"文档数"`
	Seotitle      string `json:"seotitle"      dc:"SEO标题"`
	Keywords      string `json:"keywords"      dc:"关键词"`
	Description   string `json:"description"   dc:"描述"`
	InnerChain    int    `json:"innerChain"    dc:"自动内链:0=关,1=开"`
	Weigh         int    `json:"weigh"         dc:"权重"`
	Status        int    `json:"status"        dc:"状态:0=禁用,1=启用"`
	UpdateTime    int64  `json:"updateTime"    dc:"修改时间"`
	CreateTime    int64  `json:"createTime"    dc:"创建时间"`
}

// TagsExportModel 导出标签管理
type TagsExportModel struct {
	Id            int    `json:"id"            dc:"ID"`
	Name          string `json:"name"          dc:"名称"`
	Type          string `json:"type"          dc:"标签类型:default=default,success=success,info=info,warning=warning,danger=danger"`
	DocumentCount int    `json:"documentCount" dc:"文档数"`
	Seotitle      string `json:"seotitle"      dc:"SEO标题"`
	Keywords      string `json:"keywords"      dc:"关键词"`
	Description   string `json:"description"   dc:"描述"`
	InnerChain    int    `json:"innerChain"    dc:"自动内链:0=关,1=开"`
	Weigh         int    `json:"weigh"         dc:"权重"`
	Status        int    `json:"status"        dc:"状态:0=禁用,1=启用"`
	UpdateTime    int64  `json:"updateTime"    dc:"修改时间"`
	CreateTime    int64  `json:"createTime"    dc:"创建时间"`
}

// TagsStatusInp 更新标签管理状态
type TagsStatusInp struct {
	Id     int `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int `json:"status" dc:"状态"`
}

func (in *TagsStatusInp) Filter(ctx context.Context) (err error) {
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

type TagsStatusModel struct{}
