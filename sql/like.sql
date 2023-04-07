DROP TABLE IF EXISTS `like`;
CREATE TABLE `like`
(
    `id`          bigint        NOT NULL UNIQUE AUTO_INCREMENT COMMENT 'id',
    `user_id`     bigint        NOT NULL COMMENT '用户id',
    `video_id`  bigint        NOT NULL COMMENT '点赞视频id',
    `create_time` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) USING BTREE
) CHARSET=utf8;
