package im

import (
	"context"
	v1 "hotgo/api/api/im/v1"
	"hotgo/internal/service"
)

// ConversationList 获取会话列表
func (c *ControllerV1) ConversationList(ctx context.Context, req *v1.ConversationListReq) (res *v1.ConversationListRes, err error) {
	return service.ApiIm().ConversationList(ctx, req)
}

// MessageList 获取聊天记录
func (c *ControllerV1) MessageList(ctx context.Context, req *v1.MessageListReq) (res *v1.MessageListRes, err error) {
	return service.ApiIm().MessageList(ctx, req)
}

// AddFriend 添加好友
func (c *ControllerV1) AddFriend(ctx context.Context, req *v1.AddFriendReq) (res *v1.AddFriendRes, err error) {
	return service.ApiIm().AddFriend(ctx, req)
}

// CreateGroup 创建群组
func (c *ControllerV1) CreateGroup(ctx context.Context, req *v1.CreateGroupReq) (res *v1.CreateGroupRes, err error) {
	return service.ApiIm().CreateGroup(ctx, req)
}

// ContactList 获取联系人列表
func (c *ControllerV1) ContactList(ctx context.Context, req *v1.ContactListReq) (res *v1.ContactListRes, err error) {
	return service.ApiIm().ContactList(ctx, req)
}

// GroupList 获取群组列表
func (c *ControllerV1) GroupList(ctx context.Context, req *v1.GroupListReq) (res *v1.GroupListRes, err error) {
	return service.ApiIm().GroupList(ctx, req)
}

// OpenConversation 发起会话
func (c *ControllerV1) OpenConversation(ctx context.Context, req *v1.OpenConversationReq) (res *v1.OpenConversationRes, err error) {
	return service.ApiIm().OpenConversation(ctx, req)
}

// InviteGroup 邀请加入群聊
func (c *ControllerV1) InviteGroup(ctx context.Context, req *v1.InviteGroupReq) (res *v1.InviteGroupRes, err error) {
	return service.ApiIm().InviteGroup(ctx, req)
}

// GroupMemberList 获取群成员列表
func (c *ControllerV1) GroupMemberList(ctx context.Context, req *v1.GroupMemberListReq) (res *v1.GroupMemberListRes, err error) {
	return service.ApiIm().GroupMemberList(ctx, req)
}
