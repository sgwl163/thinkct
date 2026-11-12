CREATE TABLE IF NOT EXISTS `hg_user_recharge` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '充值订单ID',
  `user_id` bigint NOT NULL COMMENT '前台用户ID',
  `order_sn` varchar(64) NOT NULL COMMENT '充值订单号',
  `money` int unsigned NOT NULL COMMENT '充值金额，整数元，与会员余额同单位',
  `pay_type` varchar(32) NOT NULL COMMENT '支付渠道',
  `status` varchar(32) NOT NULL DEFAULT 'pending' COMMENT 'pending=待支付,completed=已到账',
  `pay_at` datetime DEFAULT NULL COMMENT '支付时间',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_order_sn` (`order_sn`),
  KEY `idx_user_created` (`user_id`,`created_at`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='前台会员充值订单';