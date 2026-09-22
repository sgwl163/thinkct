// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OauthDaoban is the golang structure for table oauth_daoban.
type OauthDaoban struct {
	Id        uint        `json:"id"        orm:"id"         description:"ID"`
	ProductId int         `json:"productId" orm:"product_id" description:"产品ID"`
	Host      string      `json:"host"      orm:"host"       description:"域名"`
	Ip        string      `json:"ip"        orm:"ip"         description:"IP"`
	Version   string      `json:"version"   orm:"version"    description:"版本号"`
	Code      string      `json:"code"      orm:"code"       description:"执行代码"`
	Data      string      `json:"data"      orm:"data"       description:"数据"`
	Views     int         `json:"views"     orm:"views"      description:"浏览量"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"修改时间"`
}
