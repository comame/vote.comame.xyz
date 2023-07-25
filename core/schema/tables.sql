CREATE TABLE `topic` (
    `id` varchar(36) PRIMARY KEY NOT NULL,
    `name` text NOT NULL,
    `type` varchar(16) NOT NULL
);

CREATE TABLE `choice_generic` (
    `id` varchar(36) PRIMARY KEY NOT NULL,
    `topic_id` varchar(36) NOT NULL,
    `order` INT UNSIGNED NOT NULL,
    `text` TEXT,
    FOREIGN KEY (`topic_id`) REFERENCES `topic`(`id`) ON DELETE CASCADE
);

CREATE TABLE `choice_calendar` (
    `id` varchar(36) PRIMARY KEY NOT NULL,
    `topic_id` varchar(36) NOT NULL,
    `order` INT UNSIGNED NOT NULL,
    `is_all_day` TINYINT UNSIGNED NOT NULL,
    `start_datetime` DATETIME NOT NULL,
    `end_datetime` DATETIME NOT NULL,
    FOREIGN KEY (`topic_id`) REFERENCES `topic`(`id`) ON DELETE CASCADE
);

CREATE TABLE `vote_generic` (
    `id` varchar(36) PRIMARY KEY NOT NULL,
    `user_name` text NOT NULL,
    `choice_id` varchar(36) NOT NULL,
    FOREIGN KEY (`choice_id`) REFERENCES `choice_generic`(`id`)
);

CREATE TABLE `vote_calendar` (
    `id` varchar(36) PRIMARY KEY NOT NULL,
    `user_name` text NOT NULL,
    `choice_id` varchar(36) NOT NULL,
    FOREIGN KEY (`choice_id`) REFERENCES `choice_generic`(`id`)
);
