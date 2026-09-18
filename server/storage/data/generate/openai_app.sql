CREATE TABLE `hg_openai_app_category` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `name` varchar(64) NOT NULL COMMENT '分类名称',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态(1:启用, 2:禁用)',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='AI应用分类';

CREATE TABLE `hg_openai_app` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '应用ID',
  `category_id` bigint(20) NOT NULL COMMENT '分类ID',
  `title` varchar(128) NOT NULL COMMENT '应用名称',
  `icon` varchar(255) DEFAULT NULL COMMENT '应用图标',
  `description` varchar(512) DEFAULT NULL COMMENT '应用描述',
  `prompt` text NOT NULL COMMENT '系统提示词',
  `use_count` int(11) NOT NULL DEFAULT '0' COMMENT '使用次数',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态(1:启用, 2:禁用)',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='AI应用';