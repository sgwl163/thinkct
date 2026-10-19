// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"hotgo/addons/appAuth/model"
	"hotgo/addons/appAuth/model/input/sysin"
	"hotgo/internal/library/hgorm/handler"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	ISysAppAuth interface {
		// Model 授权管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取授权管理列表
		List(ctx context.Context, in *sysin.AppAuthListInp) (list []*sysin.AppAuthListModel, totalCount int, err error)
		// Export 导出授权管理
		Export(ctx context.Context, in *sysin.AppAuthListInp) (err error)
		// Edit 修改/新增授权管理
		Edit(ctx context.Context, in *sysin.AppAuthEditInp) (err error)
		// Delete 删除授权管理
		Delete(ctx context.Context, in *sysin.AppAuthDeleteInp) (err error)
		// View 获取授权管理指定信息
		View(ctx context.Context, in *sysin.AppAuthViewInp) (res *sysin.AppAuthViewModel, err error)
		// Status 更新授权管理状态
		Status(ctx context.Context, in *sysin.AppAuthStatusInp) (err error)
	}
	ISysAppDaoban interface {
		// Model 盗版管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取盗版管理列表
		List(ctx context.Context, in *sysin.AppDaobanListInp) (list []*sysin.AppDaobanListModel, totalCount int, err error)
		// Export 导出盗版管理
		Export(ctx context.Context, in *sysin.AppDaobanListInp) (err error)
		// Edit 修改/新增盗版管理
		Edit(ctx context.Context, in *sysin.AppDaobanEditInp) (err error)
		// Delete 删除盗版管理
		Delete(ctx context.Context, in *sysin.AppDaobanDeleteInp) (err error)
		// View 获取盗版管理指定信息
		View(ctx context.Context, in *sysin.AppDaobanViewInp) (res *sysin.AppDaobanViewModel, err error)
	}
	ISysAppLog interface {
		// Model 升级记录ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取升级记录列表
		List(ctx context.Context, in *sysin.AppLogListInp) (list []*sysin.AppLogListModel, totalCount int, err error)
		// Export 导出升级记录
		Export(ctx context.Context, in *sysin.AppLogListInp) (err error)
		// Edit 修改/新增升级记录
		Edit(ctx context.Context, in *sysin.AppLogEditInp) (err error)
		// Delete 删除升级记录
		Delete(ctx context.Context, in *sysin.AppLogDeleteInp) (err error)
		// View 获取升级记录指定信息
		View(ctx context.Context, in *sysin.AppLogViewInp) (res *sysin.AppLogViewModel, err error)
	}
	ISysAppProducts interface {
		// Model 产品管理ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取产品管理列表
		List(ctx context.Context, in *sysin.AppProductsListInp) (list []*sysin.AppProductsListModel, totalCount int, err error)
		// Export 导出产品管理
		Export(ctx context.Context, in *sysin.AppProductsListInp) (err error)
		// Edit 修改/新增产品管理
		Edit(ctx context.Context, in *sysin.AppProductsEditInp) (err error)
		// Delete 删除产品管理
		Delete(ctx context.Context, in *sysin.AppProductsDeleteInp) (err error)
		// View 获取产品管理指定信息
		View(ctx context.Context, in *sysin.AppProductsViewInp) (res *sysin.AppProductsViewModel, err error)
		// Status 更新产品管理状态
		Status(ctx context.Context, in *sysin.AppProductsStatusInp) (err error)
	}
	ISysAppVersion interface {
		// Model 版本更新ORM模型
		Model(ctx context.Context, option ...*handler.Option) *gdb.Model
		// List 获取版本更新列表
		List(ctx context.Context, in *sysin.AppVersionListInp) (list []*sysin.AppVersionListModel, totalCount int, err error)
		// Export 导出版本更新
		Export(ctx context.Context, in *sysin.AppVersionListInp) (err error)
		// Edit 修改/新增版本更新
		Edit(ctx context.Context, in *sysin.AppVersionEditInp) (err error)
		// Delete 删除版本更新
		Delete(ctx context.Context, in *sysin.AppVersionDeleteInp) (err error)
		// View 获取版本更新指定信息
		View(ctx context.Context, in *sysin.AppVersionViewInp) (res *sysin.AppVersionViewModel, err error)
		// Status 更新版本更新状态
		Status(ctx context.Context, in *sysin.AppVersionStatusInp) (err error)
	}
	ISysConfig interface {
		// GetBasic 获取基础配置
		GetBasic(ctx context.Context) (conf *model.BasicConfig, err error)
		// GetConfigByGroup 获取指定分组配置
		GetConfigByGroup(ctx context.Context, in *sysin.GetConfigInp) (res *sysin.GetConfigModel, err error)
		// UpdateConfigByGroup 更新指定分组的配置
		UpdateConfigByGroup(ctx context.Context, in *sysin.UpdateConfigInp) error
	}
	ISysIndex interface {
		// Test 测试
		Test(ctx context.Context, in *sysin.IndexTestInp) (res *sysin.IndexTestModel, err error)
	}
)

var (
	localSysAppAuth     ISysAppAuth
	localSysAppDaoban   ISysAppDaoban
	localSysAppLog      ISysAppLog
	localSysAppProducts ISysAppProducts
	localSysAppVersion  ISysAppVersion
	localSysConfig      ISysConfig
	localSysIndex       ISysIndex
)

func SysAppAuth() ISysAppAuth {
	if localSysAppAuth == nil {
		panic("implement not found for interface ISysAppAuth, forgot register?")
	}
	return localSysAppAuth
}

func RegisterSysAppAuth(i ISysAppAuth) {
	localSysAppAuth = i
}

func SysAppDaoban() ISysAppDaoban {
	if localSysAppDaoban == nil {
		panic("implement not found for interface ISysAppDaoban, forgot register?")
	}
	return localSysAppDaoban
}

func RegisterSysAppDaoban(i ISysAppDaoban) {
	localSysAppDaoban = i
}

func SysAppLog() ISysAppLog {
	if localSysAppLog == nil {
		panic("implement not found for interface ISysAppLog, forgot register?")
	}
	return localSysAppLog
}

func RegisterSysAppLog(i ISysAppLog) {
	localSysAppLog = i
}

func SysAppProducts() ISysAppProducts {
	if localSysAppProducts == nil {
		panic("implement not found for interface ISysAppProducts, forgot register?")
	}
	return localSysAppProducts
}

func RegisterSysAppProducts(i ISysAppProducts) {
	localSysAppProducts = i
}

func SysAppVersion() ISysAppVersion {
	if localSysAppVersion == nil {
		panic("implement not found for interface ISysAppVersion, forgot register?")
	}
	return localSysAppVersion
}

func RegisterSysAppVersion(i ISysAppVersion) {
	localSysAppVersion = i
}

func SysConfig() ISysConfig {
	if localSysConfig == nil {
		panic("implement not found for interface ISysConfig, forgot register?")
	}
	return localSysConfig
}

func RegisterSysConfig(i ISysConfig) {
	localSysConfig = i
}

func SysIndex() ISysIndex {
	if localSysIndex == nil {
		panic("implement not found for interface ISysIndex, forgot register?")
	}
	return localSysIndex
}

func RegisterSysIndex(i ISysIndex) {
	localSysIndex = i
}
