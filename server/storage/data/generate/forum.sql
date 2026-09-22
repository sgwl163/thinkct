CREATE TABLE IF NOT EXISTS `hg_forum_category` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '板块ID',
    `pid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '上级板块ID',
    `name` varchar(100) NOT NULL COMMENT '板块名称',
    `description` varchar(500) NOT NULL DEFAULT '' COMMENT '板块简介',
    `image` varchar(255) NOT NULL DEFAULT '' COMMENT '板块图标',
    `sort` int NOT NULL DEFAULT '0' COMMENT '排序',
    `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态:1正常,2禁用',
    `topic_count` int unsigned NOT NULL DEFAULT '0' COMMENT '主题数量',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY `idx_pid_sort` (`pid`, `sort`),
    KEY `idx_status_sort` (`status`, `sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='社区论坛板块表';

CREATE TABLE IF NOT EXISTS `hg_forum_topic` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主题ID',
    `category_id` bigint unsigned NOT NULL COMMENT '板块ID',
    `user_id` bigint unsigned NOT NULL COMMENT '发布用户ID',
    `title` varchar(255) NOT NULL COMMENT '主题标题',
    `keywords` varchar(500) NOT NULL DEFAULT '' COMMENT '主题关键词，英文逗号分隔',
    `description` varchar(500) NOT NULL DEFAULT '' COMMENT '主题摘要',
    `content` longtext NOT NULL COMMENT '主题内容',
    `images` text NOT NULL COMMENT '主题图片，英文逗号分隔',
    `status` varchar(32) NOT NULL DEFAULT 'normal' COMMENT '状态:normal,pending,rejected,hidden,closed',
    `is_top` tinyint NOT NULL DEFAULT '0' COMMENT '是否置顶:0否,1是',
    `is_essence` tinyint NOT NULL DEFAULT '0' COMMENT '是否精华:0否,1是',
    `comment_count` int unsigned NOT NULL DEFAULT '0' COMMENT '评论数量',
    `like_count` int unsigned NOT NULL DEFAULT '0' COMMENT '点赞数量',
    `favorite_count` int unsigned NOT NULL DEFAULT '0' COMMENT '收藏数量',
    `view_count` int unsigned NOT NULL DEFAULT '0' COMMENT '浏览数量',
    `last_reply_user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '最后回复用户ID',
    `last_reply_at` datetime DEFAULT NULL COMMENT '最后回复时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_category_status_top_created_at` (`category_id`, `status`, `is_top`, `created_at`),
    KEY `idx_user_created_at` (`user_id`, `created_at`),
    KEY `idx_status_last_reply_at` (`status`, `last_reply_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='社区论坛主题表';

CREATE TABLE IF NOT EXISTS `hg_forum_comment` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '评论ID',
    `topic_id` bigint unsigned NOT NULL COMMENT '主题ID',
    `user_id` bigint unsigned NOT NULL COMMENT '评论用户ID',
    `parent_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '父评论ID',
    `reply_user_id` bigint unsigned NOT NULL DEFAULT '0' COMMENT '被回复用户ID',
    `content` varchar(2000) NOT NULL COMMENT '评论内容',
    `status` varchar(32) NOT NULL DEFAULT 'normal' COMMENT '状态:normal,pending,rejected,hidden',
    `like_count` int unsigned NOT NULL DEFAULT '0' COMMENT '点赞数量',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY `idx_topic_status_created_at` (`topic_id`, `status`, `created_at`),
    KEY `idx_parent_created_at` (`parent_id`, `created_at`),
    KEY `idx_user_created_at` (`user_id`, `created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='社区论坛评论表';

CREATE TABLE IF NOT EXISTS `hg_forum_interaction` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '互动ID',
    `target_type` varchar(32) NOT NULL COMMENT '互动对象类型:topic,comment',
    `target_id` bigint unsigned NOT NULL COMMENT '互动对象ID',
    `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
    `action_type` varchar(32) NOT NULL COMMENT '互动类型:like,favorite',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_target_user_action` (`target_type`, `target_id`, `user_id`, `action_type`),
    KEY `idx_target_action` (`target_type`, `target_id`, `action_type`),
    KEY `idx_user_created_at` (`user_id`, `created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='社区论坛互动表';

CREATE TABLE IF NOT EXISTS `hg_forum_user_follow` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '关注ID',
    `user_id` bigint unsigned NOT NULL COMMENT '关注用户ID',
    `follow_user_id` bigint unsigned NOT NULL COMMENT '被关注用户ID',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '关注时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_user_follow_user` (`user_id`, `follow_user_id`),
    KEY `idx_follow_user_created_at` (`follow_user_id`, `created_at`),
    KEY `idx_user_created_at` (`user_id`, `created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='社区论坛用户关注表';

CREATE TABLE IF NOT EXISTS `hg_forum_report` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '举报ID',
    `target_type` varchar(32) NOT NULL COMMENT '举报对象类型:topic,comment',
    `target_id` bigint unsigned NOT NULL COMMENT '举报对象ID',
    `user_id` bigint unsigned NOT NULL COMMENT '举报用户ID',
    `reason` varchar(100) NOT NULL COMMENT '举报原因',
    `description` varchar(1000) NOT NULL DEFAULT '' COMMENT '补充说明',
    `status` varchar(32) NOT NULL DEFAULT 'pending' COMMENT '处理状态:pending,processed,rejected',
    `handled_by` bigint unsigned NOT NULL DEFAULT '0' COMMENT '处理管理员ID',
    `handled_at` datetime DEFAULT NULL COMMENT '处理时间',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    KEY `idx_target` (`target_type`, `target_id`),
    KEY `idx_user_created_at` (`user_id`, `created_at`),
    KEY `idx_status_created_at` (`status`, `created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='社区论坛举报表';
