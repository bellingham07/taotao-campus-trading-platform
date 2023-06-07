CREATE TABLE IF NOT EXISTS `user_info`
(
    `id`         bigint       NOT NULL COMMENT 'id(snowflake)',
    `username`   VARCHAR(20)  NOT NULL COMMENT '账户',
    `password`   VARCHAR(20)  NOT NULL COMMENT '密码',
    `name`       VARCHAR(20)  NOT NULL COMMENT '姓名x',
    `gender`     tinyint      NOT NULL COMMENT '性别',
    `phone`      VARCHAR(11)  NOT NULL COMMENT '手机号',
    `avatar`     VARCHAR(255) NOT NULL COMMENT '头像url',
    `intro`      VARCHAR(200) NOT NULL COMMENT '个人简介',
    `location`   VARCHAR(50)  NOT NULL COMMENT '住址',
    `last_login` datetime     NOT NULL COMMENT '上次登录时间',
    `like`       bigint       NOT NULL COMMENT '获赞数',
    `status`     tinyint      NOT NULL COMMENT '用户账户状态',
    `done`       bigint       NOT NULL COMMENT '成交数',
    `call`       VARCHAR(20)  NOT NULL COMMENT '称号',
    `fans`       bigint       NOT NULL COMMENT '粉丝数',
    `follow`     bigint       NOT NULL COMMENT '关注数',
    `positive`   bigint       NOT NULL COMMENT '好评次数',
    `negative`   bigint       NOT NULL COMMENT '差评次数',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `user_call`
(
    `id`        int         NOT NULL AUTO_INCREMENT COMMENT 'id',
    `name`      VARCHAR(20) NOT NULL COMMENT '称号名字',
    `create_by` datetime    NOT NULL COMMENT '管理员的id',
    `create_at` datetime    NOT NULL COMMENT '创建时间',
    `update_by` datetime    NOT NULL COMMENT '管理员的id',
    `update_at` datetime    NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `user_follow`
(
    `id`             bigint   NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`        bigint   NOT NULL COMMENT '用户id',
    `follow_user_id` bigint   NOT NULL COMMENT '收藏的用户id',
    `create_at`      DATETIME NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `user_dorm`
(
    `id`        int         NOT NULL COMMENT 'id',
    `name`      VARCHAR(20) NOT NULL COMMENT '地址名',
    `create_by` datetime    NOT NULL COMMENT '管理员的id',
    `create_at` datetime    NOT NULL COMMENT '创建时间',
    `update_by` datetime    NOT NULL COMMENT '管理员的id',
    `update_at` datetime    NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `user_opt`
(
    `id`       BIGINT      NOT NULL,
    `username` VARCHAR(20) NOT NULL,
    `password` VARCHAR(20) NOT NULL,
    `name`     VARCHAR(20) NOT NULL,
    `status`   TINYINT     NOT NULL,
    `level`    TINYINT     NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;