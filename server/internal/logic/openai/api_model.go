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
	"hotgo/internal/library/hgorm"
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

type sOpenaiApiModel struct{}

func NewOpenaiApiModel() *sOpenaiApiModel {
	return &sOpenaiApiModel{}
}

func init() {
	service.RegisterOpenaiApiModel(NewOpenaiApiModel())
}

// Model AI模型列表ORM模型
func (s *sOpenaiApiModel) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.OpenaiApiModel.Ctx(ctx), option...)
}

// List 获取AI模型列表列表
func (s *sOpenaiApiModel) List(ctx context.Context, in *openaiin.ApiModelListInp) (list []*openaiin.ApiModelListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.FieldsPrefix(dao.OpenaiApiModel.Table(), openaiin.ApiModelListModel{})
	joinFields := hgorm.JoinFields(ctx, openaiin.ApiModelListModel{}, &dao.OpenaiApiGroup, "openaiApiGroup")
	if joinFields != "" {
		mod = mod.Fields(joinFields)
	}

	// 关联表字段
	mod = mod.LeftJoinOnFields(dao.OpenaiApiGroup.Table(), dao.OpenaiApiModel.Columns().GroupId, "=", dao.OpenaiApiGroup.Columns().Id)

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.OpenaiApiModel.Columns().Id, in.Id)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.OpenaiApiModel.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.OpenaiApiModel.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderAsc(dao.OpenaiApiModel.Table() + "." + dao.OpenaiApiModel.Columns().Sort).OrderDesc(dao.OpenaiApiModel.Table() + "." + dao.OpenaiApiModel.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取AI模型列表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出AI模型列表
func (s *sOpenaiApiModel) Export(ctx context.Context, in *openaiin.ApiModelListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(openaiin.ApiModelExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出AI模型列表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []openaiin.ApiModelExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增AI模型列表
func (s *sOpenaiApiModel) Edit(ctx context.Context, in *openaiin.ApiModelEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(openaiin.ApiModelUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改AI模型列表失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(openaiin.ApiModelInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增AI模型列表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除AI模型列表
func (s *sOpenaiApiModel) Delete(ctx context.Context, in *openaiin.ApiModelDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除AI模型列表失败，请稍后重试！")
		return
	}
	return
}

// MaxSort 获取AI模型列表最大排序
func (s *sOpenaiApiModel) MaxSort(ctx context.Context, in *openaiin.ApiModelMaxSortInp) (res *openaiin.ApiModelMaxSortModel, err error) {
	if err = dao.OpenaiApiModel.Ctx(ctx).Fields(dao.OpenaiApiModel.Columns().Sort).OrderDesc(dao.OpenaiApiModel.Columns().Sort).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取AI模型列表最大排序，请稍后重试！")
		return
	}

	if res == nil {
		res = new(openaiin.ApiModelMaxSortModel)
	}

	res.Sort = form.DefaultMaxSort(res.Sort)
	return
}

// View 获取AI模型列表指定信息
func (s *sOpenaiApiModel) View(ctx context.Context, in *openaiin.ApiModelViewInp) (res *openaiin.ApiModelViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取AI模型列表信息，请稍后重试！")
		return
	}
	return
}

// Status 更新AI模型列表状态
func (s *sOpenaiApiModel) Status(ctx context.Context, in *openaiin.ApiModelStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.OpenaiApiModel.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新AI模型列表状态失败，请稍后重试！")
		return
	}
	return
}
