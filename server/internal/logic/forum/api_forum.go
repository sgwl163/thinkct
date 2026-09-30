package forum

import (
	"context"
	"sort"

	v1 "hotgo/api/api/forum/v1"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sApiForum struct{}

type userActivityRecord struct {
	activity *v1.UserActivity
	topicID  uint64
	userID   uint64
}

func init() {
	service.RegisterApiForum(&sApiForum{})
}

func (s *sApiForum) CategoryList(ctx context.Context, req *v1.CategoryListReq) (res *v1.CategoryListRes, err error) {
	list := make([]*entity.ForumCategory, 0)
	cols := dao.ForumCategory.Columns()
	if err = dao.ForumCategory.Ctx(ctx).
		Where(cols.Status, consts.ForumCategoryStatusNormal).
		OrderAsc(cols.Sort).
		OrderAsc(cols.Id).
		Scan(&list); err != nil {
		return nil, gerror.Wrap(err, "获取论坛板块失败，请稍后重试")
	}
	return &v1.CategoryListRes{List: list}, nil
}

func (s *sApiForum) TopicList(ctx context.Context, req *v1.TopicListReq) (res *v1.TopicListRes, err error) {
	cols := dao.ForumTopic.Columns()
	mod := dao.ForumTopic.Ctx(ctx).Where(cols.Status, consts.ForumTopicStatusNormal)
	if req.CategoryId > 0 {
		mod = mod.Where(cols.CategoryId, req.CategoryId)
	}
	if req.Keyword != "" {
		keyword := "%" + req.Keyword + "%"
		mod = mod.Where("("+cols.Title+" LIKE ? OR "+cols.Keywords+" LIKE ? OR "+cols.Description+" LIKE ?)", keyword, keyword, keyword)
	}
	switch req.Sort {
	case "hot":
		mod = mod.OrderDesc(cols.IsTop).OrderDesc(cols.IsEssence).OrderDesc(cols.LikeCount).OrderDesc(cols.CommentCount).OrderDesc(cols.ViewCount).OrderDesc(cols.Id)
	case "recommend":
		mod = mod.OrderDesc(cols.IsTop).OrderDesc(cols.IsEssence).OrderDesc(cols.FavoriteCount).OrderDesc(cols.LikeCount).OrderDesc(cols.CommentCount).OrderDesc(cols.ViewCount).OrderDesc(cols.LastReplyAt).OrderDesc(cols.Id)
	default:
		mod = mod.OrderDesc(cols.IsTop).OrderDesc(cols.LastReplyAt).OrderDesc(cols.CreatedAt).OrderDesc(cols.Id)
	}

	total, err := mod.Clone().Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取论坛主题数量失败，请稍后重试")
	}
	topics := make([]*entity.ForumTopic, 0)
	if total > 0 {
		if err = mod.Page(req.Page, req.PerPage).Scan(&topics); err != nil {
			return nil, gerror.Wrap(err, "获取论坛主题失败，请稍后重试")
		}
	}
	list, err := s.buildTopics(ctx, topics)
	if err != nil {
		return nil, err
	}
	res = &v1.TopicListRes{List: list}
	res.PageRes.Pack(req, total)
	return res, nil
}

func (s *sApiForum) TopicDetail(ctx context.Context, req *v1.TopicDetailReq) (res *v1.TopicDetailRes, err error) {
	cols := dao.ForumTopic.Columns()
	var topic *entity.ForumTopic
	if err = dao.ForumTopic.Ctx(ctx).
		WherePri(req.Id).
		Where(cols.Status, consts.ForumTopicStatusNormal).
		Scan(&topic); err != nil {
		return nil, gerror.Wrap(err, "获取论坛主题失败，请稍后重试")
	}
	if topic == nil {
		return nil, gerror.New("主题不存在或已不可见")
	}
	if _, err = dao.ForumTopic.Ctx(ctx).WherePri(topic.Id).Increment(cols.ViewCount, 1); err != nil {
		return nil, gerror.Wrap(err, "更新主题浏览量失败，请稍后重试")
	}
	topic.ViewCount++
	list, err := s.buildTopics(ctx, []*entity.ForumTopic{topic})
	if err != nil {
		return nil, err
	}
	prev, next, err := s.loadTopicNavigation(ctx, topic)
	if err != nil {
		return nil, err
	}
	return &v1.TopicDetailRes{Topic: list[0], Prev: prev, Next: next}, nil
}

func (s *sApiForum) UserProfile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	if err = s.ensureUser(ctx, req.UserId); err != nil {
		return nil, err
	}

	authors, err := s.loadAuthors(ctx, []uint64{req.UserId})
	if err != nil {
		return nil, err
	}
	var user *entity.User
	userCols := dao.User.Columns()
	if err = dao.User.Ctx(ctx).Fields(userCols.Motto, userCols.JoinedAt).WherePri(req.UserId).Scan(&user); err != nil {
		return nil, gerror.Wrap(err, "获取作者资料失败，请稍后重试")
	}

	topicCols := dao.ForumTopic.Columns()
	topicModel := dao.ForumTopic.Ctx(ctx).
		Where(topicCols.UserId, req.UserId).
		Where(topicCols.Status, consts.ForumTopicStatusNormal).
		WhereNull(topicCols.DeletedAt)
	articleCount, err := topicModel.Clone().Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取作者主题数量失败，请稍后重试")
	}
	viewCount, err := topicModel.Sum(topicCols.ViewCount)
	if err != nil {
		return nil, gerror.Wrap(err, "获取作者阅读数量失败，请稍后重试")
	}

	followCols := dao.ForumUserFollow.Columns()
	followingCount, err := dao.ForumUserFollow.Ctx(ctx).Where(followCols.UserId, req.UserId).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取作者关注数量失败，请稍后重试")
	}
	followerCount, err := dao.ForumUserFollow.Ctx(ctx).Where(followCols.FollowUserId, req.UserId).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取作者粉丝数量失败，请稍后重试")
	}

	interactionCols := dao.ForumInteraction.Columns()
	interactionModel := dao.ForumInteraction.Ctx(ctx).
		Where(interactionCols.UserId, req.UserId).
		Where(interactionCols.TargetType, consts.ForumInteractionTargetTypeTopic)
	likeCount, err := interactionModel.Clone().Where(interactionCols.ActionType, consts.ForumInteractionActionTypeLike).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取作者点赞数量失败，请稍后重试")
	}
	favoriteCount, err := interactionModel.Where(interactionCols.ActionType, consts.ForumInteractionActionTypeFavorite).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取作者收藏数量失败，请稍后重试")
	}

	currentUserID := contexts.GetUserId(ctx)
	isSelf := currentUserID > 0 && uint64(currentUserID) == req.UserId
	isFollowed := false
	if currentUserID > 0 && !isSelf {
		followCount, followErr := dao.ForumUserFollow.Ctx(ctx).
			Where(followCols.UserId, currentUserID).
			Where(followCols.FollowUserId, req.UserId).
			Count()
		if followErr != nil {
			return nil, gerror.Wrap(followErr, "获取关注状态失败，请稍后重试")
		}
		isFollowed = followCount > 0
	}

	profile := &v1.UserProfile{
		Author:         authors[req.UserId],
		ArticleCount:   uint(articleCount),
		ViewCount:      uint(viewCount),
		FollowingCount: followingCount,
		FollowerCount:  followerCount,
		FavoriteCount:  favoriteCount,
		LikeCount:      likeCount,
		IsFollowed:     isFollowed,
		IsSelf:         isSelf,
	}
	if user != nil {
		profile.Motto = user.Motto
		profile.JoinedAt = user.JoinedAt
	}
	return &v1.UserProfileRes{Profile: profile}, nil
}

