# 前台商城订单接口文档

本文档说明商城订单接口。所有接口基础路径为 `/api`，且均需用户登录。请求头应携带：

```text
Authorization: Bearer <token>
```

所有接口返回统一结构：

```json
{"code":0,"message":"操作成功","timestamp":1234567890,"traceID":"xxx","data":{}}
```

## 数据表

- `hg_shop_order`：订单主表，保存订单号、用户、状态和金额。
- `hg_shop_order_item`：订单商品明细，保存商品名称、图片和成交价格快照。

订单创建时，商品名称、图片和价格均由服务端从 `hg_web_products` 读取，客户端提交的商品 ID 与数量不会影响订单金额。商品必须处于 `normal` 状态。

订单状态：

- `pending_payment`：待支付
- `paid`：已支付
- `completed`：已完成
- `cancelled`：已取消
- `closed`：已关闭

当前版本只创建待支付订单，不包含支付、发货、库存和售后处理。

## 创建订单

- 路径：`POST /api/shop/order/create`
- 鉴权：需要登录
- 内容类型：`application/json`

请求参数：

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| products | array | 是 | 商品列表，最多 50 种商品 |
| products[].productId | uint64 | 是 | 商品 ID |
| products[].quantity | uint | 是 | 购买数量，范围 1 至 999 |
| remark | string | 否 | 订单备注，最多 500 个字符 |

请求示例：

```json
{
  "products": [
    {"productId": 12, "quantity": 2},
    {"productId": 27, "quantity": 1}
  ],
  "remark": "请在工作日配送"
}
```

响应示例：

```json
{
  "code": 0,
  "message": "操作成功",
  "data": {
    "order": {
      "id": 1,
      "orderSn": "SHOP20260919193000ABCD1234",
      "userId": 8,
      "status": "pending_payment",
      "totalAmount": 299.00,
      "discountAmount": 0,
      "freightAmount": 0,
      "payableAmount": 299.00,
      "paidAmount": 0,
      "remark": "请在工作日配送",
      "createTime": 1789817400,
      "items": [
        {
          "id": 1,
          "orderId": 1,
          "productId": 12,
          "title": "商品名称快照",
          "image": "https://example.com/product.jpg",
          "unitPrice": 99.50,
          "quantity": 2,
          "totalAmount": 199.00
        }
      ]
    }
  }
}
```

## 获取我的订单列表

- 路径：`GET /api/shop/order/list`
- 鉴权：需要登录

请求参数：

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| page | int | 否 | 页码，默认 1 |
| pageSize | int | 否 | 每页数量，默认 10，最大 200 |
| status | string | 否 | 按订单状态筛选 |

响应中的 `list` 为订单列表，每个订单均包含 `items` 商品明细数组；只会返回当前登录用户的订单。

## 获取订单详情

- 路径：`GET /api/shop/order/detail`
- 鉴权：需要登录

请求参数：

| 字段 | 类型 | 必填 | 说明 |
| --- | --- | --- | --- |
| orderSn | string | 是 | 订单号 |

响应字段与创建订单返回的 `order` 一致。订单详情按当前登录用户过滤，其他用户无法通过订单号读取订单。
