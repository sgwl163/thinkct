// Package forumin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package forumin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ReportUpdateFields 修改社区论坛举报表字段过滤
type ReportUpdateFields struct {
	TargetType  string      `json:"targetType"  dc:"举报对象类型:topic,comment"`
	TargetId    int64       `json:"targetId"    dc:"举报对象ID"`
	UserId      int64       `json:"userId"      dc:"举报用户ID"`
	Reason      string      `json:"reason"      dc:"举报原因"`
	Description string      `json:"description" dc:"补充说明"`
	Status      string      `json:"status"      dc:"处理状态:pending,processed,rejected"`
	HandledBy   int64       `json:"handledBy"   dc:"处理管理员ID"`
	HandledAt   *gtime.Time `json:"handledAt"   dc:"处理时间"`
}

// ReportInsertFields 新增社区论坛举报表字段过滤
type ReportInsertFields struct {
	TargetType  string      `json:"targetType"  dc:"举报对象类型:topic,comment"`
	TargetId    int64       `json:"targetId"    dc:"举报对象ID"`
	UserId      int64       `json:"userId"      dc:"举报用户ID"`
	Reason      string      `json:"reason"      dc:"举报原因"`
	Description string      `json:"description" dc:"补充说明"`
	Status      string      `json:"status"      dc:"处理状态:pending,processed,rejected"`
	HandledBy   int64       `json:"handledBy"   dc:"处理管理员ID"`
	HandledAt   *gtime.Time `json:"handledAt"   dc:"处理时间"`
}

// ReportEditInp 修改/新增社区论坛举报表
type ReportEditInp struct {
	entity.ForumReport
}

func (in *ReportEditInp) Filter(ctx context.Context) (err error) {
	// 验证举报对象类型:topic,comment
	if err := g.Validator().Rules("required").Data(in.TargetType).Messages("举报对象类型:topic,comment不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证举报对象ID
	if err := g.Validator().Rules("required").Data(in.TargetId).Messages("举报对象ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证举报用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("举报用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证举报原因
	if err := g.Validator().Rules("required").Data(in.Reason).Messages("举报原因不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证补充说明
	if err := g.Validator().Rules("required").Data(in.Description).Messages("补充说明不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证处理状态:pending,processed,rejected
	if err := g.Validator().Rules("required").Data(in.Status).Messages("处理状态:pending,processed,rejected不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证处理管理员ID
	if err := g.Validator().Rules("required").Data(in.HandledBy).Messages("处理管理员ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type ReportEditModel struct{}

// ReportDeleteInp 删除社区论坛举报表
type ReportDeleteInp struct {
	Id interface{} `json:"id" v:"required#举报ID不能为空" dc:"举报ID"`
}

func (in *ReportDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ReportDeleteModel struct{}

// ReportViewInp 获取指定社区论坛举报表信息
type ReportViewInp struct {
	Id int64 `json:"id" v:"required#举报ID不能为空" dc:"举报ID"`
}

func (in *ReportViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ReportViewModel struct {
	entity.ForumReport
}

// ReportListInp 获取社区论坛举报表列表
type ReportListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"举报ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *ReportListInp) Filter(ctx context.Context) (err error) {
	return
}

type ReportListModel struct {
	Id         int64       `json:"id"         dc:"举报ID"`
	TargetType string      `json:"targetType" dc:"举报对象类型:topic,comment"`
	TargetId   int64       `json:"targetId"   dc:"举报对象ID"`
	UserId     int64       `json:"userId"     dc:"举报用户ID"`
	Reason     string      `json:"reason"     dc:"举报原因"`
	Status     string      `json:"status"     dc:"处理状态:pending,processed,rejected"`
	HandledBy  int64       `json:"handledBy"  dc:"处理管理员ID"`
	HandledAt  *gtime.Time `json:"handledAt"  dc:"处理时间"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
}

// ReportExportModel 导出社区论坛举报表
type ReportExportModel struct {
	Id         int64       `json:"id"         dc:"举报ID"`
	TargetType string      `json:"targetType" dc:"举报对象类型:topic,comment"`
	TargetId   int64       `json:"targetId"   dc:"举报对象ID"`
	UserId     int64       `json:"userId"     dc:"举报用户ID"`
	Reason     string      `json:"reason"     dc:"举报原因"`
	Status     string      `json:"status"     dc:"处理状态:pending,processed,rejected"`
	HandledBy  int64       `json:"handledBy"  dc:"处理管理员ID"`
	HandledAt  *gtime.Time `json:"handledAt"  dc:"处理时间"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
}