func (s *sApiForum) UserActivityList(ctx context.Context, req *v1.UserActivityListReq) (res *v1.UserActivityListRes, err error) {
	if err = s.ensureUser(ctx, req.UserId); err != nil {
		return nil, err
	}

	records := make([]*userActivityRecord, 0)
	topicEntities := make(map[uint64]*entity.ForumTopic)
	topicCols := dao.ForumTopic.Columns()
	if req.Type == "" || req.Type == consts.ForumUserActivityTypeTopic {
		var topics []*entity.ForumTopic
		if err = dao.ForumTopic.Ctx(ctx).
			Where(topicCols.UserId, req.UserId).
			Where(topicCols.Status, consts.ForumTopicStatusNormal).
			WhereNull(topicCols.DeletedAt).
			Scan(&topics); err != nil {
			return nil, gerror.Wrap(err, "获取用户发布动态失败，请稍后重试")
		}
		for _, topic := range topics {
			topicEntities[topic.Id] = topic
			records = append(records, &userActivityRecord{
				activity: &v1.UserActivity{Id: topic.Id, Type: consts.ForumUserActivityTypeTopic, CreatedAt: topic.CreatedAt},
				topicID:  topic.Id,
			})
		}
	}

	interactionTopicIDs := make([]uint64, 0)
	interactionRecords := make([]*userActivityRecord, 0)
	if req.Type == "" || req.Type == consts.ForumUserActivityTypeLike || req.Type == consts.ForumUserActivityTypeFavorite {
		interactionCols := dao.ForumInteraction.Columns()
		interactionModel := dao.ForumInteraction.Ctx(ctx).
			Where(interactionCols.UserId, req.UserId).
			Where(interactionCols.TargetType, consts.ForumInteractionTargetTypeTopic)
		if req.Type == consts.ForumUserActivityTypeLike {
			interactionModel = interactionModel.Where(interactionCols.ActionType, consts.ForumInteractionActionTypeLike)
		} else if req.Type == consts.ForumUserActivityTypeFavorite {
			interactionModel = interactionModel.Where(interactionCols.ActionType, consts.ForumInteractionActionTypeFavorite)
		} else {
			interactionModel = interactionModel.WhereIn(interactionCols.ActionType, []string{consts.ForumInteractionActionTypeLike, consts.ForumInteractionActionTypeFavorite})
		}
		var interactions []*entity.ForumInteraction
		if err = interactionModel.Scan(&interactions); err != nil {
			return nil, gerror.Wrap(err, "获取用户互动动态失败，请稍后重试")
		}
		for _, interaction := range interactions {
			activityType := consts.ForumUserActivityTypeLike
			if interaction.ActionType == consts.ForumInteractionActionTypeFavorite {
				activityType = consts.ForumUserActivityTypeFavorite
			}
			interactionTopicIDs = append(interactionTopicIDs, interaction.TargetId)
			interactionRecords = append(interactionRecords, &userActivityRecord{
				activity: &v1.UserActivity{Id: interaction.Id, Type: activityType, CreatedAt: interaction.CreatedAt},
				topicID:  interaction.TargetId,
			})
		}
	}
	if len(interactionTopicIDs) > 0 {
		var topics []*entity.ForumTopic
		if err = dao.ForumTopic.Ctx(ctx).
			WhereIn(topicCols.Id, interactionTopicIDs).
			Where(topicCols.Status, consts.ForumTopicStatusNormal).
			WhereNull(topicCols.DeletedAt).
			Scan(&topics); err != nil {
			return nil, gerror.Wrap(err, "获取用户互动主题失败，请稍后重试")
		}
		for _, topic := range topics {
			topicEntities[topic.Id] = topic
		}
		for _, record := range interactionRecords {
			if topicEntities[record.topicID] != nil {
				records = append(records, record)
			}
		}
	}

	followUserIDs := make([]uint64, 0)
	if req.Type == "" || req.Type == consts.ForumUserActivityTypeFollow {
		followCols := dao.ForumUserFollow.Columns()
		var follows []*entity.ForumUserFollow
		if err = dao.ForumUserFollow.Ctx(ctx).Where(followCols.UserId, req.UserId).Scan(&follows); err != nil {
			return nil, gerror.Wrap(err, "获取用户关注动态失败，请稍后重试")
		}
		for _, follow := range follows {
			followUserIDs = append(followUserIDs, follow.FollowUserId)
			records = append(records, &userActivityRecord{
				activity: &v1.UserActivity{Id: follow.Id, Type: consts.ForumUserActivityTypeFollow, CreatedAt: follow.CreatedAt},
				userID:   follow.FollowUserId,
			})
		}
	}

	topicEntitiesList := make([]*entity.ForumTopic, 0, len(topicEntities))
	for _, topic := range topicEntities {
		topicEntitiesList = append(topicEntitiesList, topic)
	}
	builtTopics, err := s.buildTopics(ctx, topicEntitiesList)
	if err != nil {
		return nil, err
	}
	activityTopics := make(map[uint64]*v1.Topic, len(builtTopics))
	for _, topic := range builtTopics {
		activityTopics[topic.Id] = topic
	}
	followUsers, err := s.loadAuthors(ctx, followUserIDs)
	if err != nil {
		return nil, err
	}
	visibleRecords := make([]*userActivityRecord, 0, len(records))
	for _, record := range records {
		if record.topicID > 0 {
			record.activity.Topic = activityTopics[record.topicID]
			if record.activity.Topic == nil {
				continue
			}
		}
		if record.userID > 0 {
			record.activity.User = followUsers[record.userID]
			if record.activity.User == nil {
				continue
			}
		}
		visibleRecords = append(visibleRecords, record)
	}
	sort.SliceStable(visibleRecords, func(i, j int) bool {
		left := visibleRecords[i].activity
		right := visibleRecords[j].activity
		if left.CreatedAt == nil {
			return false
		}
		if right.CreatedAt == nil {
			return true
		}
		if left.CreatedAt.Timestamp() == right.CreatedAt.Timestamp() {
			return left.Id > right.Id
		}
		return left.CreatedAt.After(right.CreatedAt)
	})

	total := len(visibleRecords)
	start := (req.Page - 1) * req.PerPage
	if start > total {
		start = total
	}
	end := start + req.PerPage
	if end > total {
		end = total
	}
	list := make([]*v1.UserActivity, 0, end-start)
	for _, record := range visibleRecords[start:end] {
		list = append(list, record.activity)
	}
	res = &v1.UserActivityListRes{List: list}
	res.PageRes.Pack(req, total)
	return res, nil
}

