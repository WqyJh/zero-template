CREATE TABLE `user` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT "用户ID",
  `nickname` varchar(128) NOT NULL DEFAULT "",
  `avatar` varchar(256) CHARSET latin1 COLLATE latin1_general_ci NOT NULL DEFAULT "",
  `gender` tinyint NOT NULL DEFAULT '0' COMMENT "性别 0:unknown 1:male 2:female",
  `status` tinyint NOT NULL DEFAULT '0' COMMENT "用户状态 0:正常 1:封禁 2:注销 3:删除",
  `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
  PRIMARY KEY (`id`),
  KEY `idx_created_at` (`created_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;