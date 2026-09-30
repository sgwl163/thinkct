// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppDaoban is the golang structure of table hg_app_daoban for DAO operations like Where/Data.
type AppDaoban struct {
	g.Meta    `orm:"table:hg_app_daoban, do:true"`
	Id        any         // ID
	ProductId any         // 产品ID
	Host      any         // 域名
	Ip        any         // IP
	Version   any         // 版本号
	Code      any         // 执行代码
	Data      any         // 数据
	Views     any         // 浏览量
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
