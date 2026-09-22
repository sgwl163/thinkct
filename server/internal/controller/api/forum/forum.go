package forum

import (
	"context"

	v1 "hotgo/api/api/forum/v1"
	"hotgo/internal/service"
)

func (c *cPublicV1) CategoryList(ctx context.Context, req *v1.CategoryListReq) (res *v1.CategoryListRes, err error) {
	return service.ApiForum().CategoryList(ctx, req)
}

func (c *cPublicV1) TopicList(ctx context.Context, req *v1.TopicListReq) (res *v1.TopicListRes, err error) {
	return service.ApiForum().TopicList(ctx, req)
}

func (c *cPublicV1) TopicDetail(ctx context.Context, req *v1.TopicDetailReq) (res *v1.TopicDetailRes, err error) {
	return service.ApiForum().TopicDetail(ctx, req)
}

func (c *cPublicV1) UserProfile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error) {
	return service.ApiForum().UserProfile(ctx, req)
}

func (c *cPublicV1) UserActivityList(ctx context.Context, req *v1.UserActivityListReq) (res *v1.UserActivityListRes, err error) {
	return service.ApiForum().UserActivityList(ctx, req)
}

func (c *cPublicV1) CommentList(ctx context.Context, req *v1.CommentListReq) (res *v1.CommentListRes, err error) {
	return service.ApiForum().CommentList(ctx, req)
}

func (c *cV1) CreateTopic(ctx context.Context, req *v1.CreateTopicReq) (res *v1.CreateTopicRes, err error) {
	return service.ApiForum().CreateTopic(ctx, req)
}

func (c *cV1) MyTopicList(ctx context.Context, req *v1.MyTopicListReq) (res *v1.MyTopicListRes, err error) {
	return service.ApiForum().MyTopicList(ctx, req)
}

func (c *cV1) MyTopicDetail(ctx context.Context, req *v1.MyTopicDetailReq) (res *v1.MyTopicDetailRes, err error) {
	return service.ApiForum().MyTopicDetail(ctx, req)
}

func (c *cV1) UpdateTopic(ctx context.Context, req *v1.UpdateTopicReq) (res *v1.UpdateTopicRes, err error) {
	return service.ApiForum().UpdateTopic(ctx, req)
}

func (c *cV1) CreateComment(ctx context.Context, req *v1.CreateCommentReq) (res *v1.CreateCommentRes, err error) {
	return service.ApiForum().CreateComment(ctx, req)
}

func (c *cV1) DeleteComment(ctx context.Context, req *v1.DeleteCommentReq) (res *v1.DeleteCommentRes, err error) {
	return service.ApiForum().DeleteComment(ctx, req)
}

func (c *cV1) ToggleTopicLike(ctx context.Context, req *v1.ToggleTopicLikeReq) (res *v1.ToggleTopicLikeRes, err error) {
	return service.ApiForum().ToggleTopicLike(ctx, req)
}

func (c *cV1) ToggleTopicFavorite(ctx context.Context, req *v1.ToggleTopicFavoriteReq) (res *v1.ToggleTopicFavoriteRes, err error) {
	return service.ApiForum().ToggleTopicFavorite(ctx, req)
}

func (c *cV1) ToggleCommentLike(ctx context.Context, req *v1.ToggleCommentLikeReq) (res *v1.ToggleCommentLikeRes, err error) {
	return service.ApiForum().ToggleCommentLike(ctx, req)
}

func (c *cV1) ToggleUserFollow(ctx context.Context, req *v1.ToggleUserFollowReq) (res *v1.ToggleUserFollowRes, err error) {
	return service.ApiForum().ToggleUserFollow(ctx, req)
}

func (c *cV1) CreateReport(ctx context.Context, req *v1.CreateReportReq) (res *v1.CreateReportRes, err error) {
	return service.ApiForum().CreateReport(ctx, req)
}
