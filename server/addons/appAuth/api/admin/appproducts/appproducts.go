// Package appproducts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package appproducts

import (
	"hotgo/addons/appAuth/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询产品管理列表
type ListReq struct {
	g.Meta `path:"/appProducts/list" method:"get" tags:"产品管理" summary:"获取产品管理列表"`
	sysin.AppProductsListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.AppProductsListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出产品管理列表
type ExportReq struct {
	g.Meta `path:"/appProducts/export" method:"get" tags:"产品管理" summary:"导出产品管理列表"`
	sysin.AppProductsListInp
}

type ExportRes struct{}

// ViewReq 获取产品管理指定信息
type ViewReq struct {
	g.Meta `path:"/appProducts/view" method:"get" tags:"产品管理" summary:"获取产品管理指定信息"`
	sysin.AppProductsViewInp
}

type ViewRes struct {
	*sysin.AppProductsViewModel
}

// EditReq 修改/新增产品管理
type EditReq struct {
	g.Meta `path:"/appProducts/edit" method:"post" tags:"产品管理" summary:"修改/新增产品管理"`
	sysin.AppProductsEditInp
}

type EditRes struct{}

// DeleteReq 删除产品管理
type DeleteReq struct {
	g.Meta `path:"/appProducts/delete" method:"post" tags:"产品管理" summary:"删除产品管理"`
	sysin.AppProductsDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新产品管理状态
type StatusReq struct {
	g.Meta `path:"/appProducts/status" method:"post" tags:"产品管理" summary:"更新产品管理状态"`
	sysin.AppProductsStatusInp
}

type StatusRes struct{}