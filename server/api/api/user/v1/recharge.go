package v1

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/payin"

	"github.com/gogf/gf/v2/frame/g"
)

type RechargeOption struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type RechargeOptionsReq struct {
	g.Meta `path:"/user/recharge/options" method:"get" tags:"用户充值" summary:"获取充值选项"`
}

type RechargeOptionsRes struct {
	Balance        uint             `json:"balance"`
	Amounts        []uint           `json:"amounts"`
	MinAmount      uint             `json:"minAmount"`
	MaxAmount      uint             `json:"maxAmount"`
	PaymentMethods []RechargeOption `json:"paymentMethods"`
	StatusOptions  []RechargeOption `json:"statusOptions"`
}

type RechargeCreateReq struct {
	g.Meta  `path:"/user/recharge/create" method:"post" tags:"用户充值" summary:"创建会员充值订单"`
	Money   string `json:"money" v:"required|regex:^[1-9][0-9]{0,5}$|between:1,100000#请输入充值金额|充值金额须为整数元|充值金额须为1至100000元"`
	PayType string `json:"payType" v:"required|in:alipay,wxpay#请选择支付方式|不支持的支付方式"`
}

type RechargeCreateRes struct {
	Recharge *entity.UserRecharge    `json:"recharge"`
	Payment  *payin.CreateOrderModel `json:"payment"`
}

type RechargeViewReq struct {
	g.Meta  `path:"/user/recharge/view" method:"get" tags:"用户充值" summary:"查询本人充值订单"`
	OrderSn string `json:"orderSn" v:"required|length:1,64#订单号不能为空|订单号长度不正确"`
}

type RechargeViewRes struct {
	Recharge *entity.UserRecharge    `json:"recharge"`
	Balance  uint                    `json:"balance"`
	Payment  *payin.CreateOrderModel `json:"payment" dc:"原订单扫码信息，不可用或已到账时为null"`
}

type RechargeListReq struct {
	g.Meta  `path:"/user/recharge/list" method:"get" tags:"用户充值" summary:"获取本人充值记录"`
	Page    int `json:"page" d:"1" v:"min:1"`
	PerPage int `json:"perPage" d:"10" v:"between:1,100"`
}

type RechargeListRes struct {
	List       []*entity.UserRecharge `json:"list"`
	Page       int                    `json:"page"`
	PerPage    int                    `json:"perPage"`
	TotalCount int                    `json:"totalCount"`
}
