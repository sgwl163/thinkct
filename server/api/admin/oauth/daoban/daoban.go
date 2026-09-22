// Package daoban
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package daoban

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/oauthin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询盗版管理列表
type ListReq struct {
	g.Meta `path:"/daoban/list" method:"get" tags:"盗版管理" summary:"获取盗版管理列表"`
	oauthin.DaobanListInp
}

type ListRes struct {
	form.PageRes
	List []*oauthin.DaobanListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出盗版管理列表
type ExportReq struct {
	g.Meta `path:"/daoban/export" method:"get" tags:"盗版管理" summary:"导出盗版管理列表"`
	oauthin.DaobanListInp
}

type ExportRes struct{}

// ViewReq 获取盗版管理指定信息
type ViewReq struct {
	g.Meta `path:"/daoban/view" method:"get" tags:"盗版管理" summary:"获取盗版管理指定信息"`
	oauthin.DaobanViewInp
}

type ViewRes struct {
	*oauthin.DaobanViewModel
}

// EditReq 修改/新增盗版管理
type EditReq struct {
	g.Meta `path:"/daoban/edit" method:"post" tags:"盗版管理" summary:"修改/新增盗版管理"`
	oauthin.DaobanEditInp
}

type EditRes struct{}

// DeleteReq 删除盗版管理
type DeleteReq struct {
	g.Meta `path:"/daoban/delete" method:"post" tags:"盗版管理" summary:"删除盗版管理"`
	oauthin.DaobanDeleteInp
}

type DeleteRes struct{}