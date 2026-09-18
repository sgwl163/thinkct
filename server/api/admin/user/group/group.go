// Package group
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package group

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/userin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询分组管理列表
type ListReq struct {
	g.Meta `path:"/group/list" method:"get" tags:"分组管理" summary:"获取分组管理列表"`
	userin.GroupListInp
}

type ListRes struct {
	form.PageRes
	List []*userin.GroupListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出分组管理列表
type ExportReq struct {
	g.Meta `path:"/group/export" method:"get" tags:"分组管理" summary:"导出分组管理列表"`
	userin.GroupListInp
}

type ExportRes struct{}

// ViewReq 获取分组管理指定信息
type ViewReq struct {
	g.Meta `path:"/group/view" method:"get" tags:"分组管理" summary:"获取分组管理指定信息"`
	userin.GroupViewInp
}

type ViewRes struct {
	*userin.GroupViewModel
}

// EditReq 修改/新增分组管理
type EditReq struct {
	g.Meta `path:"/group/edit" method:"post" tags:"分组管理" summary:"修改/新增分组管理"`
	userin.GroupEditInp
}

type EditRes struct{}

// DeleteReq 删除分组管理
type DeleteReq struct {
	g.Meta `path:"/group/delete" method:"post" tags:"分组管理" summary:"删除分组管理"`
	userin.GroupDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新分组管理状态
type StatusReq struct {
	g.Meta `path:"/group/status" method:"post" tags:"分组管理" summary:"更新分组管理状态"`
	userin.GroupStatusInp
}

type StatusRes struct{}