func (s *sApiForum) CommentList(ctx context.Context, req *v1.CommentListReq) (res *v1.CommentListRes, err error) {
	if err = s.ensureTopic(ctx, req.TopicId); err != nil {
		return nil, err
	}
	cols := dao.ForumComment.Columns()
	rootModel := dao.ForumComment.Ctx(ctx).
		Where(cols.TopicId, req.TopicId).
		Where(cols.ParentId, 0).
		Where(cols.Status, consts.ForumCommentStatusNormal).
		WhereNull(cols.DeletedAt)
	if req.Sort == "hot" {
		rootModel = rootModel.OrderDesc(cols.LikeCount).OrderDesc(cols.CreatedAt).OrderDesc(cols.Id)
	} else {
		rootModel = rootModel.OrderDesc(cols.CreatedAt).OrderDesc(cols.Id)
	}
	total, err := rootModel.Clone().Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取评论数量失败，请稍后重试")
	}
	roots := make([]*entity.ForumComment, 0)
	if total > 0 {
		if err = rootModel.Page(req.Page, req.PerPage).Scan(&roots); err != nil {
			return nil, gerror.Wrap(err, "获取评论失败，请稍后重试")
		}
	}
	comments, err := s.loadCommentReplies(ctx, req.TopicId, req.Sort)
	if err != nil {
		return nil, err
	}
	comments = append(roots, comments...)
	list, err := s.buildCommentTree(ctx, roots, comments, req.Sort)
	if err != nil {
		return nil, err
	}
	res = &v1.CommentListRes{List: list}
	res.PageRes.Pack(req, total)
	return res, nil
}

func (s *sApiForum) CreateTopic(ctx context.Context, req *v1.CreateTopicReq) (res *v1.CreateTopicRes, err error) {
	userID, err := requireUser(ctx)
	if err != nil {
		return nil, err
	}
	category, err := s.ensureCategory(ctx, req.CategoryId)
	if err != nil {
		return nil, err
	}
	topic := &entity.ForumTopic{
		CategoryId:  req.CategoryId,
		UserId:      uint64(userID),
		Title:       req.Title,
		Keywords:    req.Keywords,
		Description: req.Description,
		Content:     req.Content,
		Images:      req.Images,
		Status:      consts.ForumTopicStatusNormal,
	}
	if err = dao.ForumTopic.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id, txErr := tx.Model(dao.ForumTopic.Table()).Data(topic).InsertAndGetId()
		if txErr != nil {
			return gerror.Wrap(txErr, "发布主题失败，请稍后重试")
		}
		topic.Id = uint64(id)
		if _, txErr = tx.Model(dao.ForumCategory.Table()).WherePri(category.Id).Increment(dao.ForumCategory.Columns().TopicCount, 1); txErr != nil {
			return gerror.Wrap(txErr, "更新板块主题数失败，请稍后重试")
		}
		return nil
	}); err != nil {
		return nil, err
	}
	if err = dao.ForumTopic.Ctx(ctx).WherePri(topic.Id).Scan(&topic); err != nil {
		return nil, gerror.Wrap(err, "获取新建主题失败，请稍后重试")
	}
	list, err := s.buildTopics(ctx, []*entity.ForumTopic{topic})
	if err != nil {
		return nil, err
	}
	return &v1.CreateTopicRes{Topic: list[0]}, nil
}

func (s *sApiForum) MyTopicList(ctx context.Context, req *v1.MyTopicListReq) (res *v1.MyTopicListRes, err error) {
	userID, err := requireUser(ctx)
	if err != nil {
		return nil, err
	}

	cols := dao.ForumTopic.Columns()
	mod := dao.ForumTopic.Ctx(ctx).
		Where(cols.UserId, userID).
		WhereNull(cols.DeletedAt)
	if req.Status != "" {
		mod = mod.Where(cols.Status, req.Status)
	}
	total, err := mod.Clone().Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取我的主题数量失败，请稍后重试")
	}
	topics := make([]*entity.ForumTopic, 0)
	if total > 0 {
		if err = mod.OrderDesc(cols.UpdatedAt).OrderDesc(cols.Id).Page(req.Page, req.PerPage).Scan(&topics); err != nil {
			return nil, gerror.Wrap(err, "获取我的主题失败，请稍后重试")
		}
	}
	list, err := s.buildTopics(ctx, topics)
	if err != nil {
		return nil, err
	}
	res = &v1.MyTopicListRes{List: list}
	res.PageRes.Pack(req, total)
	return res, nil
}

