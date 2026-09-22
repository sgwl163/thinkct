// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserCreditsLog is the golang structure for table user_credits_log.
type UserCreditsLog struct {
	Id         int64       `json:"id"         orm:"id"          description:"流水ID"`
	UserId     int64       `json:"userId"     orm:"user_id"     description:"用户ID"`
	CreditType string      `json:"creditType" orm:"credit_type" description:"资产类型:money=余额,score=积分"`
	BeforeNum  int64       `json:"beforeNum"  orm:"before_num"  description:"变动前数量"`
	Num        int64       `json:"num"        orm:"num"         description:"变动数量"`
	AfterNum   int64       `json:"afterNum"   orm:"after_num"   description:"变动后数量"`
	Remark     string      `json:"remark"     orm:"remark"      description:"备注"`
	MapId      int64       `json:"mapId"      orm:"map_id"      description:"关联ID"`
	Status     int         `json:"status"     orm:"status"      description:"状态"`
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at"  description:"创建时间"`
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at"  description:"更新时间"`
}
