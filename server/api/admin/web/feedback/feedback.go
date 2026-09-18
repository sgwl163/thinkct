// Package feedback
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package feedback

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/webin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询留言管理列表
type ListReq struct {
	g.Meta `path:"/feedback/list" method:"get" tags:"留言管理" summary:"获取留言管理列表"`
	webin.FeedbackListInp
}

type ListRes struct {
	form.PageRes
	List []*webin.FeedbackListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出留言管理列表
type ExportReq struct {
	g.Meta `path:"/feedback/export" method:"get" tags:"留言管理" summary:"导出留言管理列表"`
	webin.FeedbackListInp
}

type ExportRes struct{}

// ViewReq 获取留言管理指定信息
type ViewReq struct {
	g.Meta `path:"/feedback/view" method:"get" tags:"留言管理" summary:"获取留言管理指定信息"`
	webin.FeedbackViewInp
}

type ViewRes struct {
	*webin.FeedbackViewModel
}

// EditReq 修改/新增留言管理
type EditReq struct {
	g.Meta `path:"/feedback/edit" method:"post" tags:"留言管理" summary:"修改/新增留言管理"`
	webin.FeedbackEditInp
}

type EditRes struct{}

// DeleteReq 删除留言管理
type DeleteReq struct {
	g.Meta `path:"/feedback/delete" method:"post" tags:"留言管理" summary:"删除留言管理"`
	webin.FeedbackDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新留言管理状态
type StatusReq struct {
	g.Meta `path:"/feedback/status" method:"post" tags:"留言管理" summary:"更新留言管理状态"`
	webin.FeedbackStatusInp
}

type StatusRes struct{}