// Package consts
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package consts

import (
	"hotgo/internal/library/dict"
	"hotgo/internal/model"
)

// 授权验证方式，值类型与授权表的字符串字段保持一致。
const (
	AppAuthModeDomain   = "0" // 单域名验证
	AppAuthModeDomainIP = "1" // IP双重验证
)

var AppAuthModeOptions = []*model.Option{
	dict.GenInfoOption(AppAuthModeDomain, "单域名验证"),
	dict.GenSuccessOption(AppAuthModeDomainIP, "IP双重验证", map[string]bool{"requireIp": true}),
}

// 授权域名范围。
const (
	AppAuthDomainSingle   = "0" // 当前单域名
	AppAuthDomainWildcard = "1" // 顶级泛域名
)

var AppAuthDomainOptions = []*model.Option{
	dict.GenPrimaryOption(AppAuthDomainSingle, "当前单域名"),
	dict.GenWarningOption(AppAuthDomainWildcard, "顶级泛域名"),
}

func init() {
	dict.RegisterEnums("AppAuthModeOptions", "授权验证方式", AppAuthModeOptions)
	dict.RegisterEnums("AppAuthDomainOptions", "授权域名范围", AppAuthDomainOptions)
}
