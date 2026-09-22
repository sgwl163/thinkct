// Package comment
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package comment

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/forumin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询评论管理列表
type ListReq struct {
	g.Meta `path:"/comment/list" method:"get" tags:"评论管理" summary:"获取评论管理列表"`
	forumin.CommentListInp
}

type ListRes struct {
	form.PageRes
	List []*forumin.CommentListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出评论管理列表
type ExportReq struct {
	g.Meta `path:"/comment/export" method:"get" tags:"评论管理" summary:"导出评论管理列表"`
	forumin.CommentListInp
}

type ExportRes struct{}

// ViewReq 获取评论管理指定信息
type ViewReq struct {
	g.Meta `path:"/comment/view" method:"get" tags:"评论管理" summary:"获取评论管理指定信息"`
	forumin.CommentViewInp
}

type ViewRes struct {
	*forumin.CommentViewModel
}

// EditReq 修改/新增评论管理
type EditReq struct {
	g.Meta `path:"/comment/edit" method:"post" tags:"评论管理" summary:"修改/新增评论管理"`
	forumin.CommentEditInp
}

type EditRes struct{}

// DeleteReq 删除评论管理
type DeleteReq struct {
	g.Meta `path:"/comment/delete" method:"post" tags:"评论管理" summary:"删除评论管理"`
	forumin.CommentDeleteInp
}

type DeleteRes struct{}