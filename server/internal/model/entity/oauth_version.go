// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OauthVersion is the golang structure for table oauth_version.
type OauthVersion struct {
	Id        int         `json:"id"        orm:"id"         description:"ID"`
	ProductId int         `json:"productId" orm:"product_id" description:"产品ID"`
	Cover     string      `json:"cover"     orm:"cover"      description:"封面"`
	Title     string      `json:"title"     orm:"title"      description:"标题"`
	Content   string      `json:"content"   orm:"content"    description:"详情"`
	Version   string      `json:"version"   orm:"version"    description:"版本号"`
	File      string      `json:"file"      orm:"file"       description:"文件"`
	Status    uint        `json:"status"    orm:"status"     description:"状态:1=启用,0=禁用"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"修改时间"`
}
