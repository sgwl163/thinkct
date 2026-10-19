// Package appdaoban
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package appdaoban

import (
	"hotgo/addons/appAuth/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询盗版管理列表
type ListReq struct {
	g.Meta `path:"/appDaoban/list" method:"get" tags:"盗版管理" summary:"获取盗版管理列表"`
	sysin.AppDaobanListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.AppDaobanListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出盗版管理列表
type ExportReq struct {
	g.Meta `path:"/appDaoban/export" method:"get" tags:"盗版管理" summary:"导出盗版管理列表"`
	sysin.AppDaobanListInp
}

type ExportRes struct{}

// ViewReq 获取盗版管理指定信息
type ViewReq struct {
	g.Meta `path:"/appDaoban/view" method:"get" tags:"盗版管理" summary:"获取盗版管理指定信息"`
	sysin.AppDaobanViewInp
}

type ViewRes struct {
	*sysin.AppDaobanViewModel
}

// EditReq 修改/新增盗版管理
type EditReq struct {
	g.Meta `path:"/appDaoban/edit" method:"post" tags:"盗版管理" summary:"修改/新增盗版管理"`
	sysin.AppDaobanEditInp
}

type EditRes struct{}

// DeleteReq 删除盗版管理
type DeleteReq struct {
	g.Meta `path:"/appDaoban/delete" method:"post" tags:"盗版管理" summary:"删除盗版管理"`
	sysin.AppDaobanDeleteInp
}

type DeleteRes struct{}