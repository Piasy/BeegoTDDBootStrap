CREATE DATABASE IF NOT EXISTS `beego_unit_test` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE beego_unit_test;

-- --------------------------------------------------
--  Table Structure for `github.com/Piasy/BeegoTDDBootStrap/models.User`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `users` (
  `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `uid` bigint NOT NULL DEFAULT 0  UNIQUE,
  `token` varchar(40) NOT NULL DEFAULT ''  UNIQUE,
  `phone` varchar(11) UNIQUE,
  `weixin` varchar(191) UNIQUE,
  `weibo` varchar(191) UNIQUE,
  `qq` varchar(191) UNIQUE,
  `password` varchar(40) NOT NULL DEFAULT '' ,
  `nickname` varchar(12) NOT NULL DEFAULT '' ,
  `qq_nickname` varchar(127) NOT NULL DEFAULT '' ,
  `weibo_nickname` varchar(127) NOT NULL DEFAULT '' ,
  `weixin_nickname` varchar(127) NOT NULL DEFAULT '' ,
  `gender` integer NOT NULL DEFAULT 0 ,
  `avatar` varchar(191) NOT NULL DEFAULT '' ,
  `create_at` bigint NOT NULL DEFAULT 0  UNIQUE,
  `update_at` bigint NOT NULL DEFAULT 0
) ENGINE=InnoDB;

-- --------------------------------------------------
--  Table Structure for `github.com/Piasy/BeegoTDDBootStrap/models.Verification`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `verifications` (
  `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
  `phone` varchar(20) NOT NULL DEFAULT ''  UNIQUE,
  `code` varchar(6) NOT NULL DEFAULT '' ,
  `expire` bigint NOT NULL DEFAULT 0
) ENGINE=InnoDB;
