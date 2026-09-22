// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// OauthAuth is the golang structure for table oauth_auth.
type OauthAuth struct {
	Id         uint        `json:"id"         orm:"id"          description:"ID"`
	UserId     int         `json:"userId"     orm:"user_id"     description:"用户ID"`
	ProductId  int         `json:"productId"  orm:"product_id"  description:"产品ID"`
	Name       string      `json:"name"       orm:"name"        description:"名称"`
	Host       string      `json:"host"       orm:"host"        description:"域名"`
	Ip         string      `json:"ip"         orm:"ip"          description:"IP"`
	Qq         string      `json:"qq"         orm:"qq"          description:"QQ"`
	Phone      string      `json:"phone"      orm:"phone"       description:"电话"`
	Syskey     string      `json:"syskey"     orm:"syskey"      description:"授权码"`
	VersionId  string      `json:"versionId"  orm:"version_id"  description:"版本"`
	AuthMode   string      `json:"authMode"   orm:"auth_mode"   description:"验证方式:0=单域名验证,1=IP双重验证"`
	AuthDomain string      `json:"authDomain" orm:"auth_domain" description:"验证域名:0=当前单域名方式,1=顶级泛域名方式"`
	Views      int         `json:"views"      orm:"views"       description:"浏览量"`
	Status     uint        `json:"status"     orm:"status"      description:"状态:1=启用,0=禁用"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"修改时间"`
	ExpiredAt  *gtime.Time `json:"expiredAt"  orm:"expired_at"  description:"到期时间"`
}
