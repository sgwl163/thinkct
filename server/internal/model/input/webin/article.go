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

// ArticleUpdateFields 修改文章管理字段过滤
type ArticleUpdateFields struct {
	UserId      int    `json:"userId"      dc:"用户ID"`
	AdminId     int    `json:"adminId"     dc:"管理员ID"`
	ChannelId   int    `json:"channelId"   dc:"频道"`
	ChannelIds  string `json:"channelIds"  dc:"副频道"`
	Title       string `json:"title"       v:"required#文章标题不能为空" dc:"文章标题"`
	Diyname     string `json:"diyname"     dc:"自定义名称"`
	TitleStyle  string `json:"titleStyle"  dc:"标题样式"`
	Flag        string `json:"flag"        dc:"标志:top=置顶,hot=热门,recommend=推荐,new=最新"`
	Images      string `json:"images"      dc:"封面图"`
	Content     string `json:"content"     dc:"内容"`
	Seotitle    string `json:"seotitle"    dc:"SEO标题"`
	Keywords    string `json:"keywords"    dc:"关键词"`
	Description string `json:"description" dc:"描述"`
	Tags        string `json:"tags"        dc:"标签"`
	Url         string `json:"url"         dc:"外部链接"`
	Target      string `json:"target"      dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Views       int    `json:"views"       dc:"浏览量"`
	Comments    int    `json:"comments"    dc:"评论量"`
	Likes       int    `json:"likes"       dc:"点赞数"`
	Dislikes    int    `json:"dislikes"    dc:"点踩数"`
	Weigh       int    `json:"weigh"       dc:"权重"`
	Status      string `json:"status"      dc:"状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线"`
	PublishTime int64  `json:"publishTime" dc:"发布时间"`
	UpdateTime  int64  `json:"updateTime"  dc:"修改时间"`
	CreateTime  int64  `json:"createTime"  dc:"创建时间"`
}

// ArticleInsertFields 新增文章管理字段过滤
type ArticleInsertFields struct {
	UserId      int    `json:"userId"      dc:"用户ID"`
	AdminId     int    `json:"adminId"     dc:"管理员ID"`
	ChannelId   int    `json:"channelId"   dc:"频道"`
	ChannelIds  string `json:"channelIds"  dc:"副频道"`
	Diyname     string `json:"diyname"     dc:"自定义名称"`
	Title       string `json:"title"       dc:"标题"`
	TitleStyle  string `json:"titleStyle"  dc:"标题样式"`
	Flag        string `json:"flag"        dc:"标志:top=置顶,hot=热门,recommend=推荐,new=最新"`
	Images      string `json:"images"      dc:"封面图片"`
	Content     string `json:"content"     dc:"内容"`
	Seotitle    string `json:"seotitle"    dc:"SEO标题"`
	Keywords    string `json:"keywords"    dc:"关键词"`
	Description string `json:"description" dc:"描述"`
	Tags        string `json:"tags"        dc:"标签"`
	Url         string `json:"url"         dc:"外部链接"`
	Target      string `json:"target"      dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Views       int    `json:"views"       dc:"浏览量"`
	Comments    int    `json:"comments"    dc:"评论量"`
	Likes       int    `json:"likes"       dc:"点赞数"`
	Dislikes    int    `json:"dislikes"    dc:"点踩数"`
	Weigh       int    `json:"weigh"       dc:"权重"`
	Status      string `json:"status"      dc:"状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线"`
	PublishTime int64  `json:"publishTime" dc:"发布时间"`
	UpdateTime  int64  `json:"updateTime"  dc:"修改时间"`
	CreateTime  int64  `json:"createTime"  dc:"创建时间"`
}

// ArticleEditInp 修改/新增文章管理
type ArticleEditInp struct {
	entity.WebArticle
}

