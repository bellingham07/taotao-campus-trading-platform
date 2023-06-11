USE taotao_trading_file;

CREATE TABLE IF NOT EXISTS `file_avatar`
(
    `id`         BIGINT AUTO_INCREMENT NOT NULL,
    `user_id`    BIGINT                NOT NULL,
    `url`        VARCHAR(255)          NOT NULL,
    `objectName` VARCHAR(255)          NOT NULL,
    `upload_at`  DATETIME              NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `file_cmdty`
(
    `id`         BIGINT AUTO_INCREMENT NOT NULL COMMENT 'bigint自增',
    `cmdty_id`   BIGINT                NOT NULL,
    `user_id`    BIGINT                NOT NULL,
    `url`        VARCHAR(255)          NOT NULL,
    `objectName` VARCHAR(255)          NOT NULL,
    `upload_at`  DATETIME              NOT NULL,
    `is_cover`   TINYINT DEFAULT 0     NOT NULL COMMENT '默认为0，封面为1',
    `order`      INT     DEFAULT 1     NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE IF NOT EXISTS `file_atcl`
(
    `id`         BIGINT AUTO_INCREMENT NOT NULL COMMENT 'bigint自增',
    `atcl_id`    BIGINT                NOT NULL,
    `user_id`    BIGINT                NOT NULL,
    `url`        VARCHAR(255)          NOT NULL,
    `objectName` VARCHAR(255)          NOT NULL,
    `upload_at`  DATETIME              NOT NULL,
    `is_cover`   TINYINT DEFAULT 0     NOT NULL COMMENT '默认为0，封面为1',
    `order`      INT     DEFAULT 1     NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;