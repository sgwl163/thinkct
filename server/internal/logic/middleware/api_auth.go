// Package middleware
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"hotgo/internal/consts"
	"hotgo/internal/library/response"
	"hotgo/internal/library/token"
	"hotgo/utility/simple"
)

// ApiAuth API鉴权中间件
func (s *sMiddleware) ApiAuth(r *ghttp.Request) {
	var (
		ctx  = r.Context()
		path = gstr.Replace(r.URL.Path, simple.RouterPrefix(ctx, consts.AppApi), "", 1)
	)

	// 不需要验证登录的路由地址
	if s.IsExceptLogin(ctx, consts.AppApi, path) {
		r.Middleware.Next()
		return
	}

	// 将用户信息传递到上下文中
	if err := s.DeliverUserContext(r); err != nil {
		response.JsonExit(r, gcode.CodeNotAuthorized.Code(), err.Error())
		return
	}

	// 验证路由访问权限
	// ...

	r.Middleware.Next()
}

// OptionalApiAuth API可选鉴权中间件。存在有效登录令牌时写入用户上下文，
// 没有令牌或令牌失效时仍按游客继续请求。
func (s *sMiddleware) OptionalApiAuth(r *ghttp.Request) {
	if token.GetAuthorization(r) != "" {
		_ = s.DeliverUserContext(r)
	}
	r.Middleware.Next()
}
