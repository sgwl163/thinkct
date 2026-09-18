// Package group
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package group

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/imin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询IM-群组表列表
type ListReq struct {
	g.Meta `path:"/group/list" method:"get" tags:"IM-群组表" summary:"获取IM-群组表列表"`
	imin.GroupListInp
}

type ListRes struct {
	form.PageRes
	List []*imin.GroupListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出IM-群组表列表
type ExportReq struct {
	g.Meta `path:"/group/export" method:"get" tags:"IM-群组表" summary:"导出IM-群组表列表"`
	imin.GroupListInp
}

type ExportRes struct{}

// ViewReq 获取IM-群组表指定信息
type ViewReq struct {
	g.Meta `path:"/group/view" method:"get" tags:"IM-群组表" summary:"获取IM-群组表指定信息"`
	imin.GroupViewInp
}

type ViewRes struct {
	*imin.GroupViewModel
}

// EditReq 修改/新增IM-群组表
type EditReq struct {
	g.Meta `path:"/group/edit" method:"post" tags:"IM-群组表" summary:"修改/新增IM-群组表"`
	imin.GroupEditInp
}

type EditRes struct{}

// DeleteReq 删除IM-群组表
type DeleteReq struct {
	g.Meta `path:"/group/delete" method:"post" tags:"IM-群组表" summary:"删除IM-群组表"`
	imin.GroupDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新IM-群组表状态
type StatusReq struct {
	g.Meta `path:"/group/status" method:"post" tags:"IM-群组表" summary:"更新IM-群组表状态"`
	imin.GroupStatusInp
}

type StatusRes struct{}