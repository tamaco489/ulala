CREATE TABLE `movie_types` (
  `type_id` INT(2) UNSIGNED NOT NULL AUTO_INCREMENT,
  `type_name` VARCHAR(50) NOT NULL,
  `title` VARCHAR(100) NOT NULL,
  `description` VARCHAR(200) NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`type_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `movie_formats` (
  `format_id` INT(2) UNSIGNED NOT NULL AUTO_INCREMENT,
  `movie_format` VARCHAR(50) NOT NULL UNIQUE,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`format_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

CREATE TABLE `movies` (
  `movie_id` INT(8) UNSIGNED NOT NULL,
  `title` VARCHAR(100) NOT NULL,
  `release_year` INT(4) UNSIGNED NOT NULL,
  `description` VARCHAR(200) NOT NULL,
  `type_id` INT(2) UNSIGNED NOT NULL,
  `format_id` INT(2) UNSIGNED NOT NULL,
  `image_id` INT(8) UNSIGNED NOT NULL UNIQUE,
  `thumbnail_id` INT(8) UNSIGNED NOT NULL UNIQUE,
  `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`movie_id`),
  FOREIGN KEY (`type_id`) REFERENCES `movie_types` (`type_id`),
  FOREIGN KEY (`format_id`) REFERENCES `movie_formats` (`format_id`)
) AUTO_INCREMENT = 10000001 ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
