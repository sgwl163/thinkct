// Package applog
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package applog

import (
	"hotgo/addons/appAuth/model/input/sysin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询升级记录列表
type ListReq struct {
	g.Meta `path:"/appLog/list" method:"get" tags:"升级记录" summary:"获取升级记录列表"`
	sysin.AppLogListInp
}

type ListRes struct {
	form.PageRes
	List []*sysin.AppLogListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出升级记录列表
type ExportReq struct {
	g.Meta `path:"/appLog/export" method:"get" tags:"升级记录" summary:"导出升级记录列表"`
	sysin.AppLogListInp
}

type ExportRes struct{}

// ViewReq 获取升级记录指定信息
type ViewReq struct {
	g.Meta `path:"/appLog/view" method:"get" tags:"升级记录" summary:"获取升级记录指定信息"`
	sysin.AppLogViewInp
}

type ViewRes struct {
	*sysin.AppLogViewModel
}

// EditReq 修改/新增升级记录
type EditReq struct {
	g.Meta `path:"/appLog/edit" method:"post" tags:"升级记录" summary:"修改/新增升级记录"`
	sysin.AppLogEditInp
}

type EditRes struct{}

// DeleteReq 删除升级记录
type DeleteReq struct {
	g.Meta `path:"/appLog/delete" method:"post" tags:"升级记录" summary:"删除升级记录"`
	sysin.AppLogDeleteInp
}

type DeleteRes struct{}