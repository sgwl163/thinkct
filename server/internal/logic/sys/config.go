// Package sys
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sys

import (
	"context"
	"fmt"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/global"
	"hotgo/internal/library/payment"
	"hotgo/internal/library/sms"
	"hotgo/internal/library/storager"
	"hotgo/internal/library/token"
	"hotgo/internal/library/wechat"
	"hotgo/internal/model"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/sysin"
	"hotgo/internal/service"
	"hotgo/utility/simple"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

type sSysConfig struct{}

func NewSysConfig() *sSysConfig {
	return &sSysConfig{}
}

func init() {
	service.RegisterSysConfig(NewSysConfig())
}

// InitConfig 初始化系统配置
func (s *sSysConfig) InitConfig(ctx context.Context) {
	if err := s.LoadConfig(ctx); err != nil {
		g.Log().Fatalf(ctx, "InitConfig fail：%+v", err)
	}
}

// LoadConfig 加载系统配置
func (s *sSysConfig) LoadConfig(ctx context.Context) (err error) {
	if err = s.ensureBasic(ctx); err != nil {
		return
	}
	if err = s.EnsureOpenai(ctx); err != nil {
		return
	}

	wx, err := s.GetWechat(ctx)
	if err != nil {
		return
	}
	wechat.SetConfig(wx)

	pay, err := s.GetPay(ctx)
	if err != nil {
		return
	}
	payment.SetConfig(pay)

	upload, err := s.GetUpload(ctx)
	if err != nil {
		return
	}
	storager.SetConfig(upload)

	sm, err := s.GetSms(ctx)
	if err != nil {
		return
	}
	sms.SetConfig(sm)

	tk, err := s.GetLoadToken(ctx)
	if err != nil {
		return
	}
	token.SetConfig(tk)

	// 更多
	// ...
	return
}

// GetLogin 获取登录配置
func (s *sSysConfig) GetLogin(ctx context.Context) (conf *model.LoginConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "login"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetWechat 获取微信配置
func (s *sSysConfig) GetWechat(ctx context.Context) (conf *model.WechatConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "wechat"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetPay 获取支付配置
func (s *sSysConfig) GetPay(ctx context.Context) (conf *model.PayConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "pay"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetSms 获取短信配置
func (s *sSysConfig) GetSms(ctx context.Context) (conf *model.SmsConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "sms"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetOpenai 获取模型配置
func (s *sSysConfig) GetOpenai(ctx context.Context) (conf *model.OpenaiConfig, err error) {
	var models []*entity.SysConfig
	cols := dao.SysConfig.Columns()
	if err = dao.SysConfig.Ctx(ctx).
		Fields(cols.Key, cols.Value, cols.Type).
		Where(cols.Group, "openai").
		Scan(&models); err != nil {
		return nil, gerror.Wrap(err, "获取模型配置失败")
	}

	conf = new(model.OpenaiConfig)
	for _, item := range models {
		value := consts.ConvType(item.Value, item.Type)
		switch item.Key {
		case "openaiApiKey":
			conf.ApiKey = gconv.String(value)
		case "openaiBaseUrl":
			conf.BaseUrl = gconv.String(value)
		case "openaiDefaultModel":
			conf.DefaultModel = gconv.String(value)
		}
	}
	return conf, nil
}

// EnsureOpenai 初始化模型配置
func (s *sSysConfig) EnsureOpenai(ctx context.Context) (err error) {
	defaults := []struct {
		key, name, value, tip string
		sort                  int
	}{
		{"openaiApiKey", "API Key", "", "用于访问模型服务的密钥", 130},
		{"openaiBaseUrl", "接口地址", "https://api.gpt.ge/v1", "OpenAI兼容接口地址", 131},
		{"openaiDefaultModel", "默认模型", "gpt-3.5-turbo", "聊天会话未指定模型时使用", 132},
	}

	// 兼容已存在的模型接口组数据，只在系统配置项为空时迁移一次。
	var legacyGroup *entity.OpenaiApiGroup
	_ = dao.OpenaiApiGroup.Ctx(ctx).
		Where(dao.OpenaiApiGroup.Columns().Status, consts.StatusEnabled).
		Where(dao.OpenaiApiGroup.Columns().DeletedAt, nil).
		OrderAsc(dao.OpenaiApiGroup.Columns().Sort).
		Scan(&legacyGroup)
	if legacyGroup != nil {
		defaults[0].value = legacyGroup.ApiKey
		if legacyGroup.BaseUrl != "" {
			defaults[1].value = legacyGroup.BaseUrl
		}
	}

	cols := dao.SysConfig.Columns()
	var models []*entity.SysConfig
	if err = dao.SysConfig.Ctx(ctx).Where(cols.Group, "openai").Scan(&models); err != nil {
		return gerror.Wrap(err, "初始化模型配置失败")
	}
	for _, item := range defaults {
		row := s.getConfigByKey(item.key, models)
		if row != nil {
			if row.Value == "" && item.value != "" {
				_, err = dao.SysConfig.Ctx(ctx).WherePri(row.Id).Data(g.Map{
					cols.Value:     item.value,
					cols.UpdatedAt: gtime.Now(),
				}).Update()
				if err != nil {
					return gerror.Wrap(err, "迁移模型配置失败")
				}
			}
			continue
		}
		_, err = dao.SysConfig.Ctx(ctx).Data(g.Map{
			cols.Group:        "openai",
			cols.Name:         item.name,
			cols.Type:         consts.ConfigTypeString,
			cols.Key:          item.key,
			cols.Value:        item.value,
			cols.DefaultValue: item.value,
			cols.Sort:         item.sort,
			cols.Tip:          item.tip,
			cols.IsDefault:    1,
			cols.Status:       consts.StatusEnabled,
			cols.CreatedAt:    gtime.Now(),
			cols.UpdatedAt:    gtime.Now(),
		}).Insert()
		if err != nil {
			return gerror.Wrap(err, "写入模型配置失败")
		}
	}
	return
}

// GetGeo 获取地理配置
func (s *sSysConfig) GetGeo(ctx context.Context) (conf *model.GeoConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "geo"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetUpload 获取上传配置
func (s *sSysConfig) GetUpload(ctx context.Context) (conf *model.UploadConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "upload"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetSmtp 获取邮件配置
func (s *sSysConfig) GetSmtp(ctx context.Context) (conf *model.EmailConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "smtp"})
	if err != nil {
		return
	}
	if err = gconv.Scan(models.List, &conf); err != nil {
		return
	}

	conf.Addr = fmt.Sprintf("%s:%d", conf.Host, conf.Port)

	return
}

// ensureBasic 初始化基础配置
func (s *sSysConfig) ensureBasic(ctx context.Context) (err error) {
	const (
		faviconKey    = "basicFavicon"
		imageTypeKey  = "uploadImageType"
		faviconSuffix = "ico"
	)

	cols := dao.SysConfig.Columns()
	var models []*entity.SysConfig
	if err = dao.SysConfig.Ctx(ctx).
		Where(cols.Group, "basic").
		Where(cols.Key, faviconKey).
		Scan(&models); err != nil {
		return gerror.Wrap(err, "检查网站图标配置失败")
	}
	if len(models) == 0 {
		_, err = dao.SysConfig.Ctx(ctx).Data(g.Map{
			cols.Group:        "basic",
			cols.Name:         "网站图标",
			cols.Type:         consts.ConfigTypeString,
			cols.Key:          faviconKey,
			cols.Value:        "",
			cols.DefaultValue: "",
			cols.Sort:         25,
			cols.Tip:          "浏览器标签页图标，建议上传 .ico 格式文件",
			cols.IsDefault:    1,
			cols.Status:       consts.StatusEnabled,
			cols.CreatedAt:    gtime.Now(),
			cols.UpdatedAt:    gtime.Now(),
		}).Insert()
		if err != nil {
			return gerror.Wrap(err, "初始化网站图标配置失败")
		}
	}

	var uploadConfig *entity.SysConfig
	if err = dao.SysConfig.Ctx(ctx).
		Where(cols.Group, "upload").
		Where(cols.Key, imageTypeKey).
		Scan(&uploadConfig); err != nil {
		return gerror.Wrap(err, "检查图片上传类型配置失败")
	}
	if uploadConfig == nil || gstr.InArray(gstr.SplitAndTrim(uploadConfig.Value, ","), faviconSuffix) {
		return nil
	}

	_, err = dao.SysConfig.Ctx(ctx).WherePri(uploadConfig.Id).Data(g.Map{
		cols.Value:     gstr.Trim(uploadConfig.Value, ",") + "," + faviconSuffix,
		cols.UpdatedAt: gtime.Now(),
	}).Update()
	if err != nil {
		return gerror.Wrap(err, "启用网站图标上传失败")
	}
	return nil
}

// ensureSite 初始化站点配置
func (s *sSysConfig) ensureSite(ctx context.Context) (err error) {
	const policeRecordKey = "sitePoliceRecord"

	cols := dao.SysConfig.Columns()
	var config *entity.SysConfig
	if err = dao.SysConfig.Ctx(ctx).
		Where(cols.Group, "site").
		Where(cols.Key, policeRecordKey).
		Scan(&config); err != nil {
		return gerror.Wrap(err, "检查公安备案号配置失败")
	}
	if config != nil {
		return nil
	}

	_, err = dao.SysConfig.Ctx(ctx).Data(g.Map{
		cols.Group:        "site",
		cols.Name:         "公安备案号",
		cols.Type:         consts.ConfigTypeString,
		cols.Key:          policeRecordKey,
		cols.Value:        "",
		cols.DefaultValue: "",
		cols.Sort:         75,
		cols.Tip:          "网站公安机关备案号",
		cols.IsDefault:    1,
		cols.Status:       consts.StatusEnabled,
		cols.CreatedAt:    gtime.Now(),
		cols.UpdatedAt:    gtime.Now(),
	}).Insert()
	if err != nil {
		return gerror.Wrap(err, "初始化公安备案号配置失败")
	}
	return nil
}

// GetBasic 获取基础配置
func (s *sSysConfig) GetBasic(ctx context.Context) (conf *model.BasicConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "basic"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetSite 获取站点配置
func (s *sSysConfig) GetSite(ctx context.Context) (conf *model.SiteConfig, err error) {
	models, err := s.GetConfigByGroup(ctx, &sysin.GetConfigInp{Group: "site"})
	if err != nil {
		return
	}
	err = gconv.Scan(models.List, &conf)
	return
}

// GetLoadTCP 获取本地tcp配置
func (s *sSysConfig) GetLoadTCP(ctx context.Context) (conf *model.TCPConfig, err error) {
	err = g.Cfg().MustGet(ctx, "tcp").Scan(&conf)
	return
}

// GetLoadGenerate 获取本地生成配置
func (s *sSysConfig) GetLoadGenerate(ctx context.Context) (conf *model.GenerateConfig, err error) {
	err = g.Cfg().MustGet(ctx, "hggen").Scan(&conf)
	return
}

// GetLoadToken 获取本地token配置
func (s *sSysConfig) GetLoadToken(ctx context.Context) (conf *model.TokenConfig, err error) {
	err = g.Cfg().MustGet(ctx, "token").Scan(&conf)
	return
}

// GetLoadLog 获取本地日志配置
func (s *sSysConfig) GetLoadLog(ctx context.Context) (conf *model.LogConfig, err error) {
	err = g.Cfg().MustGet(ctx, "system.log").Scan(&conf)
	return
}

// GetLoadServeLog 获取本地服务日志配置
func (s *sSysConfig) GetLoadServeLog(ctx context.Context) (conf *model.ServeLogConfig, err error) {
	err = g.Cfg().MustGet(ctx, "system.serveLog").Scan(&conf)
	return
}

// GetConfigByGroup 获取指定分组的配置
func (s *sSysConfig) GetConfigByGroup(ctx context.Context, in *sysin.GetConfigInp) (res *sysin.GetConfigModel, err error) {
	if in.Group == "" {
		err = gerror.New("分组不能为空")
		return
	}
	if in.Group == "basic" {
		if err = s.ensureBasic(ctx); err != nil {
			return
		}
	}
	if in.Group == "site" {
		if err = s.ensureSite(ctx); err != nil {
			return
		}
	}

	var models []*entity.SysConfig
	cols := dao.SysConfig.Columns()
	if err = dao.SysConfig.Ctx(ctx).Fields(cols.Key, cols.Value, cols.Type).Where(cols.Group, in.Group).Scan(&models); err != nil {
		err = gerror.Wrapf(err, "获取配置分组[ %v ]失败，请稍后重试！", in.Group)
		return
	}

	res = new(sysin.GetConfigModel)
	if len(models) > 0 {
		res.List = make(g.Map, len(models))
		for _, v := range models {
			val, err := s.ConversionType(ctx, v)
			if err != nil {
				return nil, err
			}
			res.List[v.Key] = val
		}
	}

	res.List = simple.FilterMaskDemo(ctx, res.List)
	return
}

// ConversionType 转换类型
func (s *sSysConfig) ConversionType(ctx context.Context, models *entity.SysConfig) (value interface{}, err error) {
	if models == nil {
		err = gerror.New("数据不存在")
		return
	}
	return consts.ConvType(models.Value, models.Type), nil
}

// UpdateConfigByGroup 更新指定分组的配置
func (s *sSysConfig) UpdateConfigByGroup(ctx context.Context, in *sysin.UpdateConfigInp) (err error) {
	if in.Group == "" {
		err = gerror.New("分组不能为空")
		return
	}
	if in.Group == "basic" {
		if err = s.ensureBasic(ctx); err != nil {
			return
		}
	}
	if in.Group == "site" {
		if err = s.ensureSite(ctx); err != nil {
			return
		}
	}
	var (
		mod    = dao.SysConfig.Ctx(ctx)
		models []*entity.SysConfig
	)

	if err = mod.Where("group", in.Group).Scan(&models); err != nil {
		return
	}

	err = dao.SysConfig.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		for k, v := range in.List {
			row := s.getConfigByKey(k, models)
			// 演示环境敏感配置会以占位文本返回，保存时保留数据库中的真实值。
			if gconv.String(v) == consts.DemoTips {
				continue
			}
			// 新增
			if row == nil {
				err = gerror.Newf("暂不支持从前台添加变量，请先在数据库表[%v]中配置变量：%v", dao.SysConfig.Table(), k)
				return
			}

			// 更新
			_, err = dao.SysConfig.Ctx(ctx).Where("id", row.Id).Data(g.Map{"value": v, "updated_at": gtime.Now()}).Update()
			if err != nil {
				return
			}
		}
		return s.syncUpdate(ctx, in)
	})

	if err != nil {
		return
	}

	global.PublishClusterSync(ctx, consts.ClusterSyncSysconfig, nil)
	return
}

func (s *sSysConfig) getConfigByKey(key string, models []*entity.SysConfig) *entity.SysConfig {
	if len(models) == 0 {
		return nil
	}

	for _, v := range models {
		if key == v.Key {
			return v
		}
	}
	return nil
}

// syncUpdate 同步更新一些加载配置
func (s *sSysConfig) syncUpdate(ctx context.Context, in *sysin.UpdateConfigInp) (err error) {
	var cfg any
	switch in.Group {
	case "wechat":
		cfg, err = s.GetWechat(ctx)
		if err == nil {
			wechat.SetConfig(cfg.(*model.WechatConfig))
		}
	case "pay":
		cfg, err = s.GetPay(ctx)
		if err == nil {
			payment.SetConfig(cfg.(*model.PayConfig))
		}
	case "upload":
		cfg, err = s.GetUpload(ctx)
		if err == nil {
			storager.SetConfig(cfg.(*model.UploadConfig))
		}
	case "sms":
		cfg, err = s.GetSms(ctx)
		if err == nil {
			sms.SetConfig(cfg.(*model.SmsConfig))
		}
	}

	if err != nil {
		err = gerror.Newf("syncUpdate %v conifg fail：%+v", in.Group, err.Error())
	}
	return
}

// ClusterSync 集群同步
func (s *sSysConfig) ClusterSync(ctx context.Context, message *gredis.Message) {
	if err := s.LoadConfig(ctx); err != nil {
		g.Log().Errorf(ctx, "ClusterSync fail：%+v", err)
	}
}
