package user

import (
	"context"
	"math"

	v1 "hotgo/api/api/user/v1"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/payin"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
)

func (s *sApiUser) RechargeOptions(ctx context.Context, _ *v1.RechargeOptionsReq) (*v1.RechargeOptionsRes, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.RechargeOptionsRes{
		Balance: user.Money, Amounts: []uint{50, 100, 200, 500, 1000, 2000}, MinAmount: 1, MaxAmount: 100000,
		PaymentMethods: rechargePaymentMethods(),
		StatusOptions: []v1.RechargeOption{
			{Value: consts.UserRechargePending, Label: "待支付"},
			{Value: consts.UserRechargeCompleted, Label: "已到账"},
		},
	}, nil
}

func (s *sApiUser) RechargeCreate(ctx context.Context, req *v1.RechargeCreateReq) (*v1.RechargeCreateRes, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	available := false
	for _, method := range rechargePaymentMethods() {
		if method.Value == req.PayType {
			available = true
			break
		}
	}
	if !available {
		return nil, gerror.New("该支付渠道尚未配置，请联系管理员")
	}
	now := gtime.Now()
	recharge := &entity.UserRecharge{
		UserId: int64(user.Id), OrderSn: "UR" + guid.S(), Money: gconv.Uint(req.Money),
		PayType: req.PayType, Status: consts.UserRechargePending, CreatedAt: now, UpdatedAt: now,
	}
	// 先持久化业务单，再调用第三方；网关失败仍保留待支付记录用于核对，不写入余额。
	id, err := dao.UserRecharge.Ctx(ctx).Data(recharge).InsertAndGetId()
	if err != nil {
		return nil, gerror.Wrap(err, "创建充值订单失败，请稍后重试")
	}
	recharge.Id = id
	order, err := s.createRechargePayment(ctx, int64(user.Id), recharge)
	if err != nil {
		return nil, gerror.Wrap(err, "获取支付信息失败，请稍后重试或联系管理员核对充值记录")
	}
	if order == nil || order.PayURL == "" {
		return nil, gerror.New("支付渠道未返回可用支付信息，请联系管理员核对充值记录")
	}
	return &v1.RechargeCreateRes{Recharge: recharge, Payment: order}, nil
}

func (s *sApiUser) RechargeView(ctx context.Context, req *v1.RechargeViewReq) (*v1.RechargeViewRes, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	cols := dao.UserRecharge.Columns()
	var recharge *entity.UserRecharge
	if err = dao.UserRecharge.Ctx(ctx).Where(cols.UserId, user.Id).Where(cols.OrderSn, req.OrderSn).Scan(&recharge); err != nil {
		return nil, gerror.Wrap(err, "获取充值订单失败")
	}
	if recharge == nil {
		return nil, gerror.New("充值订单不存在")
	}
	res := &v1.RechargeViewRes{Recharge: recharge, Balance: user.Money}
	if recharge.Status == consts.UserRechargePending {
		payCols := dao.PayLog.Columns()
		var pay *entity.PayLog
		if err = dao.PayLog.Ctx(ctx).
			Where(payCols.OrderSn, recharge.OrderSn).
			Where(payCols.OrderGroup, consts.OrderGroupUserRecharge).
			Where(payCols.AppId, consts.AppApi).
			Where(payCols.MemberId, user.Id).
			Where(payCols.PayType, recharge.PayType).
			OrderDesc(payCols.Id).Scan(&pay); err != nil {
			return nil, gerror.Wrap(err, "获取原订单支付信息失败")
		}
		res.Payment = rechargePayment(pay)
	}
	return res, nil
}

func (s *sApiUser) RechargeList(ctx context.Context, req *v1.RechargeListReq) (*v1.RechargeListRes, error) {
	user, err := s.currentUser(ctx)
	if err != nil {
		return nil, err
	}
	cols := dao.UserRecharge.Columns()
	mod := dao.UserRecharge.Ctx(ctx).Where(cols.UserId, user.Id)
	total, err := mod.Clone().Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取充值记录失败")
	}
	list := make([]*entity.UserRecharge, 0)
	if err = mod.Page(req.Page, req.PerPage).OrderDesc(cols.Id).Scan(&list); err != nil {
		return nil, gerror.Wrap(err, "获取充值记录失败")
	}
	return &v1.RechargeListRes{List: list, Page: req.Page, PerPage: req.PerPage, TotalCount: total}, nil
}

// RechargeNotify 仅由验签后的支付网关调用；订单锁保证重复通知不重复入账。
func (s *sApiUser) RechargeNotify(ctx context.Context, in *payin.NotifyCallFuncInp) error {
	pay := in.Pay
	if pay == nil || pay.OrderGroup != consts.OrderGroupUserRecharge || pay.AppId != consts.AppApi || pay.PayStatus != consts.PayStatusOk {
		return gerror.New("充值支付记录不正确")
	}
	return dao.UserRecharge.Transaction(ctx, func(ctx context.Context, _ gdb.TX) error {
		cols := dao.UserRecharge.Columns()
		var recharge *entity.UserRecharge
		if err := dao.UserRecharge.Ctx(ctx).Where(cols.OrderSn, pay.OrderSn).LockUpdate().Scan(&recharge); err != nil {
			return err
		}
		if recharge == nil || recharge.UserId != pay.MemberId || recharge.PayType != pay.PayType || float64(recharge.Money) != pay.PayAmount || math.Abs(pay.ActualAmount-pay.PayAmount) > 0.005 {
			return gerror.New("充值订单与支付记录不一致，请人工核对")
		}
		if recharge.Status == consts.UserRechargeCompleted {
			return nil
		}
		if recharge.Status != consts.UserRechargePending {
			return gerror.New("充值订单状态不正确")
		}
		var user *entity.User
		if err := dao.User.Ctx(ctx).WherePri(recharge.UserId).LockUpdate().Scan(&user); err != nil {
			return err
		}
		if user == nil {
			return gerror.New("充值账户不存在，请人工核对")
		}
		after := uint64(user.Money) + uint64(recharge.Money)
		if after > math.MaxUint32 {
			return gerror.New("账户余额超出上限，请人工核对")
		}
		now := gtime.Now()
		userCols := dao.User.Columns()
		if _, err := dao.User.Ctx(ctx).WherePri(user.Id).Data(g.Map{userCols.Money: after, userCols.UpdatedAt: now}).Update(); err != nil {
			return err
		}
		logCols := dao.UserCreditsLog.Columns()
		if _, err := dao.UserCreditsLog.Ctx(ctx).Data(g.Map{
			logCols.UserId: user.Id, logCols.CreditType: consts.UserCreditTypeMoney,
			logCols.BeforeNum: user.Money, logCols.Num: recharge.Money, logCols.AfterNum: after,
			logCols.Remark: "在线充值 " + recharge.OrderSn, logCols.MapId: recharge.Id,
			logCols.Status: consts.StatusEnabled, logCols.CreatedAt: now, logCols.UpdatedAt: now,
		}).Insert(); err != nil {
			return err
		}
		_, err := dao.UserRecharge.Ctx(ctx).WherePri(recharge.Id).Data(g.Map{
			cols.Status: consts.UserRechargeCompleted, cols.PayAt: pay.PayAt, cols.UpdatedAt: now,
		}).Update()
		return err
	})
}
