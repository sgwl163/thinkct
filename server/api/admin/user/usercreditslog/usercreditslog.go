// Package usercreditslog
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package usercreditslog

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/userin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询资产明细列表
type ListReq struct {
	g.Meta `path:"/userCreditsLog/list" method:"get" tags:"资产明细" summary:"获取资产明细列表"`
	userin.UserCreditsLogListInp
}

type ListRes struct {
	form.PageRes
	List []*userin.UserCreditsLogListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出资产明细列表
type ExportReq struct {
	g.Meta `path:"/userCreditsLog/export" method:"get" tags:"资产明细" summary:"导出资产明细列表"`
	userin.UserCreditsLogListInp
}

type ExportRes struct{}

// ViewReq 获取资产明细指定信息
type ViewReq struct {
	g.Meta `path:"/userCreditsLog/view" method:"get" tags:"资产明细" summary:"获取资产明细指定信息"`
	userin.UserCreditsLogViewInp
}

type ViewRes struct {
	*userin.UserCreditsLogViewModel
}

// EditReq 修改/新增资产明细
type EditReq struct {
	g.Meta `path:"/userCreditsLog/edit" method:"post" tags:"资产明细" summary:"修改/新增资产明细"`
	userin.UserCreditsLogEditInp
}

type EditRes struct{}

// DeleteReq 删除资产明细
type DeleteReq struct {
	g.Meta `path:"/userCreditsLog/delete" method:"post" tags:"资产明细" summary:"删除资产明细"`
	userin.UserCreditsLogDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新资产明细状态
type StatusReq struct {
	g.Meta `path:"/userCreditsLog/status" method:"post" tags:"资产明细" summary:"更新资产明细状态"`
	userin.UserCreditsLogStatusInp
}

type StatusRes struct{}