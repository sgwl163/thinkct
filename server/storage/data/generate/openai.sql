CREATE TABLE IF NOT EXISTS `hg_openai_chat_folder` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `name` varchar(255) NOT NULL COMMENT '文件夹名称',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='AI聊天文件夹';

CREATE TABLE IF NOT EXISTS `hg_openai_chat_session` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `folder_id` bigint(20) DEFAULT '0' COMMENT '文件夹ID',
  `title` varchar(255) NOT NULL COMMENT '会话标题',
  `prompt` varchar(1000) DEFAULT '' COMMENT '系统提示词',
  `model` varchar(255) DEFAULT '' COMMENT '使用的模型',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_folder_id` (`folder_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='AI聊天会话';

CREATE TABLE IF NOT EXISTS `hg_openai_chat_message` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `user_id` bigint(20) NOT NULL COMMENT '用户ID',
  `session_id` bigint(20) NOT NULL COMMENT '会话ID',
  `role` varchar(50) NOT NULL COMMENT '角色(user/assistant/system)',
  `content` longtext NOT NULL COMMENT '消息内容',
  `model` varchar(255) DEFAULT '' COMMENT '使用的模型',
  `req_tokens` int(11) DEFAULT '0' COMMENT '请求token数',
  `res_tokens` int(11) DEFAULT '0' COMMENT '响应token数',
  `total_tokens` int(11) DEFAULT '0' COMMENT '总token数',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_session_id` (`session_id`),
  KEY `idx_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='AI聊天消息';
