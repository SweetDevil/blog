/*
Navicat MySQL Data Transfer

Source Server         : 本机
Source Server Version : 50617
Source Host           : 127.0.0.1:3306
Source Database       : th

Target Server Type    : MYSQL
Target Server Version : 50617
File Encoding         : 65001

Date: 2015-05-06 10:54:48
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',
  `title` char(255) NOT NULL COMMENT 'article title',
  `ctime` int(10) unsigned NOT NULL COMMENT 'the article created time',
  `content` text NOT NULL COMMENT 'article content',
  `etime` int(10) unsigned NOT NULL COMMENT 'article last edit time',
  `thumb` varchar(255) DEFAULT NULL COMMENT 'article thumb image',
  `uid` int(10) unsigned NOT NULL COMMENT 'editor of this article',
  `catid` int(10) unsigned DEFAULT NULL COMMENT 'category id ,which category this article belongs to',
  `type` tinyint(3) unsigned DEFAULT NULL COMMENT 'article type, 1 means normal, 2 means gallery',
  PRIMARY KEY (`id`),
  KEY `title` (`title`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of article
-- ----------------------------

-- ----------------------------
-- Table structure for attachment
-- ----------------------------
DROP TABLE IF EXISTS `attachment`;
CREATE TABLE `attachment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pk` int(10) unsigned NOT NULL COMMENT 'foreign table primary key',
  `table` varchar(50) NOT NULL COMMENT 'foreign table name',
  `path` varchar(255) DEFAULT NULL COMMENT 'attachment sotrage path',
  `alias` varchar(100) DEFAULT NULL COMMENT 'alias of attachment, for download name',
  `ext` varchar(20) DEFAULT NULL COMMENT 'extension name of attachment',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of attachment
-- ----------------------------

-- ----------------------------
-- Table structure for category
-- ----------------------------
DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'category id ,primary key',
  `name` varchar(255) NOT NULL COMMENT 'name of category',
  `pid` int(10) unsigned DEFAULT NULL COMMENT 'the parent category id, for make tree list',
  `type` tinyint(3) unsigned DEFAULT '1' COMMENT 'type for category, 1 means normal, 2 means single page',
  `url` varchar(255) DEFAULT NULL COMMENT 'custom  link for category',
  `sort` smallint(5) unsigned DEFAULT '0',
  `backgroud` varchar(255) DEFAULT NULL COMMENT 'background css code for this category, default is white color',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of category
-- ----------------------------

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',
  `uid` int(10) unsigned NOT NULL COMMENT 'user''s id of this comment',
  `content` text NOT NULL COMMENT 'content of this comment',
  `pid` int(10) unsigned NOT NULL COMMENT 'parent id, mean which comment this comment reply to',
  `reply_id` int(10) unsigned DEFAULT NULL COMMENT 'used to make level in level tree list',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of comment
-- ----------------------------

-- ----------------------------
-- Table structure for config
-- ----------------------------
DROP TABLE IF EXISTS `config`;
CREATE TABLE `config` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'primary key',
  `name` varchar(80) NOT NULL COMMENT 'name of config',
  `value` varchar(255) NOT NULL COMMENT 'value of config',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of config
-- ----------------------------

-- ----------------------------
-- Table structure for tag
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL COMMENT 'name of tag',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of tag
-- ----------------------------

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'user id, primary key',
  `name` char(50) NOT NULL COMMENT 'user name',
  `password` char(50) NOT NULL COMMENT 'user''s password',
  `nickname` varchar(50) DEFAULT NULL COMMENT 'user''s nickname',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT 'user status,  1 means not verify,just registed,2 means normal actived,3 means banded',
  `email` varchar(255) DEFAULT NULL COMMENT 'user''s email, used to verify',
  `regtime` int(10) unsigned NOT NULL COMMENT 'user registe time',
  `points` mediumint(8) unsigned DEFAULT NULL COMMENT 'user points ,used to get attachment which need points',
  PRIMARY KEY (`id`),
  KEY `username` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
