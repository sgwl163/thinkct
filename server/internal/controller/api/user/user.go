package user

import (
	"context"

	v1 "hotgo/api/api/user/v1"
	"hotgo/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	return service.ApiUser().Login(ctx, req)
}

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	return service.ApiUser().Info(ctx, req)
}

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	return service.ApiUser().Register(ctx, req)
}

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	return service.ApiUser().Logout(ctx, req)
}

func (c *ControllerV1) UpdateInfo(ctx context.Context, req *v1.UpdateInfoReq) (res *v1.UpdateInfoRes, err error) {
	return service.ApiUser().UpdateInfo(ctx, req)
}

func (c *ControllerV1) UpdatePwd(ctx context.Context, req *v1.UpdatePwdReq) (res *v1.UpdatePwdRes, err error) {
	return service.ApiUser().UpdatePwd(ctx, req)
}

func (c *ControllerV1) SendResetCode(ctx context.Context, req *v1.SendResetCodeReq) (res *v1.SendResetCodeRes, err error) {
	return service.ApiUser().SendResetCode(ctx, req)
}

func (c *ControllerV1) ResetPwd(ctx context.Context, req *v1.ResetPwdReq) (res *v1.ResetPwdRes, err error) {
	return service.ApiUser().ResetPwd(ctx, req)
}

func (c *ControllerV1) SendMobileBindCode(ctx context.Context, req *v1.SendMobileBindCodeReq) (res *v1.SendMobileBindCodeRes, err error) {
	return service.ApiUser().SendMobileBindCode(ctx, req)
}

func (c *ControllerV1) BindMobile(ctx context.Context, req *v1.BindMobileReq) (res *v1.BindMobileRes, err error) {
	return service.ApiUser().BindMobile(ctx, req)
}

func (c *ControllerV1) SendEmailBindCode(ctx context.Context, req *v1.SendEmailBindCodeReq) (res *v1.SendEmailBindCodeRes, err error) {
	return service.ApiUser().SendEmailBindCode(ctx, req)
}

func (c *ControllerV1) BindEmail(ctx context.Context, req *v1.BindEmailReq) (res *v1.BindEmailRes, err error) {
	return service.ApiUser().BindEmail(ctx, req)
}

func (c *ControllerV1) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	return service.ApiUser().Upload(ctx, req)
}

func (c *ControllerV1) CreditsLogList(ctx context.Context, req *v1.CreditsLogListReq) (res *v1.CreditsLogListRes, err error) {
	return service.ApiUser().CreditsLogList(ctx, req)
}

func (c *ControllerV1) LoginLogList(ctx context.Context, req *v1.LoginLogListReq) (res *v1.LoginLogListRes, err error) {
	return service.ApiUser().LoginLogList(ctx, req)
}

func (c *ControllerV1) Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error) {
	return service.ApiUser().Cancel(ctx, req)
}
