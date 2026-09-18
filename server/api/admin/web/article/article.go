// Package article
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package article

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/webin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询文章管理列表
type ListReq struct {
	g.Meta `path:"/article/list" method:"get" tags:"文章管理" summary:"获取文章管理列表"`
	webin.ArticleListInp
}

type ListRes struct {
	form.PageRes
	List []*webin.ArticleListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出文章管理列表
type ExportReq struct {
	g.Meta `path:"/article/export" method:"get" tags:"文章管理" summary:"导出文章管理列表"`
	webin.ArticleListInp
}

type ExportRes struct{}

// ViewReq 获取文章管理指定信息
type ViewReq struct {
	g.Meta `path:"/article/view" method:"get" tags:"文章管理" summary:"获取文章管理指定信息"`
	webin.ArticleViewInp
}

type ViewRes struct {
	*webin.ArticleViewModel
}

// EditReq 修改/新增文章管理
type EditReq struct {
	g.Meta `path:"/article/edit" method:"post" tags:"文章管理" summary:"修改/新增文章管理"`
	webin.ArticleEditInp
}

type EditRes struct{}

// DeleteReq 删除文章管理
type DeleteReq struct {
	g.Meta `path:"/article/delete" method:"post" tags:"文章管理" summary:"删除文章管理"`
	webin.ArticleDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新文章管理状态
type StatusReq struct {
	g.Meta `path:"/article/status" method:"post" tags:"文章管理" summary:"更新文章管理状态"`
	webin.ArticleStatusInp
}

type StatusRes struct{}