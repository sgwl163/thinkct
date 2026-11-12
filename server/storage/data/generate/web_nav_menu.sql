-- 网站_导航菜单 hg_web_nav_menu
-- 可重复执行: 表不存在时创建, 仅在表完全为空(含软删除行)时写入初始导航数据, 不覆盖已有业务数据
-- 初始数据来源 next_thinkct/config/NavMenu.ts 的全部有效节点(已排除被注释节点), 共100条
-- 原始重复ID已按前序遍历顺序重新分配为 1..100, pid/level/tree 依新ID重建, sort 为同级原始顺序, URL 保持原样
CREATE TABLE IF NOT EXISTS `hg_web_nav_menu` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '导航ID',
  `pid` bigint NOT NULL DEFAULT 0 COMMENT '上级ID',
  `level` int NOT NULL DEFAULT 1 COMMENT '关系树级别',
  `tree` varchar(512) NOT NULL DEFAULT '' COMMENT '关系树',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '导航名称',
  `url` varchar(512) NOT NULL DEFAULT '' COMMENT '链接地址',
  `image` varchar(512) NOT NULL DEFAULT '' COMMENT '图片',
  `image_alt` varchar(255) NOT NULL DEFAULT '' COMMENT '图片替代文本',
  `sort` int NOT NULL DEFAULT 0 COMMENT '排序',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  KEY `idx_pid_sort` (`pid`,`sort`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='网站_导航菜单';

INSERT INTO `hg_web_nav_menu` (`id`, `pid`, `level`, `tree`, `name`, `url`, `image`, `image_alt`, `sort`, `status`, `created_at`, `updated_at`)
SELECT s.`id`, s.`pid`, s.`level`, s.`tree`, s.`name`, s.`url`, s.`image`, s.`image_alt`, s.`sort`, 1, NOW(), NOW()
FROM (
  SELECT 1 AS `id`, 0 AS `pid`, 1 AS `level`, '' AS `tree`, '产品' AS `name`, '/products' AS `url`, '' AS `image`, '' AS `image_alt`, 1 AS `sort`
  UNION ALL SELECT 2, 1, 2, 'tr_1 ', '硬件', '/products-overview', '', '', 1
  UNION ALL SELECT 3, 2, 3, 'tr_1 tr_2 ', '硬件新动向', '/products-trends', '', '', 1
  UNION ALL SELECT 4, 2, 3, 'tr_1 tr_2 ', '电路和负载保护', '/products-circuit-protection', '', '', 2
  UNION ALL SELECT 5, 2, 3, 'tr_1 tr_2 ', '状态监测', '/products-condition-monitoring', '', '', 3
  UNION ALL SELECT 6, 2, 3, 'tr_1 tr_2 ', '连接设备', '/products-connection-devices', '', '', 4
  UNION ALL SELECT 7, 2, 3, 'tr_1 tr_2 ', '分布式控制系统', '/products-dcs', '', '', 5
  UNION ALL SELECT 8, 2, 3, 'tr_1 tr_2 ', '变频器', '/products', '', '', 6
  UNION ALL SELECT 9, 2, 3, 'tr_1 tr_2 ', '能源监测', '/products', '', '', 7
  UNION ALL SELECT 10, 2, 3, 'tr_1 tr_2 ', '人机界面 (HMI)', '/products', '', '', 8
  UNION ALL SELECT 11, 2, 3, 'tr_1 tr_2 ', '独立推车技术', '/products', '', '', 9
  UNION ALL SELECT 12, 2, 3, 'tr_1 tr_2 ', '工业计算机和监视器', '/products', '', '', 10
  UNION ALL SELECT 13, 2, 3, 'tr_1 tr_2 ', '工业控制产品', '/products', '', '', 11
  UNION ALL SELECT 14, 2, 3, 'tr_1 tr_2 ', '输入/输出模块', '/products', '', '', 12
  UNION ALL SELECT 15, 2, 3, 'tr_1 tr_2 ', '照明控制', '/products', '', '', 13
  UNION ALL SELECT 16, 2, 3, 'tr_1 tr_2 ', '运动控制', '/products', '', '', 14
  UNION ALL SELECT 17, 2, 3, 'tr_1 tr_2 ', '马达控制', '/products', '', '', 15
  UNION ALL SELECT 18, 2, 3, 'tr_1 tr_2 ', '网络安全与基础架构', '/products', '', '', 16
  UNION ALL SELECT 19, 2, 3, 'tr_1 tr_2 ', '成套解决方案', '/products', '', '', 17
  UNION ALL SELECT 20, 2, 3, 'tr_1 tr_2 ', '电源', '/products', '', '', 18
  UNION ALL SELECT 21, 2, 3, 'tr_1 tr_2 ', '照明控制', '/products', '', '', 19
  UNION ALL SELECT 22, 2, 3, 'tr_1 tr_2 ', '可编程控制器', '/products', '', '', 20
  UNION ALL SELECT 23, 2, 3, 'tr_1 tr_2 ', '照明控制', '/products', '', '', 21
  UNION ALL SELECT 24, 2, 3, 'tr_1 tr_2 ', '按钮和信号设备', '/products', '', '', 22
  UNION ALL SELECT 25, 2, 3, 'tr_1 tr_2 ', '继电器和计时器', '/products', '', '', 23
  UNION ALL SELECT 26, 2, 3, 'tr_1 tr_2 ', '安全仪表系统', '/products', '', '', 24
  UNION ALL SELECT 27, 2, 3, 'tr_1 tr_2 ', '安全产品', '/products', '', '', 25
  UNION ALL SELECT 28, 2, 3, 'tr_1 tr_2 ', '传感器和开关', '/products', '', '', 26
  UNION ALL SELECT 29, 2, 3, 'tr_1 tr_2 ', '信号接口', '/products', '', '', 27
  UNION ALL SELECT 30, 1, 2, 'tr_1 ', '机房与控制系统', '/products/control-systems', '', '', 2
  UNION ALL SELECT 31, 30, 3, 'tr_1 tr_30 ', '软件新动向', '/products/', '', '', 1
  UNION ALL SELECT 32, 30, 3, 'tr_1 tr_30 ', '云软件', '/products/', '', '', 2
  UNION ALL SELECT 33, 30, 3, 'tr_1 tr_30 ', '软件订阅', '/products/', '', '', 3
  UNION ALL SELECT 34, 30, 3, 'tr_1 tr_30 ', '关键电源管理', '/products', '', '', 4
  UNION ALL SELECT 35, 34, 4, 'tr_1 tr_30 tr_34 ', '交流不间断电源（UPS）', '/ups-power-monitor', '', '', 1
  UNION ALL SELECT 36, 34, 4, 'tr_1 tr_30 tr_34 ', '直流电源系统', '/dc-power-system-monitor', '', '', 2
  UNION ALL SELECT 37, 34, 4, 'tr_1 tr_30 tr_34 ', 'EPS电源', '/eps-power-monitor', '', '', 3
  UNION ALL SELECT 38, 34, 4, 'tr_1 tr_30 tr_34 ', '稳压电源', '/industrial-ups-monitor', '', '', 4
  UNION ALL SELECT 39, 34, 4, 'tr_1 tr_30 tr_34 ', '通信电源', '/communication-power-monitor', '', '', 5
  UNION ALL SELECT 40, 34, 4, 'tr_1 tr_30 tr_34 ', '蓄电池', '/battery-monitor', '', '', 6
  UNION ALL SELECT 41, 30, 3, 'tr_1 tr_30 ', '配电', '/products/', '', '', 5
  UNION ALL SELECT 42, 41, 4, 'tr_1 tr_30 tr_41 ', '智能配电柜', '/smart-power-distribution-monitor', '', '', 1
  UNION ALL SELECT 43, 41, 4, 'tr_1 tr_30 tr_41 ', '列头柜', '/：column-head-cabinet', '', '', 2
  UNION ALL SELECT 44, 41, 4, 'tr_1 tr_30 tr_41 ', 'PDU', '/pdu-monitor', '', '', 3
  UNION ALL SELECT 45, 30, 3, 'tr_1 tr_30 ', '监控', '/products', '', '', 6
  UNION ALL SELECT 46, 45, 4, 'tr_1 tr_30 tr_45 ', 'UPS监控', '/ups-monitor', '', '', 1
  UNION ALL SELECT 47, 45, 4, 'tr_1 tr_30 tr_45 ', '机房监控', '/server-room-monitor', '', '', 2
  UNION ALL SELECT 48, 0, 1, '', '服务', '/services', '//vcdon.oss-cn-beijing.aliyuncs.com/images/services-featured.jpg', '', 2
  UNION ALL SELECT 49, 48, 2, 'tr_48 ', '服务内容', '/services-details', '', '', 1
  UNION ALL SELECT 50, 49, 3, 'tr_48 tr_49 ', '咨询服务', '/services-details', '', '', 1
  UNION ALL SELECT 51, 49, 3, 'tr_48 tr_49 ', '维护与支持', '/services-details', '', '', 2
  UNION ALL SELECT 52, 49, 3, 'tr_48 tr_49 ', '设备升级', '/services-details', '', '', 3
  UNION ALL SELECT 53, 48, 2, 'tr_48 ', '生命周期服务', '/services-lifecycle', '', '', 2
  UNION ALL SELECT 54, 53, 3, 'tr_48 tr_53 ', '安装与调试', '/services-lifecycle', '', '', 1
  UNION ALL SELECT 55, 53, 3, 'tr_48 tr_53 ', '退役与回收', '/services-lifecycle', '', '', 2
  UNION ALL SELECT 56, 48, 2, 'tr_48 ', '定制化服务', '/custom-development', '', '', 3
  UNION ALL SELECT 57, 56, 3, 'tr_48 tr_56 ', '定制开发', '/custom-development', '', '', 1
  UNION ALL SELECT 58, 56, 3, 'tr_48 tr_56 ', '系统集成', '/custom-integration', '', '', 2
  UNION ALL SELECT 59, 0, 1, '', '解决方案与行业', '/solutions', '//vcdon.oss-cn-beijing.aliyuncs.com/images/solutions-featured.jpg', '', 3
  UNION ALL SELECT 60, 59, 2, 'tr_59 ', '行业解决方案', '/solutions', '', '', 1
  UNION ALL SELECT 61, 60, 3, 'tr_59 tr_60 ', '交通行业', '/solutions-transportation', '', '', 1
  UNION ALL SELECT 62, 60, 3, 'tr_59 tr_60 ', '教育行业', '/solutions-education', '', '', 2
  UNION ALL SELECT 63, 60, 3, 'tr_59 tr_60 ', '医疗行业', '/solutions-healthcare', '', '', 3
  UNION ALL SELECT 64, 60, 3, 'tr_59 tr_60 ', '制造行业', '/solutions-manufacturing', '', '', 4
  UNION ALL SELECT 65, 60, 3, 'tr_59 tr_60 ', '地产行业', '/solutions-real-estate', '', '', 5
  UNION ALL SELECT 66, 59, 2, 'tr_59 ', '技术方案', '/solutions-technologies', '', '', 2
  UNION ALL SELECT 67, 66, 3, 'tr_59 tr_66 ', '工业物联网', '/solutions-technologies-iiot', '', '', 1
  UNION ALL SELECT 68, 66, 3, 'tr_59 tr_66 ', '数字化转型', '/solutions-digital-transformation', '', '', 2
  UNION ALL SELECT 69, 66, 3, 'tr_59 tr_66 ', '人工智能', '/solutions-technologies-ai', '', '', 3
  UNION ALL SELECT 70, 59, 2, 'tr_59 ', '智能制造', '/solutions-smart-manufacturing', '', '', 3
  UNION ALL SELECT 71, 70, 3, 'tr_59 tr_70 ', '智能工厂', '/solutions-smart-factory', '', '', 1
  UNION ALL SELECT 72, 70, 3, 'tr_59 tr_70 ', '数据分析', '/solutions-smart-analytics', '', '', 2
  UNION ALL SELECT 73, 0, 1, '', '支持', '/support-center', '//vcdon.oss-cn-beijing.aliyuncs.com/images/support-featured.jpg', '', 4
  UNION ALL SELECT 74, 73, 2, 'tr_73 ', '技术支持', '/technical-support', '', '', 1
  UNION ALL SELECT 75, 74, 3, 'tr_73 tr_74 ', '在线支持', '/technical-support', '', '', 1
  UNION ALL SELECT 76, 74, 3, 'tr_73 tr_74 ', '电话支持', '/technical-support', '', '', 2
  UNION ALL SELECT 77, 74, 3, 'tr_73 tr_74 ', '远程诊断', '/technical-support', '', '', 3
  UNION ALL SELECT 78, 73, 2, 'tr_73 ', '培训与资源', '/training-resources', '', '', 2
  UNION ALL SELECT 79, 78, 3, 'tr_73 tr_78 ', '培训课程', '/training-resources', '', '', 1
  UNION ALL SELECT 80, 78, 3, 'tr_73 tr_78 ', '技术文档', '/training-resources', '', '', 2
  UNION ALL SELECT 81, 78, 3, 'tr_73 tr_78 ', '案例研究', '/training-resources', '', '', 3
  UNION ALL SELECT 82, 78, 3, 'tr_73 tr_78 ', '下载中心', '/support/downloads', '', '', 4
  UNION ALL SELECT 83, 73, 2, 'tr_73 ', '客户支持', '/customer-support', '', '', 3
  UNION ALL SELECT 84, 83, 3, 'tr_73 tr_83 ', '故障排除', '/customer-support', '', '', 1
  UNION ALL SELECT 85, 83, 3, 'tr_73 tr_83 ', '客户反馈', '/customer-support', '', '', 2
  UNION ALL SELECT 86, 83, 3, 'tr_73 tr_83 ', '多元化、道德与完整性', '/company/diversity', '', '', 3
  UNION ALL SELECT 87, 83, 3, 'tr_73 tr_83 ', '关于我们', '/company/about', '', '', 4
  UNION ALL SELECT 88, 83, 3, 'tr_73 tr_83 ', '联系我们', '/company/contact', '', '', 5
  UNION ALL SELECT 89, 0, 1, '', '新闻中心', '/news', '//vcdon.oss-cn-beijing.aliyuncs.com/images/news-featured.jpg', '', 5
  UNION ALL SELECT 90, 89, 2, 'tr_89 ', '公司新闻', '/news', '', '', 1
  UNION ALL SELECT 91, 90, 3, 'tr_89 tr_90 ', '最新动态', '/news', '', '', 1
  UNION ALL SELECT 92, 90, 3, 'tr_89 tr_90 ', '行业动态', '/industry-trends', '', '', 2
  UNION ALL SELECT 93, 90, 3, 'tr_89 tr_90 ', '企业公告', '/company-announcements', '', '', 3
  UNION ALL SELECT 94, 89, 2, 'tr_89 ', '行业动态', '/news', '', '', 2
  UNION ALL SELECT 95, 94, 3, 'tr_89 tr_94 ', '市场趋势', '/market-trends', '', '', 1
  UNION ALL SELECT 96, 94, 3, 'tr_89 tr_94 ', '技术洞察', '/market-trends', '', '', 2
  UNION ALL SELECT 97, 94, 3, 'tr_89 tr_94 ', '行业分析', '/market-trends', '', '', 3
  UNION ALL SELECT 98, 89, 2, 'tr_89 ', '活动与展览', '/events-exhibitions', '', '', 3
  UNION ALL SELECT 99, 98, 3, 'tr_89 tr_98 ', '展会报道', '/exhibition-report', '', '', 1
  UNION ALL SELECT 100, 98, 3, 'tr_89 tr_98 ', '网络研讨会', '/exhibition-report', '', '', 2
) AS s
WHERE NOT EXISTS (SELECT 1 FROM `hg_web_nav_menu` AS e)
ORDER BY s.`id`;