func (s *sApiForum) MyTopicDetail(ctx context.Context, req *v1.MyTopicDetailReq) (res *v1.MyTopicDetailRes, err error) {
	userID, err := requireUser(ctx)
	if err != nil {
		return nil, err
	}

	cols := dao.ForumTopic.Columns()
	var topic *entity.ForumTopic
	if err = dao.ForumTopic.Ctx(ctx).
		WherePri(req.Id).
		Where(cols.UserId, userID).
		WhereNull(cols.DeletedAt).
		Scan(&topic); err != nil {
		return nil, gerror.Wrap(err, "获取我的主题失败，请稍后重试")
	}
	if topic == nil {
		return nil, gerror.New("主题不存在或无权查看")
	}
	list, err := s.buildTopics(ctx, []*entity.ForumTopic{topic})
	if err != nil {
		return nil, err
	}
	return &v1.MyTopicDetailRes{Topic: list[0]}, nil
}

func (s *sApiForum) UpdateTopic(ctx context.Context, req *v1.UpdateTopicReq) (res *v1.UpdateTopicRes, err error) {
	userID, err := requireUser(ctx)
	if err != nil {
		return nil, err
	}
	if _, err = s.ensureCategory(ctx, req.CategoryId); err != nil {
		return nil, err
	}

	cols := dao.ForumTopic.Columns()
	categoryCols := dao.ForumCategory.Columns()
	var topic *entity.ForumTopic
	if err = dao.ForumTopic.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		if txErr := tx.Model(dao.ForumTopic.Table()).
			WherePri(req.Id).
			Where(cols.UserId, userID).
			WhereNull(cols.DeletedAt).
			Scan(&topic); txErr != nil {
			return gerror.Wrap(txErr, "获取主题失败，请稍后重试")
		}
		if topic == nil {
			return gerror.New("主题不存在或无权编辑")
		}
		if _, txErr := tx.Model(dao.ForumTopic.Table()).
			WherePri(topic.Id).
			Where(cols.UserId, userID).
			Data(g.Map{
				cols.CategoryId:  req.CategoryId,
				cols.Title:       req.Title,
				cols.Keywords:    req.Keywords,
				cols.Description: req.Description,
				cols.Content:     req.Content,
				cols.Images:      req.Images,
			}).
			Update(); txErr != nil {
			return gerror.Wrap(txErr, "编辑主题失败，请稍后重试")
		}
		if topic.CategoryId != req.CategoryId {
			if _, txErr := tx.Model(dao.ForumCategory.Table()).WherePri(topic.CategoryId).WhereGT(categoryCols.TopicCount, 0).Decrement(categoryCols.TopicCount, 1); txErr != nil {
				return gerror.Wrap(txErr, "更新原板块主题数失败，请稍后重试")
			}
			if _, txErr := tx.Model(dao.ForumCategory.Table()).WherePri(req.CategoryId).Increment(categoryCols.TopicCount, 1); txErr != nil {
				return gerror.Wrap(txErr, "更新新板块主题数失败，请稍后重试")
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	if err = dao.ForumTopic.Ctx(ctx).WherePri(req.Id).Where(cols.UserId, userID).Scan(&topic); err != nil {
		return nil, gerror.Wrap(err, "获取编辑后的主题失败，请稍后重试")
	}
	list, err := s.buildTopics(ctx, []*entity.ForumTopic{topic})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateTopicRes{Topic: list[0]}, nil
}

func (s *sApiForum) CreateComment(ctx context.Context, req *v1.CreateCommentReq) (res *v1.CreateCommentRes, err error) {
	userID, err := requireUser(ctx)
	if err != nil {
		return nil, err
	}
	if err = s.ensureTopic(ctx, req.TopicId); err != nil {
		return nil, err
	}
	var replyUserID uint64
	if req.ParentId > 0 {
		parent, parentErr := s.ensureComment(ctx, req.TopicId, req.ParentId)
		if parentErr != nil {
			return nil, parentErr
		}
		replyUserID = parent.UserId
	}
	now := gtime.Now()
	comment := &entity.ForumComment{TopicId: req.TopicId, UserId: uint64(userID), ParentId: req.ParentId, ReplyUserId: replyUserID, Content: req.Content, Status: consts.ForumCommentStatusNormal}
	if err = dao.ForumComment.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		id, txErr := tx.Model(dao.ForumComment.Table()).Data(comment).InsertAndGetId()
		if txErr != nil {
			return gerror.Wrap(txErr, "发布评论失败，请稍后重试")
		}
		comment.Id = uint64(id)
		if _, txErr = tx.Model(dao.ForumTopic.Table()).WherePri(req.TopicId).Increment(dao.ForumTopic.Columns().CommentCount, 1); txErr != nil {
			return gerror.Wrap(txErr, "更新主题评论数失败，请稍后重试")
		}
		if _, txErr = tx.Model(dao.ForumTopic.Table()).WherePri(req.TopicId).Data(g.Map{dao.ForumTopic.Columns().LastReplyUserId: userID, dao.ForumTopic.Columns().LastReplyAt: now}).Update(); txErr != nil {
			return gerror.Wrap(txErr, "更新主题回复信息失败，请稍后重试")
		}
		return nil
	}); err != nil {
		return nil, err
	}
	if err = dao.ForumComment.Ctx(ctx).WherePri(comment.Id).Scan(&comment); err != nil {
		return nil, gerror.Wrap(err, "获取新建评论失败，请稍后重试")
	}
	list, err := s.buildComments(ctx, []*entity.ForumComment{comment})
	if err != nil {
		return nil, err
	}
	return &v1.CreateCommentRes{Comment: list[0]}, nil
}

