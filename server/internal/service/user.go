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
	IUserUserGroup interface {
		// Model 用户分组ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取用户分组列表
		List(ctx context.Context, in *userin.UserGroupListInp) (list []*userin.UserGroupListModel, totalCount int, err error)
		// Export 导出用户分组
		Export(ctx context.Context, in *userin.UserGroupListInp) (err error)
		// Edit 修改/新增用户分组
		Edit(ctx context.Context, in *userin.UserGroupEditInp) (err error)
		// Delete 删除用户分组
		Delete(ctx context.Context, in *userin.UserGroupDeleteInp) (err error)
		// View 获取用户分组指定信息
		View(ctx context.Context, in *userin.UserGroupViewInp) (res *userin.UserGroupViewModel, err error)
	}
	IUserUserLog interface {
		// Model 操作记录ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取操作记录列表
		List(ctx context.Context, in *userin.UserLogListInp) (list []*userin.UserLogListModel, totalCount int, err error)
		// Export 导出操作记录
		Export(ctx context.Context, in *userin.UserLogListInp) (err error)
		// Edit 修改/新增操作记录
		Edit(ctx context.Context, in *userin.UserLogEditInp) (err error)
		// Delete 删除操作记录
		Delete(ctx context.Context, in *userin.UserLogDeleteInp) (err error)
		// View 获取操作记录指定信息
		View(ctx context.Context, in *userin.UserLogViewInp) (res *userin.UserLogViewModel, err error)
		// Status 更新操作记录状态
		Status(ctx context.Context, in *userin.UserLogStatusInp) (err error)
	}
	IUserUserRule interface {
		// Model 权限规则ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取权限规则列表
		List(ctx context.Context, in *userin.UserRuleListInp) (list []*userin.UserRuleListModel, totalCount int, err error)
		// Export 导出权限规则
		Export(ctx context.Context, in *userin.UserRuleListInp) (err error)
		// Edit 修改/新增权限规则
		Edit(ctx context.Context, in *userin.UserRuleEditInp) (err error)
		// Delete 删除权限规则
		Delete(ctx context.Context, in *userin.UserRuleDeleteInp) (err error)
		// View 获取权限规则指定信息
		View(ctx context.Context, in *userin.UserRuleViewInp) (res *userin.UserRuleViewModel, err error)
	}
)

var (
	localApiUser       IApiUser
	localUserUser      IUserUser
	localUserUserGroup IUserUserGroup
	localUserUserLog   IUserUserLog
	localUserUserRule  IUserUserRule
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

func UserUser() IUserUser {
	if localUserUser == nil {
		panic("implement not found for interface IUserUser, forgot register?")
	}
	return localUserUser
}

func RegisterUserUser(i IUserUser) {
	localUserUser = i
}

func UserUserGroup() IUserUserGroup {
	if localUserUserGroup == nil {
		panic("implement not found for interface IUserUserGroup, forgot register?")
	}
	return localUserUserGroup
}

func RegisterUserUserGroup(i IUserUserGroup) {
	localUserUserGroup = i
}

func UserUserLog() IUserUserLog {
	if localUserUserLog == nil {
		panic("implement not found for interface IUserUserLog, forgot register?")
	}
	return localUserUserLog
}

func RegisterUserUserLog(i IUserUserLog) {
	localUserUserLog = i
}

func UserUserRule() IUserUserRule {
	if localUserUserRule == nil {
		panic("implement not found for interface IUserUserRule, forgot register?")
	}
	return localUserUserRule
}

func RegisterUserUserRule(i IUserUserRule) {
	localUserUserRule = i
}
