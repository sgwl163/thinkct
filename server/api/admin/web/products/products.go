// Package products
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package products

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/webin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询产品管理列表
type ListReq struct {
	g.Meta `path:"/products/list" method:"get" tags:"产品管理" summary:"获取产品管理列表"`
	webin.ProductsListInp
}

type ListRes struct {
	form.PageRes
	List []*webin.ProductsListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出产品管理列表
type ExportReq struct {
	g.Meta `path:"/products/export" method:"get" tags:"产品管理" summary:"导出产品管理列表"`
	webin.ProductsListInp
}

type ExportRes struct{}

// ViewReq 获取产品管理指定信息
type ViewReq struct {
	g.Meta `path:"/products/view" method:"get" tags:"产品管理" summary:"获取产品管理指定信息"`
	webin.ProductsViewInp
}

type ViewRes struct {
	*webin.ProductsViewModel
}

// EditReq 修改/新增产品管理
type EditReq struct {
	g.Meta `path:"/products/edit" method:"post" tags:"产品管理" summary:"修改/新增产品管理"`
	webin.ProductsEditInp
}

type EditRes struct{}

// DeleteReq 删除产品管理
type DeleteReq struct {
	g.Meta `path:"/products/delete" method:"post" tags:"产品管理" summary:"删除产品管理"`
	webin.ProductsDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新产品管理状态
type StatusReq struct {
	g.Meta `path:"/products/status" method:"post" tags:"产品管理" summary:"更新产品管理状态"`
	webin.ProductsStatusInp
}

type StatusRes struct{}