// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// WebLinks is the golang structure for table web_links.
type WebLinks struct {
	Id         uint   `json:"id"         orm:"id"          description:"ID"`
	UserId     uint   `json:"userId"     orm:"user_id"     description:"用户ID"`
	Title      string `json:"title"      orm:"title"       description:"标题"`
	Link       string `json:"link"       orm:"link"        description:"链接"`
	Target     string `json:"target"     orm:"target"      description:"跳转方式:_blank=新标签页,_self=当前标签页,_top=当前窗体,_parent=父窗口"`
	Logo       string `json:"logo"       orm:"logo"        description:"LOGO"`
	Weigh      int    `json:"weigh"      orm:"weigh"       description:"权重"`
	Status     string `json:"status"     orm:"status"      description:"状态:disable=禁用,enable=启用,pending_trial=待审"`
	Remark     string `json:"remark"     orm:"remark"      description:"备注"`
	UpdateTime uint64 `json:"updateTime" orm:"update_time" description:"修改时间"`
	CreateTime uint64 `json:"createTime" orm:"create_time" description:"创建时间"`
}
