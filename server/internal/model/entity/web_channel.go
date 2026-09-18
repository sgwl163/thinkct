// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// WebChannel is the golang structure for table web_channel.
type WebChannel struct {
	Id               uint   `json:"id"               orm:"id"                 description:"ID"`
	Pid              uint   `json:"pid"              orm:"pid"                description:"上级频道"`
	Level            int    `json:"level"            orm:"level"              description:"tree"`
	Tree             string `json:"tree"             orm:"tree"               description:"关系树"`
	Model            string `json:"model"            orm:"model"              description:"模型"`
	Type             string `json:"type"             orm:"type"               description:"类型:cover=封面频道,list=列表频道,link=跳转链接"`
	Name             string `json:"name"             orm:"name"               description:"频道名称"`
	Diyname          string `json:"diyname"          orm:"diyname"            description:"自定义名称"`
	Template         string `json:"template"         orm:"template"           description:"模板地址"`
	Url              string `json:"url"              orm:"url"                description:"跳转链接"`
	Target           string `json:"target"           orm:"target"             description:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Seotitle         string `json:"seotitle"         orm:"seotitle"           description:"SEO标题"`
	Keywords         string `json:"keywords"         orm:"keywords"           description:"关键词"`
	Description      string `json:"description"      orm:"description"        description:"描述"`
	AllowVisitGroups string `json:"allowVisitGroups" orm:"allow_visit_groups" description:"访问策略:all=全部,user=仅会员"`
	Weigh            int    `json:"weigh"            orm:"weigh"              description:"权重"`
	Status           uint   `json:"status"           orm:"status"             description:"状态:0=禁用,1=启用"`
	UpdateTime       uint64 `json:"updateTime"       orm:"update_time"        description:"修改时间"`
	CreateTime       uint64 `json:"createTime"       orm:"create_time"        description:"创建时间"`
}
