// Package appauth
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package appauth

import (
	"hotgo/internal/model/input/appin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询授权管理列表
type ListReq struct {
	g.Meta `path:"/appAuth/list" method:"get" tags:"授权管理" summary:"获取授权管理列表"`
	appin.AppAuthListInp
}

type ListRes struct {
	form.PageRes
	List []*appin.AppAuthListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出授权管理列表
type ExportReq struct {
	g.Meta `path:"/appAuth/export" method:"get" tags:"授权管理" summary:"导出授权管理列表"`
	appin.AppAuthListInp
}

type ExportRes struct{}

// ViewReq 获取授权管理指定信息
type ViewReq struct {
	g.Meta `path:"/appAuth/view" method:"get" tags:"授权管理" summary:"获取授权管理指定信息"`
	appin.AppAuthViewInp
}

type ViewRes struct {
	*appin.AppAuthViewModel
}

// EditReq 修改/新增授权管理
type EditReq struct {
	g.Meta `path:"/appAuth/edit" method:"post" tags:"授权管理" summary:"修改/新增授权管理"`
	appin.AppAuthEditInp
}

type EditRes struct{}

// DeleteReq 删除授权管理
type DeleteReq struct {
	g.Meta `path:"/appAuth/delete" method:"post" tags:"授权管理" summary:"删除授权管理"`
	appin.AppAuthDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新授权管理状态
type StatusReq struct {
	g.Meta `path:"/appAuth/status" method:"post" tags:"授权管理" summary:"更新授权管理状态"`
	appin.AppAuthStatusInp
}

type StatusRes struct{}