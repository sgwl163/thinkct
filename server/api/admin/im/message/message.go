// Package message
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package message

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/imin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询IM-聊天记录表列表
type ListReq struct {
	g.Meta `path:"/message/list" method:"get" tags:"IM-聊天记录表" summary:"获取IM-聊天记录表列表"`
	imin.MessageListInp
}

type ListRes struct {
	form.PageRes
	List []*imin.MessageListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出IM-聊天记录表列表
type ExportReq struct {
	g.Meta `path:"/message/export" method:"get" tags:"IM-聊天记录表" summary:"导出IM-聊天记录表列表"`
	imin.MessageListInp
}

type ExportRes struct{}

// ViewReq 获取IM-聊天记录表指定信息
type ViewReq struct {
	g.Meta `path:"/message/view" method:"get" tags:"IM-聊天记录表" summary:"获取IM-聊天记录表指定信息"`
	imin.MessageViewInp
}

type ViewRes struct {
	*imin.MessageViewModel
}

// EditReq 修改/新增IM-聊天记录表
type EditReq struct {
	g.Meta `path:"/message/edit" method:"post" tags:"IM-聊天记录表" summary:"修改/新增IM-聊天记录表"`
	imin.MessageEditInp
}

type EditRes struct{}

// DeleteReq 删除IM-聊天记录表
type DeleteReq struct {
	g.Meta `path:"/message/delete" method:"post" tags:"IM-聊天记录表" summary:"删除IM-聊天记录表"`
	imin.MessageDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新IM-聊天记录表状态
type StatusReq struct {
	g.Meta `path:"/message/status" method:"post" tags:"IM-聊天记录表" summary:"更新IM-聊天记录表状态"`
	imin.MessageStatusInp
}

type StatusRes struct{}