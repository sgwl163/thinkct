// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserRechargeDao is the data access object for the table hg_user_recharge.
type UserRechargeDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  UserRechargeColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// UserRechargeColumns defines and stores column names for the table hg_user_recharge.
type UserRechargeColumns struct {
	Id        string // 充值订单ID
	UserId    string // 前台用户ID
	OrderSn   string // 充值订单号
	Money     string // 充值金额，整数元，与会员余额同单位
	PayType   string // 支付渠道
	Status    string // pending=待支付,completed=已到账
	PayAt     string // 支付时间
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// userRechargeColumns holds the columns for the table hg_user_recharge.
var userRechargeColumns = UserRechargeColumns{
	Id:        "id",
	UserId:    "user_id",
	OrderSn:   "order_sn",
	Money:     "money",
	PayType:   "pay_type",
	Status:    "status",
	PayAt:     "pay_at",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewUserRechargeDao creates and returns a new DAO object for table data access.
func NewUserRechargeDao(handlers ...gdb.ModelHandler) *UserRechargeDao {
	return &UserRechargeDao{
		group:    "default",
		table:    "hg_user_recharge",
		columns:  userRechargeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserRechargeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserRechargeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserRechargeDao) Columns() UserRechargeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserRechargeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserRechargeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserRechargeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
