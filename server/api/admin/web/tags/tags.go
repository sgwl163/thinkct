// Package tags
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package tags

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/webin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询标签管理列表
type ListReq struct {
	g.Meta `path:"/tags/list" method:"get" tags:"标签管理" summary:"获取标签管理列表"`
	webin.TagsListInp
}

type ListRes struct {
	form.PageRes
	List []*webin.TagsListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出标签管理列表
type ExportReq struct {
	g.Meta `path:"/tags/export" method:"get" tags:"标签管理" summary:"导出标签管理列表"`
	webin.TagsListInp
}

type ExportRes struct{}

// ViewReq 获取标签管理指定信息
type ViewReq struct {
	g.Meta `path:"/tags/view" method:"get" tags:"标签管理" summary:"获取标签管理指定信息"`
	webin.TagsViewInp
}

type ViewRes struct {
	*webin.TagsViewModel
}

// EditReq 修改/新增标签管理
type EditReq struct {
	g.Meta `path:"/tags/edit" method:"post" tags:"标签管理" summary:"修改/新增标签管理"`
	webin.TagsEditInp
}

type EditRes struct{}

// DeleteReq 删除标签管理
type DeleteReq struct {
	g.Meta `path:"/tags/delete" method:"post" tags:"标签管理" summary:"删除标签管理"`
	webin.TagsDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新标签管理状态
type StatusReq struct {
	g.Meta `path:"/tags/status" method:"post" tags:"标签管理" summary:"更新标签管理状态"`
	webin.TagsStatusInp
}

type StatusRes struct{}