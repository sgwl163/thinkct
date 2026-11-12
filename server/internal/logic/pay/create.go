// Package pay
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package pay

// 订单创建

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "hotgo/api/api/pay/v1"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/library/hgorm/handler"
	"hotgo/internal/library/location"
	"hotgo/internal/library/payment"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/payin"
	"hotgo/internal/service"
	"hotgo/utility/validate"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gmeta"
)

// Create 创建支付订单和日志
func (s *sPay) Create(ctx context.Context, in payin.PayCreateInp) (res *payin.PayCreateModel, err error) {
	request := ghttp.RequestFromCtx(ctx)
	if in.TradeType == "" {
		in.TradeType = payment.AutoTradeType(in.PayType, request.UserAgent())
	}

	if in.Openid == "" {
		if in.Openid, err = service.CommonWechat().GetOpenId(ctx); err != nil {
			return
		}
	}

	if in.TradeType == consts.TradeTypeWxMP {
		if in.Openid == "" {
			err = gerror.New("微信公众号支付必须设置openid")
			return
		}
		if in.ReturnUrl == "" {
			err = gerror.New("微信公众号支付必须设置同步通知地址")
			return
		}
	}

	notifyURL, err := s.GenNotifyURL(ctx, in)
	if err != nil {
		return
	}

	config := payment.GetConfig()
	mchId := ""
	switch in.PayType {
	case consts.PayTypeAliPay:
		mchId = config.AliPayAppId
	case consts.PayTypeWxPay:
		mchId = config.WxPayMchId
	case consts.PayTypeQQPay:
		mchId = config.QQPayMchId
	}

	data := &entity.PayLog{
		MemberId:      contexts.GetUserId(ctx),
		AppId:         contexts.GetModule(ctx),
		AddonsName:    contexts.GetAddonName(ctx),
		OrderSn:       in.OrderSn,
		OrderGroup:    in.OrderGroup,
		Openid:        in.Openid,
		MchId:         mchId,
		Subject:       in.Subject,
		Detail:        in.Detail,
		OutTradeNo:    payment.GenOutTradeNo(ctx),
		TransactionId: "",
		PayType:       in.PayType,
		PayAmount:     in.PayAmount,
		PayStatus:     consts.PayStatusWait,
		TradeType:     in.TradeType,
		IsRefund:      consts.RefundStatusNo,
		CreateIp:      location.GetClientIp(request),
		NotifyUrl:     notifyURL,
		ReturnUrl:     in.ReturnUrl,
		TraceIds:      gjson.New([]string{gctx.CtxId(ctx)}),
		Status:        consts.StatusEnabled,
	}

	// 前台用户没有后台角色，充值支付日志不能应用后台 member_id 数据权限过滤。
	model := s.Model(ctx)
	if in.OrderGroup == consts.OrderGroupUserRecharge {
		model = s.Model(ctx, &handler.Option{FilterAuth: false})
	}
	if _, err = model.Data(data).OmitEmptyData().Insert(); err != nil {
		return
	}

	// 创建第三方平台支付订单
	order, err := payment.New(in.PayType).CreateOrder(ctx, payin.CreateOrderInp{Pay: data})
	if err != nil {
		return
	}

	// 保存原支付单的扫码信息，刷新页面只恢复原单，不再向渠道重复下单。
	if in.OrderGroup == consts.OrderGroupUserRecharge && order != nil {
		snapshot, encodeErr := json.Marshal(order)
		if encodeErr != nil {
			return nil, gerror.Wrap(encodeErr, "保存充值支付信息失败")
		}
		cols := dao.PayLog.Columns()
		if _, err = dao.PayLog.Ctx(ctx).
			Where(cols.OutTradeNo, data.OutTradeNo).
			Where(cols.OrderGroup, consts.OrderGroupUserRecharge).
			Where(cols.AppId, data.AppId).
			Where(cols.MemberId, data.MemberId).
			Data(g.Map{cols.Custom: string(snapshot)}).Update(); err != nil {
			return nil, gerror.Wrap(err, "保存充值支付信息失败，请在充值记录中核对订单")
		}
	}

	res = new(payin.PayCreateModel)
	res.Order = order
	return
}

// GenNotifyURL 生成支付通知地址
func (s *sPay) GenNotifyURL(ctx context.Context, in payin.PayCreateInp) (notifyURL string, err error) {
	basic, err := service.SysConfig().GetBasic(ctx)
	if err != nil {
		return
	}

	if basic.Domain == "" {
		err = gerror.New("请先到后台【系统设置】-【配置管理】中设置网站域名！")
		return
	}

	if !validate.IsURL(basic.Domain) {
		err = gerror.New("网站域名格式有误，请检查！")
		return
	}

	var object interface{}
	switch in.PayType {
	case consts.PayTypeAliPay:
		object = v1.NotifyAliPayReq{}
	case consts.PayTypeWxPay:
		object = v1.NotifyWxPayReq{}
	case consts.PayTypeQQPay:
		object = v1.NotifyQQPayReq{}
	default:
		err = gerror.Newf("未被支持的支付方式：%v", in.PayType)
		return
	}

	notifyURL = fmt.Sprintf("%s%s%s",
		basic.Domain,
		g.Cfg().MustGet(ctx, "router.api.prefix", "/api").String(),
		gmeta.Get(object, "path").String(),
	)
	return
}
