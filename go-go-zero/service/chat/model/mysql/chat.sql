CREATE TABLE IF NOT EXISTS `chat_room`
(
    `id`        BIGINT       NOT NULL COMMENT 'id',
    `cmdty_id`  BIGINT       NOT NULL COMMENT '对应的商品信息',
    `seller_id` BIGINT       NOT NULL COMMENT '卖家id',
    `seller`    VARCHAR(20)  NOT NULL,
    `buyer_id`  BIGINT       NOT NULL,
    `buyer`     VARCHAR(20)  NOT NULL,
    `cover`     VARCHAR(255) NOT NULL,
    `create_at` DATETIME     NOT NULL,
    `status`    TINYINT      NOT NULL default 1,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE UNIQUE INDEX `cid_sid_bid_idx` ON `chat_room` (cmdty_id, seller_id, buyer_id);