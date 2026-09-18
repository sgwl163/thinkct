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

type sOpenaiAppCategory struct{}

func NewOpenaiAppCategory() *sOpenaiAppCategory {
	return &sOpenaiAppCategory{}
}

func init() {
	service.RegisterOpenaiAppCategory(NewOpenaiAppCategory())
}

// Model AI应用分类ORM模型
func (s *sOpenaiAppCategory) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.OpenaiAppCategory.Ctx(ctx), option...)
}

// List 获取AI应用分类列表
func (s *sOpenaiAppCategory) List(ctx context.Context, in *openaiin.OpenaiAppCategoryListInp) (list []*openaiin.OpenaiAppCategoryListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(openaiin.OpenaiAppCategoryListModel{})

	// 查询分类ID
	if in.Id > 0 {
		mod = mod.Where(dao.OpenaiAppCategory.Columns().Id, in.Id)
	}

	// 查询状态(1:启用, 2:禁用)
	if in.Status > 0 {
		mod = mod.Where(dao.OpenaiAppCategory.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.OpenaiAppCategory.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderAsc(dao.OpenaiAppCategory.Columns().Sort).OrderDesc(dao.OpenaiAppCategory.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取AI应用分类列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出AI应用分类
func (s *sOpenaiAppCategory) Export(ctx context.Context, in *openaiin.OpenaiAppCategoryListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(openaiin.OpenaiAppCategoryExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出AI应用分类-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []openaiin.OpenaiAppCategoryExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增AI应用分类
func (s *sOpenaiAppCategory) Edit(ctx context.Context, in *openaiin.OpenaiAppCategoryEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(openaiin.OpenaiAppCategoryUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改AI应用分类失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(openaiin.OpenaiAppCategoryInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增AI应用分类失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除AI应用分类
func (s *sOpenaiAppCategory) Delete(ctx context.Context, in *openaiin.OpenaiAppCategoryDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除AI应用分类失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取AI应用分类最大排序
func (s *sOpenaiAppCategory) MaxSort(ctx context.Context, in *openaiin.OpenaiAppCategoryMaxSortInp) (res *openaiin.OpenaiAppCategoryMaxSortModel, err error) {
	if err = dao.OpenaiAppCategory.Ctx(ctx).Fields(dao.OpenaiAppCategory.Columns().Sort).OrderDesc(dao.OpenaiAppCategory.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取AI应用分类最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(openaiin.OpenaiAppCategoryMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取AI应用分类指定信息
func (s *sOpenaiAppCategory) View(ctx context.Context, in *openaiin.OpenaiAppCategoryViewInp) (res *openaiin.OpenaiAppCategoryViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取AI应用分类信息，请稍后重试！")
		return
	}
	return
}

// Status 更新AI应用分类状态
func (s *sOpenaiAppCategory) Status(ctx context.Context, in *openaiin.OpenaiAppCategoryStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.OpenaiAppCategory.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新AI应用分类状态失败，请稍后重试！")
		return
	}
	return
}