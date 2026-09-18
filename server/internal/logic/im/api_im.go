package im

import (
	"context"
	v1 "hotgo/api/api/im/v1"
	"hotgo/internal/library/contexts"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sApiIm struct{}

func init() {
	service.RegisterApiIm(&sApiIm{})
}

func (s *sApiIm) ConversationList(ctx context.Context, req *v1.ConversationListReq) (res *v1.ConversationListRes, err error) {
	userCtx := contexts.GetUser(ctx)
	if userCtx == nil {
		err = gerror.New("请先登录")
		return
	}

	var list []*v1.ConversationItem
	err = g.DB().Model("hg_im_conversation").
		Where("user_id", userCtx.Id).
		Order("is_top desc, last_msg_time desc").
		Scan(&list)

	for _, v := range list {
		if v.ChatType == 1 { // 单聊
			var u struct {
				Nickname string
				Username string
				Avatar   string
			}
			g.DB().Model("hg_user").Where("id", v.TargetId).Scan(&u)
			v.TargetName = u.Nickname
			if v.TargetName == "" {
				v.TargetName = u.Username
			}
			v.TargetAvatar = u.Avatar
		} else if v.ChatType == 2 { // 群聊
			var ginfo struct {
				Name   string
				Avatar string
			}
			g.DB().Model("hg_im_group").Where("id", v.TargetId).Scan(&ginfo)
			v.TargetName = ginfo.Name
			v.TargetAvatar = ginfo.Avatar
		}
		// 统一格式化一下创建时间，避免有些时间返回为空字符串的情况
		var ctime struct {
			CreatedAt *gtime.Time `json:"created_at"`
		}
		g.DB().Model("hg_im_conversation").Where("id", v.Id).Scan(&ctime)
		if ctime.CreatedAt != nil {
			v.CreatedAt = ctime.CreatedAt.Format("Y-m-d H:i:s")
		}
	}

	res = &v1.ConversationListRes{
		List: list,
	}
	return
}

func (s *sApiIm) AddFriend(ctx context.Context, req *v1.AddFriendReq) (res *v1.AddFriendRes, err error) {
	userCtx := contexts.GetUser(ctx)
	if userCtx == nil {
		err = gerror.New("请先登录")
		return
	}

	userId := userCtx.Id

	// 检查不能添加自己
	if userId == req.FriendId {
		err = gerror.New("不能添加自己为好友")
		return
	}

	// 检查目标用户是否存在
	count, _ := g.DB().Model("hg_user").Where("id", req.FriendId).Count()
	if count == 0 {
		err = gerror.New("该用户不存在")
		return
	}

	// 检查是否已经是好友
	isFriend, _ := g.DB().Model("hg_im_contact").
		Where("user_id", userId).
		Where("friend_id", req.FriendId).
		Count()

	if isFriend > 0 {
		err = gerror.New("对方已经是您的好友了")
		return
	}

	now := gtime.Now()
	// 互相添加为好友 (双向关系)
	_, err = g.DB().Model("hg_im_contact").Data(g.List{
		{
			"user_id":    userId,
			"friend_id":  req.FriendId,
			"remark":     req.Remark,
			"status":     1,
			"created_at": now,
			"updated_at": now,
		},
		{
			"user_id":    req.FriendId,
			"friend_id":  userId,
			"remark":     "",
			"status":     1,
			"created_at": now,
			"updated_at": now,
		},
	}).Insert()

	return
}

func (s *sApiIm) CreateGroup(ctx context.Context, req *v1.CreateGroupReq) (res *v1.CreateGroupRes, err error) {
	userCtx := contexts.GetUser(ctx)
	if userCtx == nil {
		err = gerror.New("请先登录")
		return
	}

	userId := userCtx.Id
	now := gtime.Now()

	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 创建群组
		groupId, e := tx.Model("hg_im_group").Data(g.Map{
			"name":       req.Name,
			"avatar":     req.Avatar,
			"owner_id":   userId,
			"notice":     req.Notice,
			"status":     1,
			"created_at": now,
			"updated_at": now,
		}).InsertAndGetId()

		if e != nil {
			return e
		}

		// 2. 将创建者加入群成员表并设置为群主
		_, e = tx.Model("hg_im_group_member").Data(g.Map{
			"group_id":   groupId,
			"user_id":    userId,
			"role":       1, // 1=群主
			"is_mute":    0,
			"created_at": now,
		}).Insert()

		if e != nil {
			return e
		}

		res = &v1.CreateGroupRes{GroupId: groupId}
		return nil
	})

	return
}

