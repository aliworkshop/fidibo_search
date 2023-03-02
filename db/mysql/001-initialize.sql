alter
schema fidibo collate utf8_general_ci;
CREATE TABLE `books`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `title`      varchar(50) NOT NULL,
    `content`    MEDIUMTEXT NULL,
    `slug`       varchar(50) NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `publishers`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `name`       varchar(50) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `book_publisher`
(
    `book_id`      bigint unsigned NOT NULL,
    `publisher_id` bigint unsigned NOT NULL,
    UNIQUE KEY `book_publisher_unq_idx` (`book_id`, `publisher_id`),
    FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON UPDATE CASCADE,
    FOREIGN KEY (`publisher_id`) REFERENCES `publishers` (`id`) ON UPDATE CASCADE
);

CREATE TABLE `authors`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `created_at` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `name`       varchar(50) NOT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE `book_author`
(
    `book_id`      bigint unsigned NOT NULL,
    `author_id` bigint unsigned NOT NULL,
    UNIQUE KEY `book_author_unq_idx` (`book_id`, `author_id`),
    FOREIGN KEY (`book_id`) REFERENCES `books` (`id`) ON UPDATE CASCADE,
    FOREIGN KEY (`author_id`) REFERENCES `authors` (`id`) ON UPDATE CASCADE
);