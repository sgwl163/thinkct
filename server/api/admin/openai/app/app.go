// Package app
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package app

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/openaiin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询AI应用列表
type ListReq struct {
	g.Meta `path:"/app/list" method:"get" tags:"AI应用" summary:"获取AI应用列表"`
	openaiin.AppListInp
}

type ListRes struct {
	form.PageRes
	List []*openaiin.AppListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出AI应用列表
type ExportReq struct {
	g.Meta `path:"/app/export" method:"get" tags:"AI应用" summary:"导出AI应用列表"`
	openaiin.AppListInp
}

type ExportRes struct{}

// ViewReq 获取AI应用指定信息
type ViewReq struct {
	g.Meta `path:"/app/view" method:"get" tags:"AI应用" summary:"获取AI应用指定信息"`
	openaiin.AppViewInp
}

type ViewRes struct {
	*openaiin.AppViewModel
}

// EditReq 修改/新增AI应用
type EditReq struct {
	g.Meta `path:"/app/edit" method:"post" tags:"AI应用" summary:"修改/新增AI应用"`
	openaiin.AppEditInp
}

type EditRes struct{}

// DeleteReq 删除AI应用
type DeleteReq struct {
	g.Meta `path:"/app/delete" method:"post" tags:"AI应用" summary:"删除AI应用"`
	openaiin.AppDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取AI应用最大排序
type MaxSortReq struct {
	g.Meta `path:"/app/maxSort" method:"get" tags:"AI应用" summary:"获取AI应用最大排序"`
	openaiin.AppMaxSortInp
}

type MaxSortRes struct {
	*openaiin.AppMaxSortModel
}

// StatusReq 更新AI应用状态
type StatusReq struct {
	g.Meta `path:"/app/status" method:"post" tags:"AI应用" summary:"更新AI应用状态"`
	openaiin.AppStatusInp
}

type StatusRes struct{}
