// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package sys

import (
	"context"
	"fmt"
	"hotgo/addons/appAuth/model/input/sysin"
	"hotgo/addons/appAuth/service"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysAppLog struct{}

func NewSysAppLog() *sSysAppLog {
	return &sSysAppLog{}
}

func init() {
	service.RegisterSysAppLog(NewSysAppLog())
}

// Model 升级记录ORM模型
func (s *sSysAppLog) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.AppLog.Ctx(ctx), option...)
}

// List 获取升级记录列表
func (s *sSysAppLog) List(ctx context.Context, in *sysin.AppLogListInp) (list []*sysin.AppLogListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(sysin.AppLogListModel{})

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.AppLog.Columns().Id, in.Id)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.AppLog.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.AppLog.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取升级记录列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出升级记录
func (s *sSysAppLog) Export(ctx context.Context, in *sysin.AppLogListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(sysin.AppLogExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出升级记录-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []sysin.AppLogExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增升级记录
func (s *sSysAppLog) Edit(ctx context.Context, in *sysin.AppLogEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(sysin.AppLogUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改升级记录失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(sysin.AppLogInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增升级记录失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除升级记录
func (s *sSysAppLog) Delete(ctx context.Context, in *sysin.AppLogDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除升级记录失败，请稍后重试！")
		return
	}
	return
}

// View 获取升级记录指定信息
func (s *sSysAppLog) View(ctx context.Context, in *sysin.AppLogViewInp) (res *sysin.AppLogViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取升级记录信息，请稍后重试！")
		return
	}
	return
}