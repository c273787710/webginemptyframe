/*
 Navicat MySQL Data Transfer

 Source Server         : 本地
 Source Server Type    : MySQL
 Source Server Version : 100137
 Source Host           : localhost:3306
 Source Schema         : ginshop

 Target Server Type    : MySQL
 Target Server Version : 100137
 File Encoding         : 65001

 Date: 12/06/2020 13:49:36
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gin_admin
-- ----------------------------
DROP TABLE IF EXISTS `gin_admin`;
CREATE TABLE `gin_admin`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `username` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '登录名',
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `salt` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码盐',
  `nickname` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `last_ip` varchar(15) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '最好登录ip',
  `is_sup` tinyint(3) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否超级管理员',
  `role_id` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '角色id',
  `loginfailure` tinyint(4) NOT NULL COMMENT '失败次数',
  `logintime` int(11) NOT NULL COMMENT '登录时间',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '最好更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of gin_admin
-- ----------------------------
INSERT INTO `gin_admin` VALUES (1, 'admin', 'c917bcc3ea4a5c87981d7ecde5fdeb30', 'BpLnfg', '', '', '::1', 1, 0, 0, 1591940146, 1591005331, 1591940146);

-- ----------------------------
-- Table structure for gin_role
-- ----------------------------
DROP TABLE IF EXISTS `gin_role`;
CREATE TABLE `gin_role`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `role_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `rule_ids` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '规则ids',
  `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父类角色',
  `add_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for gin_rule
-- ----------------------------
DROP TABLE IF EXISTS `gin_rule`;
CREATE TABLE `gin_rule`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `rule_name` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '节点名称,对应前台name',
  `rule_path` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '对应后台path',
  `title` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '规则名称',
  `pid` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '父类id',
  `auth` tinyint(3) UNSIGNED NOT NULL DEFAULT 1 COMMENT '是否需要鉴权',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of gin_rule
-- ----------------------------
INSERT INTO `gin_rule` VALUES (1, 'rule', 'v1/admin/rule', '鉴权规则管理', 0, 1, 1591353334, 1591434697);
INSERT INTO `gin_rule` VALUES (2, 'ruleadd', 'v1/admin/rule/add', '规则添加', 0, 1, 1591432519, 1591432519);

SET FOREIGN_KEY_CHECKS = 1;
