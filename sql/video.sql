DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`
(
    `id`             bigint          NOT NULL UNIQUE primary key AUTO_INCREMENT,
    `video_id`        bigint          NOT NULL UNIQUE COMMENT '视频id',
    `user_id`        bigint          NOT NULL UNIQUE COMMENT '发布作者id',
    `title`          char(32)  NOT NULL COMMENT '视频标题',
    `play_url`       varchar(500) NOT NULL COMMENT '视频播放地址',
    `cover_url`      varchar(500) NOT NULL COMMENT '封面地址',
    `favorite_count` int          NOT NULL DEFAULT '0' COMMENT '点赞数量',
    `comment_count`  int          NOT NULL DEFAULT '0' COMMENT '评论数量',
    `create_time`    timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time`    timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    KEY              `idx_video_id` (`video_id`) USING BTREE,
    KEY              `idx_user_id` (`user_id`) USING BTREE,
    KEY              `idx_create_time` (`create_time`) USING BTREE
)  CHARSET=utf8;