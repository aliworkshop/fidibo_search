CREATE TABLE `users`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `first_name` varchar(50) NULL,
    `last_name`  varchar(50) NULL,
    `email`      varchar(50) NULL,
    `username`   varchar(50) NOT NULL,
    `password`   varchar(65) NOT NULL,
    `active`     tinyint(1),
    UNIQUE KEY `users_username_unq_idx` (`username`),
    PRIMARY KEY (`id`)
);