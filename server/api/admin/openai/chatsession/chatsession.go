// Package chatsession
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package chatsession

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/openaiin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询聊天会话列表
type ListReq struct {
	g.Meta `path:"/chatSession/list" method:"get" tags:"聊天会话" summary:"获取聊天会话列表"`
	openaiin.ChatSessionListInp
}

type ListRes struct {
	form.PageRes
	List []*openaiin.ChatSessionListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出聊天会话列表
type ExportReq struct {
	g.Meta `path:"/chatSession/export" method:"get" tags:"聊天会话" summary:"导出聊天会话列表"`
	openaiin.ChatSessionListInp
}

type ExportRes struct{}

// ViewReq 获取聊天会话指定信息
type ViewReq struct {
	g.Meta `path:"/chatSession/view" method:"get" tags:"聊天会话" summary:"获取聊天会话指定信息"`
	openaiin.ChatSessionViewInp
}

type ViewRes struct {
	*openaiin.ChatSessionViewModel
}

// EditReq 修改/新增聊天会话
type EditReq struct {
	g.Meta `path:"/chatSession/edit" method:"post" tags:"聊天会话" summary:"修改/新增聊天会话"`
	openaiin.ChatSessionEditInp
}

type EditRes struct{}

// DeleteReq 删除聊天会话
type DeleteReq struct {
	g.Meta `path:"/chatSession/delete" method:"post" tags:"聊天会话" summary:"删除聊天会话"`
	openaiin.ChatSessionDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新聊天会话状态
type StatusReq struct {
	g.Meta `path:"/chatSession/status" method:"post" tags:"聊天会话" summary:"更新聊天会话状态"`
	openaiin.ChatSessionStatusInp
}

type StatusRes struct{}