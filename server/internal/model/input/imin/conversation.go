// Package imin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package imin

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ConversationUpdateFields 修改会话列表管理字段过滤
type ConversationUpdateFields struct {
	UserId         int64       `json:"userId"         dc:"归属用户ID"`
	ChatType       int         `json:"chatType"       dc:"会话类型: 1=单聊, 2=群聊, 3=系统通知"`
	TargetId       int64       `json:"targetId"       dc:"目标ID(好友ID或群组ID)"`
	LastMsgId      int64       `json:"lastMsgId"      dc:"最后一条消息ID"`
	LastMsgContent string      `json:"lastMsgContent" dc:"最后消息内容摘要(用于列表展示)"`
	LastMsgTime    *gtime.Time `json:"lastMsgTime"    dc:"最后消息时间"`
	UnreadCount    int         `json:"unreadCount"    dc:"未读消息数"`
	IsTop          int         `json:"isTop"          dc:"是否置顶: 0=否, 1=是"`
	IsDisturb      int         `json:"isDisturb"      dc:"消息免打扰: 0=否, 1=是"`
}

// ConversationInsertFields 新增会话列表管理字段过滤
type ConversationInsertFields struct {
	UserId         int64       `json:"userId"         dc:"归属用户ID"`
	ChatType       int         `json:"chatType"       dc:"会话类型: 1=单聊, 2=群聊, 3=系统通知"`
	TargetId       int64       `json:"targetId"       dc:"目标ID(好友ID或群组ID)"`
	LastMsgId      int64       `json:"lastMsgId"      dc:"最后一条消息ID"`
	LastMsgContent string      `json:"lastMsgContent" dc:"最后消息内容摘要(用于列表展示)"`
	LastMsgTime    *gtime.Time `json:"lastMsgTime"    dc:"最后消息时间"`
	UnreadCount    int         `json:"unreadCount"    dc:"未读消息数"`
	IsTop          int         `json:"isTop"          dc:"是否置顶: 0=否, 1=是"`
	IsDisturb      int         `json:"isDisturb"      dc:"消息免打扰: 0=否, 1=是"`
}

// ConversationEditInp 修改/新增会话列表管理
type ConversationEditInp struct {
	entity.ImConversation
}

func (in *ConversationEditInp) Filter(ctx context.Context) (err error) {
	// 验证归属用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("归属用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证会话类型: 1=单聊, 2=群聊, 3=系统通知
	if err := g.Validator().Rules("required").Data(in.ChatType).Messages("会话类型: 1=单聊, 2=群聊, 3=系统通知不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证目标ID(好友ID或群组ID)
	if err := g.Validator().Rules("required").Data(in.TargetId).Messages("目标ID(好友ID或群组ID)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证未读消息数
	if err := g.Validator().Rules("required").Data(in.UnreadCount).Messages("未读消息数不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证是否置顶: 0=否, 1=是
	if err := g.Validator().Rules("required").Data(in.IsTop).Messages("是否置顶: 0=否, 1=是不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证消息免打扰: 0=否, 1=是
	if err := g.Validator().Rules("required").Data(in.IsDisturb).Messages("消息免打扰: 0=否, 1=是不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type ConversationEditModel struct{}

// ConversationDeleteInp 删除会话列表管理
type ConversationDeleteInp struct {
	Id interface{} `json:"id" v:"required#会话ID不能为空" dc:"会话ID"`
}

func (in *ConversationDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ConversationDeleteModel struct{}

// ConversationViewInp 获取指定会话列表管理信息
type ConversationViewInp struct {
	Id int64 `json:"id" v:"required#会话ID不能为空" dc:"会话ID"`
}

func (in *ConversationViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ConversationViewModel struct {
	entity.ImConversation
}

// ConversationListInp 获取会话列表管理列表
type ConversationListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"会话ID"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *ConversationListInp) Filter(ctx context.Context) (err error) {
	return
}

type ConversationListModel struct {
	Id          int64       `json:"id"          dc:"会话ID"`
	UserId      int64       `json:"userId"      dc:"归属用户ID"`
	ChatType    int         `json:"chatType"    dc:"会话类型: 1=单聊, 2=群聊, 3=系统通知"`
	TargetId    int64       `json:"targetId"    dc:"目标ID(好友ID或群组ID)"`
	LastMsgId   int64       `json:"lastMsgId"   dc:"最后一条消息ID"`
	LastMsgTime *gtime.Time `json:"lastMsgTime" dc:"最后消息时间"`
	UnreadCount int         `json:"unreadCount" dc:"未读消息数"`
	IsTop       int         `json:"isTop"       dc:"是否置顶: 0=否, 1=是"`
	IsDisturb   int         `json:"isDisturb"   dc:"消息免打扰: 0=否, 1=是"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"更新时间"`
}

// ConversationExportModel 导出会话列表管理
type ConversationExportModel struct {
	Id          int64       `json:"id"          dc:"会话ID"`
	UserId      int64       `json:"userId"      dc:"归属用户ID"`
	ChatType    int         `json:"chatType"    dc:"会话类型: 1=单聊, 2=群聊, 3=系统通知"`
	TargetId    int64       `json:"targetId"    dc:"目标ID(好友ID或群组ID)"`
	LastMsgId   int64       `json:"lastMsgId"   dc:"最后一条消息ID"`
	LastMsgTime *gtime.Time `json:"lastMsgTime" dc:"最后消息时间"`
	UnreadCount int         `json:"unreadCount" dc:"未读消息数"`
	IsTop       int         `json:"isTop"       dc:"是否置顶: 0=否, 1=是"`
	IsDisturb   int         `json:"isDisturb"   dc:"消息免打扰: 0=否, 1=是"`
	CreatedAt   *gtime.Time `json:"createdAt"   dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   dc:"更新时间"`
}