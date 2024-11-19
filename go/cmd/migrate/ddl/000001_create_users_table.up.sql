CREATE TABLE `users` (
  `uid` INT(8) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) DEFAULT 'anonymous',
  `email` VARCHAR(50) NOT NULL UNIQUE,
  `auth_uuid` VARCHAR(50) NOT NULL UNIQUE,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`uid`),
  INDEX `uid_index` (`uid`),
  INDEX `auth_uuid_index` (`auth_uuid`)
) AUTO_INCREMENT = 10000001 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
