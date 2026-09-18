-- hotgo自动生成菜单权限SQL 通常情况下只在首次生成代码时自动执行一次
-- 如需再次执行请先手动删除生成的菜单权限和SQL文件：P:\go\hotgo-2.0\server\storage\data\generate\api_model_menu.sql
-- Version: 2.18.9
-- Date: 2026-06-13 16:16:59
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
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, '2461', 'AI模型列表', 'apiModel', '/apiModel', 'MenuOutlined', '1', '/openai/apiModel/index', '', '', 'ParentLayout', '1', '', '0', '0', '', '0', '0', '0', '2', 'tr_2461 ', '0', '', '1', @now, @now);


SET @dirId = LAST_INSERT_ID();


-- 菜单页面
-- 列表
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @dirId, 'AI模型列表列表', 'apiModelIndex', 'index', '', '2', '', '/apiModel/list', '', '/apiModel/index', '1', 'apiModel', '0', '0', '', '0', '1', '0', '3', CONCAT('tr_2461 tr_', @dirId,' '), '10', '', '1', @now, @now);


SET @listId = LAST_INSERT_ID();

-- 详情
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, 'AI模型列表详情', 'apiModelView', '', '', '3', '', '/apiModel/view', '', '', '1', '', '0', '0', '', '0', '1', '0', '4', CONCAT('tr_2461 tr_', @dirId, ' tr_', @listId,' '), '10', '', '1', @now, @now);


-- 菜单按钮

-- 编辑
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '编辑/新增AI模型列表', 'apiModelEdit', '', '', '3', '', '/apiModel/edit', '', '', '1', '', '0', '0', '', '0', '1', '0', '4', CONCAT('tr_2461 tr_', @dirId, ' tr_', @listId,' '), '20', '', '1', @now, @now);


SET @editId = LAST_INSERT_ID();

-- 获取最大排序
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @editId, '获取AI模型列表最大排序', 'apiModelMaxSort', '', '', '3', '', '/apiModel/maxSort', '', '', '1', '', '0', '0', '', '0', '0', '0', '5', CONCAT('tr_2461 tr_', @dirId, ' tr_', @listId, ' tr_', @editId,' '), '30', '', '1', @now, @now);


-- 删除
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '删除AI模型列表', 'apiModelDelete', '', '', '3', '', '/apiModel/delete', '', '', '1', '', '0', '0', '', '0', '0', '0', '4', CONCAT('tr_2461 tr_', @dirId, ' tr_', @listId,' '), '40', '', '1', @now, @now);


-- 更新状态
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '修改AI模型列表状态', 'apiModelStatus', '', '', '3', '', '/apiModel/status', '', '', '1', '', '0', '0', '', '0', '0', '0', '4', CONCAT('tr_2461 tr_', @dirId, ' tr_', @listId,' '), '50', '', '1', @now, @now);



-- 导出
INSERT INTO `hg_admin_menu` (`id`, `pid`, `title`, `name`, `path`, `icon`, `type`, `redirect`, `permissions`, `permission_name`, `component`, `always_show`, `active_menu`, `is_root`, `is_frame`, `frame_src`, `keep_alive`, `hidden`, `affix`, `level`, `tree`, `sort`, `remark`, `status`, `created_at`, `updated_at`) VALUES (NULL, @listId, '导出AI模型列表', 'apiModelExport', '', '', '3', '', '/apiModel/export', '', '', '1', '', '0', '0', '', '0', '0', '0', '4', CONCAT('tr_2461 tr_', @dirId, ' tr_', @listId,' '), '70', '', '1', @now, @now);


COMMIT;