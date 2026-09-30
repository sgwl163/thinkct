CREATE TABLE IF NOT EXISTS `hg_user_credits_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '流水ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `credit_type` varchar(16) NOT NULL COMMENT '资产类型:money=余额,score=积分',
  `before_num` bigint(20) NOT NULL DEFAULT '0' COMMENT '变动前数量',
  `num` bigint(20) NOT NULL COMMENT '变动数量',
  `after_num` bigint(20) NOT NULL DEFAULT '0' COMMENT '变动后数量',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `map_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '关联ID',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_type_created_at` (`user_id`, `credit_type`, `created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户_资产变动流水';

CREATE TABLE IF NOT EXISTS `hg_user_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '日志ID',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户ID',
  `username` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `log_type` varchar(32) NOT NULL DEFAULT '' COMMENT '记录类型:login=登录,operation=操作',
  `action` varchar(64) NOT NULL DEFAULT '' COMMENT '操作标识',
  `action_name` varchar(128) NOT NULL DEFAULT '' COMMENT '操作名称',
  `event_at` datetime DEFAULT NULL COMMENT '事件时间',
  `ip` varchar(128) NOT NULL DEFAULT '' COMMENT 'IP地址',
  `user_agent` varchar(512) NOT NULL DEFAULT '' COMMENT 'UA信息',
  `message` varchar(1000) NOT NULL DEFAULT '' COMMENT '结果消息',
  `extra` text COMMENT '扩展数据',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_type_event_at` (`user_id`, `log_type`, `event_at`),
  KEY `idx_action_event_at` (`action`, `event_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户_统一行为记录';
