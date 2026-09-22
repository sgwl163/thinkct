CREATE TABLE IF NOT EXISTS `hg_shop_order` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '订单ID',
    `order_sn` varchar(64) NOT NULL COMMENT '订单号',
    `user_id` bigint unsigned NOT NULL COMMENT '下单用户ID',
    `status` varchar(32) NOT NULL DEFAULT 'pending_payment' COMMENT '订单状态:pending_payment=待支付,paid=已支付,completed=已完成,cancelled=已取消,closed=已关闭',
    `total_amount` decimal(12,2) NOT NULL DEFAULT '0.00' COMMENT '商品总金额',
    `discount_amount` decimal(12,2) NOT NULL DEFAULT '0.00' COMMENT '优惠金额',
    `freight_amount` decimal(12,2) NOT NULL DEFAULT '0.00' COMMENT '运费金额',
    `payable_amount` decimal(12,2) NOT NULL DEFAULT '0.00' COMMENT '应付金额',
    `paid_amount` decimal(12,2) NOT NULL DEFAULT '0.00' COMMENT '实付金额',
    `remark` varchar(500) NOT NULL DEFAULT '' COMMENT '用户备注',
    `payment_type` varchar(32) NOT NULL DEFAULT '' COMMENT '支付方式',
    `payment_no` varchar(128) NOT NULL DEFAULT '' COMMENT '第三方支付单号',
    `payment_time` bigint unsigned NOT NULL DEFAULT '0' COMMENT '支付时间',
    `completed_time` bigint unsigned NOT NULL DEFAULT '0' COMMENT '完成时间',
    `cancelled_time` bigint unsigned NOT NULL DEFAULT '0' COMMENT '取消时间',
    `update_time` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    `create_time` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_order_sn` (`order_sn`),
    KEY `idx_user_create_time` (`user_id`, `create_time`),
    KEY `idx_status_create_time` (`status`, `create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城订单主表';

CREATE TABLE IF NOT EXISTS `hg_shop_order_item` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '订单明细ID',
    `order_id` bigint unsigned NOT NULL COMMENT '订单ID',
    `product_id` bigint unsigned NOT NULL COMMENT '商品ID',
    `title` varchar(255) NOT NULL DEFAULT '' COMMENT '商品名称快照',
    `image` varchar(255) NOT NULL DEFAULT '' COMMENT '商品图片快照',
    `unit_price` decimal(12,2) NOT NULL DEFAULT '0.00' COMMENT '成交单价',
    `quantity` int unsigned NOT NULL DEFAULT '1' COMMENT '购买数量',
    `total_amount` decimal(12,2) NOT NULL DEFAULT '0.00' COMMENT '商品小计',
    `create_time` bigint unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
    `update_time` bigint unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
    PRIMARY KEY (`id`),
    KEY `idx_order_id` (`order_id`),
    KEY `idx_product_id` (`product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城订单商品明细表';
