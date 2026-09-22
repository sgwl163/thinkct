package user

import (
	"context"

	v1 "hotgo/api/api/user/v1"
)

type IUserV1 interface {
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error)
	Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
	UpdateInfo(ctx context.Context, req *v1.UpdateInfoReq) (res *v1.UpdateInfoRes, err error)
	UpdatePwd(ctx context.Context, req *v1.UpdatePwdReq) (res *v1.UpdatePwdRes, err error)
	SendResetCode(ctx context.Context, req *v1.SendResetCodeReq) (res *v1.SendResetCodeRes, err error)
	ResetPwd(ctx context.Context, req *v1.ResetPwdReq) (res *v1.ResetPwdRes, err error)
	SendMobileBindCode(ctx context.Context, req *v1.SendMobileBindCodeReq) (res *v1.SendMobileBindCodeRes, err error)
	BindMobile(ctx context.Context, req *v1.BindMobileReq) (res *v1.BindMobileRes, err error)
	SendEmailBindCode(ctx context.Context, req *v1.SendEmailBindCodeReq) (res *v1.SendEmailBindCodeRes, err error)
	BindEmail(ctx context.Context, req *v1.BindEmailReq) (res *v1.BindEmailRes, err error)
	Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error)
	CreditsLogList(ctx context.Context, req *v1.CreditsLogListReq) (res *v1.CreditsLogListRes, err error)
	LoginLogList(ctx context.Context, req *v1.LoginLogListReq) (res *v1.LoginLogListRes, err error)
	Cancel(ctx context.Context, req *v1.CancelReq) (res *v1.CancelRes, err error)
}
