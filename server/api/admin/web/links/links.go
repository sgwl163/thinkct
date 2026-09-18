// Package links
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package links

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/webin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询友情链接列表
type ListReq struct {
	g.Meta `path:"/links/list" method:"get" tags:"友情链接" summary:"获取友情链接列表"`
	webin.LinksListInp
}

type ListRes struct {
	form.PageRes
	List []*webin.LinksListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出友情链接列表
type ExportReq struct {
	g.Meta `path:"/links/export" method:"get" tags:"友情链接" summary:"导出友情链接列表"`
	webin.LinksListInp
}

type ExportRes struct{}

// ViewReq 获取友情链接指定信息
type ViewReq struct {
	g.Meta `path:"/links/view" method:"get" tags:"友情链接" summary:"获取友情链接指定信息"`
	webin.LinksViewInp
}

type ViewRes struct {
	*webin.LinksViewModel
}

// EditReq 修改/新增友情链接
type EditReq struct {
	g.Meta `path:"/links/edit" method:"post" tags:"友情链接" summary:"修改/新增友情链接"`
	webin.LinksEditInp
}

type EditRes struct{}

// DeleteReq 删除友情链接
type DeleteReq struct {
	g.Meta `path:"/links/delete" method:"post" tags:"友情链接" summary:"删除友情链接"`
	webin.LinksDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新友情链接状态
type StatusReq struct {
	g.Meta `path:"/links/status" method:"post" tags:"友情链接" summary:"更新友情链接状态"`
	webin.LinksStatusInp
}

type StatusRes struct{}