func (s *sApiIm) ContactList(ctx context.Context, req *v1.ContactListReq) (res *v1.ContactListRes, err error) {
	userCtx := contexts.GetUser(ctx)
	if userCtx == nil {
		err = gerror.New("请先登录")
		return
	}

	res = new(v1.ContactListRes)
	res.List = make([]*v1.ContactItem, 0)

	// 查询联系人表，并连表查询 user 信息获取昵称和头像
	err = g.DB().Model("hg_im_contact c").
		LeftJoin("hg_user u", "c.friend_id=u.id").
		Fields("c.friend_id, c.remark, c.status, u.nickname, u.username, u.avatar").
		Where("c.user_id", userCtx.Id).
		Scan(&res.List)

	// 如果没有昵称，则用 username 顶替
	for _, v := range res.List {
		if v.Nickname == "" {
			// 在上面查询时我们可以稍微做点手脚或者在这里处理
			// 为了简便，我们假定前端可以直接显示 nickname，如果是空就需要在 SQL 里 IFNULL
			// 或者我们在循环里处理
			var u struct{ Username string }
			g.DB().Model("hg_user").Where("id", v.FriendId).Scan(&u)
			if v.Nickname == "" {
				v.Nickname = u.Username
			}
		}
	}

	return
}

func (s *sApiIm) GroupList(ctx context.Context, req *v1.GroupListReq) (res *v1.GroupListRes, err error) {
	userCtx := contexts.GetUser(ctx)
	if userCtx == nil {
		err = gerror.New("请先登录")
		return
	}

	res = new(v1.GroupListRes)
	res.List = make([]*v1.GroupItem, 0)

	// 查询我加入的群
	err = g.DB().Model("hg_im_group_member gm").
		LeftJoin("hg_im_group g", "gm.group_id=g.id").
		Fields("gm.group_id, gm.role, g.name, g.avatar").
		Where("gm.user_id", userCtx.Id).
		Where("g.status", 1). // 1=正常群组
		Scan(&res.List)

	return
}

func (s *sApiIm) OpenConversation(ctx context.Context, req *v1.OpenConversationReq) (res *v1.OpenConversationRes, err error) {
	userCtx := contexts.GetUser(ctx)
	if userCtx == nil {
		err = gerror.New("请先登录")
		return
	}

	userId := userCtx.Id

	// 查询是否已经存在会话
	var conv struct {
		Id int64
	}
	err = g.DB().Model("hg_im_conversation").
		Where("user_id", userId).
		Where("chat_type", req.ChatType).
		Where("target_id", req.TargetId).
		Scan(&conv)

	// Scan 在查不到数据时会返回 sql.ErrNoRows，在这里这是正常逻辑，不是真正的错误
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}
	// 重置 err，防止外层被当作错误处理
	err = nil

	res = new(v1.OpenConversationRes)

	if conv.Id > 0 {
		res.ConversationId = conv.Id
	} else {
		// 如果不存在，则创建一条空会话
		now := gtime.Now()
		insertId, _ := g.DB().Model("hg_im_conversation").Data(g.Map{
			"user_id":          userId,
			"chat_type":        req.ChatType,
			"target_id":        req.TargetId,
			"last_msg_id":      0,
			"last_msg_content": "",
			"last_msg_time":    now,
			"unread_count":     0,
			"is_top":           0,
			"is_disturb":       0,
			"created_at":       now,
			"updated_at":       now,
		}).InsertAndGetId()

		res.ConversationId = insertId
	}

	return
}

