DROP TABLE IF EXISTS `sys_tokens`;
CREATE TABLE `sys_tokens` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT '雪花算法ID（64位整型）',
    `status` TINYINT UNSIGNED NULL DEFAULT 1 COMMENT '状态:1正常/2禁用',
    `user_id` BIGINT UNSIGNED NOT NULL COMMENT '用户的id',
    `username` VARCHAR(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT 'unknown' COMMENT '用户名',
    `token` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Token 字符串',
    `source` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT 'Token 来源（local, 第三方如github等）',
    `expired_at` DATETIME(3) NOT NULL COMMENT '过期时间',
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE INDEX `idx_token` (`token`) USING BTREE,
    INDEX `idx_user_id` (`user_id`) USING BTREE,
    INDEX `idx_expired_at` (`expired_at`) USING BTREE,
    INDEX `idx_user_token` (`user_id`, `token`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '令牌信息表' ROW_FORMAT = Dynamic;