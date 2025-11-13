
CREATE TABLE `t_users` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `name` varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
                           `age` int NOT NULL,
                           `gender` varchar(10) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           `create_at` datetime DEFAULT NULL,
                           `update_at` datetime DEFAULT NULL,
                           `num` int DEFAULT NULL,
                           `data` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';