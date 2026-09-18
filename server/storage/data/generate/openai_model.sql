CREATE TABLE IF NOT EXISTS `hg_openai_api_group` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) NOT NULL COMMENT '组名称(如OpenAI、文心一言)',
  `base_url` varchar(255) DEFAULT '' COMMENT '接口地址',
  `api_key` varchar(255) DEFAULT '' COMMENT 'API Key',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='AI接口分组';

CREATE TABLE IF NOT EXISTS `hg_openai_api_model` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `group_id` bigint(20) NOT NULL COMMENT '分组ID',
  `name` varchar(255) NOT NULL COMMENT '模型别名(如GPT-4)',
  `model` varchar(255) NOT NULL COMMENT '模型标识(如gpt-4-turbo)',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态',
  `sort` int(11) DEFAULT '0' COMMENT '排序',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '修改时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_group_id` (`group_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='AI模型列表';

-- 插入测试数据
INSERT INTO `hg_openai_api_group` (`id`, `name`, `base_url`, `api_key`, `created_at`, `updated_at`) VALUES 
(1, '默认OpenAI组', 'https://api.openai.com/v1', 'sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx', NOW(), NOW());

INSERT INTO `hg_openai_api_model` (`group_id`, `name`, `model`, `created_at`, `updated_at`) VALUES 
(1, 'GPT-3.5 Turbo', 'gpt-3.5-turbo', NOW(), NOW()), 
(1, 'GPT-4 Vision', 'gpt-4-vision-preview', NOW(), NOW()),
(1, 'GPT-4 Turbo', 'gpt-4-turbo', NOW(), NOW()),
(1, 'GPT-4o', 'gpt-4o', NOW(), NOW());
