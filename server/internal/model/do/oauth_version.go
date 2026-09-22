// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// OauthVersion is the golang structure of table hg_oauth_version for DAO operations like Where/Data.
type OauthVersion struct {
	g.Meta    `orm:"table:hg_oauth_version, do:true"`
	Id        any         // ID
	ProductId any         // 产品ID
	Cover     any         // 封面
	Title     any         // 标题
	Content   any         // 详情
	Version   any         // 版本号
	File      any         // 文件
	Status    any         // 状态:1=启用,0=禁用
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
