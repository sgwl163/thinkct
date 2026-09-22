// Package forum
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package forum

import (
	"context"
	"hotgo/api/admin/forum/report"
	"hotgo/internal/model/input/forumin"
	"hotgo/internal/service"
)

var (
	Report = cReport{}
)

type cReport struct{}

// List 查看社区论坛举报表列表
func (c *cReport) List(ctx context.Context, req *report.ListReq) (res *report.ListRes, err error) {
	list, totalCount, err := service.ForumReport().List(ctx, &req.ReportListInp)
	if err != nil {
		return
	}

	if list == nil {
		list = []*forumin.ReportListModel{}
	}

	res = new(report.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// Export 导出社区论坛举报表列表
func (c *cReport) Export(ctx context.Context, req *report.ExportReq) (res *report.ExportRes, err error) {
	err = service.ForumReport().Export(ctx, &req.ReportListInp)
	return
}

// Edit 更新社区论坛举报表
func (c *cReport) Edit(ctx context.Context, req *report.EditReq) (res *report.EditRes, err error) {
	err = service.ForumReport().Edit(ctx, &req.ReportEditInp)
	return
}

// View 获取指定社区论坛举报表信息
func (c *cReport) View(ctx context.Context, req *report.ViewReq) (res *report.ViewRes, err error) {
	data, err := service.ForumReport().View(ctx, &req.ReportViewInp)
	if err != nil {
		return
	}

	res = new(report.ViewRes)
	res.ReportViewModel = data
	return
}

// Delete 删除社区论坛举报表
func (c *cReport) Delete(ctx context.Context, req *report.DeleteReq) (res *report.DeleteRes, err error) {
	err = service.ForumReport().Delete(ctx, &req.ReportDeleteInp)
	return
}