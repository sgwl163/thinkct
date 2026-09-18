// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WebChannel is the golang structure of table hg_web_channel for DAO operations like Where/Data.
type WebChannel struct {
	g.Meta           `orm:"table:hg_web_channel, do:true"`
	Id               any // ID
	Pid              any // 上级频道
	Level            any // tree
	Tree             any // 关系树
	Model            any // 模型
	Type             any // 类型:cover=封面频道,list=列表频道,link=跳转链接
	Name             any // 频道名称
	Diyname          any // 自定义名称
	Template         any // 模板地址
	Url              any // 跳转链接
	Target           any // 跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口
	Seotitle         any // SEO标题
	Keywords         any // 关键词
	Description      any // 描述
	AllowVisitGroups any // 访问策略:all=全部,user=仅会员
	Weigh            any // 权重
	Status           any // 状态:0=禁用,1=启用
	UpdateTime       any // 修改时间
	CreateTime       any // 创建时间
}