func (s *sApiForum) DeleteComment(ctx context.Context, req *v1.DeleteCommentReq) (res *v1.DeleteCommentRes, err error) {
	userID, err := requireUser(ctx)
	if err != nil {
		return nil, err
	}

	commentCols := dao.ForumComment.Columns()
	topicCols := dao.ForumTopic.Columns()
	interactionCols := dao.ForumInteraction.Columns()
	commentCount := uint(0)
	err = dao.ForumComment.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var comment *entity.ForumComment
		if txErr := tx.Model(dao.ForumComment.Table()).
			WherePri(req.CommentId).
			Where(commentCols.UserId, userID).
			Where(commentCols.Status, consts.ForumCommentStatusNormal).
			WhereNull(commentCols.DeletedAt).
			Scan(&comment); txErr != nil {
			return gerror.Wrap(txErr, "获取评论失败，请稍后重试")
		}
		if comment == nil {
			return gerror.New("评论不存在或无权删除")
		}

		childCount, txErr := tx.Model(dao.ForumComment.Table()).
			Where(commentCols.ParentId, comment.Id).
			WhereNull(commentCols.DeletedAt).
			Count()
		if txErr != nil {
			return gerror.Wrap(txErr, "检查评论回复失败，请稍后重试")
		}
		if childCount > 0 {
			return gerror.New("评论存在回复，不能删除")
		}

		if _, txErr = tx.Model(dao.ForumComment.Table()).
			WherePri(comment.Id).
			Where(commentCols.UserId, userID).
			WhereNull(commentCols.DeletedAt).
			Data(g.Map{commentCols.DeletedAt: gtime.Now()}).
			Update(); txErr != nil {
			return gerror.Wrap(txErr, "删除评论失败，请稍后重试")
		}
		if _, txErr = tx.Model(dao.ForumInteraction.Table()).
			Where(interactionCols.TargetType, consts.ForumInteractionTargetTypeComment).
			Where(interactionCols.TargetId, comment.Id).
			Delete(); txErr != nil {
			return gerror.Wrap(txErr, "清理评论点赞失败，请稍后重试")
		}
		if _, txErr = tx.Model(dao.ForumTopic.Table()).
			WherePri(comment.TopicId).
			WhereGT(topicCols.CommentCount, 0).
			Decrement(topicCols.CommentCount, 1); txErr != nil {
			return gerror.Wrap(txErr, "更新主题评论数失败，请稍后重试")
		}

		var latestComment *entity.ForumComment
		if txErr = tx.Model(dao.ForumComment.Table()).
			Fields(commentCols.UserId, commentCols.CreatedAt).
			Where(commentCols.TopicId, comment.TopicId).
			Where(commentCols.Status, consts.ForumCommentStatusNormal).
			WhereNull(commentCols.DeletedAt).
			OrderDesc(commentCols.CreatedAt).
			OrderDesc(commentCols.Id).
			Limit(1).
			Scan(&latestComment); txErr != nil {
			return gerror.Wrap(txErr, "获取主题最新回复失败，请稍后重试")
		}
		lastReplyData := g.Map{topicCols.LastReplyUserId: 0, topicCols.LastReplyAt: nil}
		if latestComment != nil {
			lastReplyData[topicCols.LastReplyUserId] = latestComment.UserId
			lastReplyData[topicCols.LastReplyAt] = latestComment.CreatedAt
		}
		if _, txErr = tx.Model(dao.ForumTopic.Table()).WherePri(comment.TopicId).Data(lastReplyData).Update(); txErr != nil {
			return gerror.Wrap(txErr, "更新主题最新回复失败，请稍后重试")
		}

		commentCountValue, txErr := tx.Model(dao.ForumTopic.Table()).WherePri(comment.TopicId).Value(topicCols.CommentCount)
		if txErr != nil {
			return gerror.Wrap(txErr, "获取主题评论数失败，请稍后重试")
		}
		commentCount = commentCountValue.Uint()
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &v1.DeleteCommentRes{CommentCount: commentCount}, nil
}

func (s *sApiForum) ToggleTopicLike(ctx context.Context, req *v1.ToggleTopicLikeReq) (res *v1.ToggleTopicLikeRes, err error) {
	userID, err := requireUser(ctx)
	if err != nil {
		return nil, err
	}
	if err = s.ensureTopic(ctx, req.TopicId); err != nil {
		return nil, err
	}
	liked, count, err := s.toggleInteraction(ctx, uint64(userID), consts.ForumInteractionTargetTypeTopic, req.TopicId, consts.ForumInteractionActionTypeLike, dao.ForumTopic.Table(), dao.ForumTopic.Columns().LikeCount)
	if err != nil {
		return nil, err
	}
	return &v1.ToggleTopicLikeRes{Liked: liked, LikeCount: count}, nil
}

func (s *sApiForum) ToggleTopicFavorite(ctx context.Context, req *v1.ToggleTopicFavoriteReq) (res *v1.ToggleTopicFavoriteRes, err error) {
	userID, err := requireUser(ctx)
	if err != nil {
		return nil, err
	}
	if err = s.ensureTopic(ctx, req.TopicId); err != nil {
		return nil, err
	}
	favorited, count, err := s.toggleInteraction(ctx, uint64(userID), consts.ForumInteractionTargetTypeTopic, req.TopicId, consts.ForumInteractionActionTypeFavorite, dao.ForumTopic.Table(), dao.ForumTopic.Columns().FavoriteCount)
	if err != nil {
		return nil, err
	}
	return &v1.ToggleTopicFavoriteRes{Favorited: favorited, FavoriteCount: count}, nil
}

func (s *sApiForum) ToggleCommentLike(ctx context.Context, req *v1.ToggleCommentLikeReq) (res *v1.ToggleCommentLikeRes, err error) {
	userID, err := requireUser(ctx)
	if err != nil {
		return nil, err
	}
	comment, err := s.ensureVisibleComment(ctx, req.CommentId)
	if err != nil {
		return nil, err
	}
	liked, count, err := s.toggleInteraction(ctx, uint64(userID), consts.ForumInteractionTargetTypeComment, comment.Id, consts.ForumInteractionActionTypeLike, dao.ForumComment.Table(), dao.ForumComment.Columns().LikeCount)
	if err != nil {
		return nil, err
	}
	return &v1.ToggleCommentLikeRes{Liked: liked, LikeCount: count}, nil
}

