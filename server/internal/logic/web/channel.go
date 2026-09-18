// Package web
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package web

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/webin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"

	"hotgo/utility/tree"
)

type sWebChannel struct{}

func NewWebChannel() *sWebChannel {
	return &sWebChannel{}
}

func init() {
	service.RegisterWebChannel(NewWebChannel())
}

// Model 栏目管理ORM模型
func (s *sWebChannel) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.WebChannel.Ctx(ctx), option...)
}

// List 获取栏目管理列表
func (s *sWebChannel) List(ctx context.Context, in *webin.ChannelListInp) (list []*webin.ChannelListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(webin.ChannelListModel{})

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.WebChannel.Columns().Id, in.Id)
	}

	// 查询上级频道
	if in.Pid > 0 {
		mod = mod.Where(dao.WebChannel.Columns().Pid, in.Pid)
	}

	// 查询状态:0=禁用,1=启用
	if in.Status > 0 {
		mod = mod.Where(dao.WebChannel.Columns().Status, in.Status)
	}

	// 树形列表判断是否需要分页
	if in.Pagination {
		mod = mod.Page(in.Page, in.PerPage)
	}

	// 排序
	mod = mod.OrderDesc(dao.WebChannel.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取栏目管理列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出栏目管理
func (s *sWebChannel) Export(ctx context.Context, in *webin.ChannelListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(webin.ChannelExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出栏目管理-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []webin.ChannelExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增栏目管理
func (s *sWebChannel) Edit(ctx context.Context, in *webin.ChannelEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		var pid int64
		pid, in.Level, in.Tree, err = hgorm.AutoUpdateTree(ctx, &dao.WebChannel, int64(in.Id), int64(in.Pid))
		if err != nil {
			return err
		}
		in.Pid = uint(pid)
		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(webin.ChannelUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改栏目管理失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(webin.ChannelInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增栏目管理失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除栏目管理
func (s *sWebChannel) Delete(ctx context.Context, in *webin.ChannelDeleteInp) (err error) {
	count, err := dao.WebChannel.Ctx(ctx).Where(dao.WebChannel.Columns().Pid, in.Id).Count()
	if err != nil {
		err = gerror.Wrap(err, consts.ErrorORM)
		return err
	}
	if count > 0 {
		return gerror.New("请先删除该栏目管理下的所有下级！")
	}

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除栏目管理失败，请稍后重试！")
		return
	}
	return
}

// View 获取栏目管理指定信息
func (s *sWebChannel) View(ctx context.Context, in *webin.ChannelViewInp) (res *webin.ChannelViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取栏目管理信息，请稍后重试！")
		return
	}
	return
}

// Status 更新栏目管理状态
func (s *sWebChannel) Status(ctx context.Context, in *webin.ChannelStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.WebChannel.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新栏目管理状态失败，请稍后重试！")
		return
	}
	return
}

// TreeOption 获取栏目管理关系树选项
func (s *sWebChannel) TreeOption(ctx context.Context, modelName string) (nodes []tree.Node, err error) {
	var models []*webin.ChannelTreeOption
	mod := s.Model(ctx).Fields(webin.ChannelTreeOption{})
	
	if modelName != "" {
		mod = mod.Where(dao.WebChannel.Columns().Model, modelName)
	}

	if err = mod.OrderAsc(dao.WebChannel.Columns().Pid).OrderDesc(dao.WebChannel.Columns().Id).Scan(&models); err != nil {
		err = gerror.Wrap(err, "获取栏目管理关系树选项失败！")
		return
	}
	nodes = make([]tree.Node, len(models))
	for i, v := range models {
		nodes[i] = v
	}
	return tree.ListToTree(0, nodes)
}