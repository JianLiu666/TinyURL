CREATE DATABASE IF NOT EXISTS tinyurl character set utf8 collate utf8_general_ci;

CREATE TABLE `url` (
    `id` INT NOT NULL AUTO_INCREMENT COMMENT '短網址',
    `md5` VARCHAR(32) DEFAULT NULL COMMENT '原始網址(md5)',
    `created_at` DATETIME NOT NULL COMMENT '短網址建立時間',
    `expired_at` DATETIME NOT NULL COMMENT '短網址有效時間',
    PRIMARY KEY (`id`)
) COMMENT = '網址連結';

USE tinyurl