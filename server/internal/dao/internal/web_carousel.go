// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WebCarouselDao is the data access object for the table hg_web_carousel.
type WebCarouselDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  WebCarouselColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// WebCarouselColumns defines and stores column names for the table hg_web_carousel.
type WebCarouselColumns struct {
	Id          string // 轮播ID
	Title       string // 标题
	Description string // 描述
	CtaText     string // 按钮文字
	CtaLink     string // 按钮链接
	Color       string // 配色主题
	Image       string // 图片
	Alt         string // 图片替代文本
	AspectRatio string // 图片宽高比
	Sort        string // 排序
	Status      string // 状态
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
	DeletedAt   string // 删除时间
}

// webCarouselColumns holds the columns for the table hg_web_carousel.
var webCarouselColumns = WebCarouselColumns{
	Id:          "id",
	Title:       "title",
	Description: "description",
	CtaText:     "cta_text",
	CtaLink:     "cta_link",
	Color:       "color",
	Image:       "image",
	Alt:         "alt",
	AspectRatio: "aspect_ratio",
	Sort:        "sort",
	Status:      "status",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
	DeletedAt:   "deleted_at",
}

// NewWebCarouselDao creates and returns a new DAO object for table data access.
func NewWebCarouselDao(handlers ...gdb.ModelHandler) *WebCarouselDao {
	return &WebCarouselDao{
		group:    "default",
		table:    "hg_web_carousel",
		columns:  webCarouselColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WebCarouselDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WebCarouselDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WebCarouselDao) Columns() WebCarouselColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WebCarouselDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WebCarouselDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *WebCarouselDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
