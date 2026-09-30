// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRule is the golang structure of table hg_user_rule for DAO operations like Where/Data.
type UserRule struct {
	g.Meta       `orm:"table:hg_user_rule, do:true"`
	Id           any         // ID
	Pid          any         // 上级菜单
	Type         any         // 类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮
	Title        any         // 标题
	Name         any         // 规则名称
	Path         any         // 路由路径
	Icon         any         // 图标
	MenuType     any         // 菜单类型:tab=选项卡,link=链接,iframe=Iframe
	Url          any         // Url
	Component    any         // 组件路径
	NoLoginValid any         // 未登录有效:0=否,1=是
	Extend       any         // 扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单
	Remark       any         // 备注
	Weigh        any         // 权重
	Status       any         // 状态:0=禁用,1=启用
	CreatedAt    *gtime.Time // 更新时间
	UpdatedAt    *gtime.Time // 创建时间
}
