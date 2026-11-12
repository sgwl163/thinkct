// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// WebNavMenu is the golang structure for table web_nav_menu.
type WebNavMenu struct {
	Id        int64       `json:"id"        orm:"id"         description:"导航ID"`
	Pid       int64       `json:"pid"       orm:"pid"        description:"上级ID"`
	Level     int         `json:"level"     orm:"level"      description:"关系树级别"`
	Tree      string      `json:"tree"      orm:"tree"       description:"关系树"`
	Name      string      `json:"name"      orm:"name"       description:"导航名称"`
	Url       string      `json:"url"       orm:"url"        description:"链接地址"`
	Image     string      `json:"image"     orm:"image"      description:"图片"`
	ImageAlt  string      `json:"imageAlt"  orm:"image_alt"  description:"图片替代文本"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`
	Status    int         `json:"status"    orm:"status"     description:"状态"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"`
}
