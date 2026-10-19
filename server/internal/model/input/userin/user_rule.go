// Package userin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package userin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRuleUpdateFields 修改权限规则字段过滤
type UserRuleUpdateFields struct {
	Pid          int    `json:"pid"          dc:"上级菜单"`
	Type         string `json:"type"         dc:"类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮"`
	Title        string `json:"title"        dc:"标题"`
	Name         string `json:"name"         dc:"规则名称"`
	Path         string `json:"path"         dc:"路由路径"`
	Icon         string `json:"icon"         dc:"图标"`
	MenuType     string `json:"menuType"     dc:"菜单类型:tab=选项卡,link=链接,iframe=Iframe"`
	Url          string `json:"url"          dc:"Url"`
	Component    string `json:"component"    dc:"组件路径"`
	NoLoginValid int    `json:"noLoginValid" dc:"未登录有效:0=否,1=是"`
	Extend       string `json:"extend"       dc:"扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单"`
	Remark       string `json:"remark"       dc:"备注"`
	Weigh        int    `json:"weigh"        dc:"权重"`
	Status       string `json:"status"       dc:"状态:0=禁用,1=启用"`
}

// UserRuleInsertFields 新增权限规则字段过滤
type UserRuleInsertFields struct {
	Pid          int    `json:"pid"          dc:"上级菜单"`
	Type         string `json:"type"         dc:"类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮"`
	Title        string `json:"title"        dc:"标题"`
	Name         string `json:"name"         dc:"规则名称"`
	Path         string `json:"path"         dc:"路由路径"`
	Icon         string `json:"icon"         dc:"图标"`
	MenuType     string `json:"menuType"     dc:"菜单类型:tab=选项卡,link=链接,iframe=Iframe"`
	Url          string `json:"url"          dc:"Url"`
	Component    string `json:"component"    dc:"组件路径"`
	NoLoginValid int    `json:"noLoginValid" dc:"未登录有效:0=否,1=是"`
	Extend       string `json:"extend"       dc:"扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单"`
	Remark       string `json:"remark"       dc:"备注"`
	Weigh        int    `json:"weigh"        dc:"权重"`
	Status       string `json:"status"       dc:"状态:0=禁用,1=启用"`
}

// UserRuleEditInp 修改/新增权限规则
type UserRuleEditInp struct {
	entity.UserRule
}

func (in *UserRuleEditInp) Filter(ctx context.Context) (err error) {
	// 验证上级菜单
	if err := g.Validator().Rules("required").Data(in.Pid).Messages("上级菜单不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮
	if err := g.Validator().Rules("required").Data(in.Type).Messages("类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证标题
	if err := g.Validator().Rules("required").Data(in.Title).Messages("标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证规则名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("规则名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证路由路径
	if err := g.Validator().Rules("required").Data(in.Path).Messages("路由路径不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证图标
	if err := g.Validator().Rules("required").Data(in.Icon).Messages("图标不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证菜单类型:tab=选项卡,link=链接,iframe=Iframe
	if err := g.Validator().Rules("required").Data(in.MenuType).Messages("菜单类型:tab=选项卡,link=链接,iframe=Iframe不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证Url
	if err := g.Validator().Rules("required").Data(in.Url).Messages("Url不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证组件路径
	if err := g.Validator().Rules("required").Data(in.Component).Messages("组件路径不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证未登录有效:0=否,1=是
	if err := g.Validator().Rules("required").Data(in.NoLoginValid).Messages("未登录有效:0=否,1=是不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单
	if err := g.Validator().Rules("required").Data(in.Extend).Messages("扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证备注
	if err := g.Validator().Rules("required").Data(in.Remark).Messages("备注不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证权重
	if err := g.Validator().Rules("required").Data(in.Weigh).Messages("权重不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态:0=禁用,1=启用
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态:0=禁用,1=启用不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type UserRuleEditModel struct{}

// UserRuleDeleteInp 删除权限规则
type UserRuleDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *UserRuleDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type UserRuleDeleteModel struct{}

// UserRuleViewInp 获取指定权限规则信息
type UserRuleViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *UserRuleViewInp) Filter(ctx context.Context) (err error) {
	return
}

type UserRuleViewModel struct {
	entity.UserRule
}

// UserRuleListInp 获取权限规则列表
type UserRuleListInp struct {
	form.PageReq
	Id        int           `json:"id"        dc:"ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"更新时间"`
}

func (in *UserRuleListInp) Filter(ctx context.Context) (err error) {
	return
}

type UserRuleListModel struct {
	Id           int         `json:"id"           dc:"ID"`
	Pid          int         `json:"pid"          dc:"上级菜单"`
	Type         string      `json:"type"         dc:"类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮"`
	Title        string      `json:"title"        dc:"标题"`
	Name         string      `json:"name"         dc:"规则名称"`
	Path         string      `json:"path"         dc:"路由路径"`
	Icon         string      `json:"icon"         dc:"图标"`
	MenuType     string      `json:"menuType"     dc:"菜单类型:tab=选项卡,link=链接,iframe=Iframe"`
	Url          string      `json:"url"          dc:"Url"`
	Component    string      `json:"component"    dc:"组件路径"`
	NoLoginValid int         `json:"noLoginValid" dc:"未登录有效:0=否,1=是"`
	Extend       string      `json:"extend"       dc:"扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单"`
	Remark       string      `json:"remark"       dc:"备注"`
	Weigh        int         `json:"weigh"        dc:"权重"`
	Status       string      `json:"status"       dc:"状态:0=禁用,1=启用"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"更新时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"创建时间"`
}

// UserRuleExportModel 导出权限规则
type UserRuleExportModel struct {
	Id           int         `json:"id"           dc:"ID"`
	Pid          int         `json:"pid"          dc:"上级菜单"`
	Type         string      `json:"type"         dc:"类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮"`
	Title        string      `json:"title"        dc:"标题"`
	Name         string      `json:"name"         dc:"规则名称"`
	Path         string      `json:"path"         dc:"路由路径"`
	Icon         string      `json:"icon"         dc:"图标"`
	MenuType     string      `json:"menuType"     dc:"菜单类型:tab=选项卡,link=链接,iframe=Iframe"`
	Url          string      `json:"url"          dc:"Url"`
	Component    string      `json:"component"    dc:"组件路径"`
	NoLoginValid int         `json:"noLoginValid" dc:"未登录有效:0=否,1=是"`
	Extend       string      `json:"extend"       dc:"扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单"`
	Remark       string      `json:"remark"       dc:"备注"`
	Weigh        int         `json:"weigh"        dc:"权重"`
	Status       string      `json:"status"       dc:"状态:0=禁用,1=启用"`
	CreatedAt    *gtime.Time `json:"createdAt"    dc:"更新时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    dc:"创建时间"`
}