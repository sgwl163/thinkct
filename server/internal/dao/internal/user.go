// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for the table hg_user.
type UserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserColumns defines and stores column names for the table hg_user.
type UserColumns struct {
	Id            string // ID
	GroupId       string // 分组ID
	Username      string // 用户名
	Nickname      string // 昵称
	Email         string // 邮箱
	Mobile        string // 手机
	Avatar        string // 头像
	Gender        string // 性别:0=未知,1=男,2=女
	Birthday      string // 生日
	Money         string // 余额
	Score         string // 积分
	LastLoginTime string // 上次登录时间
	LastLoginIp   string // 上次登录IP
	LoginFailure  string // 登录失败次数
	JoinIp        string // 加入IP
	JoinTime      string // 加入时间
	Motto         string // 签名
	Password      string // 密码
	Salt          string // 密码盐（废弃待删）
	Status        string // 状态:enable=启用,disable=禁用
	UpdateTime    string // 更新时间
	CreateTime    string // 创建时间
}

// userColumns holds the columns for the table hg_user.
var userColumns = UserColumns{
	Id:            "id",
	GroupId:       "group_id",
	Username:      "username",
	Nickname:      "nickname",
	Email:         "email",
	Mobile:        "mobile",
	Avatar:        "avatar",
	Gender:        "gender",
	Birthday:      "birthday",
	Money:         "money",
	Score:         "score",
	LastLoginTime: "last_login_time",
	LastLoginIp:   "last_login_ip",
	LoginFailure:  "login_failure",
	JoinIp:        "join_ip",
	JoinTime:      "join_time",
	Motto:         "motto",
	Password:      "password",
	Salt:          "salt",
	Status:        "status",
	UpdateTime:    "update_time",
	CreateTime:    "create_time",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao(handlers ...gdb.ModelHandler) *UserDao {
	return &UserDao{
		group:    "default",
		table:    "hg_user",
		columns:  userColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
