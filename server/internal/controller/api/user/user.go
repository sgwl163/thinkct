package user

import (
	"context"

	v1 "hotgo/api/api/user/v1"
	"hotgo/internal/service"
)

// Login 用户登录
func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	return service.ApiUser().Login(ctx, req)
}

// Info 获取用户信息
func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	return service.ApiUser().Info(ctx, req)
}

// Register 用户注册
func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	return service.ApiUser().Register(ctx, req)
}

// Logout 退出登录
func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	return service.ApiUser().Logout(ctx, req)
}

// UpdateInfo 修改用户信息
func (c *ControllerV1) UpdateInfo(ctx context.Context, req *v1.UpdateInfoReq) (res *v1.UpdateInfoRes, err error) {
	return service.ApiUser().UpdateInfo(ctx, req)
}
