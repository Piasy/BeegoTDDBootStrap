CREATE DATABASE IF NOT EXISTS `test` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- --------------------------------------------------
--  Table Structure for `github.com/Piasy/BeegoTDDBootStrap/models.User`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `uid` bigint NOT NULL DEFAULT 0 ,
  `username` varchar(255) NOT NULL DEFAULT '' ,
  `password` varchar(255) NOT NULL DEFAULT '' ,
  `nickname` varchar(255) NOT NULL DEFAULT '' ,
  `token` varchar(255) NOT NULL DEFAULT '' ,
  `phone` varchar(255) NOT NULL DEFAULT ''
) ENGINE=InnoDB;

-- --------------------------------------------------
--  Table Structure for `github.com/Piasy/BeegoTDDBootStrap/models.Verification`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `verifications` (
  `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `phone` varchar(255) NOT NULL DEFAULT '' ,
  `code` varchar(255) NOT NULL DEFAULT '' ,
  `expire` bigint NOT NULL DEFAULT 0
) ENGINE=InnoDB;
