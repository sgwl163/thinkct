// Package report
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package report

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/forumin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询社区论坛举报表列表
type ListReq struct {
	g.Meta `path:"/report/list" method:"get" tags:"社区论坛举报表" summary:"获取社区论坛举报表列表"`
	forumin.ReportListInp
}

type ListRes struct {
	form.PageRes
	List []*forumin.ReportListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出社区论坛举报表列表
type ExportReq struct {
	g.Meta `path:"/report/export" method:"get" tags:"社区论坛举报表" summary:"导出社区论坛举报表列表"`
	forumin.ReportListInp
}

type ExportRes struct{}

// ViewReq 获取社区论坛举报表指定信息
type ViewReq struct {
	g.Meta `path:"/report/view" method:"get" tags:"社区论坛举报表" summary:"获取社区论坛举报表指定信息"`
	forumin.ReportViewInp
}

type ViewRes struct {
	*forumin.ReportViewModel
}

// EditReq 修改/新增社区论坛举报表
type EditReq struct {
	g.Meta `path:"/report/edit" method:"post" tags:"社区论坛举报表" summary:"修改/新增社区论坛举报表"`
	forumin.ReportEditInp
}

type EditRes struct{}

// DeleteReq 删除社区论坛举报表
type DeleteReq struct {
	g.Meta `path:"/report/delete" method:"post" tags:"社区论坛举报表" summary:"删除社区论坛举报表"`
	forumin.ReportDeleteInp
}

type DeleteRes struct{}