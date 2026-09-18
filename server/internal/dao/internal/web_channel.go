// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WebChannelDao is the data access object for the table hg_web_channel.
type WebChannelDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  WebChannelColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// WebChannelColumns defines and stores column names for the table hg_web_channel.
type WebChannelColumns struct {
	Id               string // ID
	Pid              string // 上级频道
	Level            string // tree
	Tree             string // 关系树
	Model            string // 模型
	Type             string // 类型:cover=封面频道,list=列表频道,link=跳转链接
	Name             string // 频道名称
	Diyname          string // 自定义名称
	Template         string // 模板地址
	Url              string // 跳转链接
	Target           string // 跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口
	Seotitle         string // SEO标题
	Keywords         string // 关键词
	Description      string // 描述
	AllowVisitGroups string // 访问策略:all=全部,user=仅会员
	Weigh            string // 权重
	Status           string // 状态:0=禁用,1=启用
	UpdateTime       string // 修改时间
	CreateTime       string // 创建时间
}

// webChannelColumns holds the columns for the table hg_web_channel.
var webChannelColumns = WebChannelColumns{
	Id:               "id",
	Pid:              "pid",
	Level:            "level",
	Tree:             "tree",
	Model:            "model",
	Type:             "type",
	Name:             "name",
	Diyname:          "diyname",
	Template:         "template",
	Url:              "url",
	Target:           "target",
	Seotitle:         "seotitle",
	Keywords:         "keywords",
	Description:      "description",
	AllowVisitGroups: "allow_visit_groups",
	Weigh:            "weigh",
	Status:           "status",
	UpdateTime:       "update_time",
	CreateTime:       "create_time",
}

// NewWebChannelDao creates and returns a new DAO object for table data access.
func NewWebChannelDao(handlers ...gdb.ModelHandler) *WebChannelDao {
	return &WebChannelDao{
		group:    "default",
		table:    "hg_web_channel",
		columns:  webChannelColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WebChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WebChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WebChannelDao) Columns() WebChannelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WebChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WebChannelDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *WebChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
