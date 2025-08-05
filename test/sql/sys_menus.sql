
DROP TABLE IF EXISTS `sys_menus`;
CREATE TABLE `sys_menus`  (
   id BIGINT AUTO_INCREMENT COMMENT 'id',
   `sort` INT UNSIGNED NOT NULL DEFAULT 1 COMMENT '排序编号',
  `menu_level` int UNSIGNED NOT NULL COMMENT '菜单层级',
  `menu_type` int UNSIGNED NOT NULL COMMENT '菜单类型 （菜单或目录）0 目录 1 菜单',
  `path` varchar(255)   NULL DEFAULT '' COMMENT '菜单路由路径',
  `name` varchar(255)   NOT NULL COMMENT '菜单名称',
  `redirect` varchar(255)   NULL DEFAULT '' COMMENT '跳转路径 （外链）',
  `component` varchar(255)   NULL DEFAULT '' COMMENT '组件路径',
  `disabled` tinyint(1) NULL DEFAULT 0 COMMENT '是否停用',
  `service_name` varchar(255)   NULL DEFAULT 'Other' COMMENT '服务名称',
  `permission` varchar(255)   NULL DEFAULT NULL COMMENT '权限标识',
  `title` varchar(255)   NOT NULL COMMENT '菜单显示标题',
  `icon` varchar(255)   NOT NULL COMMENT '菜单图标',
  `hide_menu` tinyint(1) NULL DEFAULT 0 COMMENT '是否隐藏菜单',
  `hide_breadcrumb` tinyint(1) NULL DEFAULT 0 COMMENT '隐藏面包屑',
  `ignore_keep_alive` tinyint(1) NULL DEFAULT 0 COMMENT '取消页面缓存',
  `hide_tab` tinyint(1) NULL DEFAULT 0 COMMENT '隐藏页头',
  `frame_src` varchar(255)   NULL DEFAULT '' COMMENT '内嵌 iframe',
  `carry_param` tinyint(1) NULL DEFAULT 0 COMMENT '携带参数',
  `hide_children_in_menu` tinyint(1) NULL DEFAULT 0 COMMENT '隐藏所有子菜单',
  `affix` tinyint(1) NULL DEFAULT 0 COMMENT 'Tab 固定',
  `dynamic_level` int UNSIGNED NULL DEFAULT 20 COMMENT 'T能打开的子TAB数',
  `real_path` varchar(255)   NULL DEFAULT '' COMMENT '菜单路由不包含参数部分',
  `parent_id` bigint UNSIGNED NULL DEFAULT 0 COMMENT '父菜单ID',
   `created_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) COMMENT '创建时间',
   `updated_at` DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3) COMMENT '修改时间',
   `deleted_at` DATETIME(3) NULL DEFAULT NULL COMMENT '软删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `menu_name`(`name`) USING BTREE,
  UNIQUE INDEX `menu_path`(`path`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin COMMENT = '菜单表' ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
