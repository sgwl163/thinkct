// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppProducts is the golang structure for table app_products.
type AppProducts struct {
	Id          uint        `json:"id"          orm:"id"          description:"ID"`
	Cover       string      `json:"cover"       orm:"cover"       description:"封面"`
	Title       string      `json:"title"       orm:"title"       description:"标题"`
	Keywords    string      `json:"keywords"    orm:"keywords"    description:"关键词"`
	Description string      `json:"description" orm:"description" description:"描述"`
	Name        string      `json:"name"        orm:"name"        description:"名称"`
	Alias       string      `json:"alias"       orm:"alias"       description:"别名"`
	Content     string      `json:"content"     orm:"content"     description:"内容"`
	Price       float64     `json:"price"       orm:"price"       description:"价格"`
	Ysweb       string      `json:"ysweb"       orm:"ysweb"       description:"演示站"`
	Unique      string      `json:"unique"      orm:"unique"      description:"特点"`
	Weigh       int         `json:"weigh"       orm:"weigh"       description:"权重"`
	Status      uint        `json:"status"      orm:"status"      description:"状态:0=禁用,1=启用"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"修改时间"`
}
