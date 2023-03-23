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