func (s *sApiIm) InviteGroup(ctx context.Context, req *v1.InviteGroupReq) (res *v1.InviteGroupRes, err error) {
	userCtx := contexts.GetUser(ctx)
	if userCtx == nil {
		err = gerror.New("请先登录")
		return
	}

	// 校验群是否存在
	count, _ := g.DB().Model("hg_im_group").Where("id", req.GroupId).Where("status", 1).Count()
	if count == 0 {
		err = gerror.New("群组不存在或已解散")
		return
	}

	now := gtime.Now()
	var list g.List

	// 遍历被邀请人，检查是否已经在群内
	for _, uid := range req.UserIds {
		c, _ := g.DB().Model("hg_im_group_member").
			Where("group_id", req.GroupId).
			Where("user_id", uid).
			Count()

		if c == 0 {
			list = append(list, g.Map{
				"group_id":   req.GroupId,
				"user_id":    uid,
				"role":       3, // 3=普通成员
				"is_mute":    0,
				"created_at": now,
			})
		}
	}

	// 批量插入成员
	if len(list) > 0 {
		_, err = g.DB().Model("hg_im_group_member").Data(list).Insert()
	}

	return
}

func (s *sApiIm) GroupMemberList(ctx context.Context, req *v1.GroupMemberListReq) (res *v1.GroupMemberListRes, err error) {
	userCtx := contexts.GetUser(ctx)
	if userCtx == nil {
		err = gerror.New("请先登录")
		return
	}

	res = new(v1.GroupMemberListRes)
	res.List = make([]*v1.GroupMemberItem, 0)

	// 连表查询成员信息和用户昵称头像
	err = g.DB().Model("hg_im_group_member gm").
		LeftJoin("hg_user u", "gm.user_id=u.id").
		Fields("gm.user_id, gm.role, gm.is_mute, u.nickname, u.username, u.avatar").
		Where("gm.group_id", req.GroupId).
		Order("gm.role asc, gm.created_at asc"). // 群主在前，管理员其次，最后是普通成员
		Scan(&res.List)

	for _, v := range res.List {
		if v.Nickname == "" {
			var u struct{ Username string }
			g.DB().Model("hg_user").Where("id", v.UserId).Scan(&u)
			v.Nickname = u.Username
		}
	}

	return
}

func (s *sApiIm) MessageList(ctx context.Context, req *v1.MessageListReq) (res *v1.MessageListRes, err error) {
	userCtx := contexts.GetUser(ctx)
	if userCtx == nil {
		err = gerror.New("请先登录")
		return
	}

	userId := userCtx.Id
	res = new(v1.MessageListRes)
	res.List = make([]*v1.MessageItem, 0)

	m := g.DB().Model("hg_im_message")

	if req.ChatType == 1 {
		// 单聊：发送者是自己接收者是对方，或者发送者是对方接收者是自己
		m = m.Where(
			m.Builder().Where(g.Map{"from_id": userId, "to_id": req.TargetId}).
				WhereOr(g.Map{"from_id": req.TargetId, "to_id": userId}),
		)
	} else if req.ChatType == 2 {
		// 群聊：接收者是指定的群组ID
		m = m.Where("chat_type", 2).Where("to_id", req.TargetId)
	}

	// 统计总数
	total, err := m.Count()
	if err != nil {
		return nil, err
	}
	res.Total = total

	// 查询记录 (倒序)
	records, err := m.Page(req.Page, req.PageSize).OrderDesc("id").All()
	if err != nil {
		return nil, err
	}

	for _, record := range records {
		item := &v1.MessageItem{}
		if err := record.Struct(item); err != nil {
			continue
		}

		// 填充发送者信息 (实际生产中可以缓存到 Redis 或者批量查询)
		var u struct {
			Nickname string
			Avatar   string
			Username string
		}
		_ = g.DB().Model("hg_user").Where("id", item.FromId).Scan(&u)

		item.FromName = u.Nickname
		if item.FromName == "" {
			item.FromName = u.Username
		}
		item.FromAvatar = u.Avatar

		// 格式化一下消息时间
		var ctime struct {
			CreatedAt *gtime.Time `json:"created_at"`
		}
		g.DB().Model("hg_im_message").Where("id", item.Id).Scan(&ctime)
		if ctime.CreatedAt != nil {
			item.CreatedAt = ctime.CreatedAt.Format("Y-m-d H:i:s")
		}

		res.List = append(res.List, item)
	}

	// 对于前端展示来说，通常需要将获取的历史记录按照时间正序排列（最老的在上面，最新的在下面）
	// 由于我们查的是 OrderDesc 获取最新的一页，这里可以翻转一下切片
	for i, j := 0, len(res.List)-1; i < j; i, j = i+1, j-1 {
		res.List[i], res.List[j] = res.List[j], res.List[i]
	}

	return
}
