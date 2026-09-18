// Package openaiappcategory
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package openaiappcategory

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/openaiin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询AI应用分类列表
type ListReq struct {
	g.Meta `path:"/openaiAppCategory/list" method:"get" tags:"AI应用分类" summary:"获取AI应用分类列表"`
	openaiin.OpenaiAppCategoryListInp
}

type ListRes struct {
	form.PageRes
	List []*openaiin.OpenaiAppCategoryListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出AI应用分类列表
type ExportReq struct {
	g.Meta `path:"/openaiAppCategory/export" method:"get" tags:"AI应用分类" summary:"导出AI应用分类列表"`
	openaiin.OpenaiAppCategoryListInp
}

type ExportRes struct{}

// ViewReq 获取AI应用分类指定信息
type ViewReq struct {
	g.Meta `path:"/openaiAppCategory/view" method:"get" tags:"AI应用分类" summary:"获取AI应用分类指定信息"`
	openaiin.OpenaiAppCategoryViewInp
}

type ViewRes struct {
	*openaiin.OpenaiAppCategoryViewModel
}

// EditReq 修改/新增AI应用分类
type EditReq struct {
	g.Meta `path:"/openaiAppCategory/edit" method:"post" tags:"AI应用分类" summary:"修改/新增AI应用分类"`
	openaiin.OpenaiAppCategoryEditInp
}

type EditRes struct{}

// DeleteReq 删除AI应用分类
type DeleteReq struct {
	g.Meta `path:"/openaiAppCategory/delete" method:"post" tags:"AI应用分类" summary:"删除AI应用分类"`
	openaiin.OpenaiAppCategoryDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取AI应用分类最大排序
type MaxSortReq struct {
	g.Meta `path:"/openaiAppCategory/maxSort" method:"get" tags:"AI应用分类" summary:"获取AI应用分类最大排序"`
	openaiin.OpenaiAppCategoryMaxSortInp
}

type MaxSortRes struct {
	*openaiin.OpenaiAppCategoryMaxSortModel
}

// StatusReq 更新AI应用分类状态
type StatusReq struct {
	g.Meta `path:"/openaiAppCategory/status" method:"post" tags:"AI应用分类" summary:"更新AI应用分类状态"`
	openaiin.OpenaiAppCategoryStatusInp
}

type StatusRes struct{}