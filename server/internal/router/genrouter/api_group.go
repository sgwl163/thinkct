// Package genrouter
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2026 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
// @AutoGenerate Version 2.18.9
package genrouter

import "hotgo/internal/controller/admin/openai"

func init() {
	LoginRequiredRouter = append(LoginRequiredRouter, openai.ApiGroup) // AI接口分组
}