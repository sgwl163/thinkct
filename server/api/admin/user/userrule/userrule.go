// Package userrule
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package userrule

import (
	"hotgo/internal/model/input/form"
	"hotgo/internal/model/input/userin"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 查询权限规则列表
type ListReq struct {
	g.Meta `path:"/userRule/list" method:"get" tags:"权限规则" summary:"获取权限规则列表"`
	userin.UserRuleListInp
}

type ListRes struct {
	form.PageRes
	List []*userin.UserRuleListModel `json:"list"   dc:"数据列表"`
}

// ExportReq 导出权限规则列表
type ExportReq struct {
	g.Meta `path:"/userRule/export" method:"get" tags:"权限规则" summary:"导出权限规则列表"`
	userin.UserRuleListInp
}

type ExportRes struct{}

// ViewReq 获取权限规则指定信息
type ViewReq struct {
	g.Meta `path:"/userRule/view" method:"get" tags:"权限规则" summary:"获取权限规则指定信息"`
	userin.UserRuleViewInp
}

type ViewRes struct {
	*userin.UserRuleViewModel
}

// EditReq 修改/新增权限规则
type EditReq struct {
	g.Meta `path:"/userRule/edit" method:"post" tags:"权限规则" summary:"修改/新增权限规则"`
	userin.UserRuleEditInp
}

type EditRes struct{}

// DeleteReq 删除权限规则
type DeleteReq struct {
	g.Meta `path:"/userRule/delete" method:"post" tags:"权限规则" summary:"删除权限规则"`
	userin.UserRuleDeleteInp
}

type DeleteRes struct{}