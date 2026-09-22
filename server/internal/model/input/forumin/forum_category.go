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

// ForumCategoryUpdateFields 修改社区板块字段过滤
type ForumCategoryUpdateFields struct {
	Pid         int64  `json:"pid"         dc:"上级板块ID"`
	Name        string `json:"name"        dc:"板块名称"`
	Description string `json:"description" dc:"板块简介"`
	Image       string `json:"image"       dc:"板块图标"`
	Sort        int    `json:"sort"        dc:"排序"`
	Status      int    `json:"status"      dc:"状态:1正常,2禁用"`
	TopicCount  int    `json:"topicCount"  dc:"主题数量"`
}

// ForumCategoryInsertFields 新增社区板块字段过滤
type ForumCategoryInsertFields struct {
	Pid         int64  `json:"pid"         dc:"上级板块ID"`
	Name        string `json:"name"        dc:"板块名称"`
	Description string `json:"description" dc:"板块简介"`
	Image       string `json:"image"       dc:"板块图标"`
	Sort        int    `json:"sort"        dc:"排序"`
	Status      int    `json:"status"      dc:"状态:1正常,2禁用"`
	TopicCount  int    `json:"topicCount"  dc:"主题数量"`
}

// ForumCategoryEditInp 修改/新增社区板块
type ForumCategoryEditInp struct {
	entity.ForumCategory
}

func (in *ForumCategoryEditInp) Filter(ctx context.Context) (err error) {
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

type ForumCategoryEditModel struct{}

// ForumCategoryDeleteInp 删除社区板块
type ForumCategoryDeleteInp struct {
	Id interface{} `json:"id" v:"required#板块ID不能为空" dc:"板块ID"`
}

func (in *ForumCategoryDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ForumCategoryDeleteModel struct{}

// ForumCategoryViewInp 获取指定社区板块信息
type ForumCategoryViewInp struct {
	Id int64 `json:"id" v:"required#板块ID不能为空" dc:"板块ID"`
}

func (in *ForumCategoryViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ForumCategoryViewModel struct {
	entity.ForumCategory
}

// ForumCategoryListInp 获取社区板块列表
type ForumCategoryListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"板块ID"`
	Status    int           `json:"status"    dc:"状态:1正常,2禁用"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *ForumCategoryListInp) Filter(ctx context.Context) (err error) {
	return
}

type ForumCategoryListModel struct {
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

// ForumCategoryExportModel 导出社区板块
type ForumCategoryExportModel struct {
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

// ForumCategoryMaxSortInp 获取社区板块最大排序
type ForumCategoryMaxSortInp struct{}

func (in *ForumCategoryMaxSortInp) Filter(ctx context.Context) (err error) {
	return
}

type ForumCategoryMaxSortModel struct {
	Sort int `json:"sort"  description:"排序"`
}

// ForumCategoryStatusInp 更新社区板块状态
type ForumCategoryStatusInp struct {
	Id     int64 `json:"id" v:"required#板块ID不能为空" dc:"板块ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *ForumCategoryStatusInp) Filter(ctx context.Context) (err error) {
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

type ForumCategoryStatusModel struct{}