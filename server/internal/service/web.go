// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/webin"
	"hotgo/utility/tree"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IWebArticle interface {
		// Model 文章管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取文章管理列表
		List(ctx context.Context, in *webin.ArticleListInp) (list []*webin.ArticleListModel, totalCount int, err error)
		// Export 导出文章管理
		Export(ctx context.Context, in *webin.ArticleListInp) (err error)
		// Edit 修改/新增文章管理
		Edit(ctx context.Context, in *webin.ArticleEditInp) (err error)
		// Delete 删除文章管理
		Delete(ctx context.Context, in *webin.ArticleDeleteInp) (err error)
		// View 获取文章管理指定信息
		View(ctx context.Context, in *webin.ArticleViewInp) (res *webin.ArticleViewModel, err error)
		// Status 更新文章管理状态
		Status(ctx context.Context, in *webin.ArticleStatusInp) (err error)
	}
	IWebChannel interface {
		// Model 栏目管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取栏目管理列表
		List(ctx context.Context, in *webin.ChannelListInp) (list []*webin.ChannelListModel, totalCount int, err error)
		// Export 导出栏目管理
		Export(ctx context.Context, in *webin.ChannelListInp) (err error)
		// Edit 修改/新增栏目管理
		Edit(ctx context.Context, in *webin.ChannelEditInp) (err error)
		// Delete 删除栏目管理
		Delete(ctx context.Context, in *webin.ChannelDeleteInp) (err error)
		// View 获取栏目管理指定信息
		View(ctx context.Context, in *webin.ChannelViewInp) (res *webin.ChannelViewModel, err error)
		// Status 更新栏目管理状态
		Status(ctx context.Context, in *webin.ChannelStatusInp) (err error)
		// TreeOption 获取栏目管理关系树选项
		TreeOption(ctx context.Context, modelName string) (nodes []tree.Node, err error)
	}
	IWebFeedback interface {
		// Model 留言管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取留言管理列表
		List(ctx context.Context, in *webin.FeedbackListInp) (list []*webin.FeedbackListModel, totalCount int, err error)
		// Export 导出留言管理
		Export(ctx context.Context, in *webin.FeedbackListInp) (err error)
		// Edit 修改/新增留言管理
		Edit(ctx context.Context, in *webin.FeedbackEditInp) (err error)
		// Delete 删除留言管理
		Delete(ctx context.Context, in *webin.FeedbackDeleteInp) (err error)
		// View 获取留言管理指定信息
		View(ctx context.Context, in *webin.FeedbackViewInp) (res *webin.FeedbackViewModel, err error)
		// Status 更新留言管理状态
		Status(ctx context.Context, in *webin.FeedbackStatusInp) (err error)
	}
	IWebLinks interface {
		// Model 友情链接ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取友情链接列表
		List(ctx context.Context, in *webin.LinksListInp) (list []*webin.LinksListModel, totalCount int, err error)
		// Export 导出友情链接
		Export(ctx context.Context, in *webin.LinksListInp) (err error)
		// Edit 修改/新增友情链接
		Edit(ctx context.Context, in *webin.LinksEditInp) (err error)
		// Delete 删除友情链接
		Delete(ctx context.Context, in *webin.LinksDeleteInp) (err error)
		// View 获取友情链接指定信息
		View(ctx context.Context, in *webin.LinksViewInp) (res *webin.LinksViewModel, err error)
		// Status 更新友情链接状态
		Status(ctx context.Context, in *webin.LinksStatusInp) (err error)
	}
	IWebProducts interface {
		// Model 产品管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取产品管理列表
		List(ctx context.Context, in *webin.ProductsListInp) (list []*webin.ProductsListModel, totalCount int, err error)
		// Export 导出产品管理
		Export(ctx context.Context, in *webin.ProductsListInp) (err error)
		// Edit 修改/新增产品管理
		Edit(ctx context.Context, in *webin.ProductsEditInp) (err error)
		// Delete 删除产品管理
		Delete(ctx context.Context, in *webin.ProductsDeleteInp) (err error)
		// View 获取产品管理指定信息
		View(ctx context.Context, in *webin.ProductsViewInp) (res *webin.ProductsViewModel, err error)
		// Status 更新产品管理状态
		Status(ctx context.Context, in *webin.ProductsStatusInp) (err error)
	}
	IWebTags interface {
		// Model 标签管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取标签管理列表
		List(ctx context.Context, in *webin.TagsListInp) (list []*webin.TagsListModel, totalCount int, err error)
		// Export 导出标签管理
		Export(ctx context.Context, in *webin.TagsListInp) (err error)
		// Edit 修改/新增标签管理
		Edit(ctx context.Context, in *webin.TagsEditInp) (err error)
		// Delete 删除标签管理
		Delete(ctx context.Context, in *webin.TagsDeleteInp) (err error)
		// View 获取标签管理指定信息
		View(ctx context.Context, in *webin.TagsViewInp) (res *webin.TagsViewModel, err error)
		// Status 更新标签管理状态
		Status(ctx context.Context, in *webin.TagsStatusInp) (err error)
	}
)

var (
	localWebArticle  IWebArticle
	localWebChannel  IWebChannel
	localWebFeedback IWebFeedback
	localWebLinks    IWebLinks
	localWebProducts IWebProducts
	localWebTags     IWebTags
)

func WebArticle() IWebArticle {
	if localWebArticle == nil {
		panic("implement not found for interface IWebArticle, forgot register?")
	}
	return localWebArticle
}

func RegisterWebArticle(i IWebArticle) {
	localWebArticle = i
}

func WebChannel() IWebChannel {
	if localWebChannel == nil {
		panic("implement not found for interface IWebChannel, forgot register?")
	}
	return localWebChannel
}

func RegisterWebChannel(i IWebChannel) {
	localWebChannel = i
}

func WebFeedback() IWebFeedback {
	if localWebFeedback == nil {
		panic("implement not found for interface IWebFeedback, forgot register?")
	}
	return localWebFeedback
}

func RegisterWebFeedback(i IWebFeedback) {
	localWebFeedback = i
}

func WebLinks() IWebLinks {
	if localWebLinks == nil {
		panic("implement not found for interface IWebLinks, forgot register?")
	}
	return localWebLinks
}

func RegisterWebLinks(i IWebLinks) {
	localWebLinks = i
}

func WebProducts() IWebProducts {
	if localWebProducts == nil {
		panic("implement not found for interface IWebProducts, forgot register?")
	}
	return localWebProducts
}

func RegisterWebProducts(i IWebProducts) {
	localWebProducts = i
}

func WebTags() IWebTags {
	if localWebTags == nil {
		panic("implement not found for interface IWebTags, forgot register?")
	}
	return localWebTags
}

func RegisterWebTags(i IWebTags) {
	localWebTags = i
}
