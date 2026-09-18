package im

import (
	"context"
	v1 "hotgo/api/api/im/v1"
)

type IImV1 interface {
	ConversationList(ctx context.Context, req *v1.ConversationListReq) (res *v1.ConversationListRes, err error)
	MessageList(ctx context.Context, req *v1.MessageListReq) (res *v1.MessageListRes, err error)
	AddFriend(ctx context.Context, req *v1.AddFriendReq) (res *v1.AddFriendRes, err error)
	CreateGroup(ctx context.Context, req *v1.CreateGroupReq) (res *v1.CreateGroupRes, err error)
	ContactList(ctx context.Context, req *v1.ContactListReq) (res *v1.ContactListRes, err error)
	GroupList(ctx context.Context, req *v1.GroupListReq) (res *v1.GroupListRes, err error)
	OpenConversation(ctx context.Context, req *v1.OpenConversationReq) (res *v1.OpenConversationRes, err error)
	InviteGroup(ctx context.Context, req *v1.InviteGroupReq) (res *v1.InviteGroupRes, err error)
	GroupMemberList(ctx context.Context, req *v1.GroupMemberListReq) (res *v1.GroupMemberListRes, err error)
}
