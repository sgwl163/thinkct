// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserRuleDao is the data access object for the table hg_user_rule.
type UserRuleDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserRuleColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserRuleColumns defines and stores column names for the table hg_user_rule.
type UserRuleColumns struct {
	Id           string // ID
	Pid          string // 上级菜单
	Type         string // 类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮
	Title        string // 标题
	Name         string // 规则名称
	Path         string // 路由路径
	Icon         string // 图标
	MenuType     string // 菜单类型:tab=选项卡,link=链接,iframe=Iframe
	Url          string // Url
	Component    string // 组件路径
	NoLoginValid string // 未登录有效:0=否,1=是
	Extend       string // 扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单
	Remark       string // 备注
	Weigh        string // 权重
	Status       string // 状态:0=禁用,1=启用
	CreatedAt    string // 更新时间
	UpdatedAt    string // 创建时间
}

// userRuleColumns holds the columns for the table hg_user_rule.
var userRuleColumns = UserRuleColumns{
	Id:           "id",
	Pid:          "pid",
	Type:         "type",
	Title:        "title",
	Name:         "name",
	Path:         "path",
	Icon:         "icon",
	MenuType:     "menu_type",
	Url:          "url",
	Component:    "component",
	NoLoginValid: "no_login_valid",
	Extend:       "extend",
	Remark:       "remark",
	Weigh:        "weigh",
	Status:       "status",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewUserRuleDao creates and returns a new DAO object for table data access.
func NewUserRuleDao(handlers ...gdb.ModelHandler) *UserRuleDao {
	return &UserRuleDao{
		group:    "default",
		table:    "hg_user_rule",
		columns:  userRuleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserRuleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserRuleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserRuleDao) Columns() UserRuleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserRuleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserRuleDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserRuleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
