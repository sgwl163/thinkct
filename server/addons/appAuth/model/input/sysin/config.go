// Package sysin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2024 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package sysin

import (
	"context"
	"hotgo/addons/appAuth/model"
	"hotgo/internal/model/input/sysin"
	"hotgo/utility/validate"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// UpdateConfigInp 更新指定配置
type UpdateConfigInp struct {
	sysin.UpdateAddonsConfigInp
}

func (in *UpdateConfigInp) Filter(ctx context.Context) error {
	if in.Group != model.ConfigGroupBasic {
		return gerror.New("请选择有效的授权配置分组")
	}
	if len(in.List) == 0 {
		return gerror.New("请填写需要更新的授权配置")
	}

	// 仅接收本插件声明的配置，避免任意新增变量或修改其他配置。
	fields := gconv.Map(model.NewBasicConfig())
	for key, raw := range in.List {
		if _, ok := fields[key]; !ok {
			return gerror.New("包含不支持的授权配置项")
		}
		if key == "authQueryCount" {
			// 在类型转换之前校验原值，避免小数被截断、空值被当作零。
			value := gconv.String(raw)
			if err := g.Validator().Rules("regex:^[0-9]+$").Data(value).
				Messages("累计授权查询次数必须是非负整数").Run(ctx); err != nil {
				return err.Current()
			}
			count, err := strconv.ParseInt(value, 10, 64)
			if err != nil || count > model.MaxAuthQueryCount {
				return gerror.New("累计授权查询次数超出允许范围")
			}
			in.List[key] = count
			continue
		}

		value, ok := raw.(string)
		if !ok {
			return gerror.New("授权提示和服务联系方式必须是文本")
		}
		value = strings.TrimSpace(value)
		maxLength := 500
		switch key {
		case "unauthorizedMessage", "expiredMessage", "disabledMessage":
			if value == "" {
				return gerror.New("授权提示信息不能为空")
			}
		case "serviceName":
			maxLength = 100
		case "servicePhone":
			maxLength = 50
		case "serviceQQ":
			maxLength = 13
			if value != "" {
				if err := g.Validator().Rules("qq").Data(value).
					Messages("客服QQ格式不正确").Run(ctx); err != nil {
					return err.Current()
				}
			}
		case "purchaseUrl", "supportUrl":
			if value != "" && (!validate.IsURL(value) ||
				(!strings.HasPrefix(strings.ToLower(value), "http://") && !strings.HasPrefix(strings.ToLower(value), "https://"))) {
				return gerror.New("购买续费地址和技术支持地址必须是有效的HTTP或HTTPS地址")
			}
		}
		if utf8.RuneCountInString(value) > maxLength {
			return gerror.Newf("授权配置文本不能超过%d个字符", maxLength)
		}
		in.List[key] = value
	}
	return nil
}

type GetConfigInp struct {
	sysin.GetAddonsConfigInp
}

func (in *GetConfigInp) Filter(_ context.Context) error {
	if in.Group != model.ConfigGroupBasic {
		return gerror.New("请选择有效的授权配置分组")
	}
	return nil
}

type GetConfigModel struct {
	List g.Map `json:"list"`
}