func (s *sApiForum) ToggleUserFollow(ctx context.Context, req *v1.ToggleUserFollowReq) (res *v1.ToggleUserFollowRes, err error) {
	userID, err := requireUser(ctx)
	if err != nil {
		return nil, err
	}
	if uint64(userID) == req.FollowUserId {
		return nil, gerror.New("不能关注自己")
	}
	if err = s.ensureUser(ctx, req.FollowUserId); err != nil {
		return nil, err
	}

	cols := dao.ForumUserFollow.Columns()
	followed := false
	followerCount := 0
	err = dao.ForumUserFollow.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		followID, txErr := tx.Model(dao.ForumUserFollow.Table()).
			Where(cols.UserId, userID).
			Where(cols.FollowUserId, req.FollowUserId).
			Value(cols.Id)
		if txErr != nil {
			return gerror.Wrap(txErr, "检查关注状态失败，请稍后重试")
		}
		if followID.Int64() > 0 {
			if _, txErr = tx.Model(dao.ForumUserFollow.Table()).Where(cols.Id, followID.Int64()).Delete(); txErr != nil {
				return gerror.Wrap(txErr, "取消关注失败，请稍后重试")
			}
		} else {
			if _, txErr = tx.Model(dao.ForumUserFollow.Table()).Data(g.Map{
				cols.UserId:       userID,
				cols.FollowUserId: req.FollowUserId,
			}).Insert(); txErr != nil {
				return gerror.Wrap(txErr, "关注用户失败，请稍后重试")
			}
			followed = true
		}
		followerCount, txErr = tx.Model(dao.ForumUserFollow.Table()).Where(cols.FollowUserId, req.FollowUserId).Count()
		if txErr != nil {
			return gerror.Wrap(txErr, "获取粉丝数量失败，请稍后重试")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &v1.ToggleUserFollowRes{Followed: followed, FollowerCount: followerCount}, nil
}

func (s *sApiForum) toggleInteraction(ctx context.Context, userID uint64, targetType string, targetID uint64, actionType string, counterTable string, counterColumn string) (active bool, count uint, err error) {
	cols := dao.ForumInteraction.Columns()
	err = dao.ForumInteraction.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		interactionID, txErr := tx.Model(dao.ForumInteraction.Table()).
			Where(cols.TargetType, targetType).
			Where(cols.TargetId, targetID).
			Where(cols.UserId, userID).
			Where(cols.ActionType, actionType).
			Value(cols.Id)
		if txErr != nil {
			return txErr
		}
		if interactionID.Int64() > 0 {
			if _, txErr = tx.Model(dao.ForumInteraction.Table()).Where(cols.Id, interactionID.Int64()).Delete(); txErr != nil {
				return gerror.Wrap(txErr, "取消互动失败，请稍后重试")
			}
			if _, txErr = tx.Model(counterTable).WherePri(targetID).WhereGT(counterColumn, 0).Decrement(counterColumn, 1); txErr != nil {
				return gerror.Wrap(txErr, "更新互动数量失败，请稍后重试")
			}
		} else {
			if _, txErr = tx.Model(dao.ForumInteraction.Table()).Data(g.Map{
				cols.TargetType: targetType,
				cols.TargetId:   targetID,
				cols.UserId:     userID,
				cols.ActionType: actionType,
			}).Insert(); txErr != nil {
				return gerror.Wrap(txErr, "提交互动失败，请稍后重试")
			}
			if _, txErr = tx.Model(counterTable).WherePri(targetID).Increment(counterColumn, 1); txErr != nil {
				return gerror.Wrap(txErr, "更新互动数量失败，请稍后重试")
			}
			active = true
		}
		value, txErr := tx.Model(counterTable).WherePri(targetID).Value(counterColumn)
		if txErr != nil {
			return txErr
		}
		count = value.Uint()
		return nil
	})
	return active, count, err
}

func (s *sApiForum) CreateReport(ctx context.Context, req *v1.CreateReportReq) (res *v1.CreateReportRes, err error) {
	userID, err := requireUser(ctx)
	if err != nil {
		return nil, err
	}
	if err = s.ensureReportTarget(ctx, req.TargetType, req.TargetId); err != nil {
		return nil, err
	}
	cols := dao.ForumReport.Columns()
	count, err := dao.ForumReport.Ctx(ctx).Where(cols.TargetType, req.TargetType).Where(cols.TargetId, req.TargetId).Where(cols.UserId, userID).Where(cols.Status, consts.ForumReportStatusPending).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "检查重复举报失败，请稍后重试")
	}
	if count > 0 {
		return nil, gerror.New("你已经举报过该内容，请等待处理")
	}
	if _, err = dao.ForumReport.Ctx(ctx).Data(g.Map{cols.TargetType: req.TargetType, cols.TargetId: req.TargetId, cols.UserId: userID, cols.Reason: req.Reason, cols.Description: req.Description, cols.Status: consts.ForumReportStatusPending}).Insert(); err != nil {
		return nil, gerror.Wrap(err, "提交举报失败，请稍后重试")
	}
	return &v1.CreateReportRes{}, nil
}

func (s *sApiForum) loadTopicNavigation(ctx context.Context, topic *entity.ForumTopic) (prev, next *v1.TopicNavigation, err error) {
	cols := dao.ForumTopic.Columns()
	var previous *entity.ForumTopic
	if err = dao.ForumTopic.Ctx(ctx).
		Fields(cols.Id, cols.Title).
		Where(cols.CategoryId, topic.CategoryId).
		Where(cols.Status, consts.ForumTopicStatusNormal).
		WhereLT(cols.Id, topic.Id).
		OrderDesc(cols.Id).
		Limit(1).
		Scan(&previous); err != nil {
		return nil, nil, gerror.Wrap(err, "获取上一篇主题失败，请稍后重试")
	}
	if previous != nil {
		prev = &v1.TopicNavigation{Id: previous.Id, Title: previous.Title}
	}

	var following *entity.ForumTopic
	if err = dao.ForumTopic.Ctx(ctx).
		Fields(cols.Id, cols.Title).
		Where(cols.CategoryId, topic.CategoryId).
		Where(cols.Status, consts.ForumTopicStatusNormal).
		WhereGT(cols.Id, topic.Id).
		OrderAsc(cols.Id).
		Limit(1).
		Scan(&following); err != nil {
		return nil, nil, gerror.Wrap(err, "获取下一篇主题失败，请稍后重试")
	}
	if following != nil {
		next = &v1.TopicNavigation{Id: following.Id, Title: following.Title}
	}
	return prev, next, nil
}

