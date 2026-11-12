// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// WebCarousel is the golang structure of table hg_web_carousel for DAO operations like Where/Data.
type WebCarousel struct {
	g.Meta      `orm:"table:hg_web_carousel, do:true"`
	Id          any         // 轮播ID
	Title       any         // 标题
	Description any         // 描述
	CtaText     any         // 按钮文字
	CtaLink     any         // 按钮链接
	Color       any         // 配色主题
	Image       any         // 图片
	Alt         any         // 图片替代文本
	AspectRatio any         // 图片宽高比
	Sort        any         // 排序
	Status      any         // 状态
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
	DeletedAt   *gtime.Time // 删除时间
}
