/*
 Navicat MySQL Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50719
 Source Host           : localhost:3306
 Source Schema         : smartgate

 Target Server Type    : MySQL
 Target Server Version : 50719
 File Encoding         : 65001

 Date: 18/08/2017 17:25:16
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sg_codec_key
-- ----------------------------
DROP TABLE IF EXISTS `sg_codec_key`;
CREATE TABLE `sg_codec_key` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `code` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `rsa_private` text COLLATE utf8mb4_bin NOT NULL,
  `rsa_public` text COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for sg_gate_info
-- ----------------------------
DROP TABLE IF EXISTS `sg_gate_info`;
CREATE TABLE `sg_gate_info` (
  `id` varchar(10) COLLATE utf8mb4_bin NOT NULL,
  `direction` tinyint(1) NOT NULL COMMENT '0-入;1-出',
  `station_id` varchar(10) COLLATE utf8mb4_bin NOT NULL,
  `station_name` varchar(10) COLLATE utf8mb4_bin NOT NULL,
  `city_id` varchar(10) COLLATE utf8mb4_bin NOT NULL,
  `city_name` varchar(20) COLLATE utf8mb4_bin NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of sg_gate_info
-- ----------------------------
BEGIN;
INSERT INTO `sg_gate_info` VALUES ('010100101', 0, '0101001', '五一广场', '0101', '长沙');
INSERT INTO `sg_gate_info` VALUES ('010100102', 1, '0101001', '王一广场', '0101', '长沙');
INSERT INTO `sg_gate_info` VALUES ('010100201', 0, '0101002', '黄兴广场', '0101', '长沙');
INSERT INTO `sg_gate_info` VALUES ('010100202', 1, '0101002', '黄兴广场', '0101', '长沙');
COMMIT;

-- ----------------------------
-- Table structure for sg_router_evidence
-- ----------------------------
DROP TABLE IF EXISTS `sg_router_evidence`;
CREATE TABLE `sg_router_evidence` (
  `evidence_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `direction` smallint(1) NOT NULL DEFAULT '0' COMMENT '0-入阐;1-出阐;3-通用',
  `create_time` datetime NOT NULL,
  `expires_at` datetime NOT NULL,
  `status` tinyint(1) NOT NULL COMMENT '1-有效;2-已使用;3-已过期;4-已废弃',
  `update_time` datetime DEFAULT NULL COMMENT '使用时间',
  PRIMARY KEY (`evidence_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for sg_router_info
-- ----------------------------
DROP TABLE IF EXISTS `sg_router_info`;
CREATE TABLE `sg_router_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(128) COLLATE utf8mb4_bin NOT NULL,
  `at_date` date NOT NULL,
  `group_no` smallint(1) NOT NULL,
  `in_station_id` varchar(10) COLLATE utf8mb4_bin DEFAULT NULL,
  `in_station_name` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL,
  `in_gate_id` varchar(10) COLLATE utf8mb4_bin DEFAULT NULL,
  `in_evidence` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL,
  `in_time` datetime DEFAULT NULL,
  `out_station_id` varchar(10) COLLATE utf8mb4_bin DEFAULT NULL,
  `out_station_name` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL,
  `out_gate_id` varchar(10) COLLATE utf8mb4_bin DEFAULT NULL,
  `out_evidence` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL,
  `out_time` datetime DEFAULT NULL,
  `money` decimal(4,2) DEFAULT NULL,
  `status` tinyint(1) NOT NULL COMMENT '1-入闸;2-出闸;3-出闸(无入闸);4-未出闸(异常);5-未入闸(异常)',
  `paid` tinyint(1) NOT NULL DEFAULT '0',
  `exception_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for sg_sys_notification
-- ----------------------------
DROP TABLE IF EXISTS `sg_sys_notification`;
CREATE TABLE `sg_sys_notification` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `category` varchar(10) COLLATE utf8mb4_bin NOT NULL,
  `content_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `received` tinyint(1) NOT NULL DEFAULT '0',
  `insert_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for sg_wallet_charge
-- ----------------------------
DROP TABLE IF EXISTS `sg_wallet_charge`;
CREATE TABLE `sg_wallet_charge` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for sg_wallet_info
-- ----------------------------
DROP TABLE IF EXISTS `sg_wallet_info`;
CREATE TABLE `sg_wallet_info` (
  `user_id` varchar(128) COLLATE utf8mb4_bin NOT NULL,
  `balance` float NOT NULL DEFAULT '0',
  `wxpay_quick` tinyint(1) NOT NULL DEFAULT '0',
  `alipay_quick` tinyint(1) NOT NULL DEFAULT '0',
  `insert_time` datetime NOT NULL,
  `update_time` datetime DEFAULT NULL,
  PRIMARY KEY (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Procedure structure for ClearData
-- ----------------------------
DROP PROCEDURE IF EXISTS `ClearData`;
delimiter ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `ClearData`()
BEGIN
  #Routine body goes here...
	TRUNCATE sg_router_evidence;
	TRUNCATE sg_router_info;
	TRUNCATE sg_sys_notification;
	TRUNCATE sg_wallet_info;

END;
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
