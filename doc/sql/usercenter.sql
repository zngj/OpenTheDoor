/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50719
 Source Host           : localhost:3306
 Source Schema         : usercenter

 Target Server Type    : MySQL
 Target Server Version : 50719
 File Encoding         : 65001

 Date: 18/08/2017 17:25:26
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for uc_login_log
-- ----------------------------
DROP TABLE IF EXISTS `uc_login_log`;
CREATE TABLE `uc_login_log` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(128) COLLATE utf8mb4_bin NOT NULL,
  `access_token` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `login_time` datetime NOT NULL,
  `expires_in` int(11) NOT NULL,
  `expires_at` datetime NOT NULL,
  `release_time` datetime DEFAULT NULL,
  `status` varchar(1) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for uc_user_info
-- ----------------------------
DROP TABLE IF EXISTS `uc_user_info`;
CREATE TABLE `uc_user_info` (
  `id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `channel` varchar(10) COLLATE utf8mb4_bin NOT NULL,
  `open_id` varchar(128) COLLATE utf8mb4_bin DEFAULT NULL,
  `nick_name` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL,
  `phone_number` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL,
  `email` varchar(50) COLLATE utf8mb4_bin DEFAULT NULL,
  `password` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `signature` varchar(100) COLLATE utf8mb4_bin DEFAULT NULL,
  `avatar` varchar(200) COLLATE utf8mb4_bin DEFAULT NULL,
  `sex` tinyint(1) DEFAULT NULL,
  `insert_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

SET FOREIGN_KEY_CHECKS = 1;
