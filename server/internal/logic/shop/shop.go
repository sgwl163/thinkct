package shop

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/frame/g"

	v1 "hotgo/api/api/shop/v1"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/library/contexts"
	"hotgo/internal/model/entity"
	"hotgo/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/grand"
)

type sApiShop struct{}

func init() {
	service.RegisterApiShop(&sApiShop{})
}

func (s *sApiShop) CreateOrder(ctx context.Context, req *v1.CreateOrderReq) (res *v1.CreateOrderRes, err error) {
	userId, err := getShopUserId(ctx)
	if err != nil {
		return nil, err
	}

	quantities := make(map[uint64]uint, len(req.Products))
	productIds := make([]uint64, 0, len(req.Products))
	for _, product := range req.Products {
		if _, exists := quantities[product.ProductId]; !exists {
			productIds = append(productIds, product.ProductId)
		}
		if product.Quantity > uint(math.MaxInt64) || product.Quantity > math.MaxUint32-quantities[product.ProductId] {
			return nil, gerror.New("商品购买数量超出范围")
		}
		quantities[product.ProductId] += product.Quantity
	}
	sort.Slice(productIds, func(i, j int) bool { return productIds[i] < productIds[j] })
	var products []*entity.ShopProduct
	if err = dao.ShopProduct.Ctx(ctx).
		WhereIn(dao.ShopProduct.Columns().Id, productIds).
		Where(dao.ShopProduct.Columns().Status, "normal").
		Scan(&products); err != nil {
		return nil, gerror.Wrap(err, "查询商品失败，请稍后重试")
	}
	if len(products) != len(productIds) {
		return nil, gerror.New("存在不存在或已下架的商品")
	}

	productMap := make(map[uint64]*entity.ShopProduct, len(products))
	for _, product := range products {
		if product.Price < 0 {
			return nil, gerror.New("商品价格异常")
		}
		productMap[uint64(product.Id)] = product
	}

	now := uint64(gtime.Now().Unix())
	var (
		order *entity.ShopOrder
		items []*entity.ShopOrderItem
	)
	err = dao.ShopOrder.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		orderSn, txErr := nextShopOrderSn(ctx, tx)
		if txErr != nil {
			return txErr
		}

		pricedItems, totalCents, txErr := priceOrder(productIds, quantities, productMap, now)
		if txErr != nil {
			return txErr
		}
		items = pricedItems
		order = &entity.ShopOrder{
			OrderSn:       orderSn,
			UserId:        uint64(userId),
			Status:        consts.ShopOrderStatusPendingPayment,
			TotalAmount:   centsToAmount(totalCents),
			PayableAmount: centsToAmount(totalCents),
			Remark:        req.Remark,
			UpdateTime:    now,
			CreateTime:    now,
		}
		orderId, txErr := tx.Model(dao.ShopOrder.Table()).Data(order).InsertAndGetId()
		if txErr != nil {
			return gerror.Wrap(txErr, "创建订单失败，请稍后重试")
		}
		order.Id = uint64(orderId)

		for _, item := range items {
			item.OrderId = order.Id
		}
		if _, txErr = tx.Model(dao.ShopOrderItem.Table()).Data(items).Insert(); txErr != nil {
			return gerror.Wrap(txErr, "创建订单明细失败，请稍后重试")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateOrderRes{Order: newOrderDetail(order, items)}, nil
}

func (s *sApiShop) OrderList(ctx context.Context, req *v1.OrderListReq) (res *v1.OrderListRes, err error) {
	userId, err := getShopUserId(ctx)
	if err != nil {
		return nil, err
	}

	model := dao.ShopOrder.Ctx(ctx).Where(dao.ShopOrder.Columns().UserId, userId)
	if req.Status != "" {
		if !isShopOrderStatus(req.Status) {
			return nil, gerror.New("订单状态不正确")
		}
		model = model.Where(dao.ShopOrder.Columns().Status, req.Status)
	}

	totalCount, err := model.Count()
	if err != nil {
		return nil, gerror.Wrap(err, "获取订单列表失败，请稍后重试")
	}

	list := make([]*entity.ShopOrder, 0)
	if totalCount > 0 {
		if err = model.Page(req.Page, req.PerPage).OrderDesc(dao.ShopOrder.Columns().Id).Scan(&list); err != nil {
			return nil, gerror.Wrap(err, "获取订单列表失败，请稍后重试")
		}
	}
	orders, err := loadOrderDetails(ctx, list)
	if err != nil {
		return nil, err
	}

	res = &v1.OrderListRes{List: orders}
	res.PageRes.Pack(req, totalCount)
	return res, nil
}

func (s *sApiShop) OrderDetail(ctx context.Context, req *v1.OrderDetailReq) (res *v1.OrderDetailRes, err error) {
	userId, err := getShopUserId(ctx)
	if err != nil {
		return nil, err
	}

	var order *entity.ShopOrder
	if err = dao.ShopOrder.Ctx(ctx).
		Where(dao.ShopOrder.Columns().OrderSn, req.OrderSn).
		Where(dao.ShopOrder.Columns().UserId, userId).
		Scan(&order); err != nil {
		return nil, gerror.Wrap(err, "获取订单详情失败，请稍后重试")
	}
	if order == nil {
		return nil, gerror.New("订单不存在")
	}

	orders, err := loadOrderDetails(ctx, []*entity.ShopOrder{order})
	if err != nil {
		return nil, err
	}
	return &v1.OrderDetailRes{Order: orders[0]}, nil
}

func getShopUserId(ctx context.Context) (int64, error) {
	userId := contexts.GetUserId(ctx)
	if userId <= 0 {
		return 0, gerror.New("请先登录")
	}
	return userId, nil
}

func nextShopOrderSn(ctx context.Context, tx gdb.TX) (string, error) {
	for i := 0; i < 10; i++ {
		orderSn := fmt.Sprintf("SHOP%s%s", gtime.Now().Format("YmdHis"), grand.S(8))
		count, err := tx.Model(dao.ShopOrder.Table()).Where(dao.ShopOrder.Columns().OrderSn, orderSn).Count()
		if err != nil {
			return "", gerror.Wrap(err, "生成订单号失败，请稍后重试")
		}
		if count == 0 {
			return orderSn, nil
		}
	}
	return "", gerror.New("生成订单号失败，请稍后重试")
}

func loadOrderDetails(ctx context.Context, orders []*entity.ShopOrder) ([]*v1.OrderDetail, error) {
	if len(orders) == 0 {
		return make([]*v1.OrderDetail, 0), nil
	}

	orderIds := make([]uint64, 0, len(orders))
	itemsByOrderId := make(map[uint64][]*entity.ShopOrderItem, len(orders))
	for _, order := range orders {
		orderIds = append(orderIds, order.Id)
		itemsByOrderId[order.Id] = make([]*entity.ShopOrderItem, 0)
	}

	var items []*entity.ShopOrderItem
	if err := dao.ShopOrderItem.Ctx(ctx).WhereIn(dao.ShopOrderItem.Columns().OrderId, orderIds).OrderAsc(dao.ShopOrderItem.Columns().Id).Scan(&items); err != nil {
		return nil, gerror.Wrap(err, "获取订单明细失败，请稍后重试")
	}
	for _, item := range items {
		itemsByOrderId[item.OrderId] = append(itemsByOrderId[item.OrderId], item)
	}

	result := make([]*v1.OrderDetail, 0, len(orders))
	for _, order := range orders {
		result = append(result, newOrderDetail(order, itemsByOrderId[order.Id]))
	}
	return result, nil
}

func newOrderDetail(order *entity.ShopOrder, items []*entity.ShopOrderItem) *v1.OrderDetail {
	if items == nil {
		items = make([]*entity.ShopOrderItem, 0)
	}
	return &v1.OrderDetail{ShopOrder: *order, Items: items}
}

func isShopOrderStatus(status string) bool {
	for _, option := range consts.ShopOrderStatusOptions {
		if option.Key == status {
			return true
		}
	}
	return false
}

const maxShopCents int64 = 999999999999

func normalizeProducts(products []v1.OrderProduct) ([]uint64, map[uint64]uint, error) {
	if len(products) == 0 || len(products) > 50 {
		return nil, nil, gerror.New("商品数量须为1至50种")
	}
	ids := make([]uint64, 0, len(products))
	quantities := make(map[uint64]uint, len(products))
	for _, product := range products {
		if product.ProductId == 0 || product.Quantity == 0 || product.Quantity > 999 || quantities[product.ProductId] > 999-product.Quantity {
			return nil, nil, gerror.New("商品购买数量须为1至999，重复商品合并后不能超过999")
		}
		if _, ok := quantities[product.ProductId]; !ok {
			ids = append(ids, product.ProductId)
		}
		quantities[product.ProductId] += product.Quantity
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	return ids, quantities, nil
}

// 按浮点字段的十进制表示校验分值，拒绝隐藏的小数分，不将其四舍五入。
func priceCents(price float64) (int64, error) {
	if math.IsNaN(price) || math.IsInf(price, 0) || price < 0 || price > centsToAmount(maxShopCents) {
		return 0, gerror.New("商品价格异常或超出订单金额范围")
	}
	text := strconv.FormatFloat(price, 'f', -1, 64)
	parts := strings.SplitN(text, ".", 2)
	whole, _ := strconv.ParseInt(parts[0], 10, 64)
	var fraction int64
	if len(parts) == 2 {
		if len(parts[1]) > 2 {
			return 0, gerror.New("商品价格精度不能超过两位小数")
		}
		fraction, _ = strconv.ParseInt(parts[1]+strings.Repeat("0", 2-len(parts[1])), 10, 64)
	}
	return whole*100 + fraction, nil
}

func priceOrder(ids []uint64, quantities map[uint64]uint, products map[uint64]*entity.ShopProduct, now uint64) ([]*entity.ShopOrderItem, int64, error) {
	items := make([]*entity.ShopOrderItem, 0, len(ids))
	var total int64
	for _, id := range ids {
		product := products[id]
		if product == nil {
			return nil, 0, gerror.New("存在不存在或已下架的商品")
		}
		cents, err := priceCents(product.Price)
		if err != nil {
			return nil, 0, err
		}
		quantity := quantities[id]
		if quantity == 0 || quantity > math.MaxUint32 || cents > maxShopCents/int64(quantity) {
			return nil, 0, gerror.New("订单金额或数量超出范围")
		}
		subtotal := cents * int64(quantity)
		if total > maxShopCents-subtotal {
			return nil, 0, gerror.New("订单金额超出范围")
		}
		total += subtotal
		items = append(items, &entity.ShopOrderItem{ProductId: id, Title: product.Title, Image: product.Image,
			UnitPrice: centsToAmount(cents), Quantity: quantity, TotalAmount: centsToAmount(subtotal), CreateTime: now, UpdateTime: now})
	}
	return items, total, nil
}

func balancePurchaseAmount(total int64, balance uint) (int64, error) {
	if total < 0 || total > maxShopCents {
		return 0, gerror.New("订单金额超出范围")
	}
	if total%100 != 0 {
		return 0, gerror.New("余额仅支持整数元支付，订单总金额含小数元，无法购买")
	}
	amount := total / 100
	if uint64(balance) < uint64(amount) {
		return 0, gerror.New("余额不足，请先充值后再购买")
	}
	return amount, nil
}

func purchaseOrderSn(userID int64, requestID string) string {
	digest := sha256.Sum256([]byte(strconv.FormatInt(userID, 10) + ":" + strings.ToLower(requestID)))
	return fmt.Sprintf("BUY%x", digest[:30])
}

func samePurchase(order *entity.ShopOrder, items []*entity.ShopOrderItem, userID int64, req *v1.PurchaseOrderReq) bool {
	if order.UserId != uint64(userID) || order.PaymentType != "balance" ||
		(order.Status != consts.ShopOrderStatusPaid && order.Status != consts.ShopOrderStatusCompleted) || order.Remark != req.Remark {
		return false
	}
	_, expected, err := normalizeProducts(req.Products)
	if err != nil {
		return false
	}
	products := make([]v1.OrderProduct, 0, len(items))
	for _, item := range items {
		products = append(products, v1.OrderProduct{ProductId: item.ProductId, Quantity: item.Quantity})
	}
	ids, actual, err := normalizeProducts(products)
	if err != nil || len(actual) != len(expected) {
		return false
	}
	for _, id := range ids {
		if actual[id] != expected[id] {
			return false
		}
	}
	return true
}

func (s *sApiShop) PurchaseOrder(ctx context.Context, req *v1.PurchaseOrderReq) (*v1.PurchaseOrderRes, error) {
	return s.purchaseOrder(ctx, req, nil)
}

// Direct purchase and cart checkout share one transaction, including cart removal.
func (s *sApiShop) purchaseOrder(ctx context.Context, req *v1.PurchaseOrderReq, checkout *v1.CartCheckoutReq) (*v1.PurchaseOrderRes, error) {
	userID, err := getShopUserId(ctx)
	if err != nil {
		return nil, err
	}
	var ids []uint64
	var quantities map[uint64]uint
	var cartIDs []uint64
	var fingerprint string
	if checkout == nil {
		if err = g.Validator().Data(req).Run(ctx); err != nil {
			return nil, err
		}
		ids, quantities, err = normalizeProducts(req.Products)
	} else {
		if err = g.Validator().Data(checkout).Run(ctx); err != nil {
			return nil, err
		}
		cartIDs, fingerprint, err = checkoutPayload(checkout.CartIds, checkout.Remark)
	}
	if err != nil {
		return nil, err
	}
	var order *entity.ShopOrder
	var items []*entity.ShopOrderItem
	err = dao.ShopOrder.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 同用户先锁余额行，再查询唯一订单号，串行化首次购买和重试。
		var user *entity.User
		if err := tx.Model(dao.User.Table()).WherePri(userID).LockUpdate().Scan(&user); err != nil {
			return gerror.Wrap(err, "读取账户失败")
		}
		if user == nil || (user.Status != "enable" && user.Status != "1") {
			return gerror.New("账户不存在或已停用")
		}
		orderSn := purchaseOrderSn(userID, req.RequestId)
		if err := tx.Model(dao.ShopOrder.Table()).Where(dao.ShopOrder.Columns().OrderSn, orderSn).LockUpdate().Scan(&order); err != nil {
			return gerror.Wrap(err, "读取订单失败")
		}
		if order != nil {
			if err := tx.Model(dao.ShopOrderItem.Table()).Where(dao.ShopOrderItem.Columns().OrderId, order.Id).OrderAsc(dao.ShopOrderItem.Columns().ProductId).Scan(&items); err != nil {
				return gerror.Wrap(err, "读取订单明细失败")
			}
			if checkout != nil {
				if !sameCheckout(order, userID, fingerprint, checkout.Remark) {
					return gerror.New("请求标识已使用，结算内容或订单状态不一致")
				}
			} else if order.PaymentNo != "" || !samePurchase(order, items, userID, req) {
				return gerror.New("请求标识已使用，购买内容或订单状态不一致")
			}
			return nil
		}
		if checkout != nil {
			var carts []*entity.ShopCart
			cc := dao.ShopCart.Columns()
			if err := tx.Model(dao.ShopCart.Table()).Where(cc.UserId, userID).WhereIn(cc.Id, cartIDs).OrderAsc(cc.Id).LockUpdate().Scan(&carts); err != nil {
				return gerror.Wrap(err, "读取购物车失败")
			}
			if len(carts) != len(cartIDs) {
				return gerror.New("所选购物车不存在或不属于当前用户")
			}
			selected := make([]v1.OrderProduct, 0, len(carts))
			for _, cart := range carts {
				selected = append(selected, v1.OrderProduct{ProductId: cart.ProductId, Quantity: cart.Quantity})
			}
			var err error
			ids, quantities, err = normalizeProducts(selected)
			if err != nil {
				return err
			}
		}
		products := make(map[uint64]*entity.ShopProduct, len(ids))
		// 按商品ID逐行锁定，保持所有购买事务的锁顺序一致。
		for _, id := range ids {
			var product *entity.ShopProduct
			if err := tx.Model(dao.ShopProduct.Table()).WherePri(id).Where(dao.ShopProduct.Columns().Status, "normal").LockUpdate().Scan(&product); err != nil {
				return gerror.Wrap(err, "读取商品失败")
			}
			products[id] = product
		}
		now := gtime.Now()
		var total int64
		var err error
		items, total, err = priceOrder(ids, quantities, products, uint64(now.Unix()))
		if err != nil {
			return err
		}
		amount, err := balancePurchaseAmount(total, user.Money)
		if err != nil {
			return err
		}
		order = &entity.ShopOrder{OrderSn: orderSn, UserId: uint64(userID), Status: consts.ShopOrderStatusPaid,
			TotalAmount: centsToAmount(total), PayableAmount: centsToAmount(total), PaidAmount: centsToAmount(total),
			Remark: req.Remark, PaymentType: "balance", PaymentNo: fingerprint, PaymentTime: uint64(now.Unix()), CreateTime: uint64(now.Unix()), UpdateTime: uint64(now.Unix())}
		orderID, err := tx.Model(dao.ShopOrder.Table()).Data(order).InsertAndGetId()
		if err != nil {
			return gerror.Wrap(err, "创建订单失败")
		}
		order.Id = uint64(orderID)
		for _, item := range items {
			item.OrderId = order.Id
		}
		if _, err = tx.Model(dao.ShopOrderItem.Table()).Data(items).Insert(); err != nil {
			return gerror.Wrap(err, "创建订单明细失败")
		}
		uc := dao.User.Columns()
		// 零元单无需UPDATE，避免MySQL未变化行数为0；仍记录零元流水。
		if amount > 0 {
			result, err := tx.Model(dao.User.Table()).WherePri(userID).Where(uc.Money, user.Money).
				WhereGTE(uc.Money, amount).Data(g.Map{uc.Money: uint64(user.Money) - uint64(amount), uc.UpdatedAt: now}).Update()
			if err != nil {
				return gerror.Wrap(err, "扣除余额失败")
			}
			rows, err := result.RowsAffected()
			if err != nil || rows != 1 {
				return gerror.New("扣除余额失败，请重试")
			}
		}
		lc := dao.UserCreditsLog.Columns()
		_, err = tx.Model(dao.UserCreditsLog.Table()).Data(g.Map{
			lc.UserId: userID, lc.CreditType: consts.UserCreditTypeMoney, lc.BeforeNum: user.Money,
			lc.Num: -amount, lc.AfterNum: uint64(user.Money) - uint64(amount), lc.MapId: order.Id,
			lc.Remark: "商城订单 " + orderSn, lc.Status: consts.StatusEnabled, lc.CreatedAt: now, lc.UpdatedAt: now,
		}).Insert()
		if err != nil {
			return gerror.Wrap(err, "记录余额流水失败")
		}
		if checkout != nil {
			cc := dao.ShopCart.Columns()
			result, err := tx.Model(dao.ShopCart.Table()).Where(cc.UserId, userID).WhereIn(cc.Id, cartIDs).Delete()
			if err != nil {
				return gerror.Wrap(err, "清理已结算购物车失败")
			}
			rows, err := result.RowsAffected()
			if err != nil || rows != int64(len(cartIDs)) {
				return gerror.New("购物车已变化，请重新结算")
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &v1.PurchaseOrderRes{Order: newOrderDetail(order, items)}, nil
}

func centsToAmount(cents int64) float64 {
	return float64(cents) / 100
}
