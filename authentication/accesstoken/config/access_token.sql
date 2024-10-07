/*
 Navicat Premium Data Transfer

 Source Server         : 127.0.0.1
 Source Server Type    : MySQL
 Source Server Version : 80400
 Source Host           : localhost:3306
 Source Schema         : access_token

 Target Server Type    : MySQL
 Target Server Version : 80400
 File Encoding         : 65001

 Date: 06/10/2024 21:26:06
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for access_token
-- ----------------------------
DROP TABLE IF EXISTS `access_token`;
CREATE TABLE `access_token`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `server_sign` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '服务标识',
  `access_token` varchar(300) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'token',
  `is_valid` tinyint NULL DEFAULT NULL COMMENT '是否有效，1是，0否',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of access_token
-- ----------------------------
INSERT INTO `access_token` VALUES (1, 's1', '8d9ca373830011efbf8b8c32231f5813', 1, '2024-10-04 17:59:47', '2024-10-05 18:03:48');
INSERT INTO `access_token` VALUES (2, 's1', 'd04f8344830411efb6788c32231f5813', 1, '2024-10-05 18:30:17', '2024-10-05 18:30:17');
INSERT INTO `access_token` VALUES (3, 's1', 'd6791c20830511efb6788c32231f5813', 1, '2024-10-05 18:37:37', '2024-10-05 18:37:37');
INSERT INTO `access_token` VALUES (4, 's1', 'dda86623832c11ef8e5a8c32231f5813', 1, '2024-10-05 23:16:59', '2024-10-05 23:16:59');
INSERT INTO `access_token` VALUES (5, 's1', 'f82955ad832c11ef8e5a8c32231f5813', 1, '2024-10-05 23:17:44', '2024-10-05 23:17:44');
INSERT INTO `access_token` VALUES (6, 's1', '0c94b419832d11ef8e5a8c32231f5813', 1, '2024-10-05 23:18:18', '2024-10-05 23:18:18');
INSERT INTO `access_token` VALUES (7, 's1', '16795487832d11ef8e5a8c32231f5813', 1, '2024-10-05 23:18:35', '2024-10-05 23:18:35');
INSERT INTO `access_token` VALUES (8, 's1', '71d052b3832d11ef8e5a8c32231f5813', 1, '2024-10-05 23:21:08', '2024-10-05 23:21:08');
INSERT INTO `access_token` VALUES (9, 's1', 'efbf1928832d11ef90278c32231f5813', 1, '2024-10-05 23:24:39', '2024-10-05 23:24:39');
INSERT INTO `access_token` VALUES (10, 's1', '8fddb935832f11ef90278c32231f5813', 1, '2024-10-05 23:36:17', '2024-10-05 23:36:17');
INSERT INTO `access_token` VALUES (11, 's1', '9d316224832f11ef90278c32231f5813', 1, '2024-10-05 23:36:40', '2024-10-05 23:36:40');
INSERT INTO `access_token` VALUES (12, 's1', '639a62a5833011ef90278c32231f5813', 1, '2024-10-05 23:42:12', '2024-10-05 23:42:12');
INSERT INTO `access_token` VALUES (13, 's1', '830e969b833011efa3218c32231f5813', 1, '2024-10-05 23:43:05', '2024-10-05 23:43:05');
INSERT INTO `access_token` VALUES (14, 's1', 'f8a000d3833111efa3218c32231f5813', 1, '2024-10-05 23:53:32', '2024-10-05 23:53:32');
INSERT INTO `access_token` VALUES (15, 's1', 'fb092007833111efa3218c32231f5813', 1, '2024-10-05 23:53:36', '2024-10-05 23:53:36');
INSERT INTO `access_token` VALUES (16, 's1', '44ffe6de833211efb3d48c32231f5813', 1, '2024-10-05 23:55:40', '2024-10-05 23:55:40');

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

SET FOREIGN_KEY_CHECKS = 1;
