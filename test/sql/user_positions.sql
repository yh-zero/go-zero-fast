DROP TABLE IF EXISTS `user_positions`;
CREATE TABLE `user_positions`  (
    id BIGINT AUTO_INCREMENT COMMENT 'id',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '用户id',
    position_id BIGINT UNSIGNED NOT NULL COMMENT '用户对应的职位id',
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_user_position` (`user_id`, `position_id`) USING BTREE,
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_position_id` (`position_id`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
