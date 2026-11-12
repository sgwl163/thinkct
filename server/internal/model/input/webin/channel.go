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

	"hotgo/utility/tree"
)

// ChannelUpdateFields 修改栏目管理字段过滤
type ChannelUpdateFields struct {
	Pid              int    `json:"pid"              dc:"上级频道"`
	Level            int    `json:"level"            dc:"tree"`
	Tree             string `json:"tree"             dc:"关系树"`
	Model            string `json:"model"            dc:"模型"`
	Type             string `json:"type"             dc:"类型:cover=封面频道,list=列表频道,link=跳转链接"`
	Name             string `json:"name"             dc:"频道名称"`
	Diyname          string `json:"diyname"          dc:"自定义名称"`
	Template         string `json:"template"         dc:"模板地址"`
	Url              string `json:"url"              dc:"跳转链接"`
	Target           string `json:"target"           dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Seotitle         string `json:"seotitle"         dc:"SEO标题"`
	Keywords         string `json:"keywords"         dc:"关键词"`
	Description      string `json:"description"      dc:"描述"`
	AllowVisitGroups string `json:"allowVisitGroups" dc:"访问策略:all=全部,user=仅会员"`
	Weigh            int    `json:"weigh"            dc:"权重"`
	Status           int    `json:"status"           dc:"状态:0=禁用,1=启用"`
	UpdateTime       int64  `json:"updateTime"       dc:"修改时间"`
	CreateTime       int64  `json:"createTime"       dc:"创建时间"`
}

// ChannelInsertFields 新增栏目管理字段过滤
type ChannelInsertFields struct {
	Pid              int    `json:"pid"              dc:"上级频道"`
	Level            int    `json:"level"            dc:"tree"`
	Tree             string `json:"tree"             dc:"关系树"`
	Model            string `json:"model"            dc:"模型"`
	Type             string `json:"type"             dc:"类型:cover=封面频道,list=列表频道,link=跳转链接"`
	Name             string `json:"name"             dc:"频道名称"`
	Diyname          string `json:"diyname"          dc:"自定义名称"`
	Template         string `json:"template"         dc:"模板地址"`
	Url              string `json:"url"              dc:"跳转链接"`
	Target           string `json:"target"           dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Seotitle         string `json:"seotitle"         dc:"SEO标题"`
	Keywords         string `json:"keywords"         dc:"关键词"`
	Description      string `json:"description"      dc:"描述"`
	AllowVisitGroups string `json:"allowVisitGroups" dc:"访问策略:all=全部,user=仅会员"`
	Weigh            int    `json:"weigh"            dc:"权重"`
	Status           int    `json:"status"           dc:"状态:0=禁用,1=启用"`
	UpdateTime       int64  `json:"updateTime"       dc:"修改时间"`
	CreateTime       int64  `json:"createTime"       dc:"创建时间"`
}

// ChannelEditInp 修改/新增栏目管理
type ChannelEditInp struct {
	entity.WebChannel
}

