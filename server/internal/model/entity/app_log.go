// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppLog is the golang structure for table app_log.
type AppLog struct {
	Id        uint        `json:"id"        orm:"id"         description:"ID"`
	UserId    uint        `json:"userId"    orm:"user_id"    description:"用户ID"`
	ProductId uint        `json:"productId" orm:"product_id" description:"产品ID"`
	Host      string      `json:"host"      orm:"host"       description:"域名"`
	Content   string      `json:"content"   orm:"content"    description:"内容"`
	Ip        string      `json:"ip"        orm:"ip"         description:"IP"`
	Version   string      `json:"version"   orm:"version"    description:"版本号"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"修改时间"`
}
