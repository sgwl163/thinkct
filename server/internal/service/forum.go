// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "hotgo/api/api/forum/v1"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/forumin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IApiForum interface {
		CategoryList(ctx context.Context, req *v1.CategoryListReq) (res *v1.CategoryListRes, err error)
		TopicList(ctx context.Context, req *v1.TopicListReq) (res *v1.TopicListRes, err error)
		TopicDetail(ctx context.Context, req *v1.TopicDetailReq) (res *v1.TopicDetailRes, err error)
		UserProfile(ctx context.Context, req *v1.UserProfileReq) (res *v1.UserProfileRes, err error)
		UserActivityList(ctx context.Context, req *v1.UserActivityListReq) (res *v1.UserActivityListRes, err error)
		CommentList(ctx context.Context, req *v1.CommentListReq) (res *v1.CommentListRes, err error)
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
	IForumCategory interface {
		// Model 社区板块ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取社区板块列表
		List(ctx context.Context, in *forumin.CategoryListInp) (list []*forumin.CategoryListModel, totalCount int, err error)
		// Export 导出社区板块
		Export(ctx context.Context, in *forumin.CategoryListInp) (err error)
		// Edit 修改/新增社区板块
		Edit(ctx context.Context, in *forumin.CategoryEditInp) (err error)
		// Delete 删除社区板块
		Delete(ctx context.Context, in *forumin.CategoryDeleteInp) (err error)
		// MaxSort 获取社区板块最大排序
		MaxSort(ctx context.Context, in *forumin.CategoryMaxSortInp) (res *forumin.CategoryMaxSortModel, err error)
		// View 获取社区板块指定信息
		View(ctx context.Context, in *forumin.CategoryViewInp) (res *forumin.CategoryViewModel, err error)
		// Status 更新社区板块状态
		Status(ctx context.Context, in *forumin.CategoryStatusInp) (err error)
	}
	IForumComment interface {
		// Model 评论管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取评论管理列表
		List(ctx context.Context, in *forumin.CommentListInp) (list []*forumin.CommentListModel, totalCount int, err error)
		// Export 导出评论管理
		Export(ctx context.Context, in *forumin.CommentListInp) (err error)
		// Edit 修改/新增评论管理
		Edit(ctx context.Context, in *forumin.CommentEditInp) (err error)
		// Delete 删除评论管理
		Delete(ctx context.Context, in *forumin.CommentDeleteInp) (err error)
		// View 获取评论管理指定信息
		View(ctx context.Context, in *forumin.CommentViewInp) (res *forumin.CommentViewModel, err error)
	}
	IForumForumCategory interface {
		// Model 社区板块ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取社区板块列表
		List(ctx context.Context, in *forumin.ForumCategoryListInp) (list []*forumin.ForumCategoryListModel, totalCount int, err error)
		// Export 导出社区板块
		Export(ctx context.Context, in *forumin.ForumCategoryListInp) (err error)
		// Edit 修改/新增社区板块
		Edit(ctx context.Context, in *forumin.ForumCategoryEditInp) (err error)
		// Delete 删除社区板块
		Delete(ctx context.Context, in *forumin.ForumCategoryDeleteInp) (err error)
		// MaxSort 获取社区板块最大排序
		MaxSort(ctx context.Context, in *forumin.ForumCategoryMaxSortInp) (res *forumin.ForumCategoryMaxSortModel, err error)
		// View 获取社区板块指定信息
		View(ctx context.Context, in *forumin.ForumCategoryViewInp) (res *forumin.ForumCategoryViewModel, err error)
		// Status 更新社区板块状态
		Status(ctx context.Context, in *forumin.ForumCategoryStatusInp) (err error)
	}
	IForumReport interface {
		// Model 社区论坛举报表ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取社区论坛举报表列表
		List(ctx context.Context, in *forumin.ReportListInp) (list []*forumin.ReportListModel, totalCount int, err error)
		// Export 导出社区论坛举报表
		Export(ctx context.Context, in *forumin.ReportListInp) (err error)
		// Edit 修改/新增社区论坛举报表
		Edit(ctx context.Context, in *forumin.ReportEditInp) (err error)
		// Delete 删除社区论坛举报表
		Delete(ctx context.Context, in *forumin.ReportDeleteInp) (err error)
		// View 获取社区论坛举报表指定信息
		View(ctx context.Context, in *forumin.ReportViewInp) (res *forumin.ReportViewModel, err error)
	}
	IForumTopic interface {
		// Model 社区主题ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取社区主题列表
		List(ctx context.Context, in *forumin.TopicListInp) (list []*forumin.TopicListModel, totalCount int, err error)
		// Export 导出社区主题
		Export(ctx context.Context, in *forumin.TopicListInp) (err error)
		// Edit 修改/新增社区主题
		Edit(ctx context.Context, in *forumin.TopicEditInp) (err error)
		// Delete 删除社区主题
		Delete(ctx context.Context, in *forumin.TopicDeleteInp) (err error)
		// View 获取社区主题指定信息
		View(ctx context.Context, in *forumin.TopicViewInp) (res *forumin.TopicViewModel, err error)
	}
)

var (
	localApiForum           IApiForum
	localForumCategory      IForumCategory
	localForumComment       IForumComment
	localForumForumCategory IForumForumCategory
	localForumReport        IForumReport
	localForumTopic         IForumTopic
)

func ApiForum() IApiForum {
	if localApiForum == nil {
		panic("implement not found for interface IApiForum, forgot register?")
	}
	return localApiForum
}

func RegisterApiForum(i IApiForum) {
	localApiForum = i
}

func ForumCategory() IForumCategory {
	if localForumCategory == nil {
		panic("implement not found for interface IForumCategory, forgot register?")
	}
	return localForumCategory
}

func RegisterForumCategory(i IForumCategory) {
	localForumCategory = i
}

func ForumComment() IForumComment {
	if localForumComment == nil {
		panic("implement not found for interface IForumComment, forgot register?")
	}
	return localForumComment
}

func RegisterForumComment(i IForumComment) {
	localForumComment = i
}

func ForumForumCategory() IForumForumCategory {
	if localForumForumCategory == nil {
		panic("implement not found for interface IForumForumCategory, forgot register?")
	}
	return localForumForumCategory
}

func RegisterForumForumCategory(i IForumForumCategory) {
	localForumForumCategory = i
}

func ForumReport() IForumReport {
	if localForumReport == nil {
		panic("implement not found for interface IForumReport, forgot register?")
	}
	return localForumReport
}

func RegisterForumReport(i IForumReport) {
	localForumReport = i
}

func ForumTopic() IForumTopic {
	if localForumTopic == nil {
		panic("implement not found for interface IForumTopic, forgot register?")
	}
	return localForumTopic
}

func RegisterForumTopic(i IForumTopic) {
	localForumTopic = i
}
