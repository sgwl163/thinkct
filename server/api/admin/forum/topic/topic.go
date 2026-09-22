// Package topic
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package topic

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/forumin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询社区主题列表
type ListReq struct {
	g.Meta `path:"/topic/list" method:"get" tags:"社区主题" summary:"获取社区主题列表"`
	forumin.TopicListInp
}

type ListRes struct {
	form.PageRes
	List []*forumin.TopicListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出社区主题列表
type ExportReq struct {
	g.Meta `path:"/topic/export" method:"get" tags:"社区主题" summary:"导出社区主题列表"`
	forumin.TopicListInp
}

type ExportRes struct{}

// ViewReq 获取社区主题指定信息
type ViewReq struct {
	g.Meta `path:"/topic/view" method:"get" tags:"社区主题" summary:"获取社区主题指定信息"`
	forumin.TopicViewInp
}

type ViewRes struct {
	*forumin.TopicViewModel
}

// EditReq 修改/新增社区主题
type EditReq struct {
	g.Meta `path:"/topic/edit" method:"post" tags:"社区主题" summary:"修改/新增社区主题"`
	forumin.TopicEditInp
}

type EditRes struct{}

// DeleteReq 删除社区主题
type DeleteReq struct {
	g.Meta `path:"/topic/delete" method:"post" tags:"社区主题" summary:"删除社区主题"`
	forumin.TopicDeleteInp
}

type DeleteRes struct{}
