package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ConversationListReq 获取会话列表请求
type ConversationListReq struct {
	g.Meta `path:"/im/conversation/list" method:"get" tags:"前台-即时通讯" summary:"获取会话列表"`
}

type ConversationListRes struct {
	List []*ConversationItem `json:"list" dc:"会话列表"`
}

type ConversationItem struct {
	Id             int64  `json:"id" dc:"会话ID"`
	ChatType       int    `json:"chatType" dc:"会话类型"`
	TargetId       int64  `json:"targetId" dc:"目标ID"`
	TargetName     string `json:"targetName" dc:"目标名称(昵称或群名)"`
	TargetAvatar   string `json:"targetAvatar" dc:"目标头像"`
	LastMsgId      int64  `json:"lastMsgId" dc:"最后一条消息ID"`
	LastMsgContent string `json:"lastMsgContent" dc:"最后一条消息内容"`
	LastMsgTime    string `json:"lastMsgTime" dc:"最后一条消息时间"`
	UnreadCount    int    `json:"unreadCount" dc:"未读数"`
	IsTop          int    `json:"isTop" dc:"是否置顶"`
	IsDisturb      int    `json:"isDisturb" dc:"是否免打扰"`
	CreatedAt      string `json:"createdAt" dc:"创建时间"` // 新增返回会话创建时间/最近时间
}

// MessageListReq 获取聊天记录请求
type MessageListReq struct {
	g.Meta   `path:"/im/message/list" method:"get" tags:"前台-即时通讯" summary:"获取聊天记录"`
	ChatType int   `json:"chatType" v:"required#会话类型不能为空" dc:"会话类型: 1=单聊, 2=群聊"`
	TargetId int64 `json:"targetId" v:"required#目标ID不能为空" dc:"目标ID"`
	Page     int   `json:"page" d:"1" dc:"页码"`
	PageSize int   `json:"pageSize" d:"20" dc:"每页数量"`
}

type MessageListRes struct {
	List  []*MessageItem `json:"list" dc:"消息列表"`
	Total int            `json:"total" dc:"总数量"`
}

type MessageItem struct {
	Id         int64  `json:"id" dc:"消息ID"`
	ChatType   int    `json:"chatType" dc:"聊天类型"`
	FromId     int64  `json:"fromId" dc:"发送者ID"`
	FromName   string `json:"fromName" dc:"发送者名称"`
	FromAvatar string `json:"fromAvatar" dc:"发送者头像"`
	ToId       int64  `json:"toId" dc:"接收者ID"`
	MsgType    int    `json:"msgType" dc:"消息类型"`
	Content    string `json:"content" dc:"消息内容"`
	CreatedAt  string `json:"createdAt" dc:"发送时间"` // 确保返回字符串时间格式
}

// AddFriendReq 添加好友请求
type AddFriendReq struct {
	g.Meta   `path:"/im/friend/add" method:"post" tags:"前台-即时通讯" summary:"添加好友"`
	FriendId int64  `json:"friendId" v:"required#好友ID不能为空" dc:"好友用户ID"`
	Remark   string `json:"remark" dc:"好友备注"`
}

type AddFriendRes struct{}

// CreateGroupReq 创建群组请求
type CreateGroupReq struct {
	g.Meta `path:"/im/group/create" method:"post" tags:"前台-即时通讯" summary:"创建群组"`
	Name   string `json:"name" v:"required#群名称不能为空" dc:"群名称"`
	Avatar string `json:"avatar" dc:"群头像"`
	Notice string `json:"notice" dc:"群公告"`
}

type CreateGroupRes struct {
	GroupId int64 `json:"groupId" dc:"群组ID"`
}

// ContactListReq 获取联系人列表请求
type ContactListReq struct {
	g.Meta `path:"/im/contact/list" method:"get" tags:"前台-即时通讯" summary:"获取联系人列表"`
}

// ContactListRes 获取联系人列表响应
type ContactListRes struct {
	List []*ContactItem `json:"list" dc:"联系人列表"`
}

type ContactItem struct {
	FriendId int64  `json:"friendId" dc:"好友ID"`
	Nickname string `json:"nickname" dc:"好友昵称"`
	Avatar   string `json:"avatar" dc:"好友头像"`
	Remark   string `json:"remark" dc:"备注名"`
	Status   int    `json:"status" dc:"状态"`
}

// GroupListReq 获取群组列表请求
type GroupListReq struct {
	g.Meta `path:"/im/group/list" method:"get" tags:"前台-即时通讯" summary:"获取群组列表"`
}

type GroupListRes struct {
	List []*GroupItem `json:"list" dc:"群组列表"`
}

type GroupItem struct {
	GroupId int64  `json:"groupId" dc:"群组ID"`
	Name    string `json:"name" dc:"群名称"`
	Avatar  string `json:"avatar" dc:"群头像"`
	Role    int    `json:"role" dc:"我的角色:1群主,2管理,3成员"`
}

// OpenConversationReq 发起/打开会话请求
type OpenConversationReq struct {
	g.Meta   `path:"/im/conversation/open" method:"post" tags:"前台-即时通讯" summary:"发起会话"`
	ChatType int   `json:"chatType" v:"required#会话类型不能为空" dc:"会话类型: 1=单聊, 2=群聊"`
	TargetId int64 `json:"targetId" v:"required#目标ID不能为空" dc:"目标ID"`
}

type OpenConversationRes struct {
	ConversationId int64 `json:"conversationId" dc:"会话ID"`
}

// InviteGroupReq 邀请加入群聊请求
type InviteGroupReq struct {
	g.Meta  `path:"/im/group/invite" method:"post" tags:"前台-即时通讯" summary:"邀请加入群聊"`
	GroupId int64   `json:"groupId" v:"required#群组ID不能为空" dc:"群组ID"`
	UserIds []int64 `json:"userIds" v:"required#被邀请用户ID不能为空" dc:"被邀请的用户ID列表"`
}

type InviteGroupRes struct{}

// GroupMemberListReq 获取群成员列表请求
type GroupMemberListReq struct {
	g.Meta  `path:"/im/group/member/list" method:"get" tags:"前台-即时通讯" summary:"获取群成员列表"`
	GroupId int64 `json:"groupId" v:"required#群组ID不能为空" dc:"群组ID"`
}

type GroupMemberListRes struct {
	List []*GroupMemberItem `json:"list" dc:"群成员列表"`
}

type GroupMemberItem struct {
	UserId   int64  `json:"userId" dc:"用户ID"`
	Nickname string `json:"nickname" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
	Role     int    `json:"role" dc:"角色:1群主,2管理,3成员"`
	IsMute   int    `json:"isMute" dc:"是否禁言"`
}
