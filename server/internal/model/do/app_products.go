// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppProducts is the golang structure of table hg_app_products for DAO operations like Where/Data.
type AppProducts struct {
	g.Meta      `orm:"table:hg_app_products, do:true"`
	Id          any         // ID
	Cover       any         // 封面
	Title       any         // 标题
	Keywords    any         // 关键词
	Description any         // 描述
	Name        any         // 名称
	Alias       any         // 别名
	Content     any         // 内容
	Price       any         // 价格
	Ysweb       any         // 演示站
	Unique      any         // 特点
	Weigh       any         // 权重
	Status      any         // 状态:0=禁用,1=启用
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 修改时间
}
