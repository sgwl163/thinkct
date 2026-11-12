// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WebCarousel is the golang structure for table web_carousel.
type WebCarousel struct {
	Id          int64       `json:"id"          orm:"id"           description:"轮播ID"`
	Title       string      `json:"title"       orm:"title"        description:"标题"`
	Description string      `json:"description" orm:"description"  description:"描述"`
	CtaText     string      `json:"ctaText"     orm:"cta_text"     description:"按钮文字"`
	CtaLink     string      `json:"ctaLink"     orm:"cta_link"     description:"按钮链接"`
	Color       string      `json:"color"       orm:"color"        description:"配色主题"`
	Image       string      `json:"image"       orm:"image"        description:"图片"`
	Alt         string      `json:"alt"         orm:"alt"          description:"图片替代文本"`
	AspectRatio string      `json:"aspectRatio" orm:"aspect_ratio" description:"图片宽高比"`
	Sort        int         `json:"sort"        orm:"sort"         description:"排序"`
	Status      int         `json:"status"      orm:"status"       description:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"   description:"更新时间"`
	DeletedAt   *gtime.Time `json:"deletedAt"   orm:"deleted_at"   description:"删除时间"`
}
