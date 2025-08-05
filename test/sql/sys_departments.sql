
DROP TABLE IF EXISTS `sys_departments`;
CREATE TABLE `sys_departments`  (
        `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '部门id-自增',
        `status` TINYINT UNSIGNED NULL DEFAULT 1 COMMENT '状态:1正常/2禁用',
        `sort` INT UNSIGNED NOT NULL DEFAULT 1 COMMENT '排序编号',
        `name` varchar(255)   NOT NULL COMMENT '部门名称',
        `parent_id` bigint UNSIGNED NULL DEFAULT 0 COMMENT '父级部门ID 0预留是一级部门',
        `ancestors` varchar(255)   NULL DEFAULT NULL COMMENT '父级列表-预留',
        `leader` varchar(255)   NULL DEFAULT NULL COMMENT '部门负责人',
        `phone` varchar(255)   NULL DEFAULT NULL COMMENT '负责人电话',
        `email` varchar(255)   NULL DEFAULT NULL COMMENT '部门负责人电子邮箱',
        `remark` varchar(255)   NULL DEFAULT NULL COMMENT '备注',
        `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
        `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
        `deleted_at` DATETIME(3) NULL DEFAULT NULL COMMENT '软删除时间',
        PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = 'Department Table | 部门表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