func (in *ArticleEditInp) Filter(ctx context.Context) (err error) {
	// 验证用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证管理员ID
	if err := g.Validator().Rules("required").Data(in.AdminId).Messages("管理员ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证频道
	if err := g.Validator().Rules("required").Data(in.ChannelId).Messages("频道不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证副频道
	if err := g.Validator().Rules("required").Data(in.ChannelIds).Messages("副频道不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证标题样式
	// if err := g.Validator().Rules("required").Data(in.TitleStyle).Messages("标题样式不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证标志:top=置顶,hot=热门,recommend=推荐,new=最新
	// if err := g.Validator().Rules("required").Data(in.Flag).Messages("标志:top=置顶,hot=热门,recommend=推荐,new=最新不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证封面图片
	if err := g.Validator().Rules("required").Data(in.Images).Messages("封面图片不能为空").Run(ctx); err != nil {
		return err.Current()
	}
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
	// 验证标签
	// if err := g.Validator().Rules("required").Data(in.Tags).Messages("标签不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证外部链接
	// if err := g.Validator().Rules("required").Data(in.Url).Messages("外部链接不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口
	// if err := g.Validator().Rules("required").Data(in.Target).Messages("跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证浏览量
	// if err := g.Validator().Rules("required").Data(in.Views).Messages("浏览量不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证评论量
	// if err := g.Validator().Rules("required").Data(in.Comments).Messages("评论量不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证点赞数
	// if err := g.Validator().Rules("required").Data(in.Likes).Messages("点赞数不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证点踩数
	// if err := g.Validator().Rules("required").Data(in.Dislikes).Messages("点踩数不能为空").Run(ctx); err != nil {
	// 	return err.Current()
	// }
	// 验证状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type ArticleEditModel struct{}

// ArticleDeleteInp 删除文章管理
type ArticleDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *ArticleDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ArticleDeleteModel struct{}

// ArticleViewInp 获取指定文章管理信息
type ArticleViewInp struct {
	Id int `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *ArticleViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ArticleViewModel struct {
	entity.WebArticle
}

// ArticleListInp 获取文章管理列表
type ArticleListInp struct {
	form.PageReq
	Id          int      `json:"id" dc:"ID"`
	ChannelId   int      `json:"channelId" dc:"频道"`
	ChannelIds  string   `json:"channelIds" dc:"副频道"`
	Title       string   `json:"title" dc:"标题"`
	Flag        string   `json:"flag" dc:"标志"`
	Status      string   `json:"status" dc:"状态"`
	CreateTime  []string `json:"createTime" dc:"创建时间"`
	UpdateTime  []string `json:"updateTime" dc:"更新时间"`
	PublishTime []string `json:"publishTime" dc:"发布时间"`
}

func (in *ArticleListInp) Filter(ctx context.Context) (err error) {
	return
}

type ArticleListModel struct {
	Id          int    `json:"id"          dc:"ID"`
	UserId      int    `json:"userId"      dc:"用户ID"`
	AdminId     int    `json:"adminId"     dc:"管理员ID"`
	ChannelId   int    `json:"channelId"   dc:"频道"`
	ChannelIds  string `json:"channelIds"  dc:"副频道"`
	Diyname     string `json:"diyname"     dc:"自定义名称"`
	Title       string `json:"title"       dc:"标题"`
	TitleStyle  string `json:"titleStyle"  dc:"标题样式"`
	Flag        string `json:"flag"        dc:"标志:top=置顶,hot=热门,recommend=推荐,new=最新"`
	Seotitle    string `json:"seotitle"    dc:"SEO标题"`
	Keywords    string `json:"keywords"    dc:"关键词"`
	Description string `json:"description" dc:"描述"`
	Tags        string `json:"tags"        dc:"标签"`
	Url         string `json:"url"         dc:"外部链接"`
	Target      string `json:"target"      dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Views       int    `json:"views"       dc:"浏览量"`
	Comments    int    `json:"comments"    dc:"评论量"`
	Likes       int    `json:"likes"       dc:"点赞数"`
	Dislikes    int    `json:"dislikes"    dc:"点踩数"`
	Weigh       int    `json:"weigh"       dc:"权重"`
	Status      string `json:"status"      dc:"状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线"`
	PublishTime int64  `json:"publishTime" dc:"发布时间"`
	UpdateTime  int64  `json:"updateTime"  dc:"修改时间"`
	CreateTime  int64  `json:"createTime"  dc:"创建时间"`
}

// ArticleExportModel 导出文章管理
type ArticleExportModel struct {
	Id          int    `json:"id"          dc:"ID"`
	UserId      int    `json:"userId"      dc:"用户ID"`
	AdminId     int    `json:"adminId"     dc:"管理员ID"`
	ChannelId   int    `json:"channelId"   dc:"频道"`
	ChannelIds  string `json:"channelIds"  dc:"副频道"`
	Diyname     string `json:"diyname"     dc:"自定义名称"`
	Title       string `json:"title"       dc:"标题"`
	TitleStyle  string `json:"titleStyle"  dc:"标题样式"`
	Flag        string `json:"flag"        dc:"标志:top=置顶,hot=热门,recommend=推荐,new=最新"`
	Seotitle    string `json:"seotitle"    dc:"SEO标题"`
	Keywords    string `json:"keywords"    dc:"关键词"`
	Description string `json:"description" dc:"描述"`
	Tags        string `json:"tags"        dc:"标签"`
	Url         string `json:"url"         dc:"外部链接"`
	Target      string `json:"target"      dc:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Views       int    `json:"views"       dc:"浏览量"`
	Comments    int    `json:"comments"    dc:"评论量"`
	Likes       int    `json:"likes"       dc:"点赞数"`
	Dislikes    int    `json:"dislikes"    dc:"点踩数"`
	Weigh       int    `json:"weigh"       dc:"权重"`
	Status      string `json:"status"      dc:"状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线"`
	PublishTime int64  `json:"publishTime" dc:"发布时间"`
	UpdateTime  int64  `json:"updateTime"  dc:"修改时间"`
	CreateTime  int64  `json:"createTime"  dc:"创建时间"`
}

// ArticleStatusInp 更新文章管理状态
type ArticleStatusInp struct {
	Id     int `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int `json:"status" dc:"状态"`
}

func (in *ArticleStatusInp) Filter(ctx context.Context) (err error) {
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

type ArticleStatusModel struct{}
