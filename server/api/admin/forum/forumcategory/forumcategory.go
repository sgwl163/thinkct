// Package forumcategory
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package forumcategory

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/forumin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询社区板块列表
type ListReq struct {
	g.Meta `path:"/forumCategory/list" method:"get" tags:"社区板块" summary:"获取社区板块列表"`
	forumin.ForumCategoryListInp
}

type ListRes struct {
	form.PageRes
	List []*forumin.ForumCategoryListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出社区板块列表
type ExportReq struct {
	g.Meta `path:"/forumCategory/export" method:"get" tags:"社区板块" summary:"导出社区板块列表"`
	forumin.ForumCategoryListInp
}

type ExportRes struct{}

// ViewReq 获取社区板块指定信息
type ViewReq struct {
	g.Meta `path:"/forumCategory/view" method:"get" tags:"社区板块" summary:"获取社区板块指定信息"`
	forumin.ForumCategoryViewInp
}

type ViewRes struct {
	*forumin.ForumCategoryViewModel
}

// EditReq 修改/新增社区板块
type EditReq struct {
	g.Meta `path:"/forumCategory/edit" method:"post" tags:"社区板块" summary:"修改/新增社区板块"`
	forumin.ForumCategoryEditInp
}

type EditRes struct{}

// DeleteReq 删除社区板块
type DeleteReq struct {
	g.Meta `path:"/forumCategory/delete" method:"post" tags:"社区板块" summary:"删除社区板块"`
	forumin.ForumCategoryDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取社区板块最大排序
type MaxSortReq struct {
	g.Meta `path:"/forumCategory/maxSort" method:"get" tags:"社区板块" summary:"获取社区板块最大排序"`
	forumin.ForumCategoryMaxSortInp
}

type MaxSortRes struct {
	*forumin.ForumCategoryMaxSortModel
}

// StatusReq 更新社区板块状态
type StatusReq struct {
	g.Meta `path:"/forumCategory/status" method:"post" tags:"社区板块" summary:"更新社区板块状态"`
	forumin.ForumCategoryStatusInp
}

type StatusRes struct{}