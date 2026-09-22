// Package router
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package router

import (
	"context"
	"hotgo/internal/consts"
	"hotgo/internal/controller/api/forum"
	"hotgo/internal/controller/api/im"
	"hotgo/internal/controller/api/member"
	"hotgo/internal/controller/api/openai"
	"hotgo/internal/controller/api/pay"
	"hotgo/internal/controller/api/shop"
	"hotgo/internal/controller/api/site"
	"hotgo/internal/controller/api/user"
	"hotgo/internal/service"
	"hotgo/utility/simple"

	"github.com/gogf/gf/v2/net/ghttp"
)

// Api 前台路由
func Api(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group(simple.RouterPrefix(ctx, consts.AppApi), func(group *ghttp.RouterGroup) {
		group.Bind(
			pay.NewV1(),  // 支付异步通知
			site.NewV1(), // 前台首页
		)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware().OptionalApiAuth)
			group.Bind(forum.NewPublicV1()) // 社区公开查询
		})
		group.Middleware(service.Middleware().ApiAuth)
		group.Bind(
			forum.NewV1(),  // 社区登录接口
			user.NewV1(),   // 用户登录等接口
			shop.NewV1(),   // 商城订单前台接口
			member.NewV1(), // 会员
			im.NewV1(),     // IM即时通讯前台接口
			&openai.Openai, // AI聊天
			&openai.App,    // OpenAI应用
		)
	})
}
