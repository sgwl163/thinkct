// Package userlog
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package userlog

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/userin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询操作记录列表
type ListReq struct {
	g.Meta `path:"/userLog/list" method:"get" tags:"操作记录" summary:"获取操作记录列表"`
	userin.UserLogListInp
}

type ListRes struct {
	form.PageRes
	List []*userin.UserLogListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出操作记录列表
type ExportReq struct {
	g.Meta `path:"/userLog/export" method:"get" tags:"操作记录" summary:"导出操作记录列表"`
	userin.UserLogListInp
}

type ExportRes struct{}

// ViewReq 获取操作记录指定信息
type ViewReq struct {
	g.Meta `path:"/userLog/view" method:"get" tags:"操作记录" summary:"获取操作记录指定信息"`
	userin.UserLogViewInp
}

type ViewRes struct {
	*userin.UserLogViewModel
}

// EditReq 修改/新增操作记录
type EditReq struct {
	g.Meta `path:"/userLog/edit" method:"post" tags:"操作记录" summary:"修改/新增操作记录"`
	userin.UserLogEditInp
}

type EditRes struct{}

// DeleteReq 删除操作记录
type DeleteReq struct {
	g.Meta `path:"/userLog/delete" method:"post" tags:"操作记录" summary:"删除操作记录"`
	userin.UserLogDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新操作记录状态
type StatusReq struct {
	g.Meta `path:"/userLog/status" method:"post" tags:"操作记录" summary:"更新操作记录状态"`
	userin.UserLogStatusInp
}

type StatusRes struct{}