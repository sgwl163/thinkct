// Package forumin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package forumin

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

// CategoryUpdateFields 修改社区板块字段过滤
type CategoryUpdateFields struct {
	Pid         int64  `json:"pid"         dc:"上级板块ID"`
	Name        string `json:"name"        dc:"板块名称"`
	Description string `json:"description" dc:"板块简介"`
	Image       string `json:"image"       dc:"板块图标"`
	Sort        int    `json:"sort"        dc:"排序"`
	Status      int    `json:"status"      dc:"状态:1正常,2禁用"`
	TopicCount  int    `json:"topicCount"  dc:"主题数量"`
}

// CategoryInsertFields 新增社区板块字段过滤
type CategoryInsertFields struct {
	Pid         int64  `json:"pid"         dc:"上级板块ID"`
	Name        string `json:"name"        dc:"板块名称"`
	Description string `json:"description" dc:"板块简介"`
	Image       string `json:"image"       dc:"板块图标"`
	Sort        int    `json:"sort"        dc:"排序"`
	Status      int    `json:"status"      dc:"状态:1正常,2禁用"`
	TopicCount  int    `json:"topicCount"  dc:"主题数量"`
}

// CategoryEditInp 修改/新增社区板块
type CategoryEditInp struct {
	entity.ForumCategory
}

func (in *CategoryEditInp) Filter(ctx context.Context) (err error) {
	// 验证上级板块ID
	if err := g.Validator().Rules("required").Data(in.Pid).Messages("上级板块ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证板块名称
	if err := g.Validator().Rules("required").Data(in.Name).Messages("板块名称不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证板块简介
	if err := g.Validator().Rules("required").Data(in.Description).Messages("板块简介不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证板块图标
	if err := g.Validator().Rules("required").Data(in.Image).Messages("板块图标不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证排序
	if err := g.Validator().Rules("required").Data(in.Sort).Messages("排序不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态:1正常,2禁用
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态:1正常,2禁用不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态:1正常,2禁用值不正确").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证主题数量
	if err := g.Validator().Rules("required").Data(in.TopicCount).Messages("主题数量不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type CategoryEditModel struct{}

// CategoryDeleteInp 删除社区板块
type CategoryDeleteInp struct {
	Id interface{} `json:"id" v:"required#板块ID不能为空" dc:"板块ID"`
}

func (in *CategoryDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type CategoryDeleteModel struct{}

// CategoryViewInp 获取指定社区板块信息
type CategoryViewInp struct {
	Id int64 `json:"id" v:"required#板块ID不能为空" dc:"板块ID"`
}

func (in *CategoryViewInp) Filter(ctx context.Context) (err error) {
	return
}

type CategoryViewModel struct {
	entity.ForumCategory
}

// CategoryListInp 获取社区板块列表
type CategoryListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"板块ID"`
	Status    int           `json:"status"    dc:"状态:1正常,2禁用"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *CategoryListInp) Filter(ctx context.Context) (err error) {
	return
}

type CategoryListModel struct {
	Id         int64       `json:"id"         dc:"板块ID"`
	Pid        int64       `json:"pid"        dc:"上级板块ID"`
	Name       string      `json:"name"       dc:"板块名称"`
	Image      string      `json:"image"      dc:"板块图标"`
	Sort       int         `json:"sort"       dc:"排序"`
	Status     int         `json:"status"     dc:"状态:1正常,2禁用"`
	TopicCount int         `json:"topicCount" dc:"主题数量"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"修改时间"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
}

// CategoryExportModel 导出社区板块
type CategoryExportModel struct {
	Id         int64       `json:"id"         dc:"板块ID"`
	Pid        int64       `json:"pid"        dc:"上级板块ID"`
	Name       string      `json:"name"       dc:"板块名称"`
	Image      string      `json:"image"      dc:"板块图标"`
	Sort       int         `json:"sort"       dc:"排序"`
	Status     int         `json:"status"     dc:"状态:1正常,2禁用"`
	TopicCount int         `json:"topicCount" dc:"主题数量"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  dc:"修改时间"`
	CreatedAt  *gtime.Time `json:"createdAt"  dc:"创建时间"`
}

// CategoryMaxSortInp 获取社区板块最大排序
type CategoryMaxSortInp struct{}

func (in *CategoryMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type CategoryMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// CategoryStatusInp 更新社区板块状态
type CategoryStatusInp struct {
	Id     int64 `json:"id" v:"required#板块ID不能为空" dc:"板块ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *CategoryStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("板块ID不能为空")
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

type CategoryStatusModel struct{}