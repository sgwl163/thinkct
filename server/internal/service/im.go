// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "hotgo/api/api/im/v1"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/imin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IApiIm interface {
		ConversationList(ctx context.Context, req *v1.ConversationListReq) (res *v1.ConversationListRes, err error)
		AddFriend(ctx context.Context, req *v1.AddFriendReq) (res *v1.AddFriendRes, err error)
		CreateGroup(ctx context.Context, req *v1.CreateGroupReq) (res *v1.CreateGroupRes, err error)
		ContactList(ctx context.Context, req *v1.ContactListReq) (res *v1.ContactListRes, err error)
		GroupList(ctx context.Context, req *v1.GroupListReq) (res *v1.GroupListRes, err error)
		OpenConversation(ctx context.Context, req *v1.OpenConversationReq) (res *v1.OpenConversationRes, err error)
		InviteGroup(ctx context.Context, req *v1.InviteGroupReq) (res *v1.InviteGroupRes, err error)
		GroupMemberList(ctx context.Context, req *v1.GroupMemberListReq) (res *v1.GroupMemberListRes, err error)
		MessageList(ctx context.Context, req *v1.MessageListReq) (res *v1.MessageListRes, err error)
	}
	IImGroup interface {
		// Model IM-群组表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取IM-群组表列表
		List(ctx context.Context, in *imin.GroupListInp) (list []*imin.GroupListModel, totalCount int, err error)
		// Export 导出IM-群组表
		Export(ctx context.Context, in *imin.GroupListInp) (err error)
		// Edit 修改/新增IM-群组表
		Edit(ctx context.Context, in *imin.GroupEditInp) (err error)
		// Delete 删除IM-群组表
		Delete(ctx context.Context, in *imin.GroupDeleteInp) (err error)
		// View 获取IM-群组表指定信息
		View(ctx context.Context, in *imin.GroupViewInp) (res *imin.GroupViewModel, err error)
		// Status 更新IM-群组表状态
		Status(ctx context.Context, in *imin.GroupStatusInp) (err error)
	}
	IImMessage interface {
		// Model IM-聊天记录表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取IM-聊天记录表列表
		List(ctx context.Context, in *imin.MessageListInp) (list []*imin.MessageListModel, totalCount int, err error)
		// Export 导出IM-聊天记录表
		Export(ctx context.Context, in *imin.MessageListInp) (err error)
		// Edit 修改/新增IM-聊天记录表
		Edit(ctx context.Context, in *imin.MessageEditInp) (err error)
		// Delete 删除IM-聊天记录表
		Delete(ctx context.Context, in *imin.MessageDeleteInp) (err error)
		// View 获取IM-聊天记录表指定信息
		View(ctx context.Context, in *imin.MessageViewInp) (res *imin.MessageViewModel, err error)
		// Status 更新IM-聊天记录表状态
		Status(ctx context.Context, in *imin.MessageStatusInp) (err error)
	}
)

var (
	localApiIm     IApiIm
	localImGroup   IImGroup
	localImMessage IImMessage
)

func ApiIm() IApiIm {
	if localApiIm == nil {
		panic("implement not found for interface IApiIm, forgot register?")
	}
	return localApiIm
}

func RegisterApiIm(i IApiIm) {
	localApiIm = i
}

func ImGroup() IImGroup {
	if localImGroup == nil {
		panic("implement not found for interface IImGroup, forgot register?")
	}
	return localImGroup
}

func RegisterImGroup(i IImGroup) {
	localImGroup = i
}

func ImMessage() IImMessage {
	if localImMessage == nil {
		panic("implement not found for interface IImMessage, forgot register?")
	}
	return localImMessage
}

func RegisterImMessage(i IImMessage) {
	localImMessage = i
}
