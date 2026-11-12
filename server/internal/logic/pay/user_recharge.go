package pay

import (
	"context"

	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/payin"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// 会员充值必须在入账事务提交后才确认收单；失败时保留待支付状态，允许渠道重新通知。
func (s *sPay) notifyUserRecharge(ctx context.Context, original *entity.PayLog, data *payin.NotifyModel) error {
	return dao.PayLog.Transaction(ctx, func(ctx context.Context, _ gdb.TX) error {
		var pay *entity.PayLog
		if err := dao.PayLog.Ctx(ctx).WherePri(original.Id).LockUpdate().Scan(&pay); err != nil {
			return err
		}
		if pay == nil || pay.OrderGroup != consts.OrderGroupUserRecharge || pay.AppId != consts.AppApi || pay.PayAmount != data.ActualAmount || data.TransactionId == "" || data.PayAt == nil {
			return gerror.New("充值收单金额或支付记录不一致，请人工核对")
		}
		if pay.PayStatus == consts.PayStatusOk {
			return service.ApiUser().RechargeNotify(ctx, &payin.NotifyCallFuncInp{Pay: pay})
		}
		if pay.PayStatus != consts.PayStatusWait {
			return gerror.New("充值支付状态不正确")
		}
		cols := dao.PayLog.Columns()
		if _, err := dao.PayLog.Ctx(ctx).WherePri(pay.Id).Data(g.Map{
			cols.TransactionId: data.TransactionId, cols.PayStatus: consts.PayStatusOk,
			cols.PayAt: data.PayAt, cols.ActualAmount: data.ActualAmount,
		}).Update(); err != nil {
			return err
		}
		pay.TransactionId = data.TransactionId
		pay.PayStatus = consts.PayStatusOk
		pay.PayAt = data.PayAt
		pay.ActualAmount = data.ActualAmount
		return service.ApiUser().RechargeNotify(ctx, &payin.NotifyCallFuncInp{Pay: pay})
	})
}
