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

// ProductsUpdateFields 修改产品管理字段过滤
type ProductsUpdateFields struct {
	UserId      int     `json:"userId"      dc:"用户ID"`
	AdminId     int     `json:"adminId"     dc:"管理员ID"`
	ChannelId   int     `json:"channelId"   dc:"频道"`
	ChannelIds  string  `json:"channelIds"  dc:"副频道"`
	Diyname     string  `json:"diyname"     dc:"自定义名称"`
	Title       string  `json:"title"       dc:"标题"`
	TitleStyle  string  `json:"titleStyle"  dc:"标题样式"`
	Flag        string  `json:"flag"        dc:"标志:top=置顶,hot=热门,recommend=推荐,new=最新"`
	Image       string  `json:"image"       dc:"图片"`
	Images      string  `json:"images"      dc:"封面图片"`
	Content     string  `json:"content"     dc:"内容"`
	Seotitle    string  `json:"seotitle"    dc:"SEO标题"`
	Keywords    string  `json:"keywords"    dc:"关键词"`
	Description string  `json:"description" dc:"描述"`
	Tags        string  `json:"tags"        dc:"标签"`
	Url         string  `json:"url"         dc:"外部链接"`
	Target      string  `json:"target"      dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Views       int     `json:"views"       dc:"浏览量"`
	Comments    int     `json:"comments"    dc:"评论量"`
	Likes       int     `json:"likes"       dc:"点赞数"`
	Dislikes    int     `json:"dislikes"    dc:"点踩数"`
	Weigh       int     `json:"weigh"       dc:"权重"`
	Price       float64 `json:"price"       dc:"价格"`
	Status      string  `json:"status"      dc:"状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线"`
	PublishTime int64   `json:"publishTime" dc:"发布时间"`
	UpdateTime  int64   `json:"updateTime"  dc:"修改时间"`
	CreateTime  int64   `json:"createTime"  dc:"创建时间"`
	Name        string  `json:"name"        dc:"型号名称"`
}

// ProductsInsertFields 新增产品管理字段过滤
type ProductsInsertFields struct {
	UserId      int     `json:"userId"      dc:"用户ID"`
	AdminId     int     `json:"adminId"     dc:"管理员ID"`
	ChannelId   int     `json:"channelId"   dc:"频道"`
	ChannelIds  string  `json:"channelIds"  dc:"副频道"`
	Diyname     string  `json:"diyname"     dc:"自定义名称"`
	Title       string  `json:"title"       dc:"标题"`
	TitleStyle  string  `json:"titleStyle"  dc:"标题样式"`
	Flag        string  `json:"flag"        dc:"标志:top=置顶,hot=热门,recommend=推荐,new=最新"`
	Image       string  `json:"image"       dc:"图片"`
	Images      string  `json:"images"      dc:"封面图片"`
	Content     string  `json:"content"     dc:"内容"`
	Seotitle    string  `json:"seotitle"    dc:"SEO标题"`
	Keywords    string  `json:"keywords"    dc:"关键词"`
	Description string  `json:"description" dc:"描述"`
	Tags        string  `json:"tags"        dc:"标签"`
	Url         string  `json:"url"         dc:"外部链接"`
	Target      string  `json:"target"      dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Views       int     `json:"views"       dc:"浏览量"`
	Comments    int     `json:"comments"    dc:"评论量"`
	Likes       int     `json:"likes"       dc:"点赞数"`
	Dislikes    int     `json:"dislikes"    dc:"点踩数"`
	Weigh       int     `json:"weigh"       dc:"权重"`
	Price       float64 `json:"price"       dc:"价格"`
	Status      string  `json:"status"      dc:"状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线"`
	PublishTime int64   `json:"publishTime" dc:"发布时间"`
	UpdateTime  int64   `json:"updateTime"  dc:"修改时间"`
	CreateTime  int64   `json:"createTime"  dc:"创建时间"`
	Name        string  `json:"name"        dc:"型号名称"`
}

// ProductsEditInp 修改/新增产品管理
type ProductsEditInp struct {
	entity.WebProducts
}

