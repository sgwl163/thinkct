// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/model/input/appin"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IAppAppAuth interface {
		// Model 授权管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取授权管理列表
		List(ctx context.Context, in *appin.AppAuthListInp) (list []*appin.AppAuthListModel, totalCount int, err error)
		// Export 导出授权管理
		Export(ctx context.Context, in *appin.AppAuthListInp) (err error)
		// Edit 修改/新增授权管理
		Edit(ctx context.Context, in *appin.AppAuthEditInp) (err error)
		// Delete 删除授权管理
		Delete(ctx context.Context, in *appin.AppAuthDeleteInp) (err error)
		// View 获取授权管理指定信息
		View(ctx context.Context, in *appin.AppAuthViewInp) (res *appin.AppAuthViewModel, err error)
		// Status 更新授权管理状态
		Status(ctx context.Context, in *appin.AppAuthStatusInp) (err error)
	}
)

var (
	localAppAppAuth IAppAppAuth
)

func AppAppAuth() IAppAppAuth {
	if localAppAppAuth == nil {
		panic("implement not found for interface IAppAppAuth, forgot register?")
	}
	return localAppAppAuth
}

func RegisterAppAppAuth(i IAppAppAuth) {
	localAppAppAuth = i
}
