CREATE DATABASE IF NOT EXISTS tinyurl character set utf8 collate utf8_general_ci;

CREATE TABLE `urls` (
    `id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'UUID',
    `hash` VARCHAR(11) NOT NULL COMMENT '短網址',
    `origin` VARCHAR(220) NOT NULL COMMENT '原始網址',
    `created_at` DATETIME NOT NULL COMMENT '短網址建立時間',
    `expires_at` DATETIME NOT NULL COMMENT '短網址有效時間',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT = '網址資訊';