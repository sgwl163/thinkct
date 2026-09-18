// Package openaiin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package openaiin

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

// ChatSessionUpdateFields 修改聊天会话字段过滤
type ChatSessionUpdateFields struct {
	UserId   int64  `json:"userId" dc:"用户ID"`
	FolderId int64  `json:"folderId" dc:"文件夹ID"`
	Title    string `json:"title"    dc:"会话标题"`
	Prompt   string `json:"prompt"   dc:"系统提示词"`
	Model    string `json:"model"    dc:"使用的模型"`
	Status   int    `json:"status"   dc:"状态"`
}

// ChatSessionInsertFields 新增聊天会话字段过滤
type ChatSessionInsertFields struct {
	UserId   int64  `json:"userId" dc:"用户ID"`
	FolderId int64  `json:"folderId" dc:"文件夹ID"`
	Title    string `json:"title"    dc:"会话标题"`
	Prompt   string `json:"prompt"   dc:"系统提示词"`
	Model    string `json:"model"    dc:"使用的模型"`
	Status   int    `json:"status"   dc:"状态"`
}

// ChatSessionEditInp 修改/新增聊天会话
type ChatSessionEditInp struct {
	entity.OpenaiChatSession
}

func (in *ChatSessionEditInp) Filter(ctx context.Context) (err error) {
	// 验证用户ID
	if err := g.Validator().Rules("required").Data(in.UserId).Messages("用户ID不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	// 验证会话标题
	if err := g.Validator().Rules("required").Data(in.Title).Messages("会话标题不能为空").Run(ctx); err != nil {
		return err.Current()
	}
	return
}

type ChatSessionEditModel struct{}

// ChatSessionDeleteInp 删除聊天会话
type ChatSessionDeleteInp struct {
	Id interface{} `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *ChatSessionDeleteInp) Filter(ctx context.Context) (err error) {
	return
}

type ChatSessionDeleteModel struct{}

// ChatSessionViewInp 获取指定聊天会话信息
type ChatSessionViewInp struct {
	Id int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
}

func (in *ChatSessionViewInp) Filter(ctx context.Context) (err error) {
	return
}

type ChatSessionViewModel struct {
	entity.OpenaiChatSession
}

// ChatSessionListInp 获取聊天会话列表
type ChatSessionListInp struct {
	form.PageReq
	Id        int64         `json:"id"        dc:"ID"`
	Status    int           `json:"status"    dc:"状态"`
	CreatedAt []*gtime.Time `json:"createdAt" dc:"创建时间"`
}

func (in *ChatSessionListInp) Filter(ctx context.Context) (err error) {
	return
}

type ChatSessionListModel struct {
	Id        int64       `json:"id"        dc:"ID"`
	UserId    int64       `json:"userId"    dc:"用户ID"`
	FolderId  int64       `json:"folderId"  dc:"文件夹ID"`
	Title     string      `json:"title"     dc:"会话标题"`
	Model     string      `json:"model"     dc:"使用的模型"`
	Status    int         `json:"status"    dc:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// ChatSessionExportModel 导出聊天会话
type ChatSessionExportModel struct {
	Id        int64       `json:"id"        dc:"ID"`
	UserId    int64       `json:"userId"    dc:"用户ID"`
	FolderId  int64       `json:"folderId"  dc:"文件夹ID"`
	Title     string      `json:"title"     dc:"会话标题"`
	Model     string      `json:"model"     dc:"使用的模型"`
	Status    int         `json:"status"    dc:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" dc:"修改时间"`
}

// ChatSessionStatusInp 更新聊天会话状态
type ChatSessionStatusInp struct {
	Id     int64 `json:"id" v:"required#ID不能为空" dc:"ID"`
	Status int   `json:"status" dc:"状态"`
}

func (in *ChatSessionStatusInp) Filter(ctx context.Context) (err error) {
	if in.Id <= 0 {
		err = gerror.New("ID不能为空")
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

type ChatSessionStatusModel struct{}
