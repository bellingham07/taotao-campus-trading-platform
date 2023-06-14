CREATE TABLE IF NOT EXISTS `trade_info`
(
    `id`             BIGINT           NOT NULL COMMENT 'id',
    `buyer_id`       BIGINT           NOT NULL COMMENT '卖家id',
    `buyer`          VARCHAR(20)      NOT NULL COMMENT '卖家名',
    `seller_id`      BIGINT           NOT NULL COMMENT '买家id',
    `seller`         VARCHAR(20)      NOT NULL COMMENT '买家名',
    `cmdty_id`       BIGINT           NOT NULL COMMENT '商品id',
    `brief_intro`    VARCHAR(20)      NOT NULL COMMENT '商品名',
    `cover`          VARCHAR(255)     NOT NULL,
    `location`       VARCHAR(20)      NOT NULL,
    `payment`        DOUBLE PRECISION NOT NULL,
    `status`         TINYINT          NOT NULL,
    `create_at`      DATETIME         NOT NULL COMMENT '创建时间',
    `is_seller_done` DATETIME         NOT NULL COMMENT '默认0，完成1',
    `is_buyer_done`  DATETIME         NOT NULL COMMENT '默认0，完成1',
    `seller_done_at` DATETIME         NOT NULL,
    `buyer_done_at`  DATETIME         NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `trade_done`
(
    `id`             BIGINT           NOT NULL COMMENT 'id',
    `buyer_id`       BIGINT           NOT NULL COMMENT '卖家id',
    `buyer`          VARCHAR(20)      NOT NULL COMMENT '卖家名',
    `seller_id`      BIGINT           NOT NULL COMMENT '买家id',
    `seller`         VARCHAR(20)      NOT NULL COMMENT '买家名',
    `cmdty_id`       BIGINT           NOT NULL COMMENT '商品id',
    `brief_intro`    VARCHAR(20)      NOT NULL COMMENT '商品名',
    `cover`          VARCHAR(255)     NOT NULL,
    `location`       VARCHAR(20)      NOT NULL,
    `payment`        DOUBLE PRECISION NOT NULL,
    `seller_done_at` DATETIME         NOT NULL,
    `buyer_done_at`  DATETIME         NOT NULL,
    `done_at`        DATETIME         NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `trade_cmt`
(
    `id`           BIGINT       NOT NULL COMMENT 'id',
    `trade_id`     BIGINT       NOT NULL,
    `buyer_id`     BIGINT       NOT NULL,
    `buyer`        VARCHAR(20)  NOT NULL,
    `buyer_cover`  VARCHAR(255) NOT NULL,
    `seller_id`    BIGINT       NOT NULL,
    `seller`       VARCHAR(20)  NOT NULL,
    `seller_cover` VARCHAR(255) NOT NULL,
    `content`      VARCHAR(200) NOT NULL COMMENT '评价内容',
    `type`         TINYINT      NOT NULL COMMENT '差评或好评，0为差评，1为好评',
    `create_at`    DATETIME     NOT NULL COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;