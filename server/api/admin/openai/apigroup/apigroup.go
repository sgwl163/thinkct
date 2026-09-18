// Package apigroup
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package apigroup

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/openaiin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询AI接口分组列表
type ListReq struct {
	g.Meta `path:"/apiGroup/list" method:"get" tags:"AI接口分组" summary:"获取AI接口分组列表"`
	openaiin.ApiGroupListInp
}

type ListRes struct {
	form.PageRes
	List []*openaiin.ApiGroupListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出AI接口分组列表
type ExportReq struct {
	g.Meta `path:"/apiGroup/export" method:"get" tags:"AI接口分组" summary:"导出AI接口分组列表"`
	openaiin.ApiGroupListInp
}

type ExportRes struct{}

// ViewReq 获取AI接口分组指定信息
type ViewReq struct {
	g.Meta `path:"/apiGroup/view" method:"get" tags:"AI接口分组" summary:"获取AI接口分组指定信息"`
	openaiin.ApiGroupViewInp
}

type ViewRes struct {
	*openaiin.ApiGroupViewModel
}

// EditReq 修改/新增AI接口分组
type EditReq struct {
	g.Meta `path:"/apiGroup/edit" method:"post" tags:"AI接口分组" summary:"修改/新增AI接口分组"`
	openaiin.ApiGroupEditInp
}

type EditRes struct{}

// DeleteReq 删除AI接口分组
type DeleteReq struct {
	g.Meta `path:"/apiGroup/delete" method:"post" tags:"AI接口分组" summary:"删除AI接口分组"`
	openaiin.ApiGroupDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取AI接口分组最大排序
type MaxSortReq struct {
	g.Meta `path:"/apiGroup/maxSort" method:"get" tags:"AI接口分组" summary:"获取AI接口分组最大排序"`
	openaiin.ApiGroupMaxSortInp
}

type MaxSortRes struct {
	*openaiin.ApiGroupMaxSortModel
}

// StatusReq 更新AI接口分组状态
type StatusReq struct {
	g.Meta `path:"/apiGroup/status" method:"post" tags:"AI接口分组" summary:"更新AI接口分组状态"`
	openaiin.ApiGroupStatusInp
}

type StatusRes struct{}