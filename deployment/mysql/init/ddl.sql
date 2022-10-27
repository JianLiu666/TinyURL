CREATE DATABASE IF NOT EXISTS `tinyurl`;

DROP TABLE IF EXISTS `urls`;
CREATE TABLE `urls` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'UUID',
    `tiny` varchar(8) NOT NULL COMMENT '短網址',
    `origin` varchar(220) NOT NULL COMMENT '原始網址',
    `created_at` datetime NOT NULL COMMENT '短網址建立時間',
    `expires_at` datetime NOT NULL COMMENT '短網址有效時間',
    PRIMARY KEY (`id`),
    UNIQUE KEY `tiny` (`tiny`),
    UNIQUE KEY `origin` (`origin`),
    UNIQUE KEY `hash` (`tiny`,`origin`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='網址資訊';

-- Prometheus metrics
CREATE USER 'exporter'@'%' IDENTIFIED BY '123456';
GRANT PROCESS, REPLICATION CLIENT ON *.* TO 'exporter'@'%';
GRANT SELECT ON *.* TO 'exporter'@'%';