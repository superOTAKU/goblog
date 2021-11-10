DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
    `user_id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `email` VARCHAR(255) NOT NULL,
    `password_hash` VARCHAR(255) NOT NULL,
    `password_solt` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP,
    UNIQUE INDEX `email_idx`(`email`)
);

DROP TABLE IF EXISTS `blogs`;
CREATE TABLE `blogs` (
    `blog_id` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `user_id` BIGINT NOT NULL,
    `content` TEXT NOT NULL,
    `created_at` TIMESTAMP NOT NULL,
    `updated_at` TIMESTAMP,
    INDEX `user_id_idx`(`user_id`)
);
