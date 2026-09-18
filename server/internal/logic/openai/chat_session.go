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

type sOpenaiChatSession struct{}

func NewOpenaiChatSession() *sOpenaiChatSession {
	return &sOpenaiChatSession{}
}

func init() {
	service.RegisterOpenaiChatSession(NewOpenaiChatSession())
}

// Model 聊天会话ORM模型
func (s *sOpenaiChatSession) Model(ctx context.Context, option ...*handler.Option) *gdb.Model {
	return handler.Model(dao.OpenaiChatSession.Ctx(ctx), option...)
}

// List 获取聊天会话列表
func (s *sOpenaiChatSession) List(ctx context.Context, in *openaiin.ChatSessionListInp) (list []*openaiin.ChatSessionListModel, totalCount int, err error) {
	mod := s.Model(ctx)

	// 字段过滤
	mod = mod.Fields(openaiin.ChatSessionListModel{})

	// 查询ID
	if in.Id > 0 {
		mod = mod.Where(dao.OpenaiChatSession.Columns().Id, in.Id)
	}

	// 查询状态
	if in.Status > 0 {
		mod = mod.Where(dao.OpenaiChatSession.Columns().Status, in.Status)
	}

	// 查询创建时间
	if len(in.CreatedAt) == 2 {
		mod = mod.WhereBetween(dao.OpenaiChatSession.Columns().CreatedAt, in.CreatedAt[0], in.CreatedAt[1])
	}

	// 分页
	mod = mod.Page(in.Page, in.PerPage)

	// 排序
	mod = mod.OrderDesc(dao.OpenaiChatSession.Columns().Id)

	// 查询数据
	if err = mod.ScanAndCount(&list, &totalCount, false); err != nil {
		err = gerror.Wrap(err, "获取聊天会话列表失败，请稍后重试！")
		return
	}
	return
}

// Export 导出聊天会话
func (s *sOpenaiChatSession) Export(ctx context.Context, in *openaiin.ChatSessionListInp) (err error) {
	list, totalCount, err := s.List(ctx, in)
	if err != nil {
		return
	}

	// 字段的排序是依据tags的字段顺序，如果你不想使用默认的排序方式，可以直接定义 tags = []string{"字段名称", "字段名称2", ...}
	tags, err := convert.GetEntityDescTags(openaiin.ChatSessionExportModel{})
	if err != nil {
		return
	}

	var (
		fileName  = "导出聊天会话-" + gctx.CtxId(ctx)
		sheetName = fmt.Sprintf("索引条件共%v行,共%v页,当前导出是第%v页,本页共%v行", totalCount, form.CalPageCount(totalCount, in.PerPage), in.Page, len(list))
		exports   []openaiin.ChatSessionExportModel
	)

	if err = gconv.Scan(list, &exports); err != nil {
		return
	}

	err = excel.ExportByStructs(ctx, tags, exports, fileName, sheetName)
	return
}

// Edit 修改/新增聊天会话
func (s *sOpenaiChatSession) Edit(ctx context.Context, in *openaiin.ChatSessionEditInp) (err error) {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {

		// 修改
		if in.Id > 0 {
			if _, err = s.Model(ctx).
				Fields(openaiin.ChatSessionUpdateFields{}).
				WherePri(in.Id).Data(in).Update(); err != nil {
				err = gerror.Wrap(err, "修改聊天会话失败，请稍后重试！")
			}
			return
		}

		// 新增
		if _, err = s.Model(ctx, &handler.Option{FilterAuth: false}).
			Fields(openaiin.ChatSessionInsertFields{}).
			Data(in).OmitEmptyData().Insert(); err != nil {
			err = gerror.Wrap(err, "新增聊天会话失败，请稍后重试！")
		}
		return
	})
}

// Delete 删除聊天会话
func (s *sOpenaiChatSession) Delete(ctx context.Context, in *openaiin.ChatSessionDeleteInp) (err error) {

	if _, err = s.Model(ctx).WherePri(in.Id).Unscoped().Delete(); err != nil {
		err = gerror.Wrap(err, "删除聊天会话失败，请稍后重试！")
		return
	}
	return
}

// View 获取聊天会话指定信息
func (s *sOpenaiChatSession) View(ctx context.Context, in *openaiin.ChatSessionViewInp) (res *openaiin.ChatSessionViewModel, err error) {
	if err = s.Model(ctx).WherePri(in.Id).Scan(&res); err != nil {
		err = gerror.Wrap(err, "获取聊天会话信息，请稍后重试！")
		return
	}
	return
}

// Status 更新聊天会话状态
func (s *sOpenaiChatSession) Status(ctx context.Context, in *openaiin.ChatSessionStatusInp) (err error) {
	if _, err = s.Model(ctx).WherePri(in.Id).Data(g.Map{
		dao.OpenaiChatSession.Columns().Status: in.Status,
	}).Update(); err != nil {
		err = gerror.Wrap(err, "更新聊天会话状态失败，请稍后重试！")
		return
	}
	return
}