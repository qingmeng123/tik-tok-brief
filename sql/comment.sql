DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
    `id`           int          NOT NULL UNIQUE AUTO_INCREMENT COMMENT 'id',
    `user_id` int          NOT NULL COMMENT '发送用户id',
    `video_id`   int          NOT NULL COMMENT '接收消息用户id',
    `content`      varchar(300) NOT NULL,
    `create_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    `update_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
    PRIMARY KEY (`id`),
    KEY `idx_video_id` (`video_id`)
) CHARSET=utf8;