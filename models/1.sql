CREATE TABLE IF NOT EXISTS `user` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT ,
  `uid` BIGINT NOT NULL ,
  `phone` VARCHAR(11) NOT NULL ,
  `username` VARCHAR(20) NOT NULL ,
  `password` VARCHAR(12) NOT NULL ,
  `nickname` VARCHAR(40) NOT NULL ,
  `token` VARCHAR(40) NOT NULL
)