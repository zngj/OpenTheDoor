CREATE DATABASE `usercenter` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

use `usercenter`;

CREATE TABLE `uc_user_info` (
  `id` varchar(32) NOT NULL,
  `channel` varchar(10) NOT NULL,
  `open_id` varchar(128) NOT NULL,
  `insert_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;


CREATE TABLE `uc_login_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(128) NOT NULL,
  `access_token` varchar(32) NOT NULL,
  `login_time` datetime NOT NULL,
  `expires_in` int(11) NOT NULL,
  `expires_time` datetime NOT NULL,
  `release_time` datetime DEFAULT NULL,
  `status` varchar(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;
