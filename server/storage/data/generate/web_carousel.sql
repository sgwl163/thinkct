-- 网站_首页轮播 hg_web_carousel
-- 可重复执行: 表不存在时创建, 仅在表完全为空(含软删除行)时写入初始轮播数据, 不覆盖已有业务数据
-- 初始数据来源 next_thinkct/components/templates/vcdon/home/Carousel.tsx 的 slides 静态数据, 共3条, 字段值保持原样
-- 字段契约: ctaText->cta_text, ctaLink->cta_link, aspectRatio->aspect_ratio, 其余同名
CREATE TABLE IF NOT EXISTS `hg_web_carousel` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '轮播ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `description` text COMMENT '描述',
  `cta_text` varchar(100) NOT NULL DEFAULT '' COMMENT '按钮文字',
  `cta_link` varchar(512) NOT NULL DEFAULT '' COMMENT '按钮链接',
  `color` varchar(50) NOT NULL DEFAULT 'emerald-blue' COMMENT '配色主题',
  `image` varchar(512) NOT NULL DEFAULT '' COMMENT '图片',
  `alt` varchar(255) NOT NULL DEFAULT '' COMMENT '图片替代文本',
  `aspect_ratio` varchar(32) NOT NULL DEFAULT '16/9' COMMENT '图片宽高比',
  `sort` int NOT NULL DEFAULT 0 COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_status_sort` (`status`,`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='网站_首页轮播';

INSERT INTO `hg_web_carousel` (`id`, `title`, `description`, `cta_text`, `cta_link`, `color`, `image`, `alt`, `aspect_ratio`, `sort`, `status`, `created_at`, `updated_at`)
SELECT s.`id`, s.`title`, s.`description`, s.`cta_text`, s.`cta_link`, s.`color`, s.`image`, s.`alt`, s.`aspect_ratio`, s.`sort`, 1, NOW(), NOW()
FROM (
  SELECT 1 AS `id`, '追求无限可能' AS `title`, '通过我们的数据驱动解决方案实现生产率和可持续发展目标。' AS `description`, '查看我们的实现方式' AS `cta_text`, '/products' AS `cta_link`, 'emerald-blue' AS `color`, '/images/vcdon/slides1.jpg' AS `image`, '服务器机房中手持 iPad 的女性工作人员' AS `alt`, '16/9' AS `aspect_ratio`, 1 AS `sort`
  UNION ALL SELECT 2, '第十版年度智能制造现状', '为您的数字战略制定基准，探索全球杰出范例，并揭示 2025 年的投资趋势。', '获取报告', '/news_hy/detail/1', 'purple-cyan', '/images/vcdon/slides2.jpg', '霓虹蓝绿波', '3000/884', 2
  UNION ALL SELECT 3, '探索未来边界', '了解我们如何帮助客户、企业和社区变得更具有弹性、敏捷性和可持续性。', '查看我们的 2025 年报告', '/industry-trends/detail/7', 'orange-pink', '/images/vcdon/slides3.jpg', '2025 年可持续发展报告封面', '16/9', 3
) AS s
WHERE NOT EXISTS (SELECT 1 FROM `hg_web_carousel` AS e)
ORDER BY s.`id`;