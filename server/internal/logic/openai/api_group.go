// Package openai
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package openai

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/openaiin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sOpenaiApiGroup struct{}

func NewOpenaiApiGroup() *sOpenaiApiGroup {
	return &sOpenaiApiGroup{}
}

func init() {
	service.RegisterOpenaiApiGroup(NewOpenaiApiGroup())
}

// Model AI接口分组ORM模型
func (s *sOpenaiApiGroup) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.OpenaiApiGroup.Ctx(ctx), option...)
}

// List 获取AI接口分组列表
func (s *sOpenaiApiGroup) List(ctx context.Context, in *openaiin.ApiGroupListInp) (list []*openaiin.ApiGroupListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(openaiin.ApiGroupListModel{})

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.OpenaiApiGroup.Columns().Id, in.Id)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.OpenaiApiGroup.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.OpenaiApiGroup.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderAsc(dao.OpenaiApiGroup.Columns().Sort).OrderDesc(dao.OpenaiApiGroup.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取AI接口分组列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出AI接口分组
func (s *sOpenaiApiGroup) Export(ctx context.Context, in *openaiin.ApiGroupListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(openaiin.ApiGroupExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出AI接口分组-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []openaiin.ApiGroupExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增AI接口分组
func (s *sOpenaiApiGroup) Edit(ctx context.Context, in *openaiin.ApiGroupEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(openaiin.ApiGroupUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改AI接口分组失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(openaiin.ApiGroupInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增AI接口分组失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除AI接口分组
func (s *sOpenaiApiGroup) Delete(ctx context.Context, in *openaiin.ApiGroupDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除AI接口分组失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取AI接口分组最大排序
func (s *sOpenaiApiGroup) MaxSort(ctx context.Context, in *openaiin.ApiGroupMaxSortInp) (res *openaiin.ApiGroupMaxSortModel, err error) {
	if err = dao.OpenaiApiGroup.Ctx(ctx).Fields(dao.OpenaiApiGroup.Columns().Sort).OrderDesc(dao.OpenaiApiGroup.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取AI接口分组最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(openaiin.ApiGroupMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取AI接口分组指定信息
func (s *sOpenaiApiGroup) View(ctx context.Context, in *openaiin.ApiGroupViewInp) (res *openaiin.ApiGroupViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取AI接口分组信息，请稍后重试！")
		return
	}
	return
}

// Status 更新AI接口分组状态
func (s *sOpenaiApiGroup) Status(ctx context.Context, in *openaiin.ApiGroupStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.OpenaiApiGroup.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新AI接口分组状态失败，请稍后重试！")
		return
	}
	return
}