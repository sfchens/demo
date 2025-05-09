/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50726 (5.7.26)
 Source Host           : localhost:3306
 Source Schema         : greasyx-admin

 Target Server Type    : MySQL
 Target Server Version : 50726 (5.7.26)
 File Encoding         : 65001

 Date: 09/05/2025 17:59:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 80 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '系统权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (71, 'p', '1', '/api/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (72, 'p', '1', '/api/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (73, 'p', '1', '/api/update/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (54, 'p', '1', '/auth/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (74, 'p', '1', '/dict/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (77, 'p', '1', '/dict/delete/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (75, 'p', '1', '/dict/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (76, 'p', '1', '/dict/update/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (61, 'p', '1', '/menu/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (64, 'p', '1', '/menu/delete/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (63, 'p', '1', '/menu/info/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (60, 'p', '1', '/menu/router', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (62, 'p', '1', '/menu/update/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (79, 'p', '1', '/record/delete/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (78, 'p', '1', '/record/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (65, 'p', '1', '/role/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (69, 'p', '1', '/role/assign/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (70, 'p', '1', '/role/delete/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (67, 'p', '1', '/role/info/:id', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (66, 'p', '1', '/role/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (68, 'p', '1', '/role/update/:id', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (56, 'p', '1', '/user/add', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (59, 'p', '1', '/user/delete/:id', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (55, 'p', '1', '/user/info', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (57, 'p', '1', '/user/list', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (58, 'p', '1', '/user/update/:id', 'PUT', '', '', '');

-- ----------------------------
-- Table structure for casbin_rules
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rules`;
CREATE TABLE `casbin_rules`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '名称',
  `ptype` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '策略p,角色G',
  `v0` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '角色或用户',
  `v1` bigint(20) NULL DEFAULT NULL COMMENT '用户或路由',
  `v2` bigint(20) NULL DEFAULT NULL COMMENT '请求方式',
  `v3` bigint(20) NULL DEFAULT NULL COMMENT '允许标识',
  `v4` bigint(20) NULL DEFAULT NULL COMMENT '请求方式',
  `v5` bigint(20) NULL DEFAULT NULL COMMENT '请求方式',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of casbin_rules
-- ----------------------------

-- ----------------------------
-- Table structure for jwt_black_lists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_black_lists`;
CREATE TABLE `jwt_black_lists`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT 'jwt',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jwt_black_lists
-- ----------------------------

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键字段',
  `parent_id` bigint(20) NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '接口管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (1, 0, 'Base组', '', '', NULL, '2025-05-08 17:15:49.000', NULL);
INSERT INTO `sys_apis` VALUES (2, 1, '用户登录', 'POST', '/auth/login', NULL, '2025-05-08 17:16:14.000', NULL);
INSERT INTO `sys_apis` VALUES (3, 1, '用户信息', 'GET', '/user/info', NULL, '2025-05-08 17:16:35.000', NULL);
INSERT INTO `sys_apis` VALUES (4, 0, '系统用户管理', '', '', NULL, '2025-05-08 17:17:29.000', NULL);
INSERT INTO `sys_apis` VALUES (5, 4, '新增管理员', 'POST', '/user/add', NULL, '2025-05-08 17:17:52.000', NULL);
INSERT INTO `sys_apis` VALUES (6, 4, '管理员列表', 'GET', '/user/list', NULL, '2025-05-08 17:18:18.000', NULL);
INSERT INTO `sys_apis` VALUES (7, 4, '修改管理员', 'PUT', '/user/update/:id', NULL, '2025-05-08 17:18:41.000', NULL);
INSERT INTO `sys_apis` VALUES (8, 4, '删除管理员', 'DELETE', '/user/delete/:id', NULL, '2025-05-08 17:19:05.000', NULL);
INSERT INTO `sys_apis` VALUES (9, 0, '系统菜单管理', '', '', NULL, '2025-05-08 17:19:31.000', NULL);
INSERT INTO `sys_apis` VALUES (10, 1, '当前用户可用的菜单', 'GET', '/menu/router', NULL, '2025-05-08 17:17:02.000', NULL);
INSERT INTO `sys_apis` VALUES (12, 9, '添加菜单', 'POST', '/menu/add', NULL, '2025-05-08 17:19:49.000', NULL);
INSERT INTO `sys_apis` VALUES (13, 9, '修改菜单', 'PUT', '/menu/update/:id', NULL, '2025-05-08 17:20:11.000', NULL);
INSERT INTO `sys_apis` VALUES (14, 9, '获取菜单数据', 'GET', '/menu/info/:id', NULL, '2025-05-08 17:20:31.000', NULL);
INSERT INTO `sys_apis` VALUES (15, 9, '删除菜单', 'DELETE', '/menu/delete/:id', NULL, '2025-05-08 17:20:50.000', NULL);
INSERT INTO `sys_apis` VALUES (16, 0, '系统角色管理', '', '', NULL, '2025-05-08 17:21:16.000', NULL);
INSERT INTO `sys_apis` VALUES (17, 16, '新增角色', 'POST', '/role/add', NULL, '2025-05-08 17:21:34.000', NULL);
INSERT INTO `sys_apis` VALUES (18, 16, '角色列表', 'GET', '/role/list', NULL, '2025-05-08 17:21:54.000', NULL);
INSERT INTO `sys_apis` VALUES (19, 16, '角色信息', 'GET', '/role/info/:id', NULL, '2025-05-08 17:22:11.000', NULL);
INSERT INTO `sys_apis` VALUES (20, 16, '角色修改', 'PUT', '/role/update/:id', NULL, '2025-05-08 17:22:37.000', NULL);
INSERT INTO `sys_apis` VALUES (21, 16, '给角色分配权限', 'PUT', '/role/assign/:id', NULL, '2025-05-08 17:23:02.000', NULL);
INSERT INTO `sys_apis` VALUES (22, 16, '删除角色', 'DELETE', '/role/delete/:id', NULL, '2025-05-08 17:23:23.000', NULL);
INSERT INTO `sys_apis` VALUES (23, 0, '系统API管理', '', '', NULL, '2025-05-08 17:23:46.000', NULL);
INSERT INTO `sys_apis` VALUES (24, 23, '新增接口', 'POST', '/api/add', NULL, '2025-05-08 17:24:07.000', NULL);
INSERT INTO `sys_apis` VALUES (25, 23, '接口列表', 'GET', '/api/list', NULL, '2025-05-08 17:24:24.000', NULL);
INSERT INTO `sys_apis` VALUES (26, 23, '修改接口', 'PUT', '/api/update/:id', NULL, '2025-05-08 17:25:15.000', NULL);
INSERT INTO `sys_apis` VALUES (28, 0, '系统字典管理', '', '', NULL, '2025-05-08 17:26:04.000', NULL);
INSERT INTO `sys_apis` VALUES (29, 28, '新增条目', 'POST', '/dict/add', NULL, '2025-05-08 17:25:39.000', NULL);
INSERT INTO `sys_apis` VALUES (30, 28, '字典列表', 'GET', '/dict/list', NULL, '2025-05-08 17:26:24.000', NULL);
INSERT INTO `sys_apis` VALUES (31, 28, '修改字典', 'PUT', '/dict/update/:id', NULL, '2025-05-08 17:26:43.000', NULL);
INSERT INTO `sys_apis` VALUES (32, 28, '删除字典', 'DELETE', '/dict/delete/:id', NULL, '2025-05-08 17:27:04.000', NULL);
INSERT INTO `sys_apis` VALUES (33, 0, '系统日志管理', '', '', NULL, '2025-05-08 17:27:58.000', NULL);
INSERT INTO `sys_apis` VALUES (34, 33, '日志列表', 'GET', '/record/list', NULL, '2025-05-08 17:28:16.000', NULL);
INSERT INTO `sys_apis` VALUES (35, 33, '日志删除', 'DELETE', '/record/delete/:id', NULL, '2025-05-08 17:28:36.000', NULL);

-- ----------------------------
-- Table structure for sys_dicts
-- ----------------------------
DROP TABLE IF EXISTS `sys_dicts`;
CREATE TABLE `sys_dicts`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键字段',
  `dict_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `dict_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `item_key` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `item_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `sort` bigint(20) NULL DEFAULT NULL,
  `status` bigint(20) NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '数据字典' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_dicts
-- ----------------------------
INSERT INTO `sys_dicts` VALUES (1, '测试', 'test', '测试', '测试', 0, 1, '', '2025-05-09 17:50:01.573', '2025-05-09 17:50:01.573', NULL);

-- ----------------------------
-- Table structure for sys_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_menus`;
CREATE TABLE `sys_menus`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键字段',
  `parent_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `tree_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `route_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `perm` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `status` bigint(20) NULL DEFAULT NULL,
  `affix_tab` bigint(20) NULL DEFAULT NULL,
  `hide_children_in_menu` bigint(20) NULL DEFAULT NULL,
  `hide_in_breadcrumb` bigint(20) NULL DEFAULT NULL,
  `hide_in_menu` bigint(20) NULL DEFAULT NULL,
  `hide_in_tab` bigint(20) NULL DEFAULT NULL,
  `keep_alive` bigint(20) NULL DEFAULT NULL,
  `sort` bigint(20) NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `redirect` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `params` json NULL COMMENT '路由参数',
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '菜单管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
INSERT INTO `sys_menus` VALUES (1, 0, NULL, 'page.dashboard.title', 'FOLDER', 'Dashboard', '/', 'BasicLayout', '', 1, 0, 0, 0, 0, 0, 0, 1, 'lucide:layout-dashboard', NULL, NULL, NULL, '2025-05-08 16:58:04.000', NULL);
INSERT INTO `sys_menus` VALUES (2, 1, NULL, 'page.dashboard.analytics', 'MENU', 'Analytics', '/analytics', '/views/dashboard/analytics/index', '', 1, 0, 0, 0, 0, 0, 0, 2, 'lucide:area-chart', NULL, NULL, NULL, '2025-05-08 16:58:18.000', NULL);
INSERT INTO `sys_menus` VALUES (3, 0, NULL, 'page.system.title', 'FOLDER', 'System', '/', 'BasicLayout', '', 1, 0, 0, 0, 0, 0, 0, 10, 'lucide:align-start-vertical', NULL, NULL, NULL, '2025-05-08 16:58:52.000', NULL);
INSERT INTO `sys_menus` VALUES (4, 3, NULL, 'page.system.user.title', 'MENU', 'User', '/system/user', '/views/system/user/index', '', 1, 0, 0, 0, 0, 0, 0, 2, 'lucide:user-round-cog', NULL, NULL, NULL, '2025-05-08 16:59:06.000', NULL);
INSERT INTO `sys_menus` VALUES (5, 3, NULL, 'page.system.menu.title', 'MENU', 'Menu', '/system/menu', '/views/system/menu/index', '', 1, 0, 0, 0, 0, 0, 0, 3, 'lucide:menu', NULL, NULL, NULL, '2025-05-08 17:01:31.000', NULL);
INSERT INTO `sys_menus` VALUES (7, 3, NULL, 'page.system.role.title', 'MENU', 'Role', '/system/role', '/views/system/role/index', '', 1, 0, 0, 0, 0, 0, 0, 4, 'lucide:anchor', NULL, NULL, NULL, '2025-05-08 17:45:20.000', NULL);
INSERT INTO `sys_menus` VALUES (8, 4, NULL, 'page.system.user.button.create', 'BUTTON', 'CreateUser', '', 'BasicLayout', 'system:user:create', 1, 0, 0, 0, 0, 0, 0, 1, '', NULL, NULL, NULL, '2025-05-08 16:59:29.000', NULL);
INSERT INTO `sys_menus` VALUES (9, 4, NULL, 'page.system.user.button.delete', 'BUTTON', 'DeleteUser', '', 'BasicLayout', 'system:user:delete', 1, 0, 0, 0, 0, 0, 0, 5, '', NULL, NULL, NULL, '2025-05-08 17:00:14.000', NULL);
INSERT INTO `sys_menus` VALUES (10, 4, NULL, 'page.system.user.button.update', 'BUTTON', 'UpdateUser', '', 'BasicLayout', 'system:user:update', 1, 0, 0, 0, 0, 0, 0, 3, '', NULL, NULL, NULL, '2025-05-08 17:00:42.000', NULL);
INSERT INTO `sys_menus` VALUES (11, 4, NULL, 'ui.button.search', 'BUTTON', 'SearchUser', '', 'BasicLayout', 'system:user:search', 1, 0, 0, 0, 0, 0, 0, 4, '', NULL, NULL, NULL, '2025-05-08 17:40:44.000', NULL);
INSERT INTO `sys_menus` VALUES (12, 5, NULL, 'page.system.menu.button.create', 'BUTTON', 'MenuCreate', '', 'BasicLayout', 'system:menu:create', 1, 0, 0, 0, 0, 0, 0, 1, '', NULL, NULL, NULL, '2025-05-08 17:02:03.000', NULL);
INSERT INTO `sys_menus` VALUES (13, 5, NULL, 'page.system.menu.button.delete', 'BUTTON', 'MenuDelete', '', 'BasicLayout', 'system:menu:delete', 1, 0, 0, 0, 0, 0, 0, 2, '', NULL, NULL, NULL, '2025-05-08 17:41:28.000', NULL);
INSERT INTO `sys_menus` VALUES (14, 1, NULL, '', '', NULL, '', NULL, NULL, 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, NULL, '2025-05-09 10:06:18.000');
INSERT INTO `sys_menus` VALUES (15, 5, NULL, 'ui.button.search', 'BUTTON', 'MenuSearch', '', 'BasicLayout', 'system:menu:search', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-08 17:02:33.000', NULL);
INSERT INTO `sys_menus` VALUES (16, 7, NULL, 'page.system.role.button.create', 'BUTTON', 'RoleCreate', '', 'BasicLayout', 'system:role:create', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-08 17:03:11.000', NULL);
INSERT INTO `sys_menus` VALUES (17, 7, NULL, 'page.system.role.button.delete', 'BUTTON', 'RoleDelete', '', 'BasicLayout', 'system:role:delete', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-08 17:45:53.000', NULL);
INSERT INTO `sys_menus` VALUES (18, 7, NULL, 'page.system.role.button.update', 'BUTTON', 'RoleUpdate', '', 'BasicLayout', 'system:role:update', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-08 17:46:20.000', NULL);
INSERT INTO `sys_menus` VALUES (19, 7, NULL, 'ui.button.search', 'BUTTON', 'RoleSearch', '', 'BasicLayout', 'system:role:search', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-08 17:46:46.000', NULL);
INSERT INTO `sys_menus` VALUES (20, 3, NULL, 'page.system.api.title', 'MENU', 'Api', '/system/api', '/views/system/api/index', '', 1, 0, 0, 0, 0, 0, 0, 0, 'lucide:atom', NULL, NULL, NULL, '2025-05-08 17:47:18.000', NULL);
INSERT INTO `sys_menus` VALUES (21, 20, NULL, 'page.system.api.button.create', 'BUTTON', 'ApiCreate', '', 'BasicLayout', 'system:api:create', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-08 17:47:38.000', NULL);
INSERT INTO `sys_menus` VALUES (22, 20, NULL, 'page.system.api.button.delete', 'BUTTON', 'ApiDelete', '', 'BasicLayout', 'system:api:delete', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-08 17:48:24.000', NULL);
INSERT INTO `sys_menus` VALUES (23, 20, NULL, 'page.system.api.button.update', 'BUTTON', 'ApiUpdate', '', 'BasicLayout', 'system:api:update', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-08 17:48:00.000', NULL);
INSERT INTO `sys_menus` VALUES (24, 20, NULL, 'ui.button.search', 'BUTTON', 'ApiSearch', '', 'BasicLayout', 'system:api:search', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-08 17:48:41.000', NULL);
INSERT INTO `sys_menus` VALUES (25, 3, NULL, 'page.system.dict.title', 'MENU', 'Dict', '/system/dict', '/views/system/dict/index', '', 1, 0, 0, 0, 0, 0, 0, 0, 'lucide:book-open-text', NULL, NULL, NULL, '2025-05-08 17:49:08.000', NULL);
INSERT INTO `sys_menus` VALUES (26, 25, NULL, 'page.system.dict.button.create', 'BUTTON', 'DictCreate', '', 'BasicLayout', 'system:dict:create', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-08 17:49:42.000', NULL);
INSERT INTO `sys_menus` VALUES (27, 25, NULL, 'page.system.dict.button.delete', 'BUTTON', 'DictDelete', '', 'BasicLayout', 'system:dict:delete', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-08 17:50:02.000', NULL);
INSERT INTO `sys_menus` VALUES (28, 25, NULL, 'page.system.dict.button.update', 'BUTTON', 'DictUpdate', '', 'BasicLayout', 'system:dict:update', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-08 17:50:19.000', NULL);
INSERT INTO `sys_menus` VALUES (29, 25, NULL, 'ui.button.search', 'BUTTON', 'DictSearch', '', 'BasicLayout', 'system:dict:search', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-08 17:50:47.000', NULL);
INSERT INTO `sys_menus` VALUES (30, 3, NULL, 'page.system.record.title', 'MENU', 'RecordList', '/system/record', '/views/system/record/index', '', 1, 0, 0, 0, 0, 0, 0, 0, 'lucide:arrow-left-right', NULL, NULL, NULL, '2025-05-08 17:51:11.000', NULL);
INSERT INTO `sys_menus` VALUES (31, 30, NULL, 'page.system.record.button.delete', 'BUTTON', 'RecordDelete', '', 'BasicLayout', 'system:record:delete', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-09 09:42:49.000', NULL);
INSERT INTO `sys_menus` VALUES (32, 1, NULL, '他俄速通', 'FOLDER', 'a', 'a', 'BasicLayout', '', 1, 0, 0, 0, 0, 0, 0, 0, 'lucide:album', NULL, NULL, NULL, '2025-05-08 16:58:33.000', '2025-05-09 10:06:23.000');
INSERT INTO `sys_menus` VALUES (34, 0, NULL, '租户管理', 'FOLDER', 'tenant', '/sys/tenant', 'BasicLayout', '', 1, 0, 0, 0, 0, 0, 0, 0, 'lucide:album', NULL, NULL, NULL, '2025-05-09 09:43:24.000', NULL);
INSERT INTO `sys_menus` VALUES (35, 0, NULL, '租户管理', 'MENU', 'Tenant', '/sys/tenant/index', 'BasicLayout', '', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-09 09:43:50.000', NULL);
INSERT INTO `sys_menus` VALUES (37, 8, NULL, '12', 'FOLDER', '12', '23', 'BasicLayout', '', 1, 0, 0, 0, 0, 0, 0, 0, '', NULL, NULL, NULL, '2025-05-08 17:40:15.000', NULL);
INSERT INTO `sys_menus` VALUES (38, 0, '', 'test', 'FOLDER', 'test111', '/test111', 'BasicLayout', '', 1, 0, 0, 0, 0, 0, 0, 0, '', '', NULL, '2025-05-09 17:45:38.795', '2025-05-09 17:45:38.795', NULL);

-- ----------------------------
-- Table structure for sys_operate_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operate_records`;
CREATE TABLE `sys_operate_records`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `user_id` bigint(20) NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `status_code` bigint(20) NULL DEFAULT NULL,
  `elapsed` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `request` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `response` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL,
  `platform` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_operate_records
-- ----------------------------
INSERT INTO `sys_operate_records` VALUES (2, '2025-05-09 17:06:55.174', '2025-05-09 17:06:55.174', NULL, '', 0, '', 'POST', '/auth/login', 7, '1.59', '密码错误', '{\"post\":\"{\\\"username\\\":\\\"admin\\\",\\\"password\\\":\\\"123456\\\",\\\"captcha\\\":true}\"}', '{}', 'Windows Google Chrome', '127.0.0.1', '');
INSERT INTO `sys_operate_records` VALUES (3, '2025-05-09 17:07:06.835', '2025-05-09 17:07:06.835', NULL, '', 0, '', 'POST', '/auth/login', 0, '77.94', '操作成功', '{\"post\":\"{\\\"username\\\":\\\"sfchen\\\",\\\"password\\\":\\\"123456\\\",\\\"captcha\\\":true}\"}', '{\"accessToken\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwiSUQiOjQsIlVzZXJuYW1lIjoic2ZjaGVuIiwiTmlja05hbWUiOiJzZmNoZW4iLCJSb2xlSWQiOjEsIkVtYWlsIjoiIiwiQnVmZmVyVGltZSI6MH0.RPn7-MX8aDJ8WjaD3cMska5bwsLMJcGc2AWf8lEFsIo\",\"id\":4,\"password\":\"\",\"realName\":\"sfchen\",\"roles\":[\"\"],\"username\":\"sfchen\"}', 'Windows Google Chrome', '127.0.0.1', '');
INSERT INTO `sys_operate_records` VALUES (4, '2025-05-09 17:07:58.250', '2025-05-09 17:07:58.250', NULL, '', 0, '', 'POST', '/auth/login', 200, '48.13', '操作成功', '{\"post\":\"{\\\"username\\\":\\\"sfchen\\\",\\\"password\\\":\\\"123456\\\",\\\"captcha\\\":true}\"}', '{\"accessToken\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiMDAwMDAwMDAtMDAwMC0wMDAwLTAwMDAtMDAwMDAwMDAwMDAwIiwiSUQiOjQsIlVzZXJuYW1lIjoic2ZjaGVuIiwiTmlja05hbWUiOiJzZmNoZW4iLCJSb2xlSWQiOjEsIkVtYWlsIjoiIiwiQnVmZmVyVGltZSI6MH0.RPn7-MX8aDJ8WjaD3cMska5bwsLMJcGc2AWf8lEFsIo\",\"id\":4,\"password\":\"\",\"realName\":\"sfchen\",\"roles\":[\"\"],\"username\":\"sfchen\"}', 'Windows Google Chrome', '127.0.0.1', '');
INSERT INTO `sys_operate_records` VALUES (5, '2025-05-09 17:32:53.436', '2025-05-09 17:32:53.436', NULL, '', 0, '', 'DELETE', '/user/delete', 200, '2.30', '操作成功', '{\"id\":1,\"post\":\"\"}', '{}', 'Windows Google Chrome', '127.0.0.1', '');
INSERT INTO `sys_operate_records` VALUES (6, '2025-05-09 17:49:04.602', '2025-05-09 17:49:04.602', NULL, '', 0, '', 'DELETE', '/api/delete', 7, '0.00', '请先删除角色API权限后再操作', '{\"id\":27,\"post\":\"\"}', '{}', 'Windows Google Chrome', '127.0.0.1', '');
INSERT INTO `sys_operate_records` VALUES (7, '2025-05-09 17:50:01.575', '2025-05-09 17:50:01.575', NULL, '', 0, '', 'POST', '/dict/add', 200, '1.51', '操作成功', '{\"post\":\"{\\\"status\\\":1,\\\"dictName\\\":\\\"测试\\\",\\\"dictType\\\":\\\"test\\\",\\\"itemKey\\\":\\\"测试\\\",\\\"itemValue\\\":\\\"测试\\\"}\"}', '{}', 'Windows Google Chrome', '127.0.0.1', '');
INSERT INTO `sys_operate_records` VALUES (8, '2025-05-09 17:53:27.679', '2025-05-09 17:53:27.679', NULL, '', 0, '', 'POST', '/user/add', 200, '79.06', '操作成功', '{\"post\":\"{\\\"status\\\":1,\\\"username\\\":\\\"alen\\\",\\\"nickname\\\":\\\"alen\\\",\\\"roleId\\\":1}\"}', '{}', 'Windows Google Chrome', '127.0.0.1', '');
INSERT INTO `sys_operate_records` VALUES (9, '2025-05-09 17:56:25.002', '2025-05-09 17:56:25.002', NULL, 'sfchen', 4, '', 'PUT', '/user/update', 200, '4.65', '操作成功', '{\"id\":5,\"post\":\"{\\\"status\\\":1,\\\"username\\\":\\\"alen\\\",\\\"nickname\\\":\\\"alen\\\",\\\"roleId\\\":1,\\\"email\\\":\\\"\\\",\\\"remark\\\":\\\"3223\\\"}\"}', '{}', 'Windows Google Chrome', '127.0.0.1', '');
INSERT INTO `sys_operate_records` VALUES (10, '2025-05-09 17:57:14.124', '2025-05-09 17:57:14.124', NULL, 'sfchen', 4, '', 'POST', '/role/add', 200, '5.10', '操作成功', '{\"post\":\"{\\\"status\\\":1,\\\"name\\\":\\\"admin\\\",\\\"code\\\":\\\"admin\\\"}\"}', '{}', 'Windows Google Chrome', '127.0.0.1', '');
INSERT INTO `sys_operate_records` VALUES (11, '2025-05-09 17:57:31.446', '2025-05-09 17:57:31.446', NULL, 'sfchen', 4, '', 'PUT', '/user/update', 200, '4.59', '操作成功', '{\"id\":5,\"post\":\"{\\\"status\\\":1,\\\"username\\\":\\\"alen\\\",\\\"nickname\\\":\\\"alen\\\",\\\"roleId\\\":2,\\\"email\\\":\\\"\\\",\\\"remark\\\":\\\"3223\\\"}\"}', '{}', 'Windows Google Chrome', '127.0.0.1', '');

-- ----------------------------
-- Table structure for sys_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_records`;
CREATE TABLE `sys_records`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键字段',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户名',
  `user_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '用户ID',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '描述',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '请求方法',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '请求路径',
  `status_code` bigint(20) NOT NULL DEFAULT 0 COMMENT '状态码',
  `elapsed` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '耗时',
  `msg` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '返回的msg',
  `request` varchar(2555) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '请求参数',
  `response` varchar(2555) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '返回参数',
  `platform` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '平台',
  `ip` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'IP',
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '地址',
  `created_at` datetime NULL DEFAULT NULL COMMENT '创建时间戳',
  `updated_at` datetime NULL DEFAULT NULL COMMENT '修改时间戳',
  `deleted_at` datetime NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '操作日志' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_records
-- ----------------------------
INSERT INTO `sys_records` VALUES (1, '', 0, '', 'POST', '/auth/logout', 200, '0.00', 'success', '{\"post\":\"{\\\"withCredentials\\\":true}\"}', 'null', 'Windows Google Chrome', '127.0.0.1', '', '2025-05-09 16:58:09', '2025-05-09 16:58:09', NULL);
INSERT INTO `sys_records` VALUES (2, '', 0, '用户登录', 'POST', '/auth/login', 200, '1.85', 'success', '{\"post\":\"{\\\"username\\\":\\\"admin\\\",\\\"password\\\":\\\"123456\\\",\\\"captcha\\\":true}\"}', '{\"accessToken\":\"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDY3ODIyOTcsImlhdCI6MTc0Njc4MTA5NywiaWQiOjEsIm5iZiI6MTc0Njc4MTA5Nywicm9sZV9pZCI6MSwidXNlcm5hbWUiOiJhZG1pbiJ9.bX3ZMTQmxqA0Xcrg0s0KHRunEQo5W8V2mGSmWm-BmuI\",\"id\":1,\"password\":\"\",\"realName\":\"admin\",\"roles\":[\"\"],\"username\":\"admin\"}', 'Windows Google Chrome', '127.0.0.1', '', '2025-05-09 16:58:18', '2025-05-09 16:58:18', NULL);

-- ----------------------------
-- Table structure for sys_role_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_apis`;
CREATE TABLE `sys_role_apis`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键字段',
  `role_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `api_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `uidx_role_id`(`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 69 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '角色接口权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_apis
-- ----------------------------
INSERT INTO `sys_role_apis` VALUES (35, 1, 1, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (36, 1, 2, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (37, 1, 3, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (38, 1, 4, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (39, 1, 5, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (40, 1, 6, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (41, 1, 7, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (42, 1, 8, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (43, 1, 9, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (44, 1, 10, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (45, 1, 12, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (46, 1, 13, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (47, 1, 14, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (48, 1, 15, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (49, 1, 16, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (50, 1, 17, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (51, 1, 18, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (52, 1, 19, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (53, 1, 20, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (54, 1, 21, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (55, 1, 22, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (56, 1, 23, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (57, 1, 24, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (58, 1, 25, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (59, 1, 26, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (60, 1, 27, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (61, 1, 28, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (62, 1, 29, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (63, 1, 30, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (64, 1, 31, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (65, 1, 32, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (66, 1, 33, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (67, 1, 34, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_apis` VALUES (68, 1, 35, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);

-- ----------------------------
-- Table structure for sys_role_auths
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_auths`;
CREATE TABLE `sys_role_auths`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键字段',
  `role_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `auth_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `uidx_role_id`(`role_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 64 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '角色菜单权限' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role_auths
-- ----------------------------
INSERT INTO `sys_role_auths` VALUES (34, 1, 2, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (35, 1, 3, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (36, 1, 4, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (37, 1, 8, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (38, 1, 37, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (39, 1, 9, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (40, 1, 10, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (41, 1, 11, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (42, 1, 5, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (43, 1, 12, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (44, 1, 13, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (45, 1, 15, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (46, 1, 7, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (47, 1, 16, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (48, 1, 17, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (49, 1, 18, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (50, 1, 19, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (51, 1, 20, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (52, 1, 21, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (53, 1, 22, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (54, 1, 23, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (55, 1, 24, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (56, 1, 25, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (57, 1, 26, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (58, 1, 27, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (59, 1, 28, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (60, 1, 29, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (61, 1, 30, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (62, 1, 31, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);
INSERT INTO `sys_role_auths` VALUES (63, 1, 1, '2025-05-09 10:06:10.000', '2025-05-09 10:06:10.000', NULL);

-- ----------------------------
-- Table structure for sys_roles
-- ----------------------------
DROP TABLE IF EXISTS `sys_roles`;
CREATE TABLE `sys_roles`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键字段',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `sort` bigint(20) NULL DEFAULT NULL,
  `status` bigint(20) NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '角色管理' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_roles
-- ----------------------------
INSERT INTO `sys_roles` VALUES (1, '超级管理员', 'super', 0, 1, '', NULL, '2025-05-08 17:07:33.000', NULL);
INSERT INTO `sys_roles` VALUES (2, 'admin', 'admin', 0, 1, '', '2025-05-09 17:57:14.120', '2025-05-09 17:57:14.120', NULL);

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键字段',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `salt` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `gender` bigint(20) NULL DEFAULT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `status` bigint(20) NULL DEFAULT NULL,
  `dept_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `role_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `create_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `update_by` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `last_login_time` bigint(20) NULL DEFAULT NULL,
  `last_login_ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `uuid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户UUID',
  `nick_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT 'https://weappimg.apm-monaco.cn/default_avatar.jpg' COMMENT '用户头像',
  `base_color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '#fff' COMMENT '基础颜色',
  `active_color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT '#1890ff' COMMENT '活跃颜色',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户手机号',
  `enable` bigint(20) NULL DEFAULT 1 COMMENT '用户是否被冻结 1正常 2冻结',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '系统用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, 'admin', 'admin', '3f29734092a7ede4be7d88642fd6b0b9', 'SAPtYl', '', 0, '', '/uploads/default/logo.png', 1, 0, 1, '', '0', '0', 0, '', '2025-05-08 16:14:56.000', '2025-05-08 16:14:56.000', '2025-05-09 17:32:53.435', NULL, '系统用户', 'dark', 'https://weappimg.apm-monaco.cn/default_avatar.jpg', '#fff', '#1890ff', NULL, 1);
INSERT INTO `sys_users` VALUES (2, 'superAdmin', 'superAdmin', 'd11f39c5e680dadc1f985bb5c528088f', 'DaJHF4', '', 0, '', '/uploads/default/logo.png', 1, 0, 1, '', '0', '0', 0, '', '2025-05-09 10:20:44.000', '2025-05-09 10:20:44.000', NULL, NULL, '系统用户', 'dark', 'https://weappimg.apm-monaco.cn/default_avatar.jpg', '#fff', '#1890ff', NULL, 1);

-- ----------------------------
-- Table structure for system_routes
-- ----------------------------
DROP TABLE IF EXISTS `system_routes`;
CREATE TABLE `system_routes`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '名称',
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '路由地址',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '请求类型',
  `status` bigint(20) NULL DEFAULT 1 COMMENT '状态',
  `type` bigint(20) NOT NULL COMMENT '1按钮2菜单3接口',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of system_routes
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
