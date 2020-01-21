/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50562
Source Host           : localhost:3306
Source Database       : dotweb_cms

Target Server Type    : MYSQL
Target Server Version : 50562
File Encoding         : 65001

Date: 2019-07-04 05:29:26
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for dt_menu
-- ----------------------------
DROP TABLE IF EXISTS `dt_menu`;
CREATE TABLE `dt_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `record_id` varchar(255) NOT NULL DEFAULT '' COMMENT '记录 Id',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `sequence` int(11) NOT NULL DEFAULT '0' COMMENT '排序数值',
  `icon` varchar(255) NOT NULL DEFAULT '' COMMENT '图标',
  `router` varchar(255) NOT NULL DEFAULT '' COMMENT '访问路由',
  `hidden` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0 不隐藏 1隐藏',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '父级Id',
  `parent_path` varchar(255) NOT NULL DEFAULT '' COMMENT '父级路径',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='菜单实体表';

-- ----------------------------
-- Records of dt_menu
-- ----------------------------

-- ----------------------------
-- Table structure for dt_menu_action
-- ----------------------------
DROP TABLE IF EXISTS `dt_menu_action`;
CREATE TABLE `dt_menu_action` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `menu_id` varchar(255) NOT NULL DEFAULT '' COMMENT '菜单id',
  `code` varchar(255) NOT NULL DEFAULT '' COMMENT '动作编号',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '动作名称',
  `created_at` datetime DEFAULT NULL COMMENT '创建时间',
  `updated_at` datetime DEFAULT NULL COMMENT '更新时间',
  `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='菜单动作关联实体表';

-- ----------------------------
-- Records of dt_menu_action
-- ----------------------------

-- ----------------------------
-- Table structure for dt_menu__resource
-- ----------------------------
DROP TABLE IF EXISTS `dt_menu__resource`;
CREATE TABLE `dt_menu__resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `menu_id` varchar(255) NOT NULL DEFAULT '' COMMENT '菜单id',
  `code` varchar(255) NOT NULL DEFAULT '' COMMENT '资源编号',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '资源名称',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `method` varchar(255) NOT NULL DEFAULT '' COMMENT '资源请求方式',
  `path` varchar(255) NOT NULL DEFAULT '' COMMENT '资源请求路径',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='菜单资源关联实体';

-- ----------------------------
-- Records of dt_menu__resource
-- ----------------------------

-- ----------------------------
-- Table structure for dt_role
-- ----------------------------
DROP TABLE IF EXISTS `dt_role`;
CREATE TABLE `dt_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `record_id` varchar(255) NOT NULL DEFAULT '' COMMENT '记录id',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '角色名称',
  `sequence` int(11) NOT NULL DEFAULT '0' COMMENT '排序值',
  `memo` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `creator` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色实体';

-- ----------------------------
-- Records of dt_role
-- ----------------------------

-- ----------------------------
-- Table structure for dt_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `dt_role_menu`;
CREATE TABLE `dt_role_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL DEFAULT '0' COMMENT '角色id',
  `menu_id` int(11) NOT NULL DEFAULT '0' COMMENT '菜单id',
  `action` varchar(255) NOT NULL DEFAULT '' COMMENT '到你个猪权限 逗号隔开',
  `resource` varchar(255) NOT NULL DEFAULT '' COMMENT '动作资源 逗号隔开',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='角色菜单关联实体';

-- ----------------------------
-- Records of dt_role_menu
-- ----------------------------

-- ----------------------------
-- Table structure for dt_user
-- ----------------------------
DROP TABLE IF EXISTS `dt_user`;
CREATE TABLE `dt_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) DEFAULT '',
  `password` varchar(255) DEFAULT NULL,
  `real_name` varchar(255) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(255) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `creator` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of dt_user
-- ----------------------------

-- ----------------------------
-- Table structure for dt_user_role
-- ----------------------------
DROP TABLE IF EXISTS `dt_user_role`;
CREATE TABLE `dt_user_role` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `role_id` int(11) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of dt_user_role
-- ----------------------------
