CREATE TABLE `tasks` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(50) NULL,
    `done` tinyint(1) DEFAULT 0,
    `created_at` DATETIME(3) NOT NULL,
    `updated_at` DATETIME(3) NOT NULL,
    `deleted_at` DATETIME(3) NULL,
    PRIMARY KEY (`id`)
)