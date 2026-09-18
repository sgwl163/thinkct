// Package imin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package imin

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MessageUpdateFields 修改IM-聊天记录表字段过滤
type MessageUpdateFields struct {
	ChatType int    `json:"chatType" dc:"聊天类型: 1=单聊, 2=群聊, 3=系统消息"`
	FromId   int64  `json:"fromId"   dc:"发送者ID"`
	ToId     int64  `json:"toId"     dc:"接收者ID(目标用户ID或群组ID)"`
	MsgType  int    `json:"msgType"  dc:"消息类型: 1=文本, 2=图片, 3=文件, 4=语音, 5=视频, 6=表情包"`
	Content  string `json:"content"  dc:"消息内容(文本内容或文件URL/JSON)"`
	Status   int    `json:"status"   dc:"状态: 1=正常, 2=已撤回, 3=已删除"`
}

// MessageInsertFields 新增IM-聊天记录表字段过滤
type MessageInsertFields struct {
	ChatType int    `json:"chatType" dc:"聊天类型: 1=单聊, 2=群聊, 3=系统消息"`
	FromId   int64  `json:"fromId"   dc:"发送者ID"`
	ToId     int64  `json:"toId"     dc:"接收者ID(目标用户ID或群组ID)"`
	MsgType  int    `json:"msgType"  dc:"消息类型: 1=文本, 2=图片, 3=文件, 4=语音, 5=视频, 6=表情包"`
	Content  string `json:"content"  dc:"消息内容(文本内容或文件URL/JSON)"`
	Status   int    `json:"status"   dc:"状态: 1=正常, 2=已撤回, 3=已删除"`
}

// MessageEditInp 修改/新增IM-聊天记录表
type MessageEditInp struct {
	entity.ImMessage
}

func (in *MessageEditInp) Filter(ctx context.Context) (err error) {
	// 验证聊天类型: 1=单聊, 2=群聊, 3=系统消息
	if err := g.Validator().Rules("required").Data(in.ChatType).Messages("聊天类型: 1=单聊, 2=群聊, 3=系统消息不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证发送者ID
	if err := g.Validator().Rules("required").Data(in.FromId).Messages("发送者ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证接收者ID(目标用户ID或群组ID)
	if err := g.Validator().Rules("required").Data(in.ToId).Messages("接收者ID(目标用户ID或群组ID)不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证消息类型: 1=文本, 2=图片, 3=文件, 4=语音, 5=视频, 6=表情包
	if err := g.Validator().Rules("required").Data(in.MsgType).Messages("消息类型: 1=文本, 2=图片, 3=文件, 4=语音, 5=视频, 6=表情包不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证状态: 1=正常, 2=已撤回, 3=已删除
	if err := g.Validator().Rules("required").Data(in.Status).Messages("状态: 1=正常, 2=已撤回, 3=已删除不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	if err := g.Validator().Rules("in:1,2").Data(in.Status).Messages("状态: 1=正常, 2=已撤回, 3=已删除值不正确").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type MessageEditModel struct{}

// MessageDeleteInp 删除IM-聊天记录表
type MessageDeleteInp struct {
	Id interface{} `json:"id" v:"required#消息ID不能为空" dc:"消息ID"`
}

func (in *MessageDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type MessageDeleteModel struct{}

// MessageViewInp 获取指定IM-聊天记录表信息
type MessageViewInp struct {
	Id int64 `json:"id" v:"required#消息ID不能为空" dc:"消息ID"`
}

func (in *MessageViewInp) Filter(ctx context.Context) (err error) {
	return
}

type MessageViewModel struct {
	entity.ImMessage
}

// MessageListInp 获取IM-聊天记录表列表
type MessageListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"消息ID"`
	Status    int           `json:"status"    dc:"状态: 1=正常, 2=已撤回, 3=已删除"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"发送时间"`
}

func (in *MessageListInp) Filter(ctx context.Context) (err error) {
	return
}

type MessageListModel struct {
	Id        int64       `json:"id"        dc:"消息ID"`
	ChatType  int         `json:"chatType"  dc:"聊天类型: 1=单聊, 2=群聊, 3=系统消息"`
	FromId    int64       `json:"fromId"    dc:"发送者ID"`
	ToId      int64       `json:"toId"      dc:"接收者ID(目标用户ID或群组ID)"`
	MsgType   int         `json:"msgType"   dc:"消息类型: 1=文本, 2=图片, 3=文件, 4=语音, 5=视频, 6=表情包"`
	Status    int         `json:"status"    dc:"状态: 1=正常, 2=已撤回, 3=已删除"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"发送时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间(如撤回时间)"`
}

// MessageExportModel 导出IM-聊天记录表
type MessageExportModel struct {
	Id        int64       `json:"id"        dc:"消息ID"`
	ChatType  int         `json:"chatType"  dc:"聊天类型: 1=单聊, 2=群聊, 3=系统消息"`
	FromId    int64       `json:"fromId"    dc:"发送者ID"`
	ToId      int64       `json:"toId"      dc:"接收者ID(目标用户ID或群组ID)"`
	MsgType   int         `json:"msgType"   dc:"消息类型: 1=文本, 2=图片, 3=文件, 4=语音, 5=视频, 6=表情包"`
	Status    int         `json:"status"    dc:"状态: 1=正常, 2=已撤回, 3=已删除"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"发送时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"更新时间(如撤回时间)"`
}

// MessageStatusInp 更新IM-聊天记录表状态
type MessageStatusInp struct {
	Id     int64 `json:"id" v:"required#消息ID不能为空" dc:"消息ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *MessageStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("消息ID不能为空")
		return
	}

	if in.Status <= 0 {
		err = gerror.New("状态不能为空")
		return
	}

	if !validate.InSlice(consts.StatusSlice, in.Status) {
		err = gerror.New("状态不正确")
		return
	}
	return
}

type MessageStatusModel struct{}