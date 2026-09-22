// Package version
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package version

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/oauthin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询版本更新列表
type ListReq struct {
	g.Meta `path:"/version/list" method:"get" tags:"版本更新" summary:"获取版本更新列表"`
	oauthin.VersionListInp
}

type ListRes struct {
	form.PageRes
	List []*oauthin.VersionListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出版本更新列表
type ExportReq struct {
	g.Meta `path:"/version/export" method:"get" tags:"版本更新" summary:"导出版本更新列表"`
	oauthin.VersionListInp
}

type ExportRes struct{}

// ViewReq 获取版本更新指定信息
type ViewReq struct {
	g.Meta `path:"/version/view" method:"get" tags:"版本更新" summary:"获取版本更新指定信息"`
	oauthin.VersionViewInp
}

type ViewRes struct {
	*oauthin.VersionViewModel
}

// EditReq 修改/新增版本更新
type EditReq struct {
	g.Meta `path:"/version/edit" method:"post" tags:"版本更新" summary:"修改/新增版本更新"`
	oauthin.VersionEditInp
}

type EditRes struct{}

// DeleteReq 删除版本更新
type DeleteReq struct {
	g.Meta `path:"/version/delete" method:"post" tags:"版本更新" summary:"删除版本更新"`
	oauthin.VersionDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新版本更新状态
type StatusReq struct {
	g.Meta `path:"/version/status" method:"post" tags:"版本更新" summary:"更新版本更新状态"`
	oauthin.VersionStatusInp
}

type StatusRes struct{}