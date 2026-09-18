// Package videotask
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package videotask

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/openaiin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询AI视频任务列表
type ListReq struct {
	g.Meta `path:"/videoTask/list" method:"get" tags:"AI视频任务" summary:"获取AI视频任务列表"`
	openaiin.VideoTaskListInp
}

type ListRes struct {
	form.PageRes
	List []*openaiin.VideoTaskListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出AI视频任务列表
type ExportReq struct {
	g.Meta `path:"/videoTask/export" method:"get" tags:"AI视频任务" summary:"导出AI视频任务列表"`
	openaiin.VideoTaskListInp
}

type ExportRes struct{}

// ViewReq 获取AI视频任务指定信息
type ViewReq struct {
	g.Meta `path:"/videoTask/view" method:"get" tags:"AI视频任务" summary:"获取AI视频任务指定信息"`
	openaiin.VideoTaskViewInp
}

type ViewRes struct {
	*openaiin.VideoTaskViewModel
}

// EditReq 修改/新增AI视频任务
type EditReq struct {
	g.Meta `path:"/videoTask/edit" method:"post" tags:"AI视频任务" summary:"修改/新增AI视频任务"`
	openaiin.VideoTaskEditInp
}

type EditRes struct{}

// DeleteReq 删除AI视频任务
type DeleteReq struct {
	g.Meta `path:"/videoTask/delete" method:"post" tags:"AI视频任务" summary:"删除AI视频任务"`
	openaiin.VideoTaskDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新AI视频任务状态
type StatusReq struct {
	g.Meta `path:"/videoTask/status" method:"post" tags:"AI视频任务" summary:"更新AI视频任务状态"`
	openaiin.VideoTaskStatusInp
}

type StatusRes struct{}