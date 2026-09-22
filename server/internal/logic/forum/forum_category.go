// Package forum
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package forum

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/forumin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sForumForumCategory struct{}

func NewForumForumCategory() *sForumForumCategory {
	return &sForumForumCategory{}
}

func init() {
	service.RegisterForumForumCategory(NewForumForumCategory())
}

// Model 社区板块ORM模型
func (s *sForumForumCategory) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.ForumCategory.Ctx(ctx), option...)
}

// List 获取社区板块列表
func (s *sForumForumCategory) List(ctx context.Context, in *forumin.ForumCategoryListInp) (list []*forumin.ForumCategoryListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(forumin.ForumCategoryListModel{})

	// 查询板块ID
	if in.Id > 0 {
		mod = mod.Where(dao.ForumCategory.Columns().Id, in.Id)
	}

	// 查询状态:1正常,2禁用
	if in.Status > 0 {
		mod = mod.Where(dao.ForumCategory.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.ForumCategory.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderAsc(dao.ForumCategory.Columns().Sort).OrderDesc(dao.ForumCategory.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取社区板块列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出社区板块
func (s *sForumForumCategory) Export(ctx context.Context, in *forumin.ForumCategoryListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(forumin.ForumCategoryExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出社区板块-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []forumin.ForumCategoryExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增社区板块
func (s *sForumForumCategory) Edit(ctx context.Context, in *forumin.ForumCategoryEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(forumin.ForumCategoryUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改社区板块失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(forumin.ForumCategoryInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增社区板块失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除社区板块
func (s *sForumForumCategory) Delete(ctx context.Context, in *forumin.ForumCategoryDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除社区板块失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取社区板块最大排序
func (s *sForumForumCategory) MaxSort(ctx context.Context, in *forumin.ForumCategoryMaxSortInp) (res *forumin.ForumCategoryMaxSortModel, err error) {
	if err = dao.ForumCategory.Ctx(ctx).Fields(dao.ForumCategory.Columns().Sort).OrderDesc(dao.ForumCategory.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取社区板块最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(forumin.ForumCategoryMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取社区板块指定信息
func (s *sForumForumCategory) View(ctx context.Context, in *forumin.ForumCategoryViewInp) (res *forumin.ForumCategoryViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取社区板块信息，请稍后重试！")
		return
	}
	return
}

// Status 更新社区板块状态
func (s *sForumForumCategory) Status(ctx context.Context, in *forumin.ForumCategoryStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.ForumCategory.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新社区板块状态失败，请稍后重试！")
		return
	}
	return
}
