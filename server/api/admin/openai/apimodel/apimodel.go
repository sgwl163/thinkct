// Package apimodel
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package apimodel

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/openaiin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询AI模型列表列表
type ListReq struct {
	g.Meta `path:"/apiModel/list" method:"get" tags:"AI模型列表" summary:"获取AI模型列表列表"`
	openaiin.ApiModelListInp
}

type ListRes struct {
	form.PageRes
	List []*openaiin.ApiModelListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出AI模型列表列表
type ExportReq struct {
	g.Meta `path:"/apiModel/export" method:"get" tags:"AI模型列表" summary:"导出AI模型列表列表"`
	openaiin.ApiModelListInp
}

type ExportRes struct{}

// ViewReq 获取AI模型列表指定信息
type ViewReq struct {
	g.Meta `path:"/apiModel/view" method:"get" tags:"AI模型列表" summary:"获取AI模型列表指定信息"`
	openaiin.ApiModelViewInp
}

type ViewRes struct {
	*openaiin.ApiModelViewModel
}

// EditReq 修改/新增AI模型列表
type EditReq struct {
	g.Meta `path:"/apiModel/edit" method:"post" tags:"AI模型列表" summary:"修改/新增AI模型列表"`
	openaiin.ApiModelEditInp
}

type EditRes struct{}

// DeleteReq 删除AI模型列表
type DeleteReq struct {
	g.Meta `path:"/apiModel/delete" method:"post" tags:"AI模型列表" summary:"删除AI模型列表"`
	openaiin.ApiModelDeleteInp
}

type DeleteRes struct{}

// MaxSortReq 获取AI模型列表最大排序
type MaxSortReq struct {
	g.Meta `path:"/apiModel/maxSort" method:"get" tags:"AI模型列表" summary:"获取AI模型列表最大排序"`
	openaiin.ApiModelMaxSortInp
}

type MaxSortRes struct {
	*openaiin.ApiModelMaxSortModel
}

// StatusReq 更新AI模型列表状态
type StatusReq struct {
	g.Meta `path:"/apiModel/status" method:"post" tags:"AI模型列表" summary:"更新AI模型列表状态"`
	openaiin.ApiModelStatusInp
}

type StatusRes struct{}