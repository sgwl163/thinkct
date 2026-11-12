package user

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	v1 "hotgo/api/api/user/v1"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/payment"
	paymentAlipay "hotgo/internal/library/payment/alipay"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/payin"
	"hotgo/internal/service"
	"hotgo/utility/validate"

	"github.com/go-pay/gopay"
	sdkAlipay "github.com/go-pay/gopay/alipay"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
)

// 前台当面付独立于后台 scan（网页支付）的语义。
const rechargeTradeTypeAliPrecreate = "precreate"

func rechargeAlipayTradeType(userAgent string) string {
	if validate.IsMobileVisit(userAgent) {
		return consts.TradeTypeAliWeb
	}
	return rechargeTradeTypeAliPrecreate
}

func rechargePrecreateParams(recharge *entity.UserRecharge, pay *entity.PayLog) gopay.BodyMap {
	return gopay.BodyMap{}.
		Set("subject", pay.Subject).
		Set("out_trade_no", pay.OutTradeNo).
		Set("total_amount", strconv.FormatUint(uint64(recharge.Money), 10)+".00").
		Set("timeout_express", "15m")
}

func rechargePrecreatePayment(response *sdkAlipay.TradePrecreateResponse, outTradeNo string) (*payin.CreateOrderModel, error) {
	if response == nil || response.Response == nil {
		return nil, gerror.New("支付宝当面付未返回预下单响应，请核对充值记录")
	}
	if response.NullResponse != nil || response.Response.Code != "10000" {
		return nil, gerror.New("支付宝当面付预下单未成功，请联系管理员核对充值记录")
	}
	if strings.TrimSpace(response.Response.QrCode) == "" {
		return nil, gerror.New("支付宝当面付未返回二维码，请核对充值记录")
	}
	if response.Response.OutTradeNo != outTradeNo {
		return nil, gerror.New("支付宝当面付返回的订单号不匹配，请核对充值记录")
	}
	return &payin.CreateOrderModel{TradeType: rechargeTradeTypeAliPrecreate, PayURL: response.Response.QrCode, OutTradeNo: outTradeNo}, nil
}

func rechargePaymentMethods() []v1.RechargeOption {
	methods := make([]v1.RechargeOption, 0, 2)
	config := payment.GetConfig()
	if config == nil {
		return methods
	}
	if config.AliPayAppId != "" && config.AliPayPrivateKey != "" && config.AliPayAppCertPublicKey != "" && config.AliPayRootCert != "" && config.AliPayCertPublicKeyRSA2 != "" {
		methods = append(methods, v1.RechargeOption{Value: consts.PayTypeAliPay, Label: consts.PayTypeNameMap[consts.PayTypeAliPay]})
	}
	if config.WxPayAppId != "" && config.WxPayMchId != "" && config.WxPaySerialNo != "" && config.WxPayAPIv3Key != "" && config.WxPayPrivateKey != "" {
		methods = append(methods, v1.RechargeOption{Value: consts.PayTypeWxPay, Label: consts.PayTypeNameMap[consts.PayTypeWxPay]})
	}
	return methods
}

// rechargePayLog 只接受已验证用户和服务端充值单，不读取后台身份或客户端支付金额。
func rechargePayLog(userID int64, recharge *entity.UserRecharge) (*entity.PayLog, error) {
	if recharge == nil || userID <= 0 || recharge.UserId != userID || recharge.Money == 0 {
		return nil, gerror.New("充值订单与当前用户不一致")
	}
	tradeType := ""
	switch recharge.PayType {
	case consts.PayTypeAliPay:
		tradeType = consts.TradeTypeAliWeb
	case consts.PayTypeWxPay:
		tradeType = consts.TradeTypeWxScan
	default:
		return nil, gerror.New("不支持的充值支付渠道")
	}
	return &entity.PayLog{
		MemberId: userID, AppId: consts.AppApi, OrderSn: recharge.OrderSn,
		OrderGroup: consts.OrderGroupUserRecharge,
		Subject:    fmt.Sprintf("会员余额充值%d元", recharge.Money),
		PayType:    recharge.PayType, PayAmount: float64(recharge.Money),
		PayStatus: consts.PayStatusWait, TradeType: tradeType,
		IsRefund: consts.RefundStatusNo, Status: consts.StatusEnabled,
	}, nil
}

// rechargeNormalizePayment 统一创建结果和历史快照的订单匹配及前台字段白名单。
// 交易方式的准入由调用处决定：新订单须匹配预期方式，历史订单保留旧 scan 兼容。
func rechargeNormalizePayment(order *payin.CreateOrderModel, outTradeNo string) *payin.CreateOrderModel {
	if order == nil || strings.TrimSpace(order.PayURL) == "" || order.OutTradeNo != outTradeNo {
		return nil
	}
	return &payin.CreateOrderModel{TradeType: order.TradeType, PayURL: order.PayURL, OutTradeNo: outTradeNo}
}

