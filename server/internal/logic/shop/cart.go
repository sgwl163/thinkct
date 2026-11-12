package shop

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	v1 "hotgo/api/api/shop/v1"
	"hotgo/internal/consts"
	"hotgo/internal/dao"
	"hotgo/internal/model/entity"
	"sort"
)

// Canonical IDs plus remark bind a retry even after the selected cart rows are deleted.
func checkoutPayload(input []uint64, remark string) ([]uint64, string, error) {
	if len(input) == 0 || len(input) > 50 {
		return nil, "", gerror.New("请选择1至50个购物车商品")
	}
	ids := append([]uint64(nil), input...)
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	for i, id := range ids {
		if id == 0 || (i > 0 && id == ids[i-1]) {
			return nil, "", gerror.New("购物车ID无效或重复")
		}
	}
	data, err := json.Marshal(struct {
		CartIds []uint64 `json:"cartIds"`
		Remark  string   `json:"remark"`
	}{ids, remark})
	if err != nil {
		return nil, "", err
	}
	digest := sha256.Sum256(data)
	return ids, fmt.Sprintf("cart:%x", digest), nil
}
func sameCheckout(order *entity.ShopOrder, userID int64, fingerprint, remark string) bool {
	return order.UserId == uint64(userID) && order.PaymentType == "balance" && order.PaymentNo == fingerprint && order.Remark == remark && (order.Status == consts.ShopOrderStatusPaid || order.Status == consts.ShopOrderStatusCompleted)
}
func (s *sApiShop) CartCheckout(ctx context.Context, req *v1.CartCheckoutReq) (*v1.CartCheckoutRes, error) {
	res, err := s.purchaseOrder(ctx, &v1.PurchaseOrderReq{RequestId: req.RequestId, Remark: req.Remark}, req)
	if err != nil {
		return nil, err
	}
	return &v1.CartCheckoutRes{Order: res.Order}, nil
}

