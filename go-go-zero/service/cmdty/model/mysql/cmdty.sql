CREATE TABLE IF NOT EXISTS `cmdty_info`
(
    `id`         BIGINT       NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`    BIGINT       NOT NULL COMMENT '用户id',
    `cover`      VARCHAR(255) NOT NULL DEFAULT '' COMMENT '封面图片',
    `tag`        VARCHAR(255) NOT NULL DEFAULT '' COMMENT '分类名',
    `price`      DOUBLE       NOT NULL DEFAULT 0 COMMENT '商品价格',
    `brand`      VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '品牌',
    `model`      VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '型号',
    `intro`      VARCHAR(100) NOT NULL DEFAULT '' COMMENT '商品介绍',
    `old`        VARCHAR(10)  NOT NULL DEFAULT '轻微使用痕迹' COMMENT '新旧程度',
    `status`     TINYINT      NOT NULL DEFAULT 1 COMMENT '商品状态，默认1为草稿，2为发布，0为下架，-1为审核未通过需修改',
    `create_at`  DATETIME     NOT NULL DEFAULT NOW() COMMENT '创建时间',
    `publish_at` DATETIME     NOT NULL DEFAULT NOW() COMMENT '发布时间',
    `view`       BIGINT       NOT NULL DEFAULT 0 COMMENT '查看数量',
    `collect`    BIGINT       NOT NULL DEFAULT 0 COMMENT '收藏数',
    `type`       TINYINT      NOT NULL COMMENT '1为售卖商品，2为收商品',
    `like`       BIGINT       NOT NULL DEFAULT 0 COMMENT '点赞数',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `cmdty_tag`
(
    `id`        INT         NOT NULL AUTO_INCREMENT COMMENT '分类ID编号',
    `name`      VARCHAR(20) NOT NULL COMMENT '分类名称',
    `create_by` BIGINT      NOT NULL DEFAULT 1 COMMENT '管理员的id',
    `create_at` DATETIME    NOT NULL DEFAULT NOW() COMMENT '创建时间',
    `update_by` BIGINT      NOT NULL DEFAULT 1 COMMENT '管理员的id',
    `update_at` DATETIME    NOT NULL DEFAULT NOW() COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `cmdty_collect`
(
    `id`        BIGINT       NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   BIGINT       NOT NULL COMMENT '用户id',
    `cmdty_id`  BIGINT       NOT NULL COMMENT '商品id',
    `intro`     VARCHAR(20)  NOT NULL DEFAULT '' COMMENT '20字的简介',
    `cover`     VARCHAR(255) NOT NULL DEFAULT '',
    `price`     DOUBLE       NOT NULL DEFAULT 0,
    `status`    TINYINT      NOT NULL DEFAULT 1 COMMENT '1存在，0失效',
    `create_at` DATETIME     NOT NULL DEFAULT NOW() COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `cmdty_cmt`
(
    `id`         BIGINT       NOT NULL AUTO_INCREMENT COMMENT 'id',
    `cmdty_id`   BIGINT       NOT NULL COMMENT '对应的商品id',
    `user_id`    BIGINT       NOT NULL COMMENT '留言的用户id',
    `content`    VARCHAR(100) NOT NULL COMMENT '留言内容',
    `root_id`    BIGINT       NOT NULL COMMENT '根留言',
    `to_user_id` BIGINT       NOT NULL COMMENT '回复给到用户',
    `create_at`  DATETIME     NOT NULL DEFAULT NOW() COMMENT '评论时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;