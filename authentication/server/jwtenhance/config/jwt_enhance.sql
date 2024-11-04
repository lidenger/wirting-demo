/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80400
 Source Host           : localhost:3306
 Source Schema         : jwt_enhance

 Target Server Type    : MySQL
 Target Server Version : 80400
 File Encoding         : 65001

 Date: 27/10/2024 15:11:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for aksk
-- ----------------------------
DROP TABLE IF EXISTS `aksk`;
CREATE TABLE `aksk`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `server_sign` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `ak` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'access key',
  `sk` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'secret key',
  `is_enable` tinyint NULL DEFAULT NULL COMMENT '是否启用',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '备注',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of aksk
-- ----------------------------
INSERT INTO `aksk` VALUES (1, 's1', 'b4a9bcc1825f11ef', 'c26c6f9c825f11efa5ae8c32', 1, 'secret1', '2024-10-04 22:49:06', '2024-10-04 22:49:08');
INSERT INTO `aksk` VALUES (2, 's1', '7dcea8db82bd11ef', '81b377ed82bd11ef826b8c32', 1, 'secret2', '2024-10-05 10:00:00', '2024-10-05 10:00:02');

-- ----------------------------
-- Table structure for jwt_deactive
-- ----------------------------
DROP TABLE IF EXISTS `jwt_deactive`;
CREATE TABLE `jwt_deactive`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `jwt_id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jwt_deactive
-- ----------------------------

-- ----------------------------
-- Table structure for order_info
-- ----------------------------
DROP TABLE IF EXISTS `order_info`;
CREATE TABLE `order_info`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `order_no` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '订单编号',
  `order_status` tinyint NULL DEFAULT NULL COMMENT '订单状态',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_info
-- ----------------------------
INSERT INTO `order_info` VALUES (1, '425c3b29832711ef8f0f8c32', 1, '2024-10-05 22:36:56', '2024-10-05 22:36:59');

-- ----------------------------
-- Table structure for server_info
-- ----------------------------
DROP TABLE IF EXISTS `server_info`;
CREATE TABLE `server_info`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `server_sign` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '服务标识',
  `server_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '服务名称',
  `server_desc` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '服务描述',
  `is_enable` tinyint NULL DEFAULT NULL COMMENT '是否启用',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of server_info
-- ----------------------------
INSERT INTO `server_info` VALUES (1, 's1', 'server1', 'test server 1', 1, '2024-09-16 15:57:02', '2024-09-16 15:57:05');

-- ----------------------------
-- Table structure for sys_secret
-- ----------------------------
DROP TABLE IF EXISTS `sys_secret`;
CREATE TABLE `sys_secret`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `secret` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密钥',
  `is_enable` tinyint NULL DEFAULT NULL COMMENT '是否启用',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_secret
-- ----------------------------
INSERT INTO `sys_secret` VALUES (1, '4d8a8dcd8e8011ef91d48c32231f5813', 1, '2024-10-20 09:12:06', '2024-10-20 09:12:09');
INSERT INTO `sys_secret` VALUES (2, '59d36f178e8011efbd618c32231f5813', 1, '2024-10-20 09:12:26', '2024-10-20 09:12:28');
INSERT INTO `sys_secret` VALUES (3, '6480b0868e8011ef9d9b8c32231f5813', 1, '2024-10-20 09:12:47', '2024-10-20 09:12:49');

SET FOREIGN_KEY_CHECKS = 1;
