// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/oauthin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IOauthAuth interface {
		// Model 授权管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取授权管理列表
		List(ctx context.Context, in *oauthin.AuthListInp) (list []*oauthin.AuthListModel, totalCount int, err error)
		// Export 导出授权管理
		Export(ctx context.Context, in *oauthin.AuthListInp) (err error)
		// Edit 修改/新增授权管理
		Edit(ctx context.Context, in *oauthin.AuthEditInp) (err error)
		// Delete 删除授权管理
		Delete(ctx context.Context, in *oauthin.AuthDeleteInp) (err error)
		// View 获取授权管理指定信息
		View(ctx context.Context, in *oauthin.AuthViewInp) (res *oauthin.AuthViewModel, err error)
		// Status 更新授权管理状态
		Status(ctx context.Context, in *oauthin.AuthStatusInp) (err error)
	}
	IOauthDaoban interface {
		// Model 盗版管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取盗版管理列表
		List(ctx context.Context, in *oauthin.DaobanListInp) (list []*oauthin.DaobanListModel, totalCount int, err error)
		// Export 导出盗版管理
		Export(ctx context.Context, in *oauthin.DaobanListInp) (err error)
		// Edit 修改/新增盗版管理
		Edit(ctx context.Context, in *oauthin.DaobanEditInp) (err error)
		// Delete 删除盗版管理
		Delete(ctx context.Context, in *oauthin.DaobanDeleteInp) (err error)
		// View 获取盗版管理指定信息
		View(ctx context.Context, in *oauthin.DaobanViewInp) (res *oauthin.DaobanViewModel, err error)
	}
	IOauthOauthLog interface {
		// Model 升级记录ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取升级记录列表
		List(ctx context.Context, in *oauthin.OauthLogListInp) (list []*oauthin.OauthLogListModel, totalCount int, err error)
		// Export 导出升级记录
		Export(ctx context.Context, in *oauthin.OauthLogListInp) (err error)
		// Edit 修改/新增升级记录
		Edit(ctx context.Context, in *oauthin.OauthLogEditInp) (err error)
		// Delete 删除升级记录
		Delete(ctx context.Context, in *oauthin.OauthLogDeleteInp) (err error)
		// View 获取升级记录指定信息
		View(ctx context.Context, in *oauthin.OauthLogViewInp) (res *oauthin.OauthLogViewModel, err error)
	}
	IOauthVersion interface {
		// Model 版本更新ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取版本更新列表
		List(ctx context.Context, in *oauthin.VersionListInp) (list []*oauthin.VersionListModel, totalCount int, err error)
		// Export 导出版本更新
		Export(ctx context.Context, in *oauthin.VersionListInp) (err error)
		// Edit 修改/新增版本更新
		Edit(ctx context.Context, in *oauthin.VersionEditInp) (err error)
		// Delete 删除版本更新
		Delete(ctx context.Context, in *oauthin.VersionDeleteInp) (err error)
		// View 获取版本更新指定信息
		View(ctx context.Context, in *oauthin.VersionViewInp) (res *oauthin.VersionViewModel, err error)
		// Status 更新版本更新状态
		Status(ctx context.Context, in *oauthin.VersionStatusInp) (err error)
	}
)

var (
	localOauthAuth     IOauthAuth
	localOauthDaoban   IOauthDaoban
	localOauthOauthLog IOauthOauthLog
	localOauthVersion  IOauthVersion
)

func OauthAuth() IOauthAuth {
	if localOauthAuth == nil {
		panic("implement not found for interface IOauthAuth, forgot register?")
	}
	return localOauthAuth
}

func RegisterOauthAuth(i IOauthAuth) {
	localOauthAuth = i
}

func OauthDaoban() IOauthDaoban {
	if localOauthDaoban == nil {
		panic("implement not found for interface IOauthDaoban, forgot register?")
	}
	return localOauthDaoban
}

func RegisterOauthDaoban(i IOauthDaoban) {
	localOauthDaoban = i
}

func OauthOauthLog() IOauthOauthLog {
	if localOauthOauthLog == nil {
		panic("implement not found for interface IOauthOauthLog, forgot register?")
	}
	return localOauthOauthLog
}

func RegisterOauthOauthLog(i IOauthOauthLog) {
	localOauthOauthLog = i
}

func OauthVersion() IOauthVersion {
	if localOauthVersion == nil {
		panic("implement not found for interface IOauthVersion, forgot register?")
	}
	return localOauthVersion
}

func RegisterOauthVersion(i IOauthVersion) {
	localOauthVersion = i
}
