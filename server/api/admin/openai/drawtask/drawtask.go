// Package drawtask
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package drawtask

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/openaiin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询AI绘画任务列表
type ListReq struct {
	g.Meta `path:"/drawTask/list" method:"get" tags:"AI绘画任务" summary:"获取AI绘画任务列表"`
	openaiin.DrawTaskListInp
}

type ListRes struct {
	form.PageRes
	List []*openaiin.DrawTaskListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出AI绘画任务列表
type ExportReq struct {
	g.Meta `path:"/drawTask/export" method:"get" tags:"AI绘画任务" summary:"导出AI绘画任务列表"`
	openaiin.DrawTaskListInp
}

type ExportRes struct{}

// ViewReq 获取AI绘画任务指定信息
type ViewReq struct {
	g.Meta `path:"/drawTask/view" method:"get" tags:"AI绘画任务" summary:"获取AI绘画任务指定信息"`
	openaiin.DrawTaskViewInp
}

type ViewRes struct {
	*openaiin.DrawTaskViewModel
}

// EditReq 修改/新增AI绘画任务
type EditReq struct {
	g.Meta `path:"/drawTask/edit" method:"post" tags:"AI绘画任务" summary:"修改/新增AI绘画任务"`
	openaiin.DrawTaskEditInp
}

type EditRes struct{}

// DeleteReq 删除AI绘画任务
type DeleteReq struct {
	g.Meta `path:"/drawTask/delete" method:"post" tags:"AI绘画任务" summary:"删除AI绘画任务"`
	openaiin.DrawTaskDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新AI绘画任务状态
type StatusReq struct {
	g.Meta `path:"/drawTask/status" method:"post" tags:"AI绘画任务" summary:"更新AI绘画任务状态"`
	openaiin.DrawTaskStatusInp
}

type StatusRes struct{}