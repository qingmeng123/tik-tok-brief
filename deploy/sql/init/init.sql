-- 授权 root 用户可以远程链接
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'root';
flush privileges;

create database tik_tok_user default character set utf8mb4 collate utf8mb4_general_ci;
create database tik_tok_video default character set utf8mb4 collate utf8mb4_general_ci;
create database tik_tok_chat default character set utf8mb4 collate utf8mb4_general_ci;
create database tik_tok_comment default character set utf8mb4 collate utf8mb4_general_ci;
create database tik_tok_follow default character set utf8mb4 collate utf8mb4_general_ci;
create database tik_tok_like default character set utf8mb4 collate utf8mb4_general_ci;



use tik_tok_user;
drop table if exists user;
create table user
(
    id               bigint auto_increment
        primary key,
    user_id          bigint                               not null,
    username         char(32)                             not null,
    password         varchar(100)                         not null,
    follow_count     bigint     default 0                 not null,
    follower_count   bigint     default 0                 not null,
    is_follow        tinyint(1) default 0                 not null,
    total_favorited  bigint     default 0                 not null,
    work_count       bigint     default 0                 not null,
    favorite_count   bigint     default 0                 not null,
    create_time      timestamp  default CURRENT_TIMESTAMP null,
    update_time      timestamp  default CURRENT_TIMESTAMP null,
    constraint user_user_id_uindex
        unique (user_id),
    constraint user_username_uindex
        unique (username)
);



use tik_tok_video;
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`
(
    `id`             bigint          NOT NULL UNIQUE AUTO_INCREMENT,
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

use tik_tok_comment;
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`
(
    `id`           bigint          NOT NULL UNIQUE AUTO_INCREMENT COMMENT 'id',
    `user_id` bigint          NOT NULL COMMENT '发送用户id',
    `video_id`   bigint          NOT NULL COMMENT '接收消息用户id',
    `content`      varchar(300) NOT NULL,
    `create_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ,
    `update_time`  timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP ,
    PRIMARY KEY (`id`),
    KEY `idx_video_id` (`video_id`)
) CHARSET=utf8;

use tik_tok_follow;
DROP TABLE IF EXISTS `follow`;
CREATE TABLE `follow`
(
    `id`          bigint        NOT NULL UNIQUE AUTO_INCREMENT COMMENT 'id',
    `user_id`     bigint        NOT NULL COMMENT '关注用户id',
    `to_user_id`  bigint        NOT NULL COMMENT '被关注用户id',
    `is_friend`   tinyint(1) NOT NULL DEFAULT '0' COMMENT '0代表没有互相关注，1代表互相关注',
    `create_time` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_time` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    KEY `idx_user_id` (`user_id`) USING BTREE,
    KEY `idx_to_user_id` (`to_user_id`)
) CHARSET=utf8;

use tik_tok_like;
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

use tik_tok_chat;
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