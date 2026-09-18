// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// WebFeedback is the golang structure for table web_feedback.
type WebFeedback struct {
	Id         uint   `json:"id"         orm:"id"          description:"ID"`
	Name       string `json:"name"       orm:"name"        description:"名称"`
	Email      string `json:"email"      orm:"email"       description:"邮箱"`
	Phone      string `json:"phone"      orm:"phone"       description:"手机号"`
	Company    string `json:"company"    orm:"company"     description:"公司名称"`
	Content    string `json:"content"    orm:"content"     description:"富文本"`
	Answer     string `json:"answer"     orm:"answer"      description:"备注"`
	Ip         string `json:"ip"         orm:"ip"          description:"IP"`
	Status     uint   `json:"status"     orm:"status"      description:"状态:0=禁用,1=启用"`
	CreateTime uint64 `json:"createTime" orm:"create_time" description:"创建时间"`
	UpdateTime uint64 `json:"updateTime" orm:"update_time" description:"修改时间"`
}
