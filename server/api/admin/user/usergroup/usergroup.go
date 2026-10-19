// Package usergroup
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package usergroup

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/userin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询用户分组列表
type ListReq struct {
	g.Meta `path:"/userGroup/list" method:"get" tags:"用户分组" summary:"获取用户分组列表"`
	userin.UserGroupListInp
}

type ListRes struct {
	form.PageRes
	List []*userin.UserGroupListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出用户分组列表
type ExportReq struct {
	g.Meta `path:"/userGroup/export" method:"get" tags:"用户分组" summary:"导出用户分组列表"`
	userin.UserGroupListInp
}

type ExportRes struct{}

// ViewReq 获取用户分组指定信息
type ViewReq struct {
	g.Meta `path:"/userGroup/view" method:"get" tags:"用户分组" summary:"获取用户分组指定信息"`
	userin.UserGroupViewInp
}

type ViewRes struct {
	*userin.UserGroupViewModel
}

// EditReq 修改/新增用户分组
type EditReq struct {
	g.Meta `path:"/userGroup/edit" method:"post" tags:"用户分组" summary:"修改/新增用户分组"`
	userin.UserGroupEditInp
}

type EditRes struct{}

// DeleteReq 删除用户分组
type DeleteReq struct {
	g.Meta `path:"/userGroup/delete" method:"post" tags:"用户分组" summary:"删除用户分组"`
	userin.UserGroupDeleteInp
}

type DeleteRes struct{}