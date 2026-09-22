package shop

import (
	"context"
	"fmt"
	"math"

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

	var products []*entity.WebProducts
	if err = dao.WebProducts.Ctx(ctx).
		WhereIn(dao.WebProducts.Columns().Id, productIds).
		Where(dao.WebProducts.Columns().Status, "normal").
		Scan(&products); err != nil {
		return nil, gerror.Wrap(err, "查询商品失败，请稍后重试")
	}
	if len(products) != len(productIds) {
		return nil, gerror.New("存在不存在或已下架的商品")
	}

	productMap := make(map[uint64]*entity.WebProducts, len(products))
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

		items = make([]*entity.ShopOrderItem, 0, len(productIds))
		var totalCents int64
		for _, productId := range productIds {
			product := productMap[productId]
			unitPriceCents := amountToCents(product.Price)
			if unitPriceCents > 999999999999 {
				return gerror.New("商品价格超出订单金额范围")
			}
			quantity := quantities[productId]
			if unitPriceCents > math.MaxInt64/int64(quantity) {
				return gerror.New("订单金额超出范围")
			}
			itemCents := unitPriceCents * int64(quantity)
			if itemCents > 999999999999 || totalCents > 999999999999-itemCents {
				return gerror.New("订单金额超出范围")
			}
			totalCents += itemCents
			items = append(items, &entity.ShopOrderItem{
				ProductId:   productId,
				Title:       product.Title,
				Image:       product.Image,
				UnitPrice:   centsToAmount(unitPriceCents),
				Quantity:    quantity,
				TotalAmount: centsToAmount(itemCents),
				CreateTime:  now,
				UpdateTime:  now,
			})
		}

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

func amountToCents(amount float64) int64 {
	return int64(math.Round(amount * 100))
}

func centsToAmount(cents int64) float64 {
	return float64(cents) / 100
}
