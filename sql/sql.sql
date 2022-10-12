DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`
(
    `id`          int(64) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`     bigint       NOT NULL default 0 comment '用户id',
    `name`        varchar(50)      NOT NULL COMMENT '用户名',
    `password`    varchar(100)              DEFAULT NULL COMMENT '密码',
    `email`       varchar(100)              DEFAULT NULL COMMENT '邮箱',
    `update_time` bigint(50)                DEFAULT NULL COMMENT '更新时间',
    `create_time` bigint(50)                DEFAULT NULL COMMENT '创建时间',
    `del_status`  tinyint(4)                DEFAULT '0' COMMENT '是否删除 -1：已删除   0：正常',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`user_id`),
    UNIQUE KEY (`name`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 11
  DEFAULT CHARSET = utf8 COMMENT ='用户表';

DROP TABLE IF EXISTS `blog`;

CREATE TABLE `blog`
(
    `id`          int(64) unsigned NOT NULL AUTO_INCREMENT,
    `blog_id`     bigint          NOT NULL COMMENT '博客id',
    `user_id`     bigint          NOT NULL default 0 comment '用户id',
    `title`       varchar(100)              DEFAULT NULL COMMENT '博客题目',
    `content`     varchar(100)              DEFAULT NULL COMMENT '博客内容',
    `update_time` bigint(50)                DEFAULT NULL COMMENT '更新时间',
    `create_time` bigint(50)                DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`blog_id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 11
  DEFAULT CHARSET = utf8 COMMENT ='博客表';

DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag`
(
    `id`          int(64) unsigned NOT NULL AUTO_INCREMENT,
    `user_id`     bigint          NOT NULL default 0 comment '用户id',
    `blog_id`     bigint          NOT NULL COMMENT '博客id',
    `tag_name`    varchar(100)              DEFAULT NULL COMMENT '标签题目',
    `update_time` bigint(50)                DEFAULT NULL COMMENT '更新时间',
    `create_time` bigint(50)                DEFAULT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  AUTO_INCREMENT = 11
  DEFAULT CHARSET = utf8 COMMENT ='标签表';


