// Package im
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package im

import (
	"context"
	"fmt"
	"hotgo/internal/dao"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/imin"
	"hotgo/internal/service"
	"hotgo/utility/convert"
	"hotgo/utility/excel"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

type sImMessage struct{}

func NewImMessage() *sImMessage {
	return &sImMessage{}
}

func init() {
	service.RegisterImMessage(NewImMessage())
}

// Model IM-聊天记录表ORM模型
func (s *sImMessage) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.ImMessage.Ctx(ctx), option...)
}

// List 获取IM-聊天记录表列表
func (s *sImMessage) List(ctx context.Context, in *imin.MessageListInp) (list []*imin.MessageListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(imin.MessageListModel{})

	// 查询消息ID
	if in.Id > 0 {
		mod = mod.Where(dao.ImMessage.Columns().Id, in.Id)
	}

	// 查询状态: 1=正常, 2=已撤回, 3=已删除
	if in.Status > 0 {
		mod = mod.Where(dao.ImMessage.Columns().Status, in.Status)
	}

	// 查询发送时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.ImMessage.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.ImMessage.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取IM-聊天记录表列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出IM-聊天记录表
func (s *sImMessage) Export(ctx context.Context, in *imin.MessageListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(imin.MessageExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出IM-聊天记录表-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []imin.MessageExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增IM-聊天记录表
func (s *sImMessage) Edit(ctx context.Context, in *imin.MessageEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(imin.MessageUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改IM-聊天记录表失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(imin.MessageInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增IM-聊天记录表失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除IM-聊天记录表
func (s *sImMessage) Delete(ctx context.Context, in *imin.MessageDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除IM-聊天记录表失败，请稍后重试！")
		return
	}
	return
}

// View 获取IM-聊天记录表指定信息
func (s *sImMessage) View(ctx context.Context, in *imin.MessageViewInp) (res *imin.MessageViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取IM-聊天记录表信息，请稍后重试！")
		return
	}
	return
}

// Status 更新IM-聊天记录表状态
func (s *sImMessage) Status(ctx context.Context, in *imin.MessageStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.ImMessage.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新IM-聊天记录表状态失败，请稍后重试！")
		return
	}
	return
}