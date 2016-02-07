CREATE DATABASE IF NOT EXISTS `beego_unit_test` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;

USE beego_unit_test;

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
