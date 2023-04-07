DROP TABLE IF EXISTS `chat`;
CREATE TABLE `chat`
(
    `id`           bigint          NOT NULL UNIQUE AUTO_INCREMENT COMMENT 'id',
    `from_user_id` bigint          NOT NULL COMMENT '发送用户id',
    `to_user_id`   bigint          NOT NULL COMMENT '接收消息用户id',
    `content`      varchar(300) NOT NULL,
    `create_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    `update_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
    PRIMARY KEY (`id`),
    KEY `idx_chat` (`from_user_id`, `to_user_id`)
) CHARSET=utf8;