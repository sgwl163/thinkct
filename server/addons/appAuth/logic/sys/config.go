// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"hotgo/addons/appAuth/global"
	"hotgo/addons/appAuth/model"
	"hotgo/addons/appAuth/model/input/sysin"
	"hotgo/addons/appAuth/service"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	isysin "hotgo/internal/model/input/sysin"
	isc "hotgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysConfig struct{}

func NewSysConfig() *sSysConfig {
	return &sSysConfig{}
}

func init() {
	service.RegisterSysConfig(NewSysConfig())
}

// GetBasic 获取产品授权基础配置。
func (s *sSysConfig) GetBasic(ctx context.Context) (conf *model.BasicConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{
		GetAddonsConfigInp: isysin.GetAddonsConfigInp{Group: model.ConfigGroupBasic},
	})
	if err != nil {
		return nil, err
	}
	conf = model.NewBasicConfig()
	if err = gconv.Struct(models.List, conf); err != nil {
		return nil, gerror.Wrap(err, "解析产品授权配置失败")
	}
	return
}

// IncrementAuthQueryCount 在数据库中原子累加，并在同一事务内读取包含本次查询的累计值。
// 更新持有行锁至事务结束，避免并发查询丢失增量；达到JS安全整数上限后保持上限。
func (s *sSysConfig) IncrementAuthQueryCount(ctx context.Context) (count int64, err error) {
	if err = s.ensureBasic(ctx); err != nil {
		return 0, err
	}
	cols := dao.SysAddonsConfig.Columns()
	err = dao.SysAddonsConfig.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		mod := dao.SysAddonsConfig.Ctx(ctx).TX(tx).
			Where(cols.AddonName, global.GetSkeleton().Name).
			Where(cols.Group, model.ConfigGroupBasic).Where(cols.Key, "authQueryCount")
		if _, updateErr := mod.Clone().WhereLT(cols.Value, model.MaxAuthQueryCount).Increment(cols.Value, 1); updateErr != nil {
			return gerror.Wrap(updateErr, "更新授权查询次数失败，请稍后重试")
		}
		value, readErr := mod.Value(cols.Value)
		if readErr != nil {
			return gerror.Wrap(readErr, "读取授权查询次数失败，请稍后重试")
		}
		if value.IsNil() {
			return gerror.New("授权查询次数配置不存在，请稍后重试")
		}
		count = value.Int64()
		return nil
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetPublic 获取公开授权配置，仅返回允许前台展示的固定字段。
func (s *sSysConfig) GetPublic(ctx context.Context, in *sysin.AuthConfigInp) (res *sysin.AuthConfigModel, err error) {
	conf, err := s.GetBasic(ctx)
	if err != nil {
		return nil, err
	}
	return &sysin.AuthConfigModel{
		UnauthorizedMessage: conf.UnauthorizedMessage,
		ExpiredMessage:      conf.ExpiredMessage,
		DisabledMessage:     conf.DisabledMessage,
		PurchaseUrl:         conf.PurchaseUrl,
		ServiceName:         conf.ServiceName,
		ServiceQQ:           conf.ServiceQQ,
		ServicePhone:        conf.ServicePhone,
		SupportUrl:          conf.SupportUrl,
		AuthQueryCount:      conf.AuthQueryCount,
	}, nil
}

// GetConfigByGroup 获取指定分组配置，兼容尚未初始化配置的旧安装。
func (s *sSysConfig) GetConfigByGroup(ctx context.Context, in *sysin.GetConfigInp) (res *sysin.GetConfigModel, err error) {
	in.GetAddonsConfigInp.AddonName = global.GetSkeleton().Name
	if err = s.ensureBasic(ctx); err != nil {
		return nil, err
	}
	models, err := isc.SysAddonsConfig().GetConfigByGroup(ctx, &in.GetAddonsConfigInp)
	if err != nil {
		return nil, gerror.Wrap(err, "获取产品授权配置失败")
	}

	res = &sysin.GetConfigModel{List: gconv.Map(model.NewBasicConfig())}
	if models != nil {
		// 只返回已声明的配置项，历史测试变量不再参与编辑和提交。
		for key := range res.List {
			if value, ok := models.List[key]; ok && value != nil {
				res.List[key] = value
			}
		}
	}
	return
}

// UpdateConfigByGroup 更新指定分组的配置。
func (s *sSysConfig) UpdateConfigByGroup(ctx context.Context, in *sysin.UpdateConfigInp) error {
	in.UpdateAddonsConfigInp.AddonName = global.GetSkeleton().Name
	if err := s.ensureBasic(ctx); err != nil {
		return err
	}
	if err := isc.SysAddonsConfig().UpdateConfigByGroup(ctx, &in.UpdateAddonsConfigInp); err != nil {
		return gerror.Wrap(err, "保存产品授权配置失败")
	}
	return nil
}

// ensureBasic 仅补充缺失的固定配置项，不覆盖用户已保存的空值、零或自定义内容。
func (s *sSysConfig) ensureBasic(ctx context.Context) error {
	defaults := model.NewBasicConfig()
	items := []struct {
		key, name, value, valueType, tip string
	}{
		{"unauthorizedMessage", "未授权提示", defaults.UnauthorizedMessage, consts.ConfigTypeString, "未找到产品授权时的提示信息"},
		{"expiredMessage", "授权到期提示", defaults.ExpiredMessage, consts.ConfigTypeString, "产品授权到期时的提示信息"},
		{"disabledMessage", "授权禁用提示", defaults.DisabledMessage, consts.ConfigTypeString, "产品授权被禁用时的提示信息"},
		{"purchaseUrl", "购买续费地址", defaults.PurchaseUrl, consts.ConfigTypeString, "购买产品授权或续费的HTTP/HTTPS地址"},
		{"serviceName", "服务商名称", defaults.ServiceName, consts.ConfigTypeString, "产品授权服务商名称"},
		{"serviceQQ", "客服QQ", defaults.ServiceQQ, consts.ConfigTypeString, "产品授权客服QQ号码"},
		{"servicePhone", "客服电话", defaults.ServicePhone, consts.ConfigTypeString, "产品授权客服手机或座机号码"},
		{"supportUrl", "技术支持地址", defaults.SupportUrl, consts.ConfigTypeString, "技术支持的HTTP/HTTPS地址"},
		{"authQueryCount", "累计授权查询次数", gconv.String(defaults.AuthQueryCount), consts.ConfigTypeInt64, "可手动调整累计次数，每次完成授权查询后自动加1，不是查询次数上限"},
	}
	cols := dao.SysAddonsConfig.Columns()
	addonName := global.GetSkeleton().Name
	var rows []*entity.SysAddonsConfig
	if err := dao.SysAddonsConfig.Ctx(ctx).Fields(cols.Key).
		Where(cols.AddonName, addonName).Where(cols.Group, model.ConfigGroupBasic).Scan(&rows); err != nil {
		return gerror.Wrap(err, "检查产品授权配置失败")
	}
	existing := make(map[string]bool, len(rows))
	for _, row := range rows {
		existing[row.Key] = true
	}
	var missing []g.Map
	for index, item := range items {
		if existing[item.key] {
			continue
		}
		missing = append(missing, g.Map{
			cols.AddonName:    addonName,
			cols.Group:        model.ConfigGroupBasic,
			cols.Name:         item.name,
			cols.Type:         item.valueType,
			cols.Key:          item.key,
			cols.Value:        item.value,
			cols.DefaultValue: item.value,
			cols.Sort:         index + 1,
			cols.Tip:          item.tip,
			cols.IsDefault:    1,
			cols.Status:       consts.StatusEnabled,
		})
	}
	if len(missing) == 0 {
		return nil
	}
	// addon_name/key 唯一索引配合 InsertIgnore，避免并发初始化覆盖已有配置。
	if _, err := dao.SysAddonsConfig.Ctx(ctx).Data(missing).InsertIgnore(); err != nil {
		return gerror.Wrap(err, "初始化产品授权配置失败")
	}
	return nil
}