// All cart writes use the same user-first lock order as checkout.
func cartWrite(ctx context.Context, fn func(gdb.TX, int64) error) error {
	userID, err := getShopUserId(ctx)
	if err != nil {
		return err
	}
	return dao.ShopCart.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		var user *entity.User
		if err := tx.Model(dao.User.Table()).WherePri(userID).LockUpdate().Scan(&user); err != nil {
			return gerror.Wrap(err, "读取账户失败")
		}
		if user == nil || (user.Status != "enable" && user.Status != "1") {
			return gerror.New("账户不存在或已停用")
		}
		return fn(tx, userID)
	})
}
func (s *sApiShop) CartAdd(ctx context.Context, req *v1.CartAddReq) (*v1.CartAddRes, error) {
	if err := g.Validator().Data(req).Run(ctx); err != nil {
		return nil, err
	}
	if _, _, err := normalizeProducts([]v1.OrderProduct{req.OrderProduct}); err != nil {
		return nil, err
	}
	var id uint64
	err := cartWrite(ctx, func(tx gdb.TX, userID int64) error {
		var product *entity.ShopProduct
		if err := tx.Model(dao.ShopProduct.Table()).WherePri(req.ProductId).Where(dao.ShopProduct.Columns().Status, "normal").Scan(&product); err != nil {
			return gerror.Wrap(err, "读取商品失败")
		}
		if product == nil {
			return gerror.New("商品不存在或已下架")
		}
		if _, err := priceCents(product.Price); err != nil {
			return err
		}
		cc := dao.ShopCart.Columns()
		var cart *entity.ShopCart
		if err := tx.Model(dao.ShopCart.Table()).Where(cc.UserId, userID).Where(cc.ProductId, req.ProductId).LockUpdate().Scan(&cart); err != nil {
			return err
		}
		now := uint64(gtime.Now().Unix())
		if cart != nil {
			if cart.Quantity > 999 || req.Quantity > 999-cart.Quantity {
				return gerror.New("商品购买数量不能超过999")
			}
			id = cart.Id
			_, err := tx.Model(dao.ShopCart.Table()).WherePri(id).Data(g.Map{cc.Quantity: cart.Quantity + req.Quantity, cc.UpdatedAt: now}).Update()
			return err
		}
		inserted, err := tx.Model(dao.ShopCart.Table()).Data(g.Map{cc.UserId: userID, cc.ProductId: req.ProductId, cc.Quantity: req.Quantity, cc.CreatedAt: now, cc.UpdatedAt: now}).InsertAndGetId()
		id = uint64(inserted)
		return err
	})
	if err != nil {
		return nil, err
	}
	return &v1.CartAddRes{Id: id}, nil
}
func (s *sApiShop) CartUpdate(ctx context.Context, req *v1.CartUpdateReq) (*v1.CartUpdateRes, error) {
	if err := g.Validator().Data(req).Run(ctx); err != nil {
		return nil, err
	}
	err := cartWrite(ctx, func(tx gdb.TX, userID int64) error {
		cc := dao.ShopCart.Columns()
		var cart *entity.ShopCart
		if err := tx.Model(dao.ShopCart.Table()).Where(cc.UserId, userID).Where(cc.Id, req.Id).LockUpdate().Scan(&cart); err != nil {
			return err
		}
		if cart == nil {
			return gerror.New("购物车商品不存在")
		}
		_, err := tx.Model(dao.ShopCart.Table()).Where(cc.UserId, userID).Where(cc.Id, req.Id).Data(g.Map{cc.Quantity: req.Quantity, cc.UpdatedAt: uint64(gtime.Now().Unix())}).Update()
		return err
	})
	if err != nil {
		return nil, err
	}
	return &v1.CartUpdateRes{}, nil
}
func (s *sApiShop) CartDelete(ctx context.Context, req *v1.CartDeleteReq) (*v1.CartDeleteRes, error) {
	if err := g.Validator().Data(req).Run(ctx); err != nil {
		return nil, err
	}
	err := cartWrite(ctx, func(tx gdb.TX, userID int64) error {
		cc := dao.ShopCart.Columns()
		_, err := tx.Model(dao.ShopCart.Table()).Where(cc.UserId, userID).Where(cc.Id, req.Id).Delete()
		return err
	})
	if err != nil {
		return nil, err
	}
	return &v1.CartDeleteRes{}, nil
}
func (s *sApiShop) CartList(ctx context.Context, req *v1.CartListReq) (*v1.CartListRes, error) {
	userID, err := getShopUserId(ctx)
	if err != nil {
		return nil, err
	}
	var user *entity.User
	if err := dao.User.Ctx(ctx).WherePri(userID).Scan(&user); err != nil {
		return nil, gerror.Wrap(err, "读取账户失败")
	}
	if user == nil || (user.Status != "enable" && user.Status != "1") {
		return nil, gerror.New("账户不存在或已停用")
	}
	var carts []*entity.ShopCart
	cc := dao.ShopCart.Columns()
	if err := dao.ShopCart.Ctx(ctx).Where(cc.UserId, userID).OrderAsc(cc.Id).Scan(&carts); err != nil {
		return nil, gerror.Wrap(err, "读取购物车失败")
	}
	res := &v1.CartListRes{Items: make([]*v1.CartItem, 0, len(carts)), Balance: user.Money}
	var total int64
	for _, cart := range carts {
		var product *entity.ShopProduct
		if err := dao.ShopProduct.Ctx(ctx).WherePri(cart.ProductId).Where(dao.ShopProduct.Columns().Status, "normal").Scan(&product); err != nil {
			return nil, gerror.Wrap(err, "读取商品失败")
		}
		if cart.Quantity == 0 || cart.Quantity > 999 {
			return nil, gerror.New("购物车数量异常，请修改或删除商品")
		}
		items, subtotal, err := priceOrder([]uint64{cart.ProductId}, map[uint64]uint{cart.ProductId: cart.Quantity}, map[uint64]*entity.ShopProduct{cart.ProductId: product}, 0)
		if err != nil {
			return nil, err
		}
		if total > maxShopCents-subtotal {
			return nil, gerror.New("购物车总金额超出范围")
		}
		total += subtotal
		item := items[0]
		res.Items = append(res.Items, &v1.CartItem{CartId: cart.Id, ProductId: cart.ProductId, Title: item.Title, Image: item.Image, UnitPrice: item.UnitPrice, Quantity: cart.Quantity, TotalAmount: item.TotalAmount})
	}
	res.TotalAmount = centsToAmount(total)
	return res, nil
}
