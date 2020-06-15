CREATE TABLE `gnix_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `mail` varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱',
  `phone` varchar(16) NOT NULL DEFAULT '' COMMENT '手机',
  `status` tinyint(4) unsigned NOT NULL DEFAULT 0 COMMENT '状态 0 - 新加 1 - 正常 2 - 禁用 3 - 删除',
  `role` bigint(20) NOT NULL DEFAULT 0 COMMENT '角色, 0 - 超级管理员',
  `name` varchar(128) NOT NULL DEFAULT '' COMMENT '姓名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(32) NOT NULL DEFAULT '' COMMENT '计算密码的盐值',
  `create_time` timestamp NOT NULL DEFAULT current_timestamp() COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp() COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unqi_mail` (`mail`),
  UNIQUE KEY `unqi_phone` (`phone`)
) ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表'