func (s *sApiForum) loadCommentReplies(ctx context.Context, topicID uint64, sort string) ([]*entity.ForumComment, error) {
	cols := dao.ForumComment.Columns()
	mod := dao.ForumComment.Ctx(ctx).
		Where(cols.TopicId, topicID).
		WhereGT(cols.ParentId, 0).
		Where(cols.Status, consts.ForumCommentStatusNormal).
		WhereNull(cols.DeletedAt)
	if sort == "hot" {
		mod = mod.OrderDesc(cols.LikeCount).OrderDesc(cols.CreatedAt).OrderDesc(cols.Id)
	} else {
		mod = mod.OrderDesc(cols.CreatedAt).OrderDesc(cols.Id)
	}
	comments := make([]*entity.ForumComment, 0)
	if err := mod.Scan(&comments); err != nil {
		return nil, gerror.Wrap(err, "获取评论回复失败，请稍后重试")
	}
	return comments, nil
}

func (s *sApiForum) buildCommentTree(ctx context.Context, roots, comments []*entity.ForumComment, sort string) ([]*v1.Comment, error) {
	items, err := s.buildComments(ctx, comments)
	if err != nil {
		return nil, err
	}
	byID := make(map[uint64]*v1.Comment, len(items))
	for _, item := range items {
		byID[item.Id] = item
	}
	for _, item := range items {
		if item.ParentId == 0 {
			continue
		}
		if parent := byID[item.ParentId]; parent != nil {
			parent.Children = append(parent.Children, item)
		}
	}
	if sort == "hot" {
		sortCommentTreeByHot(items)
	}
	result := make([]*v1.Comment, 0, len(roots))
	for _, root := range roots {
		if item := byID[root.Id]; item != nil {
			result = append(result, item)
		}
	}
	return result, nil
}

func sortCommentTreeByHot(comments []*v1.Comment) {
	for _, comment := range comments {
		for index := 1; index < len(comment.Children); index++ {
			for current := index; current > 0 && comment.Children[current].LikeCount > comment.Children[current-1].LikeCount; current-- {
				comment.Children[current], comment.Children[current-1] = comment.Children[current-1], comment.Children[current]
			}
		}
		sortCommentTreeByHot(comment.Children)
	}
}

