USE `taotao_trading_user`;

DROP TABLE IF EXISTS `user_info`;
DROP TABLE IF EXISTS `user_call`;
DROP TABLE IF EXISTS `user_location`;
DROP TABLE IF EXISTS `user_follow`;
DROP TABLE IF EXISTS `user_opt`;

CREATE TABLE IF NOT EXISTS `user_info`
(
    `id`       BIGINT       NOT NULL COMMENT 'id(snowflake)',
    `username` VARCHAR(20)  NOT NULL UNIQUE COMMENT '账户',
    `password` VARCHAR(255) NOT NULL COMMENT '密码',
    `name`     VARCHAR(20)  NOT NULL COMMENT '姓名x',
    `gender`   TINYINT      NOT NULL DEFAULT 3 COMMENT '性别',
    `phone`    VARCHAR(11)  NOT NULL DEFAULT '' COMMENT '手机号',
    `avatar`   VARCHAR(255) NOT NULL DEFAULT '' COMMENT '头像url',
    `intro`    VARCHAR(200) NOT NULL DEFAULT '' COMMENT '个人简介',
    `location` VARCHAR(50)  NOT NULL DEFAULT '' COMMENT '住址',
    `like`     BIGINT       NOT NULL DEFAULT 0 COMMENT '获赞数',
    `status`   TINYINT      NOT NULL DEFAULT 0 COMMENT '用户账户状态',
    `done`     BIGINT       NOT NULL DEFAULT 0 COMMENT '成交数',
    `call`     VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '称号',
    `fans`     BIGINT       NOT NULL DEFAULT 0 COMMENT '粉丝数',
    `follow`   BIGINT       NOT NULL DEFAULT 0 COMMENT '关注数',
    `positive` BIGINT       NOT NULL DEFAULT 0 COMMENT '好评次数',
    `negative` BIGINT       NOT NULL DEFAULT 0 COMMENT '差评次数',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `user_call`
(
    `id`        INT         NOT NULL AUTO_INCREMENT COMMENT 'id',
    `name`      VARCHAR(20) NOT NULL COMMENT '称号名字',
    `create_by` BIGINT      NOT NULL DEFAULT 1 COMMENT '管理员的id',
    `create_at` DATETIME    NOT NULL DEFAULT NOW() COMMENT '创建时间',
    `update_by` BIGINT      NOT NULL DEFAULT 1 COMMENT '管理员的id',
    `update_at` DATETIME    NOT NULL DEFAULT NOW() COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `user_follow`
(
    `id`             BIGINT   NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`        BIGINT   NOT NULL COMMENT '用户id',
    `follow_user_id` BIGINT   NOT NULL COMMENT '收藏的用户id',
    `create_at`      DATETIME NOT NULL DEFAULT NOW() COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `user_location`
(
    `id`        INT         NOT NULL COMMENT 'id',
    `name`      VARCHAR(20) NOT NULL COMMENT '地址名',
    `create_by` BIGINT      NOT NULL DEFAULT 1 COMMENT '管理员的id',
    `create_at` DATETIME    NOT NULL DEFAULT NOW() COMMENT '创建时间',
    `update_by` BIGINT      NOT NULL DEFAULT 1 COMMENT '管理员的id',
    `update_at` DATETIME    NOT NULL DEFAULT NOW() COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `user_opt`
(
    `id`       BIGINT            NOT NULL,
    `username` VARCHAR(20)       NOT NULL,
    `password` VARCHAR(255)      NOT NULL,
    `name`     VARCHAR(20)       NOT NULL,
    `status`   TINYINT DEFAULT 1 NOT NULL,
    `level`    TINYINT DEFAULT 1 NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;