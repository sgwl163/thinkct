// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserGroup is the golang structure for table user_group.
type UserGroup struct {
	Id         uint   `json:"id"         orm:"id"          description:"ID"`
	Name       string `json:"name"       orm:"name"        description:"组名"`
	Rules      string `json:"rules"      orm:"rules"       description:"权限节点"`
	Status     string `json:"status"     orm:"status"      description:"状态:0=禁用,1=启用"`
	UpdateTime uint64 `json:"updateTime" orm:"update_time" description:"更新时间"`
	CreateTime uint64 `json:"createTime" orm:"create_time" description:"创建时间"`
}
