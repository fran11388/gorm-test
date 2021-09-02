DROP DATABASE IF EXISTS gorm;
CREATE DATABASE gorm CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

USE gorm;

/*
 Navicat Premium Data Transfer

 Source Server         : local-docker-mysql1
 Source Server Type    : MySQL
 Source Server Version : 80026
 Source Host           : localhost:3306
 Source Schema         : gorm

 Target Server Type    : MySQL
 Target Server Version : 80026
 File Encoding         : 65001

 Date: 02/09/2021 15:00:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
                         `id` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
                         `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
                         `balance` decimal(18, 2) NOT NULL,
                         PRIMARY KEY (`id`) USING BTREE,
                         UNIQUE INDEX `nane`(`name`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('11', 'frank', 100.00);

SET FOREIGN_KEY_CHECKS = 1;
