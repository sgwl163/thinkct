package v1

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type Author struct {
	Id       uint64 `json:"id" dc:"用户ID"`
	Username string `json:"username" dc:"用户名"`
	Nickname string `json:"nickname" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
}

type UserProfile struct {
	Author         *Author `json:"author" dc:"作者信息"`
	Motto          string  `json:"motto" dc:"个性签名"`
	JoinTime       uint64  `json:"joinTime" dc:"加入时间"`
	ArticleCount   uint    `json:"articleCount" dc:"已发布主题数"`
	ViewCount      uint    `json:"viewCount" dc:"主题总阅读数"`
	FollowingCount int     `json:"followingCount" dc:"关注数"`
	FollowerCount  int     `json:"followerCount" dc:"粉丝数"`
	FavoriteCount  int     `json:"favoriteCount" dc:"收藏主题数"`
	LikeCount      int     `json:"likeCount" dc:"点赞主题数"`
	IsFollowed     bool    `json:"isFollowed" dc:"当前用户是否已关注"`
	IsSelf         bool    `json:"isSelf" dc:"是否为当前用户本人"`
}

type UserActivity struct {
	Id        uint64      `json:"id" dc:"动态来源ID"`
	Type      string      `json:"type" dc:"动态类型:topic,like,favorite,follow"`
	CreatedAt *gtime.Time `json:"createdAt" dc:"动态时间"`
	Topic     *Topic      `json:"topic" dc:"关联主题"`
	User      *Author     `json:"user" dc:"关联用户"`
}

type Topic struct {
	entity.ForumTopic
	Author      *Author `json:"author" dc:"主题作者"`
	IsLiked     bool    `json:"isLiked" dc:"当前用户是否已点赞"`
	IsFavorited bool    `json:"isFavorited" dc:"当前用户是否已收藏"`
}

type TopicNavigation struct {
	Id    uint64 `json:"id" dc:"主题ID"`
	Title string `json:"title" dc:"主题标题"`
}

type Comment struct {
	entity.ForumComment
	Author    *Author    `json:"author" dc:"评论作者"`
	ReplyUser *Author    `json:"replyUser" dc:"被回复用户"`
	Children  []*Comment `json:"children" dc:"子回复"`
	IsLiked   bool       `json:"isLiked" dc:"当前用户是否已点赞"`
}

type CategoryListReq struct {
	g.Meta `path:"/forum/category/list" method:"get" tags:"社区论坛" summary:"获取论坛板块列表"`
}

type CategoryListRes struct {
	List []*entity.ForumCategory `json:"list" dc:"板块列表"`
}

type TopicListReq struct {
	g.Meta `path:"/forum/topic/list" method:"get" tags:"社区论坛" summary:"获取论坛主题列表"`
	form.PageReq
	CategoryId uint64 `json:"categoryId" dc:"板块ID，可选"`
	Keyword    string `json:"keyword" v:"length:0,100#搜索关键词不能超过100个字符" dc:"标题、关键词或摘要关键词，可选"`
	Sort       string `json:"sort" v:"in:latest,hot,recommend#排序方式不正确" d:"latest" dc:"排序:latest最新,hot热门,recommend推荐"`
}

type TopicListRes struct {
	form.PageRes
	List []*Topic `json:"list" dc:"主题列表"`
}

type TopicDetailReq struct {
	g.Meta `path:"/forum/topic/detail" method:"get" tags:"社区论坛" summary:"获取论坛主题详情"`
	Id     uint64 `json:"id" v:"required|min:1#主题ID不能为空|主题ID不正确" dc:"主题ID"`
}

type TopicDetailRes struct {
	Topic *Topic           `json:"topic" dc:"主题详情"`
	Prev  *TopicNavigation `json:"prev" dc:"上一篇主题"`
	Next  *TopicNavigation `json:"next" dc:"下一篇主题"`
}

type UserProfileReq struct {
	g.Meta `path:"/forum/user/profile" method:"get" tags:"社区论坛" summary:"获取论坛作者资料"`
	UserId uint64 `json:"userId" v:"required|min:1#用户ID不能为空|用户ID不正确" dc:"用户ID"`
}

type UserProfileRes struct {
	Profile *UserProfile `json:"profile" dc:"作者资料"`
}

type UserActivityListReq struct {
	g.Meta `path:"/forum/user/activity" method:"get" tags:"社区论坛" summary:"获取论坛用户动态"`
	form.PageReq
	UserId uint64 `json:"userId" v:"required|min:1#用户ID不能为空|用户ID不正确" dc:"用户ID"`
	Type   string `json:"type" v:"in:topic,like,favorite,follow#动态类型不正确" dc:"动态类型，可选"`
}

type UserActivityListRes struct {
	form.PageRes
	List []*UserActivity `json:"list" dc:"动态列表"`
}

type CommentListReq struct {
	g.Meta `path:"/forum/comment/list" method:"get" tags:"社区论坛" summary:"获取主题评论列表"`
	form.PageReq
	TopicId uint64 `json:"topicId" v:"required|min:1#主题ID不能为空|主题ID不正确" dc:"主题ID"`
	Sort    string `json:"sort" v:"in:latest,hot#排序方式不正确" d:"latest" dc:"排序:latest最新,hot最热"`
}

type CommentListRes struct {
	form.PageRes
	List []*Comment `json:"list" dc:"评论列表"`
}

type CreateTopicReq struct {
	g.Meta      `path:"/forum/topic/create" method:"post" tags:"社区论坛" summary:"发布论坛主题"`
	CategoryId  uint64 `json:"categoryId" v:"required|min:1#板块ID不能为空|板块ID不正确" dc:"板块ID"`
	Title       string `json:"title" v:"required|length:2,255#主题标题不能为空|主题标题长度需在2到255个字符之间" dc:"主题标题"`
	Keywords    string `json:"keywords" v:"length:0,500#主题关键词不能超过500个字符" dc:"主题关键词，英文逗号分隔"`
	Description string `json:"description" v:"length:0,500#主题摘要不能超过500个字符" dc:"主题摘要"`
	Content     string `json:"content" v:"required|length:2,50000#主题内容不能为空|主题内容长度需在2到50000个字符之间" dc:"主题内容"`
	Images      string `json:"images" v:"length:0,5000#主题图片不能超过5000个字符" dc:"主题图片，英文逗号分隔"`
}

type CreateTopicRes struct {
	Topic *Topic `json:"topic" dc:"主题详情"`
}

type MyTopicListReq struct {
	g.Meta `path:"/forum/topic/myList" method:"get" tags:"社区论坛" summary:"获取我的论坛主题列表"`
	form.PageReq
	Status string `json:"status" v:"in:normal,pending,rejected,hidden,closed#主题状态不正确" dc:"主题状态，可选"`
}

type MyTopicListRes struct {
	form.PageRes
	List []*Topic `json:"list" dc:"主题列表"`
}

type MyTopicDetailReq struct {
	g.Meta `path:"/forum/topic/myDetail" method:"get" tags:"社区论坛" summary:"获取我的论坛主题详情"`
	Id     uint64 `json:"id" v:"required|min:1#主题ID不能为空|主题ID不正确" dc:"主题ID"`
}

type MyTopicDetailRes struct {
	Topic *Topic `json:"topic" dc:"主题详情"`
}

type UpdateTopicReq struct {
	g.Meta      `path:"/forum/topic/update" method:"post" tags:"社区论坛" summary:"编辑我的论坛主题"`
	Id          uint64 `json:"id" v:"required|min:1#主题ID不能为空|主题ID不正确" dc:"主题ID"`
	CategoryId  uint64 `json:"categoryId" v:"required|min:1#板块ID不能为空|板块ID不正确" dc:"板块ID"`
	Title       string `json:"title" v:"required|length:2,255#主题标题不能为空|主题标题长度需在2到255个字符之间" dc:"主题标题"`
	Keywords    string `json:"keywords" v:"length:0,500#主题关键词不能超过500个字符" dc:"主题关键词，英文逗号分隔"`
	Description string `json:"description" v:"length:0,500#主题摘要不能超过500个字符" dc:"主题摘要"`
	Content     string `json:"content" v:"required|length:2,50000#主题内容不能为空|主题内容长度需在2到50000个字符之间" dc:"主题内容"`
	Images      string `json:"images" v:"length:0,5000#主题图片不能超过5000个字符" dc:"主题图片，英文逗号分隔"`
}

type UpdateTopicRes struct {
	Topic *Topic `json:"topic" dc:"主题详情"`
}

type CreateCommentReq struct {
	g.Meta   `path:"/forum/comment/create" method:"post" tags:"社区论坛" summary:"发布论坛评论"`
	TopicId  uint64 `json:"topicId" v:"required|min:1#主题ID不能为空|主题ID不正确" dc:"主题ID"`
	ParentId uint64 `json:"parentId" dc:"父评论ID，回复时填写"`
	Content  string `json:"content" v:"required|length:1,2000#评论内容不能为空|评论内容不能超过2000个字符" dc:"评论内容"`
}

type CreateCommentRes struct {
	Comment *Comment `json:"comment" dc:"评论详情"`
}

type DeleteCommentReq struct {
	g.Meta    `path:"/forum/comment/delete" method:"post" tags:"社区论坛" summary:"删除论坛评论"`
	CommentId uint64 `json:"commentId" v:"required|min:1#评论ID不能为空|评论ID不正确" dc:"评论ID"`
}

type DeleteCommentRes struct {
	CommentCount uint `json:"commentCount" dc:"主题最新评论数"`
}

type ToggleTopicLikeReq struct {
	g.Meta  `path:"/forum/topic/toggleLike" method:"post" tags:"社区论坛" summary:"点赞或取消点赞主题"`
	TopicId uint64 `json:"topicId" v:"required|min:1#主题ID不能为空|主题ID不正确" dc:"主题ID"`
}

type ToggleTopicLikeRes struct {
	Liked     bool `json:"liked" dc:"是否已点赞"`
	LikeCount uint `json:"likeCount" dc:"最新点赞数"`
}

type ToggleTopicFavoriteReq struct {
	g.Meta  `path:"/forum/topic/toggleFavorite" method:"post" tags:"社区论坛" summary:"收藏或取消收藏主题"`
	TopicId uint64 `json:"topicId" v:"required|min:1#主题ID不能为空|主题ID不正确" dc:"主题ID"`
}

type ToggleTopicFavoriteRes struct {
	Favorited     bool `json:"favorited" dc:"是否已收藏"`
	FavoriteCount uint `json:"favoriteCount" dc:"最新收藏数"`
}

type ToggleCommentLikeReq struct {
	g.Meta    `path:"/forum/comment/toggleLike" method:"post" tags:"社区论坛" summary:"点赞或取消点赞评论"`
	CommentId uint64 `json:"commentId" v:"required|min:1#评论ID不能为空|评论ID不正确" dc:"评论ID"`
}

type ToggleCommentLikeRes struct {
	Liked     bool `json:"liked" dc:"是否已点赞"`
	LikeCount uint `json:"likeCount" dc:"最新点赞数"`
}

type ToggleUserFollowReq struct {
	g.Meta       `path:"/forum/user/toggleFollow" method:"post" tags:"社区论坛" summary:"关注或取消关注用户"`
	FollowUserId uint64 `json:"followUserId" v:"required|min:1#被关注用户ID不能为空|被关注用户ID不正确" dc:"被关注用户ID"`
}

type ToggleUserFollowRes struct {
	Followed      bool `json:"followed" dc:"是否已关注"`
	FollowerCount int  `json:"followerCount" dc:"被关注用户最新粉丝数"`
}

type CreateReportReq struct {
	g.Meta      `path:"/forum/report/create" method:"post" tags:"社区论坛" summary:"提交论坛内容举报"`
	TargetType  string `json:"targetType" v:"required|in:topic,comment#举报对象类型不能为空|举报对象类型不正确" dc:"举报对象类型:topic,comment"`
	TargetId    uint64 `json:"targetId" v:"required|min:1#举报对象ID不能为空|举报对象ID不正确" dc:"举报对象ID"`
	Reason      string `json:"reason" v:"required|length:1,100#举报原因不能为空|举报原因不能超过100个字符" dc:"举报原因"`
	Description string `json:"description" v:"length:0,1000#补充说明不能超过1000个字符" dc:"补充说明"`
}

type CreateReportRes struct{}
