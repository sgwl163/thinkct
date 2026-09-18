// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WebArticleDao is the data access object for the table hg_web_article.
type WebArticleDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  WebArticleColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// WebArticleColumns defines and stores column names for the table hg_web_article.
type WebArticleColumns struct {
	Id          string // ID
	UserId      string // 用户ID
	AdminId     string // 管理员ID
	ChannelId   string // 频道
	ChannelIds  string // 副频道
	Diyname     string // 自定义名称
	Title       string // 标题
	TitleStyle  string // 标题样式
	Flag        string // 标志:top=置顶,hot=热门,recommend=推荐,new=最新
	Images      string // 封面图片
	Content     string // 内容
	Seotitle    string // SEO标题
	Keywords    string // 关键词
	Description string // 描述
	Tags        string // 标签
	Url         string // 外部链接
	Target      string // 跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口
	Views       string // 浏览量
	Comments    string // 评论量
	Likes       string // 点赞数
	Dislikes    string // 点踩数
	Weigh       string // 权重
	Status      string // 状态:normal=正常,unaudited=待审核,refused=已拒绝,offline=已下线
	PublishTime string // 发布时间
	UpdateTime  string // 修改时间
	CreateTime  string // 创建时间
}

// webArticleColumns holds the columns for the table hg_web_article.
var webArticleColumns = WebArticleColumns{
	Id:          "id",
	UserId:      "user_id",
	AdminId:     "admin_id",
	ChannelId:   "channel_id",
	ChannelIds:  "channel_ids",
	Diyname:     "diyname",
	Title:       "title",
	TitleStyle:  "title_style",
	Flag:        "flag",
	Images:      "images",
	Content:     "content",
	Seotitle:    "seotitle",
	Keywords:    "keywords",
	Description: "description",
	Tags:        "tags",
	Url:         "url",
	Target:      "target",
	Views:       "views",
	Comments:    "comments",
	Likes:       "likes",
	Dislikes:    "dislikes",
	Weigh:       "weigh",
	Status:      "status",
	PublishTime: "publish_time",
	UpdateTime:  "update_time",
	CreateTime:  "create_time",
}

// NewWebArticleDao creates and returns a new DAO object for table data access.
func NewWebArticleDao(handlers ...gdb.ModelHandler) *WebArticleDao {
	return &WebArticleDao{
		group:    "default",
		table:    "hg_web_article",
		columns:  webArticleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WebArticleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WebArticleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WebArticleDao) Columns() WebArticleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WebArticleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WebArticleDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *WebArticleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
