package consts

import (
	"hotgo/internal/library/dict"
	"hotgo/internal/model"
)

const (
	ForumCategoryStatusNormal   = 1
	ForumCategoryStatusDisabled = 2

	ForumTopicStatusNormal   = "normal"
	ForumTopicStatusPending  = "pending"
	ForumTopicStatusRejected = "rejected"
	ForumTopicStatusHidden   = "hidden"
	ForumTopicStatusClosed   = "closed"

	ForumCommentStatusNormal   = "normal"
	ForumCommentStatusPending  = "pending"
	ForumCommentStatusRejected = "rejected"
	ForumCommentStatusHidden   = "hidden"

	ForumInteractionTargetTypeTopic   = "topic"
	ForumInteractionTargetTypeComment = "comment"

	ForumInteractionActionTypeLike     = "like"
	ForumInteractionActionTypeFavorite = "favorite"

	ForumUserActivityTypeTopic    = "topic"
	ForumUserActivityTypeLike     = "like"
	ForumUserActivityTypeFavorite = "favorite"
	ForumUserActivityTypeFollow   = "follow"

	ForumReportStatusPending   = "pending"
	ForumReportStatusProcessed = "processed"
	ForumReportStatusRejected  = "rejected"
)

var ForumTopicStatusOptions = []*model.Option{
	dict.GenSuccessOption(ForumTopicStatusNormal, "正常"),
	dict.GenWarningOption(ForumTopicStatusPending, "待审核"),
	dict.GenErrorOption(ForumTopicStatusRejected, "已驳回"),
	dict.GenErrorOption(ForumTopicStatusHidden, "已隐藏"),
	dict.GenDefaultOption(ForumTopicStatusClosed, "已关闭"),
}

var ForumCommentStatusOptions = []*model.Option{
	dict.GenSuccessOption(ForumCommentStatusNormal, "正常"),
	dict.GenWarningOption(ForumCommentStatusPending, "待审核"),
	dict.GenErrorOption(ForumCommentStatusRejected, "已驳回"),
	dict.GenErrorOption(ForumCommentStatusHidden, "已隐藏"),
}

var ForumReportStatusOptions = []*model.Option{
	dict.GenWarningOption(ForumReportStatusPending, "待处理"),
	dict.GenSuccessOption(ForumReportStatusProcessed, "已处理"),
	dict.GenDefaultOption(ForumReportStatusRejected, "已驳回"),
}

func init() {
	dict.RegisterEnums("ForumTopicStatusOptions", "论坛主题状态选项", ForumTopicStatusOptions)
	dict.RegisterEnums("ForumCommentStatusOptions", "论坛评论状态选项", ForumCommentStatusOptions)
	dict.RegisterEnums("ForumReportStatusOptions", "论坛举报状态选项", ForumReportStatusOptions)
}
