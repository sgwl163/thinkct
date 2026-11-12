CREATE TABLE IF NOT EXISTS `hg_shop_cart` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '购物车ID',
  `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
  `product_id` bigint unsigned NOT NULL COMMENT '商品ID',
  `quantity` int unsigned NOT NULL COMMENT '购买数量',
  `created_at` bigint unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `updated_at` bigint unsigned NOT NULL DEFAULT 0 COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_user_product` (`user_id`,`product_id`),
  KEY `idx_user_id_id` (`user_id`,`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='商城购物车';