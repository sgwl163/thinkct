// Package appAuth
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package appAuth

import (
	"context"
	_ "hotgo/addons/appAuth/crons"
	"hotgo/addons/appAuth/global"
	_ "hotgo/addons/appAuth/logic"
	_ "hotgo/addons/appAuth/queues"
	"hotgo/addons/appAuth/router"
	"hotgo/internal/library/addons"
	"hotgo/internal/service"
	"sync"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

type module struct {
	skeleton *addons.Skeleton
	ctx      context.Context
	sync.Mutex
}

func init() {
	newModule()
}

func newModule() {
	m := &module{
		skeleton: &addons.Skeleton{
			Label:       `应用授权系统`,
			Name:        `appAuth`,
			Group:       1,
			Logo:        "",
			Brief:       ``,
			Description: `应用授权系统`,
			Author:      `时光网络`,
			Version:     `v1.0.0`, // 当该版本号高于已安装的版本号时，会提示可以更新
		},
		ctx: gctx.New(),
	}
	addons.RegisterModule(m)
}

// Start 启动模块
func (m *module) Start(option *addons.Option) (err error) {
	// 初始化模块
	global.Init(m.ctx, m.skeleton)

	// 注册插件路由
	option.Server.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Addon)
		router.Admin(m.ctx, group)
		router.Api(m.ctx, group)
		router.Home(m.ctx, group)
		router.WebSocket(m.ctx, group)
	})
	return
}

// Stop 停止模块
func (m *module) Stop() (err error) {
	return
}

// Ctx 上下文
func (m *module) Ctx() context.Context {
	return m.ctx
}

// GetSkeleton 获取模块
func (m *module) GetSkeleton() *addons.Skeleton {
	return m.skeleton
}

// Install 安装模块
func (m *module) Install(ctx context.Context) (err error) {
	// ...
	return
}

// Upgrade 更新模块
func (m *module) Upgrade(ctx context.Context) (err error) {
	// ...
	return
}

// UnInstall 卸载模块
func (m *module) UnInstall(ctx context.Context) (err error) {
	// ...
	return
}
