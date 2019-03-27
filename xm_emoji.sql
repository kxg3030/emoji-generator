/*
Navicat MySQL Data Transfer

Source Server         : 本地连接
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : xm_emoji

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-03-27 22:11:15
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for xm_sys_emoji_file
-- ----------------------------
DROP TABLE IF EXISTS `xm_sys_emoji_file`;
CREATE TABLE `xm_sys_emoji_file` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `extension` varchar(20) NOT NULL,
  `base_path` varchar(100) NOT NULL,
  `path` varchar(100) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `sentence_count` tinyint(3) NOT NULL,
  `sentence` text NOT NULL,
  `image_url` varchar(255) NOT NULL,
  `cover_url` varchar(100) NOT NULL,
  `md5_encode` varchar(38) NOT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for xm_user_emoji_file
-- ----------------------------
DROP TABLE IF EXISTS `xm_user_emoji_file`;
CREATE TABLE `xm_user_emoji_file` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `open_id` varchar(100) NOT NULL,
  `status` tinyint(1) NOT NULL DEFAULT '1',
  `image_url` varchar(100) DEFAULT NULL,
  `create_time` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for xm_user_list
-- ----------------------------
DROP TABLE IF EXISTS `xm_user_list`;
CREATE TABLE `xm_user_list` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `open_id` varchar(100) NOT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `avatar` varchar(100) NOT NULL,
  `nick_name` varchar(50) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;
