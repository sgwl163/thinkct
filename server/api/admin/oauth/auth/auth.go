// Package auth
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package auth

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/oauthin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询授权管理列表
type ListReq struct {
	g.Meta `path:"/auth/list" method:"get" tags:"授权管理" summary:"获取授权管理列表"`
	oauthin.AuthListInp
}

type ListRes struct {
	form.PageRes
	List []*oauthin.AuthListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出授权管理列表
type ExportReq struct {
	g.Meta `path:"/auth/export" method:"get" tags:"授权管理" summary:"导出授权管理列表"`
	oauthin.AuthListInp
}

type ExportRes struct{}

// ViewReq 获取授权管理指定信息
type ViewReq struct {
	g.Meta `path:"/auth/view" method:"get" tags:"授权管理" summary:"获取授权管理指定信息"`
	oauthin.AuthViewInp
}

type ViewRes struct {
	*oauthin.AuthViewModel
}

// EditReq 修改/新增授权管理
type EditReq struct {
	g.Meta `path:"/auth/edit" method:"post" tags:"授权管理" summary:"修改/新增授权管理"`
	oauthin.AuthEditInp
}

type EditRes struct{}

// DeleteReq 删除授权管理
type DeleteReq struct {
	g.Meta `path:"/auth/delete" method:"post" tags:"授权管理" summary:"删除授权管理"`
	oauthin.AuthDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新授权管理状态
type StatusReq struct {
	g.Meta `path:"/auth/status" method:"post" tags:"授权管理" summary:"更新授权管理状态"`
	oauthin.AuthStatusInp
}

type StatusRes struct{}