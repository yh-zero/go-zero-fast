/*
 Navicat Premium Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 80034
 Source Host           : localhost:3306
 Source Schema         : go-zero-fast

 Target Server Type    : MySQL
 Target Server Version : 80034
 File Encoding         : 65001

 Date: 04/08/2025 13:00:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_departments
-- ----------------------------
DROP TABLE IF EXISTS `sys_departments`;
CREATE TABLE `sys_departments`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '部门id-自增',
  `status` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '状态:1正常/2禁用',
  `sort` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '排序编号',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '部门名称',
  `parent_id` bigint UNSIGNED NULL DEFAULT 0 COMMENT '父级部门ID 0预留是一级部门',
  `ancestors` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '父级列表-预留',
  `leader` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '部门负责人',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '负责人电话',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '部门负责人电子邮箱',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '部门表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_departments
-- ----------------------------
INSERT INTO `sys_departments` VALUES (1, 1, 1, '部门1', 0, NULL, '部门负责人', '10086', NULL, NULL, '2025-08-04 12:52:50.487', '2025-08-04 12:53:03.970', NULL);

-- ----------------------------
-- Table structure for sys_positions
-- ----------------------------
DROP TABLE IF EXISTS `sys_positions`;
CREATE TABLE `sys_positions`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `status` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '状态:1正常/2禁用',
  `sort` int UNSIGNED NOT NULL DEFAULT 1 COMMENT '排序编号',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '职位名称',
  `code` int UNSIGNED NOT NULL COMMENT '职位编码',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `position_code`(`code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '职位信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_positions
-- ----------------------------
INSERT INTO `sys_positions` VALUES (2, 1, 1, '职位1', 1, '职位说明', '2025-08-04 11:52:09.816', '2025-08-04 11:52:09.816', NULL);

-- ----------------------------
-- Table structure for sys_roles
-- ----------------------------
DROP TABLE IF EXISTS `sys_roles`;
CREATE TABLE `sys_roles`  (
  `id` bigint UNSIGNED NOT NULL COMMENT '雪花算法ID（64位整型）',
  `status` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '状态:1正常/2禁用',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '角色名',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '角色码，用于前端权限控制',
  `default_router` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'dashboard' COMMENT '默认登录页面',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '备注',
  `sort` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序编号',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `role_code`(`code`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '角色信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_roles
-- ----------------------------
INSERT INTO `sys_roles` VALUES (2001, 1, '角色1', '2200', 'dashboard', '', 0, '2025-07-31 09:24:13.575', '2025-07-31 09:24:13.575');
INSERT INTO `sys_roles` VALUES (20021, 11, '角色1', '2201', 'dashboard', '', 0, '2025-07-31 09:24:13.575', '2025-07-31 09:57:12.621');

-- ----------------------------
-- Table structure for sys_tokens
-- ----------------------------
DROP TABLE IF EXISTS `sys_tokens`;
CREATE TABLE `sys_tokens`  (
  `id` bigint UNSIGNED NOT NULL COMMENT '雪花算法ID（64位整型）',
  `status` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '状态:1正常/2禁用',
  `user_id` bigint UNSIGNED NOT NULL COMMENT '用户的id',
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'unknown' COMMENT '用户名',
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Token 字符串',
  `source` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Token 来源（local, 第三方如github等）',
  `expired_at` datetime(3) NOT NULL COMMENT '过期时间',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_token`(`token`) USING BTREE,
  INDEX `idx_user_id`(`user_id`) USING BTREE,
  INDEX `idx_expired_at`(`expired_at`) USING BTREE,
  INDEX `idx_user_token`(`user_id`, `token`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '令牌信息表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_tokens
-- ----------------------------
INSERT INTO `sys_tokens` VALUES (1, 1, 1001, 'unknown', '1', 'local', '2025-07-30 15:08:43.000', '2025-07-30 15:08:44.778', '2025-07-30 15:08:44.778');
INSERT INTO `sys_tokens` VALUES (226274307367583744, 1, 1001, 'yanghao', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3ODU0ODM5OTgsImlhdCI6MTc1Mzk0Nzk5OCwiand0RGF0YSI6eyJSb2xlSWRzIjpbMjAwMSwyMDAyXSwiUm9sZUlkIjoyMDAxLCJEZXBhcnRtZW50SWQiOjEsIlVzZXJJZCI6MTAwMX19.9YcTJj9GBziKUm5pH-p-rheeQyip94-63ipi9T-Zr98', 'local', '2026-07-31 07:46:38.847', '2025-07-31 15:46:38.851', '2025-07-31 15:46:38.851');
INSERT INTO `sys_tokens` VALUES (226275280060235776, 1, 1001, 'yanghao', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3ODU0ODQyMzAsImlhdCI6MTc1Mzk0ODIzMCwiand0RGF0YSI6eyJSb2xlSWRzIjpbMjAwMSwyMDAyXSwiUm9sZUlkIjoyMDAxLCJEZXBhcnRtZW50SWQiOjEsIlVzZXJJZCI6MTAwMX19.8EszRuvj4fNUAGdf6wYQjQZ0XT5pQFwJc7We7ZHomWM', 'local', '2026-07-31 07:50:30.756', '2025-07-31 15:50:30.758', '2025-07-31 15:50:30.758');
INSERT INTO `sys_tokens` VALUES (226276004122935296, 1, 1001, 'yanghao', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQwMTMyMDMsImlhdCI6MTc1Mzk0ODQwMywiand0RGF0YSI6eyJSb2xlSWRzIjpbMjAwMSwyMDAyXSwiUm9sZUlkIjoyMDAxLCJEZXBhcnRtZW50SWQiOjEsIlVzZXJJZCI6MTAwMX19.m2vuE4yKcvZsoOVJHmRvjoRqG9kLA3uyPhvxRAdOu6Y', 'local', '2025-08-01 01:53:23.386', '2025-07-31 15:53:23.388', '2025-07-31 15:53:23.388');
INSERT INTO `sys_tokens` VALUES (226276015166537728, 1, 1001, 'yanghao', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQwMTMyMDYsImlhdCI6MTc1Mzk0ODQwNiwiand0RGF0YSI6eyJSb2xlSWRzIjpbMjAwMSwyMDAyXSwiUm9sZUlkIjoyMDAxLCJEZXBhcnRtZW50SWQiOjEsIlVzZXJJZCI6MTAwMX19.UrLiyVikjW5gePck5JTwFTThe5oBheggp7Rx5BoZ5FA', 'local', '2025-08-01 01:53:26.020', '2025-07-31 15:53:26.021', '2025-07-31 15:53:26.021');
INSERT INTO `sys_tokens` VALUES (226276023601283072, 1, 1001, 'yanghao', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQwMTMyMDgsImlhdCI6MTc1Mzk0ODQwOCwiand0RGF0YSI6eyJSb2xlSWRzIjpbMjAwMSwyMDAyXSwiUm9sZUlkIjoyMDAxLCJEZXBhcnRtZW50SWQiOjEsIlVzZXJJZCI6MTAwMX19.XmxePyQXspgg8j5sc5LLHJXiMtcpu_VrilyBKZ0u3Q8', 'local', '2025-08-01 01:53:28.030', '2025-07-31 15:53:28.031', '2025-07-31 15:53:28.031');
INSERT INTO `sys_tokens` VALUES (226276961930002432, 1, 1001, 'yanghao', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQwMTM0MzEsImlhdCI6MTc1Mzk0ODYzMSwiand0RGF0YSI6eyJSb2xlSWRzIjpbMjAwMSwyMDAyXSwiUm9sZUlkIjoyMDAxLCJEZXBhcnRtZW50SWQiOjEsIlVzZXJJZCI6MTAwMX19.BbnxnZ1SrfOW-zgairtgX1QfgkZuQe-KOZhEmY9BPmg', 'local', '2025-08-01 01:57:11.743', '2025-07-31 15:57:11.746', '2025-07-31 15:57:11.746');
INSERT INTO `sys_tokens` VALUES (226277953887092736, 1, 1001, 'yanghao', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQwMTM2NjgsImlhdCI6MTc1Mzk0ODg2OCwiand0RGF0YSI6eyJSb2xlSWRzIjpbMjAwMSwyMDAyXSwiUm9sZUlkIjoyMDAxLCJEZXBhcnRtZW50SWQiOjEsIlVzZXJJZCI6MTAwMX19.uuh23ylVS1dVEB39Zm7kgruxcWmw-tcy4mkqFulSZfA', 'local', '2025-08-01 02:01:08.246', '2025-07-31 16:01:08.285', '2025-07-31 16:01:08.285');
INSERT INTO `sys_tokens` VALUES (226278273488863232, 1, 1001, 'yanghao', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQwMTM3NDQsImlhdCI6MTc1Mzk0ODk0NCwiand0RGF0YSI6eyJSb2xlSWRzIjpbMjAwMSwyMDAyXSwiUm9sZUlkIjoyMDAxLCJEZXBhcnRtZW50SWQiOjEsIlVzZXJJZCI6MTAwMX19.94bwZCinYeIPRuMMD35l_ZpP5_JUiSdHBI5xixlwV5E', 'local', '2025-08-01 02:02:24.446', '2025-07-31 16:02:24.447', '2025-07-31 16:02:24.447');
INSERT INTO `sys_tokens` VALUES (227648111272607744, 1, 1001, 'yanghao', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTQzNDAzMzksImlhdCI6MTc1NDI3NTUzOSwiand0RGF0YSI6eyJSb2xlSWRzIjpbMjAwMSwyMDAyXSwiUm9sZUlkIjoyMDAxLCJEZXBhcnRtZW50SWQiOjEsIlVzZXJJZCI6MTAwMX19.OoaXeRdfCm94aXkfAOOzClhf0zRSpKFzBZYN-NtKZus', 'local', '2025-08-04 20:45:39.221', '2025-08-04 10:45:39.226', '2025-08-04 10:45:39.226');

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint UNSIGNED NOT NULL COMMENT '雪花算法ID（64位整型）',
  `status` tinyint UNSIGNED NULL DEFAULT 1 COMMENT '状态:1正常/2禁用',
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '登录名',
  `password` char(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '密码（bcrypt加密）',
  `nickname` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '昵称',
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '个人简介',
  `home_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT '/dashboard' COMMENT '登录首页',
  `mobile` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'AES加密手机号',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'AES加密邮箱',
  `avatar` varchar(512) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '头像URL',
  `department_id` bigint UNSIGNED NULL DEFAULT 1 COMMENT '部门ID',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  `updated_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
  `deleted_at` datetime(3) NULL DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_username`(`username`) USING BTREE,
  UNIQUE INDEX `idx_nickname`(`nickname`) USING BTREE,
  INDEX `idx_department`(`department_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户信息表（雪花ID主键）' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1001, 1, 'yanghao', '$2a$10$37E7DBjAGkqu.qaxaDqgj.QB0WO4wem6BHfz1HeyPXWavii8kbhiC', '1', '1沙发发', '/dashboard', '1', '1', '1', 1, '2025-07-29 17:23:24.684', '2025-07-31 10:44:44.654', NULL);

-- ----------------------------
-- Table structure for user_positions
-- ----------------------------
DROP TABLE IF EXISTS `user_positions`;
CREATE TABLE `user_positions`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint UNSIGNED NOT NULL COMMENT '用户id',
  `position_id` bigint UNSIGNED NOT NULL COMMENT '用户对应的职位id',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_user_position`(`user_id`, `position_id`) USING BTREE,
  INDEX `idx_user_id`(`user_id`) USING BTREE,
  INDEX `idx_position_id`(`position_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_positions
-- ----------------------------
INSERT INTO `user_positions` VALUES (1, 1001, 2);

-- ----------------------------
-- Table structure for user_roles
-- ----------------------------
DROP TABLE IF EXISTS `user_roles`;
CREATE TABLE `user_roles`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `role_id` bigint NOT NULL COMMENT '用户对应的角色id',
  `created_at` datetime(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_user_role`(`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user_roles
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
