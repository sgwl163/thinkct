// Package conversation
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package conversation

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/imin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询会话列表管理列表
type ListReq struct {
	g.Meta `path:"/conversation/list" method:"get" tags:"会话列表管理" summary:"获取会话列表管理列表"`
	imin.ConversationListInp
}

type ListRes struct {
	form.PageRes
	List []*imin.ConversationListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出会话列表管理列表
type ExportReq struct {
	g.Meta `path:"/conversation/export" method:"get" tags:"会话列表管理" summary:"导出会话列表管理列表"`
	imin.ConversationListInp
}

type ExportRes struct{}

// ViewReq 获取会话列表管理指定信息
type ViewReq struct {
	g.Meta `path:"/conversation/view" method:"get" tags:"会话列表管理" summary:"获取会话列表管理指定信息"`
	imin.ConversationViewInp
}

type ViewRes struct {
	*imin.ConversationViewModel
}

// EditReq 修改/新增会话列表管理
type EditReq struct {
	g.Meta `path:"/conversation/edit" method:"post" tags:"会话列表管理" summary:"修改/新增会话列表管理"`
	imin.ConversationEditInp
}

type EditRes struct{}

// DeleteReq 删除会话列表管理
type DeleteReq struct {
	g.Meta `path:"/conversation/delete" method:"post" tags:"会话列表管理" summary:"删除会话列表管理"`
	imin.ConversationDeleteInp
}

type DeleteRes struct{}