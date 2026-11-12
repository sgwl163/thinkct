// Package model
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package model

const (
	ConfigGroupBasic = "basic"
	// MaxAuthQueryCount 与前端 JavaScript 可精确表示的最大整数保持一致。
	MaxAuthQueryCount int64 = 9007199254740991
)

// BasicConfig 产品授权配置，累计查询次数可手动调整并随授权查询自动递增，不是查询上限。
type BasicConfig struct {
	UnauthorizedMessage string `json:"unauthorizedMessage"`
	ExpiredMessage      string `json:"expiredMessage"`
	DisabledMessage     string `json:"disabledMessage"`
	PurchaseUrl         string `json:"purchaseUrl"`
	ServiceName         string `json:"serviceName"`
	ServiceQQ           string `json:"serviceQQ"`
	ServicePhone        string `json:"servicePhone"`
	SupportUrl          string `json:"supportUrl"`
	AuthQueryCount      int64  `json:"authQueryCount"`
}

// NewBasicConfig 默认配置，每次返回独立实例，避免不同请求相互修改。
func NewBasicConfig() *BasicConfig {
	return &BasicConfig{
		UnauthorizedMessage: "当前域名未获得产品授权，请联系服务商",
		ExpiredMessage:      "产品授权已到期，请联系服务商续费",
		DisabledMessage:     "产品授权已被禁用，请联系服务商",
	}
}