func (in *ProductsEditInp) Filter(ctx context.Context) (err error) {
	// 验证标题
	if err := g.Validator().Rules("required").Data(in.Title).Messages("标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证标志
	// if err := g.Validator().Rules("required").Data(in.Flag).Messages("标志不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// if err := g.Validator().Rules("required").Data(in.Name).Messages("型号名称不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证频道
	if err := g.Validator().Rules("min:1").Data(in.ChannelId).Messages("所属频道不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证标志:top=置顶,hot=热门,recommend=推荐,new=最新
	// if err := g.Validator().Rules("required").Data(in.Flag).Messages("标志不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证价格
	if err := g.Validator().Rules("regex:(^[0-9]{1,10}$)|(^[0-9]{1,10}[\\.]{1}[0-9]{1,2}$)").Data(in.Price).Messages("价格最多允许输入10位整数及2位小数").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type ProductsEditModel struct{}

// ProductsDeleteInp 删除产品管理
type ProductsDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *ProductsDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ProductsDeleteModel struct{}

// ProductsViewInp 获取指定产品管理信息
type ProductsViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *ProductsViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ProductsViewModel struct {
	entity.WebProducts
}

// ProductsListInp 获取产品管理列表
type ProductsListInp struct {
	form.PageReq
	Id int `json:"id" dc:"ID"`
}

func (in *ProductsListInp) Filter(ctx context.Context) (err error) {
	return
}

type ProductsListModel struct {
	Id          int     `json:"id"          dc:"ID"`
	UserId      int     `json:"userId"      dc:"用户ID"`
	AdminId     int     `json:"adminId"     dc:"管理员ID"`
	ChannelId   int     `json:"channelId"   dc:"频道"`
	ChannelIds  string  `json:"channelIds"  dc:"副频道"`
	Diyname     string  `json:"diyname"     dc:"自定义名称"`
	Title       string  `json:"title"       dc:"标题"`
	TitleStyle  string  `json:"titleStyle"  dc:"标题样式"`
	Flag        string  `json:"flag"        dc:"标志:top=置顶,hot=热门,recommend=推荐,new=最新"`
	Image       string  `json:"image"       dc:"图片"`
	Images      string  `json:"images"      dc:"多图"`
	Seotitle    string  `json:"seotitle"    dc:"SEO标题"`
	Keywords    string  `json:"keywords"    dc:"关键词"`
	Description string  `json:"description" dc:"描述"`
	Tags        string  `json:"tags"        dc:"标签"`
	Url         string  `json:"url"         dc:"外部链接"`
	Target      string  `json:"target"      dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Views       int     `json:"views"       dc:"浏览量"`
	Comments    int     `json:"comments"    dc:"评论量"`
	Likes       int     `json:"likes"       dc:"点赞数"`
	Dislikes    int     `json:"dislikes"    dc:"点踩数"`
	Weigh       int     `json:"weigh"       dc:"权重"`
	Price       float64 `json:"price"       dc:"价格"`
	Status      string  `json:"status"      dc:"状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线"`
	PublishTime int64   `json:"publishTime" dc:"发布时间"`
	UpdateTime  int64   `json:"updateTime"  dc:"修改时间"`
	CreateTime  int64   `json:"createTime"  dc:"创建时间"`
	Name        string  `json:"name"        dc:"型号名称"`
}

// ProductsExportModel 导出产品管理
type ProductsExportModel struct {
	Id          int     `json:"id"          dc:"ID"`
	UserId      int     `json:"userId"      dc:"用户ID"`
	AdminId     int     `json:"adminId"     dc:"管理员ID"`
	ChannelId   int     `json:"channelId"   dc:"频道"`
	ChannelIds  string  `json:"channelIds"  dc:"副频道"`
	Diyname     string  `json:"diyname"     dc:"自定义名称"`
	Title       string  `json:"title"       dc:"标题"`
	TitleStyle  string  `json:"titleStyle"  dc:"标题样式"`
	Flag        string  `json:"flag"        dc:"标志:top=置顶,hot=热门,recommend=推荐,new=最新"`
	Image       string  `json:"image"       dc:"图片"`
	Images      string  `json:"images"      dc:"多图"`
	Seotitle    string  `json:"seotitle"    dc:"SEO标题"`
	Keywords    string  `json:"keywords"    dc:"关键词"`
	Description string  `json:"description" dc:"描述"`
	Tags        string  `json:"tags"        dc:"标签"`
	Url         string  `json:"url"         dc:"外部链接"`
	Target      string  `json:"target"      dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Views       int     `json:"views"       dc:"浏览量"`
	Comments    int     `json:"comments"    dc:"评论量"`
	Likes       int     `json:"likes"       dc:"点赞数"`
	Dislikes    int     `json:"dislikes"    dc:"点踩数"`
	Weigh       int     `json:"weigh"       dc:"权重"`
	Price       float64 `json:"price"       dc:"价格"`
	Status      string  `json:"status"      dc:"状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线"`
	PublishTime int64   `json:"publishTime" dc:"发布时间"`
	UpdateTime  int64   `json:"updateTime"  dc:"修改时间"`
	CreateTime  int64   `json:"createTime"  dc:"创建时间"`
	Name        string  `json:"name"        dc:"型号名称"`
}

// ProductsStatusInp 更新产品管理状态
type ProductsStatusInp struct {
	Id     int `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int `json:"status" dc:"状态"`
}

func (in *ProductsStatusInp) Filter(ctx context.Context) (err error) {
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

type ProductsStatusModel struct{}
