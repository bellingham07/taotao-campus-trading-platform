CREATE TABLE IF NOT EXISTS `atcl_content`
(
    `id`        BIGINT       NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`   BIGINT       NOT NULL COMMENT '用户的id',
    `title`     VARCHAR(50)  NOT NULL,
    `content`   VARCHAR(200) NOT NULL COMMENT '帖子内容',
    `cover`     VARCHAR(255) NOT NULL COMMENT '封面',
    `status`    TINYINT      NOT NULL COMMENT '1为草稿，2为发布，-1为审核不通过',
    `create_at` DATETIME     NOT NULL COMMENT '创建时间',
    `update_at` DATETIME     NOT NULL COMMENT '更新时间',
    `collect`   BIGINT       NOT NULL,
    `like`      BIGINT       NOT NULL,
    `view`      BIGINT       NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `atcl_bulletin`
(
    `id`        INT          NOT NULL AUTO_INCREMENT COMMENT 'id',
    `content`   VARCHAR(255) NOT NULL COMMENT '内容',
    `admin_id`  BIGINT       NOT NULL COMMENT '发布管理员id',
    `create_at` DATETIME     NOT NULL COMMENT '发布时间',
    `update_at` DATETIME     NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `atcl_collect`
(
    `id`         BIGINT   NOT NULL AUTO_INCREMENT COMMENT 'id',
    `user_id`    BIGINT   NOT NULL COMMENT '用户id',
    `article_id` BIGINT   NOT NULL COMMENT '商品id',
    `status`     tinyint  NOT NULL COMMENT '1存在，0失效',
    `create_at`  datetime NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `atcl_cmt`
(
    `id`         BIGINT       NOT NULL AUTO_INCREMENT COMMENT 'id',
    `article_id` BIGINT       NOT NULL,
    `user_id`    BIGINT       NOT NULL COMMENT '留言的用户id',
    `content`    VARCHAR(200) NOT NULL COMMENT '留言内容',
    `root_id`    BIGINT       NOT NULL COMMENT '根留言',
    `to_user_id` BIGINT       NOT NULL COMMENT '回复给到用户',
    `create_at`  datetime     NOT NULL COMMENT '评论时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;