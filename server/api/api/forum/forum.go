package forum

import (
	"context"

	v1 "hotgo/api/api/forum/v1"
)

type IForumPublicV1 interface {
	CategoryList(ctx context.Context, req *v1.CategoryListReq) (res *v1.CategoryListRes, err error)
	TopicList(ctx context.Context, req *v1.TopicListReq) (res *v1.TopicListRes, err error)
	TopicDetail(ctx context.Context, req *v1.TopicDetailReq) (res *v1.TopicDetailRes, err error)
	UserProfile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error)
	UserActivityList(ctx context.Context, req *v1.UserActivityListReq) (res *v1.UserActivityListRes, err error)
	CommentList(ctx context.Context, req *v1.CommentListReq) (res *v1.CommentListRes, err error)
}

type IForumV1 interface {
	CreateTopic(ctx context.Context, req *v1.CreateTopicReq) (res *v1.CreateTopicRes, err error)
	MyTopicList(ctx context.Context, req *v1.MyTopicListReq) (res *v1.MyTopicListRes, err error)
	MyTopicDetail(ctx context.Context, req *v1.MyTopicDetailReq) (res *v1.MyTopicDetailRes, err error)
	UpdateTopic(ctx context.Context, req *v1.UpdateTopicReq) (res *v1.UpdateTopicRes, err error)
	CreateComment(ctx context.Context, req *v1.CreateCommentReq) (res *v1.CreateCommentRes, err error)
	DeleteComment(ctx context.Context, req *v1.DeleteCommentReq) (res *v1.DeleteCommentRes, err error)
	ToggleTopicLike(ctx context.Context, req *v1.ToggleTopicLikeReq) (res *v1.ToggleTopicLikeRes, err error)
	ToggleTopicFavorite(ctx context.Context, req *v1.ToggleTopicFavoriteReq) (res *v1.ToggleTopicFavoriteRes, err error)
	ToggleCommentLike(ctx context.Context, req *v1.ToggleCommentLikeReq) (res *v1.ToggleCommentLikeRes, err error)
	ToggleUserFollow(ctx context.Context, req *v1.ToggleUserFollowReq) (res *v1.ToggleUserFollowRes, err error)
	CreateReport(ctx context.Context, req *v1.CreateReportReq) (res *v1.CreateReportRes, err error)
}
