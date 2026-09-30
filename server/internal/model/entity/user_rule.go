// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserRule is the golang structure for table user_rule.
type UserRule struct {
	Id           uint        `json:"id"           orm:"id"             description:"ID"`
	Pid          uint        `json:"pid"          orm:"pid"            description:"上级菜单"`
	Type         string      `json:"type"         orm:"type"           description:"类型:route=路由,menu_dir=菜单目录,menu=菜单项,nav_user_menu=顶栏会员菜单下拉项,nav=顶栏菜单项,button=页面按钮"`
	Title        string      `json:"title"        orm:"title"          description:"标题"`
	Name         string      `json:"name"         orm:"name"           description:"规则名称"`
	Path         string      `json:"path"         orm:"path"           description:"路由路径"`
	Icon         string      `json:"icon"         orm:"icon"           description:"图标"`
	MenuType     string      `json:"menuType"     orm:"menu_type"      description:"菜单类型:tab=选项卡,link=链接,iframe=Iframe"`
	Url          string      `json:"url"          orm:"url"            description:"Url"`
	Component    string      `json:"component"    orm:"component"      description:"组件路径"`
	NoLoginValid uint        `json:"noLoginValid" orm:"no_login_valid" description:"未登录有效:0=否,1=是"`
	Extend       string      `json:"extend"       orm:"extend"         description:"扩展属性:none=无,add_rules_only=只添加为路由,add_menu_only=只添加为菜单"`
	Remark       string      `json:"remark"       orm:"remark"         description:"备注"`
	Weigh        int         `json:"weigh"        orm:"weigh"          description:"权重"`
	Status       string      `json:"status"       orm:"status"         description:"状态:0=禁用,1=启用"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"     description:"更新时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"     description:"创建时间"`
}
