// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "hotgo/api/api/user/v1"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/userin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IApiUser interface {
		Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
		Info(ctx context.Context, _ *v1.InfoReq) (res *v1.InfoRes, err error)
		Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error)
		Logout(ctx context.Context, _ *v1.LogoutReq) (res *v1.LogoutRes, err error)
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
	IUsergroup interface {
		// Model 分组管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取分组管理列表
		List(ctx context.Context, in *userin.GroupListInp) (list []*userin.GroupListModel, totalCount int, err error)
		// Export 导出分组管理
		Export(ctx context.Context, in *userin.GroupListInp) (err error)
		// Edit 修改/新增分组管理
		Edit(ctx context.Context, in *userin.GroupEditInp) (err error)
		// Delete 删除分组管理
		Delete(ctx context.Context, in *userin.GroupDeleteInp) (err error)
		// View 获取分组管理指定信息
		View(ctx context.Context, in *userin.GroupViewInp) (res *userin.GroupViewModel, err error)
		// Status 更新分组管理状态
		Status(ctx context.Context, in *userin.GroupStatusInp) (err error)
	}
	IUserUser interface {
		// Model 会员管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取会员管理列表
		List(ctx context.Context, in *userin.UserListInp) (list []*userin.UserListModel, totalCount int, err error)
		// Export 导出会员管理
		Export(ctx context.Context, in *userin.UserListInp) (err error)
		// Edit 修改/新增会员管理
		Edit(ctx context.Context, in *userin.UserEditInp) (err error)
		// Delete 删除会员管理
		Delete(ctx context.Context, in *userin.UserDeleteInp) (err error)
		// View 获取会员管理指定信息
		View(ctx context.Context, in *userin.UserViewInp) (res *userin.UserViewModel, err error)
		// Status 更新会员管理状态
		Status(ctx context.Context, in *userin.UserStatusInp) (err error)
		// Switch 更新会员管理开关
		Switch(ctx context.Context, in *userin.UserSwitchInp) (err error)
	}
	IUserRule interface {
		// Model 规则管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取规则管理列表
		List(ctx context.Context, in *userin.UserRuleListInp) (list []*userin.UserRuleListModel, totalCount int, err error)
		// Export 导出规则管理
		Export(ctx context.Context, in *userin.UserRuleListInp) (err error)
		// Edit 修改/新增规则管理
		Edit(ctx context.Context, in *userin.UserRuleEditInp) (err error)
		// Delete 删除规则管理
		Delete(ctx context.Context, in *userin.UserRuleDeleteInp) (err error)
		// View 获取规则管理指定信息
		View(ctx context.Context, in *userin.UserRuleViewInp) (res *userin.UserRuleViewModel, err error)
		// Status 更新规则管理状态
		Status(ctx context.Context, in *userin.UserRuleStatusInp) (err error)
	}
)

var (
	localApiUser   IApiUser
	localUsergroup IUsergroup
	localUserUser  IUserUser
	localUserRule  IUserRule
)

func ApiUser() IApiUser {
	if localApiUser == nil {
		panic("implement not found for interface IApiUser, forgot register?")
	}
	return localApiUser
}

func RegisterApiUser(i IApiUser) {
	localApiUser = i
}

func Usergroup() IUsergroup {
	if localUsergroup == nil {
		panic("implement not found for interface IUsergroup, forgot register?")
	}
	return localUsergroup
}

func RegisterUsergroup(i IUsergroup) {
	localUsergroup = i
}

func UserUser() IUserUser {
	if localUserUser == nil {
		panic("implement not found for interface IUserUser, forgot register?")
	}
	return localUserUser
}

func RegisterUserUser(i IUserUser) {
	localUserUser = i
}

func UserRule() IUserRule {
	if localUserRule == nil {
		panic("implement not found for interface IUserRule, forgot register?")
	}
	return localUserRule
}

func RegisterUserRule(i IUserRule) {
	localUserRule = i
}
