-- 授权 root 用户可以远程链接
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'root';
flush privileges;

create database tik_tok_user default character set utf8mb4 collate utf8mb4_general_ci;
create database tik_tok_video default character set utf8mb4 collate utf8mb4_general_ci;


use tik_tok_user;
create table user
(
    id               bigint auto_increment
        primary key,
    user_id          bigint                               not null,
    username         char(32)                             not null,
    password         varchar(100)                         not null,
    name             varchar(50)                          null,
    follow_count     bigint                               null,
    follower_count   bigint                               null,
    is_follow        tinyint(1) default 0                 null,
    avatar           varchar(500)                         null,
    background_image varchar(500)                         null,
    signature        varchar(100)                         null,
    total_favorited  bigint     default 0                 null,
    work_count       bigint     default 0                 null,
    favorite_count   bigint     default 0                 null,
    create_time      timestamp  default CURRENT_TIMESTAMP null,
    update_time      timestamp  default CURRENT_TIMESTAMP null,
    constraint user_user_id_uindex
        unique (user_id),
    constraint user_username_uindex
        unique (username)
);

use tik_tok_video;
CREATE TABLE `video`
(
    `id`             bigint          NOT NULL UNIQUE primary key AUTO_INCREMENT,
    `video_id`        bigint          NOT NULL UNIQUE COMMENT '视频id',
    `user_id`        bigint          NOT NULL  COMMENT '发布作者id',
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