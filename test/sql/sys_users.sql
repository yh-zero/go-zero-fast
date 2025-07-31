DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users` (
    `id` BIGINT UNSIGNED NOT NULL COMMENT '雪花算法ID（64位整型）',
    `status` TINYINT UNSIGNED DEFAULT 1 COMMENT '状态:1正常/2禁用',
    `username` VARCHAR(64) NOT NULL COMMENT '登录名',
    `password` CHAR(64) NOT NULL COMMENT '密码（bcrypt加密）',
    `nickname` VARCHAR(64) NOT NULL COMMENT '昵称',
    `description` VARCHAR(255) DEFAULT NULL COMMENT '个人简介',
    `home_path` VARCHAR(255) DEFAULT '/dashboard' COMMENT '登录首页',
    `mobile` VARCHAR(255)DEFAULT NULL COMMENT 'AES加密手机号',
    `email` VARCHAR(255) DEFAULT NULL COMMENT 'AES加密邮箱',
    `avatar` VARCHAR(512) DEFAULT NULL COMMENT '头像URL',
    `department_id` BIGINT UNSIGNED DEFAULT 1 COMMENT '部门ID',
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
    `deleted_at` DATETIME(3) NULL DEFAULT NULL COMMENT '软删除时间',
    PRIMARY KEY (`id`),
    UNIQUE INDEX `idx_username` (`username`),
    UNIQUE INDEX `idx_nickname` (`nickname`),
    INDEX `idx_department` (`department_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin
COMMENT='用户信息表（雪花ID主键）'
ROW_FORMAT=DYNAMIC;