func (s *sApiForum) buildTopics(ctx context.Context, topics []*entity.ForumTopic) ([]*v1.Topic, error) {
	if len(topics) == 0 {
		return []*v1.Topic{}, nil
	}
	userIDs := make([]uint64, 0, len(topics))
	topicIDs := make([]uint64, 0, len(topics))
	for _, topic := range topics {
		userIDs = append(userIDs, topic.UserId)
		topicIDs = append(topicIDs, topic.Id)
	}
	authors, err := s.loadAuthors(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	likedTopicIDs, favoritedTopicIDs, err := s.loadTopicRelations(ctx, topicIDs, contexts.GetUserId(ctx))
	if err != nil {
		return nil, err
	}
	result := make([]*v1.Topic, 0, len(topics))
	for _, topic := range topics {
		result = append(result, &v1.Topic{
			ForumTopic:  *topic,
			Author:      authors[topic.UserId],
			IsLiked:     likedTopicIDs[topic.Id],
			IsFavorited: favoritedTopicIDs[topic.Id],
		})
	}
	return result, nil
}

func (s *sApiForum) buildComments(ctx context.Context, comments []*entity.ForumComment) ([]*v1.Comment, error) {
	if len(comments) == 0 {
		return []*v1.Comment{}, nil
	}
	userIDs := make([]uint64, 0, len(comments)*2)
	commentIDs := make([]uint64, 0, len(comments))
	for _, comment := range comments {
		userIDs = append(userIDs, comment.UserId)
		commentIDs = append(commentIDs, comment.Id)
		if comment.ReplyUserId > 0 {
			userIDs = append(userIDs, comment.ReplyUserId)
		}
	}
	authors, err := s.loadAuthors(ctx, userIDs)
	if err != nil {
		return nil, err
	}
	likedCommentIDs, err := s.loadCommentLikes(ctx, commentIDs, contexts.GetUserId(ctx))
	if err != nil {
		return nil, err
	}
	result := make([]*v1.Comment, 0, len(comments))
	for _, comment := range comments {
		result = append(result, &v1.Comment{
			ForumComment: *comment,
			Author:       authors[comment.UserId],
			ReplyUser:    authors[comment.ReplyUserId],
			Children:     make([]*v1.Comment, 0),
			IsLiked:      likedCommentIDs[comment.Id],
		})
	}
	return result, nil
}

func (s *sApiForum) loadAuthors(ctx context.Context, userIDs []uint64) (map[uint64]*v1.Author, error) {
	result := make(map[uint64]*v1.Author)
	if len(userIDs) == 0 {
		return result, nil
	}
	var users []*entity.User
	cols := dao.User.Columns()
	if err := dao.User.Ctx(ctx).Fields(cols.Id, cols.Username, cols.Nickname, cols.Avatar).WhereIn(cols.Id, userIDs).Scan(&users); err != nil {
		return nil, gerror.Wrap(err, "获取用户信息失败，请稍后重试")
	}
	for _, user := range users {
		result[uint64(user.Id)] = &v1.Author{Id: uint64(user.Id), Username: user.Username, Nickname: user.Nickname, Avatar: user.Avatar}
	}
	return result, nil
}

func (s *sApiForum) loadCommentLikes(ctx context.Context, commentIDs []uint64, userID int64) (map[uint64]bool, error) {
	likedCommentIDs := make(map[uint64]bool)
	if userID <= 0 || len(commentIDs) == 0 {
		return likedCommentIDs, nil
	}

	cols := dao.ForumInteraction.Columns()
	var interactions []*entity.ForumInteraction
	if err := dao.ForumInteraction.Ctx(ctx).
		Fields(cols.TargetId).
		Where(cols.TargetType, consts.ForumInteractionTargetTypeComment).
		WhereIn(cols.TargetId, commentIDs).
		Where(cols.UserId, userID).
		Where(cols.ActionType, consts.ForumInteractionActionTypeLike).
		Scan(&interactions); err != nil {
		return nil, gerror.Wrap(err, "获取评论点赞状态失败，请稍后重试")
	}
	for _, interaction := range interactions {
		likedCommentIDs[interaction.TargetId] = true
	}
	return likedCommentIDs, nil
}

func (s *sApiForum) loadTopicRelations(ctx context.Context, topicIDs []uint64, userID int64) (map[uint64]bool, map[uint64]bool, error) {
	likedTopicIDs := make(map[uint64]bool)
	favoritedTopicIDs := make(map[uint64]bool)
	if userID <= 0 {
		return likedTopicIDs, favoritedTopicIDs, nil
	}

	cols := dao.ForumInteraction.Columns()
	var interactions []*entity.ForumInteraction
	if err := dao.ForumInteraction.Ctx(ctx).
		Fields(cols.TargetId, cols.ActionType).
		Where(cols.TargetType, consts.ForumInteractionTargetTypeTopic).
		WhereIn(cols.TargetId, topicIDs).
		Where(cols.UserId, userID).
		WhereIn(cols.ActionType, []string{consts.ForumInteractionActionTypeLike, consts.ForumInteractionActionTypeFavorite}).
		Scan(&interactions); err != nil {
		return nil, nil, gerror.Wrap(err, "获取主题互动状态失败，请稍后重试")
	}
	for _, interaction := range interactions {
		switch interaction.ActionType {
		case consts.ForumInteractionActionTypeLike:
			likedTopicIDs[interaction.TargetId] = true
		case consts.ForumInteractionActionTypeFavorite:
			favoritedTopicIDs[interaction.TargetId] = true
		}
	}
	return likedTopicIDs, favoritedTopicIDs, nil
}

func (s *sApiForum) ensureUser(ctx context.Context, userID uint64) error {
	var user *entity.User
	cols := dao.User.Columns()
	if err := dao.User.Ctx(ctx).Fields(cols.Id).WherePri(userID).Scan(&user); err != nil {
		return gerror.Wrap(err, "获取被关注用户失败，请稍后重试")
	}
	if user == nil {
		return gerror.New("被关注用户不存在")
	}
	return nil
}

func (s *sApiForum) ensureCategory(ctx context.Context, id uint64) (*entity.ForumCategory, error) {
	var category *entity.ForumCategory
	cols := dao.ForumCategory.Columns()
	if err := dao.ForumCategory.Ctx(ctx).WherePri(id).Where(cols.Status, consts.ForumCategoryStatusNormal).Scan(&category); err != nil {
		return nil, gerror.Wrap(err, "获取论坛板块失败，请稍后重试")
	}
	if category == nil {
		return nil, gerror.New("板块不存在或已禁用")
	}
	return category, nil
}

func (s *sApiForum) ensureTopic(ctx context.Context, id uint64) error {
	var topic *entity.ForumTopic
	cols := dao.ForumTopic.Columns()
	if err := dao.ForumTopic.Ctx(ctx).WherePri(id).Where(cols.Status, consts.ForumTopicStatusNormal).Scan(&topic); err != nil {
		return gerror.Wrap(err, "获取论坛主题失败，请稍后重试")
	}
	if topic == nil {
		return gerror.New("主题不存在或已不可见")
	}
	return nil
}

func (s *sApiForum) ensureVisibleComment(ctx context.Context, commentID uint64) (*entity.ForumComment, error) {
	var comment *entity.ForumComment
	cols := dao.ForumComment.Columns()
	if err := dao.ForumComment.Ctx(ctx).WherePri(commentID).Where(cols.Status, consts.ForumCommentStatusNormal).WhereNull(cols.DeletedAt).Scan(&comment); err != nil {
		return nil, gerror.Wrap(err, "获取评论失败，请稍后重试")
	}
	if comment == nil {
		return nil, gerror.New("评论不存在或已不可见")
	}
	return comment, nil
}

func (s *sApiForum) ensureComment(ctx context.Context, topicID, commentID uint64) (*entity.ForumComment, error) {
	var comment *entity.ForumComment
	cols := dao.ForumComment.Columns()
	if err := dao.ForumComment.Ctx(ctx).WherePri(commentID).Where(cols.TopicId, topicID).Where(cols.Status, consts.ForumCommentStatusNormal).WhereNull(cols.DeletedAt).Scan(&comment); err != nil {
		return nil, gerror.Wrap(err, "获取父评论失败，请稍后重试")
	}
	if comment == nil {
		return nil, gerror.New("父评论不存在或已不可见")
	}
	return comment, nil
}

func (s *sApiForum) ensureReportTarget(ctx context.Context, targetType string, targetID uint64) error {
	if targetType == "topic" {
		return s.ensureTopic(ctx, targetID)
	}
	if targetType == "comment" {
		var comment *entity.ForumComment
		cols := dao.ForumComment.Columns()
		if err := dao.ForumComment.Ctx(ctx).WherePri(targetID).Where(cols.Status, consts.ForumCommentStatusNormal).WhereNull(cols.DeletedAt).Scan(&comment); err != nil {
			return gerror.Wrap(err, "获取举报评论失败，请稍后重试")
		}
		if comment == nil {
			return gerror.New("评论不存在或已不可见")
		}
		return nil
	}
	return gerror.New("举报对象类型不正确")
}

func requireUser(ctx context.Context) (int64, error) {
	userID := contexts.GetUserId(ctx)
	if userID <= 0 {
		return 0, gerror.New("请先登录")
	}
	return userID, nil
}