func (in *ChannelEditInp) Filter(ctx context.Context) (err error) {
	// 验证类型:cover=封面频道,list=列表频道,link=跳转链接
	if err := g.Validator().Rules("required").Data(in.Type).Messages("类型:cover=封面频道,list=列表频道,link=跳转链接不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 非跳转链接栏目必须配置路由别名；跳转链接栏目可不配置。
	if in.Type != "link" || in.Diyname != "" {
		if err := g.Validator().Rules("required").Data(in.Diyname).Messages("栏目别名不能为空").Run(ctx); err != nil {
			return err.Current()
		}
		if err := g.Validator().Rules("max-length:64").Data(in.Diyname).Messages("栏目别名长度不能超过64个字符").Run(ctx); err != nil {
			return err.Current()
		}
		switch in.Diyname {
		case "about", "auth", "api", "forum", "login", "register", "forgot-password", "user", "storage", "assets", "images", "_next", "testauth", "search":
			return gerror.New("栏目别名与系统保留路径冲突，请更换其他别名")
		}
		if err := g.Validator().Rules("regex:^[a-z]([a-z0-9-]*[a-z0-9])?$").Data(in.Diyname).Messages("栏目别名须以小写字母开头，只能包含小写字母、数字和连字符，不能以连字符结尾或为纯数字").Run(ctx); err != nil {
			return err.Current()
		}
	}
	// 验证频道名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("频道名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证模型
	if err := g.Validator().Rules("required").Data(in.Model).Messages("模型不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证模板地址
	if err := g.Validator().Rules("required").Data(in.Template).Messages("模板地址不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 跳转链接、跳转方式和访问策略为可选项
	// 验证SEO标题
	if err := g.Validator().Rules("required").Data(in.Seotitle).Messages("SEO标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证关键词
	if err := g.Validator().Rules("required").Data(in.Keywords).Messages("关键词不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证描述
	if err := g.Validator().Rules("required").Data(in.Description).Messages("描述不能为空").Run(ctx); err != nil {
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

type ChannelEditModel struct{}

// ChannelDeleteInp 删除栏目管理
type ChannelDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *ChannelDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ChannelDeleteModel struct{}

// ChannelViewInp 获取指定栏目管理信息
type ChannelViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *ChannelViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ChannelViewModel struct {
	entity.WebChannel
}

// ChannelListInp 获取栏目管理列表
type ChannelListInp struct {
	form.PageReq
	Id     int `json:"id"     dc:"ID"`
	Pid    int `json:"pid"    dc:"上级频道"`
	Status int `json:"status" dc:"状态:0=禁用,1=启用"`
}

func (in *ChannelListInp) Filter(ctx context.Context) (err error) {
	return
}

type ChannelListModel struct {
	Id               int    `json:"id"               dc:"ID"`
	Pid              int    `json:"pid"              dc:"上级频道"`
	Model            string `json:"model"            dc:"模型"`
	Type             string `json:"type"             dc:"类型:cover=封面频道,list=列表频道,link=跳转链接"`
	Name             string `json:"name"             dc:"频道名称"`
	Diyname          string `json:"diyname"          dc:"自定义名称"`
	Template         string `json:"template"         dc:"模板地址"`
	Url              string `json:"url"              dc:"跳转链接"`
	Target           string `json:"target"           dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Seotitle         string `json:"seotitle"         dc:"SEO标题"`
	Keywords         string `json:"keywords"         dc:"关键词"`
	Description      string `json:"description"      dc:"描述"`
	AllowVisitGroups string `json:"allowVisitGroups" dc:"访问策略:all=全部,user=仅会员"`
	Weigh            int    `json:"weigh"            dc:"权重"`
	Status           int    `json:"status"           dc:"状态:0=禁用,1=启用"`
	UpdateTime       int64  `json:"updateTime"       dc:"修改时间"`
	CreateTime       int64  `json:"createTime"       dc:"创建时间"`
}

// ChannelExportModel 导出栏目管理
type ChannelExportModel struct {
	Id               int    `json:"id"               dc:"ID"`
	Pid              int    `json:"pid"              dc:"上级频道"`
	Model            string `json:"model"            dc:"模型"`
	Type             string `json:"type"             dc:"类型:cover=封面频道,list=列表频道,link=跳转链接"`
	Name             string `json:"name"             dc:"频道名称"`
	Diyname          string `json:"diyname"          dc:"自定义名称"`
	Template         string `json:"template"         dc:"模板地址"`
	Url              string `json:"url"              dc:"跳转链接"`
	Target           string `json:"target"           dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Seotitle         string `json:"seotitle"         dc:"SEO标题"`
	Keywords         string `json:"keywords"         dc:"关键词"`
	Description      string `json:"description"      dc:"描述"`
	AllowVisitGroups string `json:"allowVisitGroups" dc:"访问策略:all=全部,user=仅会员"`
	Weigh            int    `json:"weigh"            dc:"权重"`
	Status           int    `json:"status"           dc:"状态:0=禁用,1=启用"`
	UpdateTime       int64  `json:"updateTime"       dc:"修改时间"`
	CreateTime       int64  `json:"createTime"       dc:"创建时间"`
}

// ChannelStatusInp 更新栏目管理状态
type ChannelStatusInp struct {
	Id     int `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int `json:"status" dc:"状态"`
}

func (in *ChannelStatusInp) Filter(ctx context.Context) (err error) {
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

type ChannelStatusModel struct{}

// ChannelTreeOption 关系树选项
type ChannelTreeOption struct {
	Id   int    `json:"id"   dc:"ID"`
	Pid  int    `json:"pid"  dc:"上级频道"`
	Name string `json:"name" dc:"频道名称"`

	Children []tree.Node `json:"children"  dc:"子节点"`
}

// ID 获取节点ID
func (t *ChannelTreeOption) ID() int64 {
	return int64(t.Id)
}

// PID 获取父级节点ID
func (t *ChannelTreeOption) PID() int64 {
	return int64(t.Pid)
}

// SetChildren 设置子节点数据
func (t *ChannelTreeOption) SetChildren(children []tree.Node) {
	t.Children = children
}
