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