// createRechargePayment 直接创建前台充值支付记录，避免使用后台角色权限模型。
func (s *sApiUser) createRechargePayment(ctx context.Context, userID int64, recharge *entity.UserRecharge) (*payin.CreateOrderModel, error) {
	pay, err := rechargePayLog(userID, recharge)
	if err != nil {
		return nil, err
	}
	if pay.PayType == consts.PayTypeAliPay {
		userAgent := ""
		if request := ghttp.RequestFromCtx(ctx); request != nil {
			userAgent = request.UserAgent()
		}
		pay.TradeType = rechargeAlipayTradeType(userAgent)
	}
	payInput := payin.PayCreateInp{PayType: pay.PayType}
	notifyURL, err := service.Pay().GenNotifyURL(ctx, payInput)
	if err != nil {
		return nil, gerror.Wrap(err, "生成充值支付通知地址失败")
	}
	config := payment.GetConfig()
	if config == nil {
		return nil, gerror.New("支付配置不可用")
	}
	// rechargePayLog 已限定渠道，仅从同一份配置选取对应商户号。
	pay.MchId = config.WxPayMchId
	if pay.PayType == consts.PayTypeAliPay {
		pay.MchId = config.AliPayAppId
	}
	pay.OutTradeNo = payment.GenOutTradeNo(ctx)
	pay.NotifyUrl = notifyURL
	pay.TraceIds = gjson.New([]string{gctx.CtxId(ctx)})
	if _, err = dao.PayLog.Ctx(ctx).Data(pay).OmitEmptyData().Insert(); err != nil {
		return nil, gerror.Wrap(err, "创建充值支付记录失败")
	}
	var order *payin.CreateOrderModel
	if pay.PayType == consts.PayTypeAliPay && pay.TradeType == rechargeTradeTypeAliPrecreate {
		client, clientErr := paymentAlipay.GetClient(config)
		if clientErr != nil {
			return nil, gerror.Wrap(clientErr, "创建支付宝当面付客户端失败")
		}
		client.SetNotifyUrl(pay.NotifyUrl)
		response, precreateErr := client.TradePrecreate(ctx, rechargePrecreateParams(recharge, pay))
		if precreateErr != nil {
			return nil, gerror.Wrap(precreateErr, "支付宝当面付预下单失败，请核对充值记录")
		}
		order, err = rechargePrecreatePayment(response, pay.OutTradeNo)
	} else {
		order, err = payment.New(recharge.PayType).CreateOrder(ctx, payin.CreateOrderInp{Pay: pay})
	}
	if err != nil {
		return nil, err
	}
	order = rechargeNormalizePayment(order, pay.OutTradeNo)
	if order == nil || order.TradeType != pay.TradeType {
		return nil, gerror.New("支付渠道未返回匹配的扫码信息，请联系管理员核对充值记录")
	}
	snapshot, err := json.Marshal(order)
	if err != nil {
		return nil, gerror.Wrap(err, "保存充值支付信息失败")
	}
	cols := dao.PayLog.Columns()
	if _, err = dao.PayLog.Ctx(ctx).
		Where(cols.OutTradeNo, pay.OutTradeNo).
		Where(cols.OrderGroup, consts.OrderGroupUserRecharge).
		Where(cols.AppId, consts.AppApi).
		Where(cols.MemberId, userID).
		Data(g.Map{cols.Custom: string(snapshot)}).Update(); err != nil {
		return nil, gerror.Wrap(err, "保存充值支付信息失败，请在充值记录中核对订单")
	}
	return order, nil
}

// 历史订单没有扫码快照时不重新下单；支付已处理的记录不再展示二维码。
func rechargePayment(pay *entity.PayLog) *payin.CreateOrderModel {
	if pay == nil || pay.PayStatus != consts.PayStatusWait || pay.Status != consts.StatusEnabled || pay.Custom == "" {
		return nil
	}
	var snapshot payin.CreateOrderModel
	if err := json.Unmarshal([]byte(pay.Custom), &snapshot); err != nil {
		return nil
	}
	if snapshot.TradeType != consts.TradeTypeWxScan && !(pay.PayType == consts.PayTypeAliPay && (snapshot.TradeType == consts.TradeTypeAliWeb || snapshot.TradeType == rechargeTradeTypeAliPrecreate)) {
		return nil
	}
	return rechargeNormalizePayment(&snapshot, pay.OutTradeNo)
}
