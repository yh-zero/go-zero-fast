
DROP TABLE IF EXISTS `sys_roles`;
CREATE TABLE `sys_roles`  (
    `id` BIGINT UNSIGNED NOT NULL COMMENT '雪花算法ID（64位整型）',
    `status` TINYINT UNSIGNED NULL DEFAULT 1 COMMENT '状态:1正常/2禁用',
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '角色名',
    `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '角色码，用于前端权限控制',
    `default_router` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'dashboard' COMMENT '默认登录页面',
    `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT '备注',
    `sort` int UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序编号',
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `role_code`(`code`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = 'Role Table | 角色信息表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
