
DROP TABLE IF EXISTS `role_menus`;
CREATE TABLE `role_menus`  (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
    `role_id` bigint UNSIGNED NOT NULL  COMMENT '角色id',
    `menu_id` bigint UNSIGNED NOT NULL  COMMENT '用户对应的菜单id',
    `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
    UNIQUE KEY uk_role_menu (role_id, menu_id)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
