// Package user
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package user

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/userin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询会员管理列表
type ListReq struct {
	g.Meta `path:"/user/list" method:"get" tags:"会员管理" summary:"获取会员管理列表"`
	userin.UserListInp
}

type ListRes struct {
	form.PageRes
	List []*userin.UserListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出会员管理列表
type ExportReq struct {
	g.Meta `path:"/user/export" method:"get" tags:"会员管理" summary:"导出会员管理列表"`
	userin.UserListInp
}

type ExportRes struct{}

// ViewReq 获取会员管理指定信息
type ViewReq struct {
	g.Meta `path:"/user/view" method:"get" tags:"会员管理" summary:"获取会员管理指定信息"`
	userin.UserViewInp
}

type ViewRes struct {
	*userin.UserViewModel
}

// EditReq 修改/新增会员管理
type EditReq struct {
	g.Meta `path:"/user/edit" method:"post" tags:"会员管理" summary:"修改/新增会员管理"`
	userin.UserEditInp
}

type EditRes struct{}

// DeleteReq 删除会员管理
type DeleteReq struct {
	g.Meta `path:"/user/delete" method:"post" tags:"会员管理" summary:"删除会员管理"`
	userin.UserDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新会员管理状态
type StatusReq struct {
	g.Meta `path:"/user/status" method:"post" tags:"会员管理" summary:"更新会员管理状态"`
	userin.UserStatusInp
}

type StatusRes struct{}

// SwitchReq 更新会员管理开关状态
type SwitchReq struct {
	g.Meta `path:"/user/switch" method:"post" tags:"会员管理" summary:"更新会员管理状态"`
	userin.UserSwitchInp
}

type SwitchRes struct{}