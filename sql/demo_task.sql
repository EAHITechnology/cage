-- 每个项目需要把用到的表在此记录一下，项目迁移时可以直接使用防止数据漏下
CREATE TABLE `demo_task` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `info` VARCHAR(256) NOT NULL DEFAULT '' COMMENT 'info',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'time',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'time',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin COMMENT='demo';