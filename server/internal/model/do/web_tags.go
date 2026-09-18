// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// WebTags is the golang structure of table hg_web_tags for DAO operations like Where/Data.
type WebTags struct {
	g.Meta        `orm:"table:hg_web_tags, do:true"`
	Id            any // ID
	Name          any // 名称
	Type          any // 标签类型:default=default,success=success,info=info,warning=warning,danger=danger
	DocumentCount any // 文档数
	Seotitle      any // SEO标题
	Keywords      any // 关键词
	Description   any // 描述
	InnerChain    any // 自动内链:0=关,1=开
	Weigh         any // 权重
	Status        any // 状态:0=禁用,1=启用
	UpdateTime    any // 修改时间
	CreateTime    any // 创建时间
}
