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

 Date: 06/08/2017 22:22:36
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
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `city_code` varchar(2) COLLATE utf8mb4_bin NOT NULL,
  `city_name` varchar(20) COLLATE utf8mb4_bin NOT NULL,
  `station_code` varchar(4) COLLATE utf8mb4_bin NOT NULL,
  `station_name` varchar(10) COLLATE utf8mb4_bin NOT NULL,
  `gate_code` varchar(6) COLLATE utf8mb4_bin NOT NULL,
  `direction` varchar(3) COLLATE utf8mb4_bin NOT NULL COMMENT 'in-入;out-出',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for sg_router_evidence
-- ----------------------------
DROP TABLE IF EXISTS `sg_router_evidence`;
CREATE TABLE `sg_router_evidence` (
  `evidence_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `user_id` varchar(32) COLLATE utf8mb4_bin NOT NULL,
  `type` smallint(1) NOT NULL DEFAULT 0 COMMENT '0-通用;1-入阐;2-出阐',
  `create_time` datetime NOT NULL,
  `expires_at` datetime NOT NULL,
  `status` smallint(1) NOT NULL COMMENT '1-已下发;2-已使用;3-已过期',
  `used_time` datetime DEFAULT NULL COMMENT '使用时间',
  PRIMARY KEY (`evidence_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for sg_router_info
-- ----------------------------
DROP TABLE IF EXISTS `sg_router_info`;
CREATE TABLE `sg_router_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` varchar(128) COLLATE utf8mb4_bin NOT NULL,
  `in_time` datetime DEFAULT NULL,
  `in_ evidence` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL,
  `in_station_code` varchar(4) COLLATE utf8mb4_bin DEFAULT NULL,
  `in_station_name` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL,
  `in_gate_code` varchar(6) COLLATE utf8mb4_bin DEFAULT NULL,
  `out_time` datetime DEFAULT NULL,
  `out_ evidence` varchar(32) COLLATE utf8mb4_bin DEFAULT NULL,
  `out_station_code` varchar(4) COLLATE utf8mb4_bin DEFAULT NULL,
  `out_station_name` varchar(20) COLLATE utf8mb4_bin DEFAULT NULL,
  `out_gate_code` varchar(6) COLLATE utf8mb4_bin DEFAULT NULL,
  `status` tinyint(1) NOT NULL COMMENT '1-进行中;2-结束;3-无出站异常;4-无入站异常',
  `exception_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
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

SET FOREIGN_KEY_CHECKS = 1;
