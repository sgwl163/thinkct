// Package userin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package userin

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserCreditsLogUpdateFields 修改资产明细字段过滤
type UserCreditsLogUpdateFields struct {
	UserId     int64  `json:"userId"     dc:"用户ID"`
	CreditType string `json:"creditType" dc:"资产类型:money=余额,score=积分"`
	BeforeNum  int64  `json:"beforeNum"  dc:"变动前数量"`
	Num        int64  `json:"num"        dc:"变动数量"`
	AfterNum   int64  `json:"afterNum"   dc:"变动后数量"`
	Remark     string `json:"remark"     dc:"备注"`
	MapId      int64  `json:"mapId"      dc:"关联ID"`
	Status     int    `json:"status"     dc:"状态"`
}

// UserCreditsLogInsertFields 新增资产明细字段过滤
type UserCreditsLogInsertFields struct {
	UserId     int64  `json:"userId"     dc:"用户ID"`
	CreditType string `json:"creditType" dc:"资产类型:money=余额,score=积分"`
	BeforeNum  int64  `json:"beforeNum"  dc:"变动前数量"`
	Num        int64  `json:"num"        dc:"变动数量"`
	AfterNum   int64  `json:"afterNum"   dc:"变动后数量"`
	Remark     string `json:"remark"     dc:"备注"`
	MapId      int64  `json:"mapId"      dc:"关联ID"`
	Status     int    `json:"status"     dc:"状态"`
}

// UserCreditsLogEditInp 修改/新增资产明细
type UserCreditsLogEditInp struct {
	entity.UserCreditsLog
}

func (in *UserCreditsLogEditInp) Filter(ctx context.Context) (err error) {
	// 验证用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证资产类型:money=余额,score=积分
	if err := g.Validator().Rules("required").Data(in.CreditType).Messages("资产类型:money=余额,score=积分不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证变动前数量
	if err := g.Validator().Rules("required").Data(in.BeforeNum).Messages("变动前数量不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证变动数量
	if err := g.Validator().Rules("required").Data(in.Num).Messages("变动数量不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证变动后数量
	if err := g.Validator().Rules("required").Data(in.AfterNum).Messages("变动后数量不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证备注
	if err := g.Validator().Rules("required").Data(in.Remark).Messages("备注不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证关联ID
	if err := g.Validator().Rules("required").Data(in.MapId).Messages("关联ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态值不正确").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type UserCreditsLogEditModel struct{}

// UserCreditsLogDeleteInp 删除资产明细
type UserCreditsLogDeleteInp struct {
	Id interface{} `json:"id" v:"required#流水ID不能为空" dc:"流水ID"`
}

func (in *UserCreditsLogDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type UserCreditsLogDeleteModel struct{}

// UserCreditsLogViewInp 获取指定资产明细信息
type UserCreditsLogViewInp struct {
	Id int64 `json:"id" v:"required#流水ID不能为空" dc:"流水ID"`
}

func (in *UserCreditsLogViewInp) Filter(ctx context.Context) (err error) {
	return
}

type UserCreditsLogViewModel struct {
	entity.UserCreditsLog
}

// UserCreditsLogListInp 获取资产明细列表
type UserCreditsLogListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"流水ID"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *UserCreditsLogListInp) Filter(ctx context.Context) (err error) {
	return
}

type UserCreditsLogListModel struct {
	Id         int64       `json:"id"         dc:"流水ID"`
	UserId     int64       `json:"userId"     dc:"用户ID"`
	CreditType string      `json:"creditType" dc:"资产类型:money=余额,score=积分"`
	BeforeNum  int64       `json:"beforeNum"  dc:"变动前数量"`
	Num        int64       `json:"num"        dc:"变动数量"`
	AfterNum   int64       `json:"afterNum"   dc:"变动后数量"`
	Remark     string      `json:"remark"     dc:"备注"`
	MapId      int64       `json:"mapId"      dc:"关联ID"`
	Status     int         `json:"status"     dc:"状态"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

// UserCreditsLogExportModel 导出资产明细
type UserCreditsLogExportModel struct {
	Id         int64       `json:"id"         dc:"流水ID"`
	UserId     int64       `json:"userId"     dc:"用户ID"`
	CreditType string      `json:"creditType" dc:"资产类型:money=余额,score=积分"`
	BeforeNum  int64       `json:"beforeNum"  dc:"变动前数量"`
	Num        int64       `json:"num"        dc:"变动数量"`
	AfterNum   int64       `json:"afterNum"   dc:"变动后数量"`
	Remark     string      `json:"remark"     dc:"备注"`
	MapId      int64       `json:"mapId"      dc:"关联ID"`
	Status     int         `json:"status"     dc:"状态"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"更新时间"`
}

// UserCreditsLogStatusInp 更新资产明细状态
type UserCreditsLogStatusInp struct {
	Id     int64 `json:"id" v:"required#流水ID不能为空" dc:"流水ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *UserCreditsLogStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("流水ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type UserCreditsLogStatusModel struct{}