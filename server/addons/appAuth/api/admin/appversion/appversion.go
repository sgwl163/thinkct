// Package appversion
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package appversion

import (
	"hotgo/addons/appAuth/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询版本更新列表
type ListReq struct {
	g.Meta `path:"/appVersion/list" method:"get" tags:"版本更新" summary:"获取版本更新列表"`
	sysin.AppVersionListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.AppVersionListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出版本更新列表
type ExportReq struct {
	g.Meta `path:"/appVersion/export" method:"get" tags:"版本更新" summary:"导出版本更新列表"`
	sysin.AppVersionListInp
}

type ExportRes struct{}

// ViewReq 获取版本更新指定信息
type ViewReq struct {
	g.Meta `path:"/appVersion/view" method:"get" tags:"版本更新" summary:"获取版本更新指定信息"`
	sysin.AppVersionViewInp
}

type ViewRes struct {
	*sysin.AppVersionViewModel
}

// EditReq 修改/新增版本更新
type EditReq struct {
	g.Meta `path:"/appVersion/edit" method:"post" tags:"版本更新" summary:"修改/新增版本更新"`
	sysin.AppVersionEditInp
}

type EditRes struct{}

// DeleteReq 删除版本更新
type DeleteReq struct {
	g.Meta `path:"/appVersion/delete" method:"post" tags:"版本更新" summary:"删除版本更新"`
	sysin.AppVersionDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新版本更新状态
type StatusReq struct {
	g.Meta `path:"/appVersion/status" method:"post" tags:"版本更新" summary:"更新版本更新状态"`
	sysin.AppVersionStatusInp
}

type StatusRes struct{}