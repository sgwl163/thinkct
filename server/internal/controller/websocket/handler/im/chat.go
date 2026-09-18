package im

import (
	"context"

	"hotgo/internal/websocket"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	Chat = cChat{}
)

type cChat struct{}

// SendMsg 发送消息
func (c *cChat) SendMsg(client *websocket.Client, req *websocket.WRequest) {
	ctx := client.Context()
	user := client.User

	if user == nil {
		websocket.SendError(client, req.Event, gerror.New("身份已失效，请重新登录"))
		return
	}

	chatType := gconv.Int(req.Data["chatType"])
	toId := gconv.Int64(req.Data["toId"])
	msgType := gconv.Int(req.Data["msgType"])
	content := gconv.String(req.Data["content"])

	if toId == 0 || content == "" {
		websocket.SendError(client, req.Event, gerror.New("参数错误：接收方或内容不能为空"))
		return
	}

	// 默认消息类型为文本
	if msgType == 0 {
		msgType = 1
	}

	now := gtime.Now()

	// 1. 保存消息到数据库
	msgId, err := g.DB().Model("hg_im_message").Data(g.Map{
		"chat_type":  chatType,
		"from_id":    user.Id,
		"to_id":      toId,
		"msg_type":   msgType,
		"content":    content,
		"status":     1,
		"created_at": now,
		"updated_at": now,
	}).InsertAndGetId()

	if err != nil {
		g.Log().Error(ctx, "保存消息失败", err)
		websocket.SendError(client, req.Event, gerror.New("发送失败，请稍后再试"))
		return
	}

	// 2. 更新会话列表的未读数和最后一条消息
	c.updateConversation(client, chatType, user.Id, toId, msgId, content, now)

	// 3. 查询发送者的昵称和头像 (为了给接收方展示)
	var u struct {
		Nickname string
		Username string
		Avatar   string
	}
	_ = g.DB().Model("hg_user").Where("id", user.Id).Scan(&u)
	fromName := u.Nickname
	if fromName == "" {
		fromName = u.Username
	}

	// 4. 组装要推送的消息格式 (补充 fromName 和 fromAvatar)
	pushData := g.Map{
		"id":         msgId,
		"chatType":   chatType,
		"fromId":     user.Id,
		"fromName":   fromName,
		"fromAvatar": u.Avatar,
		"toId":       toId,
		"msgType":    msgType,
		"content":    content,
		"createdAt":  now.Format("Y-m-d H:i:s"),
	}

	res := &websocket.WResponse{
		Event: "im/chat/receive",
		Data:  gjson.New(pushData),
	}

	// 5. 发送给目标
	if chatType == 1 { // 单聊
		// 推送给对方
		websocket.SendToUser(toId, res)
	} else if chatType == 2 { // 群聊
		// 查询群成员
		memberIds, _ := g.DB().Model("hg_im_group_member").
			Where("group_id", toId).
			Array("user_id")

		for _, memberId := range memberIds {
			mId := gconv.Int64(memberId)
			// 不给自己重复推送接收事件
			if mId != user.Id {
				websocket.SendToUser(mId, res)
			}
		}
	}

	// 6. 给发送者自己返回发送成功的确认
	websocket.SendSuccess(client, req.Event, g.Map{
		"msgId": msgId,
		"time":  now.Format("Y-m-d H:i:s"),
	})
}

// updateConversation 更新会话列表
func (c *cChat) updateConversation(client *websocket.Client, chatType int, fromId, toId, msgId int64, content string, now *gtime.Time) {
	ctx := client.Context()

	// 如果是文本或者其他类型，可以截取部分作为摘要
	summary := content
	if len(summary) > 50 {
		summary = summary[:50] + "..."
	}

	// 更新发送者的会话列表
	c.saveConversation(ctx, fromId, chatType, toId, msgId, summary, now, 0)

	// 更新接收者的会话列表
	if chatType == 1 {
		// 单聊：给对方更新未读数
		c.saveConversation(ctx, toId, chatType, fromId, msgId, summary, now, 1)
	} else if chatType == 2 {
		// 群聊：给所有群成员更新未读数
		memberIds, _ := g.DB().Model("hg_im_group_member").Where("group_id", toId).Array("user_id")
		for _, memberId := range memberIds {
			mId := gconv.Int64(memberId)
			if mId != fromId {
				c.saveConversation(ctx, mId, chatType, toId, msgId, summary, now, 1)
			}
		}
	}
}

func (c *cChat) saveConversation(ctx context.Context, userId int64, chatType int, targetId, msgId int64, content string, now *gtime.Time, unreadIncr int) {
	count, _ := g.DB().Model("hg_im_conversation").
		Where("user_id", userId).
		Where("chat_type", chatType).
		Where("target_id", targetId).
		Count()

	if count > 0 {
		// 更新
		g.DB().Model("hg_im_conversation").
			Where("user_id", userId).
			Where("chat_type", chatType).
			Where("target_id", targetId).
			Increment("unread_count", unreadIncr)

		g.DB().Model("hg_im_conversation").
			Where("user_id", userId).
			Where("chat_type", chatType).
			Where("target_id", targetId).
			Data(g.Map{
				"last_msg_id":      msgId,
				"last_msg_content": content,
				"last_msg_time":    now,
				"updated_at":       now,
			}).Update()
	} else {
		// 插入
		g.DB().Model("hg_im_conversation").Data(g.Map{
			"user_id":          userId,
			"chat_type":        chatType,
			"target_id":        targetId,
			"last_msg_id":      msgId,
			"last_msg_content": content,
			"last_msg_time":    now,
			"unread_count":     unreadIncr,
			"is_top":           0,
			"is_disturb":       0,
			"created_at":       now,
			"updated_at":       now,
		}).Insert()
	}
}
