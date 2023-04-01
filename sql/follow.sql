DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow`
(
    `id`          int        NOT NULL UNIQUE AUTO_INCREMENT COMMENT 'id',
    `user_id`     int        NOT NULL COMMENT '关注用户id',
    `to_user_id`  int        NOT NULL COMMENT '被关注用户id',
    `is_friend`   tinyint(1) NOT NULL DEFAULT '0' COMMENT '0代表没有互相关注，1代表互相关注',
    `create_time` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) USING BTREE,
    KEY `idx_to_user_id` (`to_user_id`)
) CHARSET=utf8;
