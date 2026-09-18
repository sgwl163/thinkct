-- hotgo自动生成菜单权限SQL 通常情况下只在首次生成代码时自动执行一次
-- 如需再次执行请先手动删除生成的菜单权限和SQL文件：P:\go\hotgo-2.0\server\storage\data\generate\group_menu.sql
-- Version: 2.18.9
-- Date: 2026-06-09 14:33:20
-- Link https://github.com/bufanyun/hotgo

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;

--
-- 数据库： `develop_hotgo`
--

-- --------------------------------------------------------

--
-- 插入表中的数据 `hg_admin_menu`
--


SET @now := now();


-- 菜单目录
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, '2446', 'IM-群组表', 'group', '/group', 'MenuOutlined', '1', '/im/group/index', '', '', 'ParentLayout', '1', '', '0', '0', '', '0', '0', '0', '2', 'tr_2446 ', '0', '', '1', @now, @now);


SET @dirId = LAST_INSERT_ID();


-- 菜单页面
-- 列表
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @dirId, 'IM-群组表列表', 'groupIndex', 'index', '', '2', '', '/group/list', '', '/group/index', '1', 'group', '0', '0', '', '0', '1', '0', '3', CONCAT('tr_2446 tr_', @dirId,' '), '10', '', '1', @now, @now);


SET @listId = LAST_INSERT_ID();

-- 详情
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, 'IM-群组表详情', 'groupView', '', '', '3', '', '/group/view', '', '', '1', '', '0', '0', '', '0', '1', '0', '4', CONCAT('tr_2446 tr_', @dirId, ' tr_', @listId,' '), '10', '', '1', @now, @now);


-- 菜单按钮

-- 编辑
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '编辑/新增IM-群组表', 'groupEdit', '', '', '3', '', '/group/edit', '', '', '1', '', '0', '0', '', '0', '1', '0', '4', CONCAT('tr_2446 tr_', @dirId, ' tr_', @listId,' '), '20', '', '1', @now, @now);


SET @editId = LAST_INSERT_ID();


-- 删除
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '删除IM-群组表', 'groupDelete', '', '', '3', '', '/group/delete', '', '', '1', '', '0', '0', '', '0', '0', '0', '4', CONCAT('tr_2446 tr_', @dirId, ' tr_', @listId,' '), '40', '', '1', @now, @now);


-- 更新状态
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '修改IM-群组表状态', 'groupStatus', '', '', '3', '', '/group/status', '', '', '1', '', '0', '0', '', '0', '0', '0', '4', CONCAT('tr_2446 tr_', @dirId, ' tr_', @listId,' '), '50', '', '1', @now, @now);



-- 导出
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '导出IM-群组表', 'groupExport', '', '', '3', '', '/group/export', '', '', '1', '', '0', '0', '', '0', '0', '0', '4', CONCAT('tr_2446 tr_', @dirId, ' tr_', @listId,' '), '70', '', '1', @now, @now);


COMMIT;