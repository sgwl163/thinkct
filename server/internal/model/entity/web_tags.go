// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// WebTags is the golang structure for table web_tags.
type WebTags struct {
	Id            uint   `json:"id"            orm:"id"             description:"ID"`
	Name          string `json:"name"          orm:"name"           description:"名称"`
	Type          string `json:"type"          orm:"type"           description:"标签类型:default=default,success=success,info=info,warning=warning,danger=danger"`
	DocumentCount int    `json:"documentCount" orm:"document_count" description:"文档数"`
	Seotitle      string `json:"seotitle"      orm:"seotitle"       description:"SEO标题"`
	Keywords      string `json:"keywords"      orm:"keywords"       description:"关键词"`
	Description   string `json:"description"   orm:"description"    description:"描述"`
	InnerChain    uint   `json:"innerChain"    orm:"inner_chain"    description:"自动内链:0=关,1=开"`
	Weigh         int    `json:"weigh"         orm:"weigh"          description:"权重"`
	Status        uint   `json:"status"        orm:"status"         description:"状态:0=禁用,1=启用"`
	UpdateTime    uint64 `json:"updateTime"    orm:"update_time"    description:"修改时间"`
	CreateTime    uint64 `json:"createTime"    orm:"create_time"    description:"创建时间"`
}
