CREATE TABLE `student` (
    `id` bigint(255) NOT NULL AUTO_INCREMENT COMMENT ' ',
    `user_name` varchar(255) NOT NULL DEFAULT '',
    `class` varchar(255) NOT NULL DEFAULT '',
    `subject` varchar(255) NOT NULL DEFAULT '',
    `score` float NOT NULL DEFAULT 0,
    PRIMARY KEY (`id`),
    KEY `class_name` (`class`,`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;