// Package channel
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package channel

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/webin"
	"hotgo/utility/tree"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询栏目管理列表
type ListReq struct {
	g.Meta `path:"/channel/list" method:"get" tags:"栏目管理" summary:"获取栏目管理列表"`
	webin.ChannelListInp
}

type ListRes struct {
	form.PageRes
	List []*webin.ChannelListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出栏目管理列表
type ExportReq struct {
	g.Meta `path:"/channel/export" method:"get" tags:"栏目管理" summary:"导出栏目管理列表"`
	webin.ChannelListInp
}

type ExportRes struct{}

// ViewReq 获取栏目管理指定信息
type ViewReq struct {
	g.Meta `path:"/channel/view" method:"get" tags:"栏目管理" summary:"获取栏目管理指定信息"`
	webin.ChannelViewInp
}

type ViewRes struct {
	*webin.ChannelViewModel
}

// EditReq 修改/新增栏目管理
type EditReq struct {
	g.Meta `path:"/channel/edit" method:"post" tags:"栏目管理" summary:"修改/新增栏目管理"`
	webin.ChannelEditInp
}

type EditRes struct{}

// DeleteReq 删除栏目管理
type DeleteReq struct {
	g.Meta `path:"/channel/delete" method:"post" tags:"栏目管理" summary:"删除栏目管理"`
	webin.ChannelDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新栏目管理状态
type StatusReq struct {
	g.Meta `path:"/channel/status" method:"post" tags:"栏目管理" summary:"更新栏目管理状态"`
	webin.ChannelStatusInp
}

type StatusRes struct{}

// TreeOptionReq 获取栏目管理关系树选项
type TreeOptionReq struct {
	g.Meta `path:"/channel/treeOption" method:"get" tags:"栏目管理" summary:"获取栏目管理关系树选项"`
	Model  string `json:"model" dc:"模型名称"`
}

type TreeOptionRes []tree.Node