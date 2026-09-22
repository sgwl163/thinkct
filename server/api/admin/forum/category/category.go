// Package category
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package category

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/forumin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询社区板块列表
type ListReq struct {
	g.Meta `path:"/category/list" method:"get" tags:"社区板块" summary:"获取社区板块列表"`
	forumin.CategoryListInp
}

type ListRes struct {
	form.PageRes
	List []*forumin.CategoryListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出社区板块列表
type ExportReq struct {
	g.Meta `path:"/category/export" method:"get" tags:"社区板块" summary:"导出社区板块列表"`
	forumin.CategoryListInp
}

type ExportRes struct{}

// ViewReq 获取社区板块指定信息
type ViewReq struct {
	g.Meta `path:"/category/view" method:"get" tags:"社区板块" summary:"获取社区板块指定信息"`
	forumin.CategoryViewInp
}

type ViewRes struct {
	*forumin.CategoryViewModel
}

// EditReq 修改/新增社区板块
type EditReq struct {
	g.Meta `path:"/category/edit" method:"post" tags:"社区板块" summary:"修改/新增社区板块"`
	forumin.CategoryEditInp
}

type EditRes struct{}

// DeleteReq 删除社区板块
type DeleteReq struct {
	g.Meta `path:"/category/delete" method:"post" tags:"社区板块" summary:"删除社区板块"`
	forumin.CategoryDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取社区板块最大排序
type MaxSortReq struct {
	g.Meta `path:"/category/maxSort" method:"get" tags:"社区板块" summary:"获取社区板块最大排序"`
	forumin.CategoryMaxSortInp
}

type MaxSortRes struct {
	*forumin.CategoryMaxSortModel
}

// StatusReq 更新社区板块状态
type StatusReq struct {
	g.Meta `path:"/category/status" method:"post" tags:"社区板块" summary:"更新社区板块状态"`
	forumin.CategoryStatusInp
}

type StatusRes struct{}