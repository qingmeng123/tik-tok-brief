# 抖音简洁版后端接口文档

*返回信息说明*： 所有status_code为状态码，0-成功，其他值-失败



## 1.API接口文档

### 1.基础接口

#### 1.user

基本的用户登录注册，查看信息

##### /douyin/user/register/ - 用户注册接口

新用户注册时提供用户名，密码，昵称即可，用户名需要保证唯一。创建成功后返回用户 id 和权限token.

**接口类型**

POST

**接口定义**

参数

| 参数名   | 必选 | 请求类型 | 数据类型 | 说明   |
| -------- | ---- | -------- | -------- | ------ |
| username | 是   | query    | string   | 用户名 |
| password | 是   | query    | string   | 密码   |

**返回响应**

数据结构

```protobuf
message douyin_user_register_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  required int64 user_id = 3; // 用户id
  required string token = 4; // 用户鉴权token
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success",
    "user_id": 1680691988,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNjgwNjkxOTg4LCJleHAiOjE2ODEyOTE5ODgsImlhdCI6MTY4MDY5MTk4OH0.pac1Q2PUvtxvrgSL7x_dTGqIwmQuRoKQ1KjiyL5HMEw"
}
```



##### /douyin/user/login/ - 用户登录接口

通过用户名和密码进行登录，登录成功后返回用户 id 和权限token.

**接口类型**

POST

**接口定义**

参数

| 参数名   | 必选 | 请求类型 | 数据类型 | 说明   |
| -------- | ---- | -------- | -------- | ------ |
| username | 是   | query    | string   | 用户名 |
| password | 是   | query    | string   | 密码   |

**返回响应**

数据结构

```protobuf
message douyin_user_login_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  required int64 user_id = 3; // 用户id
  required string token = 4; // 用户鉴权token
}
```



示例

```
{
    "status_code": 0,
    "status_msg": "success",
    "user_id": 1680691988,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxNjgwNjkxOTg4LCJleHAiOjE2ODEyOTIwMzgsImlhdCI6MTY4MDY5MjAzOH0.yrPWbcH3InaPgtVjqFR65uhL_eEfQzhXoyXZOg7Aw7U"
}
```



##### /douyin/user/ - 用户信息接口

登陆后获取user_id用户的详细信息，如昵称，关注数和粉丝数等，若非token本人，还会返回是否关注。

**接口类型**

GET

**接口定义**

参数

| 参数名  | 必选 | 请求类型 | 数据类型 | 说明          |
| ------- | ---- | -------- | -------- | ------------- |
| user_id | 是   | query    | int64    | 用户ID        |
| token   | 是   | query    | string   | 用户鉴权token |

**返回响应**

数据结构

```protobuf
message douyin_user_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  required User user = 3; // 用户信息
}

message User {
  required int64 id = 1; // 用户id
  required string name = 2; // 用户名称
  optional int64 follow_count = 3; // 关注总数
  optional int64 follower_count = 4; // 粉丝总数
  required bool is_follow = 5; // true-已关注，false-未关注
  optional string avatar = 6; //用户头像
  optional string background_image = 7; //用户个人页顶部大图
  optional string signature = 8; //个人简介
  optional int64 total_favorited = 9; //获赞数量
  optional int64 work_count = 10; //作品数量
  optional int64 favorite_count = 11; //点赞数量
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success",
    "user": {
        "id": 1680688646,
        "name": "duryun",
        "follow_count": 0,
        "follower_count": 0,
        "is_follow": false,
        "total_favorited": 0,
        "work_count": 0,
        "favorite_count": 0
    }
}
```



#### 2.video

视频上传，获取视频流，指定用户视频信息功能

##### /douyin/publish/action/ - 视频投稿

登录用户选择视频文件上传。

**接口类型**

POST

**接口定义**

参数

| 参数名 | 必选 | 请求类型  | 数据类型 | 说明          |
| ------ | ---- | --------- | -------- | ------------- |
| data   | 是   | form-data | bytes    | 视频数据      |
| token  | 是   | form-data | string   | 用户鉴权token |
| title  | 是   | query     | string   | 视频标题      |

**返回响应**

数据结构

```protobuf
message douyin_publish_action_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success"
}
```

##### /douyin/feed/ - 视频流接口

不限制登录状态，返回按投稿时间倒序的视频列表，视频数由服务端控制，单次最多30个。

**接口类型**

GET

**接口定义**

参数

| 参数名      | 必选 | 请求类型 | 数据类型 | 说明                                   |
| ----------- | ---- | -------- | -------- | -------------------------------------- |
| latest_time | 否   | query    | int64    | 限制返回视频的最新投稿时间戳，精确到秒 |
| token       | 否   | query    | string   | 用户鉴权token                          |

**返回响应**

数据结构

```protobuf
message douyin_feed_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  repeated Video video_list = 3; // 视频列表
  optional int64 next_time = 4; // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

message Video {
  required int64 id = 1; // 视频唯一标识
  required User author = 2; // 视频作者信息
  required string play_url = 3; // 视频播放地址
  required string cover_url = 4; // 视频封面地址
  required int64 favorite_count = 5; // 视频的点赞总数
  required int64 comment_count = 6; // 视频的评论总数
  required bool is_favorite = 7; // true-已点赞，false-未点赞
  required string title = 8; // 视频标题
}

message User {
  required int64 id = 1; // 用户id
  required string name = 2; // 用户名称
  optional int64 follow_count = 3; // 关注总数
  optional int64 follower_count = 4; // 粉丝总数
  required bool is_follow = 5; // true-已关注，false-未关注
  optional string avatar = 6; //用户头像
  optional string background_image = 7; //用户个人页顶部大图
  optional string signature = 8; //个人简介
  optional int64 total_favorited = 9; //获赞数量
  optional int64 work_count = 10; //作品数量
  optional int64 favorite_count = 11; //点赞数量
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success",
    "video_list": [
        {
            "id": 1680703713,
            "author": {
                "id": 0,
                "name": "",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": false
            },
            "play_url": "https://tik-tok-1317220115.cos.ap-chongqing.myqcloud.com/video/test.mp4",
            "cover_url": "https://tik-tok-1317220115.cos.ap-chongqing.myqcloud.com/cover/test_tik_tok.jpg",
            "favorite_count": 0,
            "comment_count": 0,
            "is_favorite": false,
            "title": "test"
        }
    ]
}
```



##### /douyin/publish/list/ - 发布列表

登录用户查看指定用户的视频发布列表，直接列出用户所有投稿过的视频。

**接口类型**

GET

**接口定义**

参数

| 参数名  | 必选 | 请求类型 | 数据类型 | 说明          |
| ------- | ---- | -------- | -------- | ------------- |
| user_id | 是   | query    | int64    | 指定用户id    |
| token   | 是   | query    | string   | 用户鉴权token |

**返回响应**

数据结构

```protobuf
message douyin_publish_list_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  repeated Video video_list = 3; // 用户发布的视频列表
}

message Video {
  required int64 id = 1; // 视频唯一标识
  required User author = 2; // 视频作者信息
  required string play_url = 3; // 视频播放地址
  required string cover_url = 4; // 视频封面地址
  required int64 favorite_count = 5; // 视频的点赞总数
  required int64 comment_count = 6; // 视频的评论总数
  required bool is_favorite = 7; // true-已点赞，false-未点赞
  required string title = 8; // 视频标题
}

message User {
  required int64 id = 1; // 用户id
  required string name = 2; // 用户名称
  optional int64 follow_count = 3; // 关注总数
  optional int64 follower_count = 4; // 粉丝总数
  required bool is_follow = 5; // true-已关注，false-未关注
  optional string avatar = 6; //用户头像
  optional string background_image = 7; //用户个人页顶部大图
  optional string signature = 8; //个人简介
  optional int64 total_favorited = 9; //获赞数量
  optional int64 work_count = 10; //作品数量
  optional int64 favorite_count = 11; //点赞数量
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success",
    "video_list": [
        {
            "id": 1680703713,
            "author": {
                "id": 0,
                "name": "",
                "follow_count": 0,
                "follower_count": 5,
                "is_follow": true
            },
            "play_url": "https://tik-tok-1317220115.cos.ap-chongqing.myqcloud.com/video/test.mp4",
            "cover_url": "https://tik-tok-1317220115.cos.ap-chongqing.myqcloud.com/cover/test_tik_tok.jpg",
            "favorite_count": 0,
            "comment_count": 0,
            "is_favorite": false,
            "title": "test"
        }
    ]
}
```

##### 

### 2.互动管理系统

每个登录用户支持点赞，同时维护用户自己的点赞视频列表，在个人信息页中查看。

登录用户能够查看视频的评论列表，对视频进行评论。



##### /douyin/favorite/action/ - 赞操作

登录用户对视频的点赞和取消点赞操作。

**接口类型**

POST

**接口定义**

参数

| 参数名      | 必选 | 请求类型 | 数据类型 | 说明               |
| ----------- | ---- | -------- | -------- | ------------------ |
| video_id    | 是   | query    | int64    | 视频id             |
| token       | 是   | query    | string   | 用户鉴权token      |
| action_type | 是   | query    | int32    | 1-点赞，2-取消点赞 |

**返回响应**

数据结构

```protobuf
message douyin_favorite_action_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success"
}
```

##### 

##### /douyin/favorite/list/ - 喜欢列表

登录用户查看指定用户的所有点赞视频。

**接口类型**

GET

**接口定义**

参数

| 参数名  | 必选 | 请求类型 | 数据类型 | 说明          |
| ------- | ---- | -------- | -------- | ------------- |
| user_id | 是   | query    | int64    | 指定用户id    |
| token   | 是   | query    | string   | 用户鉴权token |

**返回响应**

数据结构

```protobuf
message douyin_favorite_list_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  repeated Video video_list = 3; // 用户点赞视频列表
}

message Video {
  required int64 id = 1; // 视频唯一标识
  required User author = 2; // 视频作者信息
  required string play_url = 3; // 视频播放地址
  required string cover_url = 4; // 视频封面地址
  required int64 favorite_count = 5; // 视频的点赞总数
  required int64 comment_count = 6; // 视频的评论总数
  required bool is_favorite = 7; // true-已点赞，false-未点赞
  required string title = 8; // 视频标题
}

message User {
  required int64 id = 1; // 用户id
  required string name = 2; // 用户名称
  optional int64 follow_count = 3; // 关注总数
  optional int64 follower_count = 4; // 粉丝总数
  required bool is_follow = 5; // true-已关注，false-未关注
  optional string avatar = 6; //用户头像
  optional string background_image = 7; //用户个人页顶部大图
  optional string signature = 8; //个人简介
  optional int64 total_favorited = 9; //获赞数量
  optional int64 work_count = 10; //作品数量
  optional int64 favorite_count = 11; //点赞数量
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success",
    "video_list": [
        {
            "id": 1680703713,
            "author": {
                "id": 0,
                "name": "",
                "follow_count": 1,
                "follower_count": 5,
                "is_follow": false
            },
            "play_url": "https://tik-tok-1317220115.cos.ap-chongqing.myqcloud.com/video/test.mp4",
            "cover_url": "https://tik-tok-1317220115.cos.ap-chongqing.myqcloud.com/cover/test_tik_tok.jpg",
            "favorite_count": 2,
            "comment_count": 0,
            "is_favorite": true,
            "title": "test"
        }
    ]
}
```

##### /douyin/comment/action/ - 评论操作

登录用户对视频进行评论。

**接口类型**

POST

**接口定义**

参数

| 参数名       | 必选 | 请求类型 | 数据类型 | 说明                                          |
| ------------ | ---- | -------- | -------- | --------------------------------------------- |
| video_id     | 是   | query    | int64    | 视频id                                        |
| token        | 是   | query    | string   | 用户鉴权token                                 |
| action_type  | 是   | query    | int32    | 1-发布评论，2-删除评论                        |
| comment_text | 否   | query    | string   | 用户填写的评论内容，在action_type=1的时候使用 |
| comment_id   | 否   | query    | int64    | 要删除的评论id，在action_type=2的时候使用     |

**返回响应**

数据结构

```protobuf
message douyin_comment_action_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  optional Comment comment = 3; // 评论成功返回评论内容，不需要重新拉取整个列表
}

message Comment {
  required int64 id = 1; // 视频评论id
  required User user =2; // 评论用户信息
  required string content = 3; // 评论内容
  required string create_date = 4; // 评论发布日期，格式 mm-dd
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success",
    "comment": {
        "id": 6,
        "user": {
            "id": 0,
            "name": "",
            "follow_count": 0,
            "follower_count": 0,
            "is_follow": false
        },
        "content": "你好",
        "create_date": "2023-04-05 22:29:01"
    }
}
```

##### 

##### /douyin/comment/list/ - 视频评论列表

查看视频的所有评论，按发布时间倒序。

**接口类型**

GET

**接口定义**

参数

| 参数名   | 必选 | 请求类型 | 数据类型 | 说明          |
| -------- | ---- | -------- | -------- | ------------- |
| video_id | 是   | query    | int64    | 视频id        |
| token    | 是   | query    | string   | 用户鉴权token |

**返回响应**

数据结构

```protobuf
message douyin_comment_list_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  repeated Comment comment_list = 3; // 评论列表
}

message Comment {
  required int64 id = 1; // 视频评论id
  required User user =2; // 评论用户信息
  required string content = 3; // 评论内容
  required string create_date = 4; // 评论发布日期，格式 mm-dd
}

message User {
  required int64 id = 1; // 用户id
  required string name = 2; // 用户名称
  optional int64 follow_count = 3; // 关注总数
  optional int64 follower_count = 4; // 粉丝总数
  required bool is_follow = 5; // true-已关注，false-未关注
  optional string avatar = 6; //用户头像
  optional string background_image = 7; //用户个人页顶部大图
  optional string signature = 8; //个人简介
  optional int64 total_favorited = 9; //获赞数量
  optional int64 work_count = 10; //作品数量
  optional int64 favorite_count = 11; //点赞数量
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success",
    "comment_list": [
        {
            "id": 6,
            "user": {
                "id": 0,
                "name": "",
                "follow_count": 0,
                "follower_count": 0,
                "is_follow": false
            },
            "content": "你好",
            "create_date": "2023-04-05 22:29:01"
        }
    ]
}
```

##### 

### 3.社交管理系统接口

实现用户之间的关注关系维护，登录用户能够关注或取关其他用户，同时查看其他用户的关注和粉丝列表。

同时彼此关注的将变为朋友关系，朋友关系的能够互发消息聊天。

##### /douyin/relation/action/ - 关系操作

登录用户对其他用户进行关注或取消关注。

**接口类型**

POST

**接口定义**

参数

| 参数名      | 必选 | 请求类型 | 数据类型 | 说明               |
| ----------- | ---- | -------- | -------- | ------------------ |
| to_user_id  | 是   | query    | int64    | 对方用户ID         |
| token       | 是   | query    | string   | 用户鉴权token      |
| action_type | 是   | query    | int32    | 1-关注，2-取消关注 |

**返回响应**

数据结构

```protobuf
message douyin_relation_action_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success"
}

{
    "status_code": 100002,
    "status_msg": "是否关注操作操作错误"
}
```

##### /douyin/relation/follow/list/ - 用户关注列表

登录用户查看user_id的所有关注用户列表，同时显示自己是否关注该列表的用户。

**接口类型**

GET

**接口定义**

参数

| 参数名  | 必选 | 请求类型 | 数据类型 | 说明          |
| ------- | ---- | -------- | -------- | ------------- |
| user_id | 是   | query    | int64    | 用户ID        |
| token   | 是   | query    | string   | 用户鉴权token |

**返回响应**

数据结构

```protobuf
message douyin_relation_follow_list_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  repeated User user_list = 3; // 用户信息列表
}

message User {
  required int64 id = 1; // 用户id
  required string name = 2; // 用户名称
  optional int64 follow_count = 3; // 关注总数
  optional int64 follower_count = 4; // 粉丝总数
  required bool is_follow = 5; // true-已关注，false-未关注
  optional string avatar = 6; //用户头像
  optional string background_image = 7; //用户个人页顶部大图
  optional string signature = 8; //个人简介
  optional int64 total_favorited = 9; //获赞数量
  optional int64 work_count = 10; //作品数量
  optional int64 favorite_count = 11; //点赞数量
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success",
    "user_list": [
        {
            "id": 1680688653,
            "name": "duryun1",
            "follow_count": 5,
            "follower_count": 1,
            "is_follow": false,
            "total_favorited": 0,
            "work_count": 0,
            "favorite_count": 0
        }
    ]
}
```



##### /douyin/relation/follower/list/ - 用户粉丝列表

登录用户查看user_id的所有粉丝列表，同时显示自己是否关注该列表的用户。

**接口类型**

GET

**接口定义**

参数

| 参数名  | 必选 | 请求类型 | 数据类型 | 说明          |
| ------- | ---- | -------- | -------- | ------------- |
| user_id | 是   | query    | int64    | 用户ID        |
| token   | 是   | query    | string   | 用户鉴权token |

**返回响应**

数据结构

```protobuf
message douyin_relation_follower_list_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  repeated User user_list = 3; // 用户列表
}

message User {
  required int64 id = 1; // 用户id
  required string name = 2; // 用户名称
  optional int64 follow_count = 3; // 关注总数
  optional int64 follower_count = 4; // 粉丝总数
  required bool is_follow = 5; // true-已关注，false-未关注
  optional string avatar = 6; //用户头像
  optional string background_image = 7; //用户个人页顶部大图
  optional string signature = 8; //个人简介
  optional int64 total_favorited = 9; //获赞数量
  optional int64 work_count = 10; //作品数量
  optional int64 favorite_count = 11; //点赞数量
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success",
    "user_list": [
        {
            "id": 1680688646,
            "name": "duryun",
            "follow_count": 1,
            "follower_count": 5,
            "is_follow": false,
            "total_favorited": 0,
            "work_count": 0,
            "favorite_count": 0
        }
    ]
}
```



##### /douyin/relation/friend/list/ - 用户好友列表

登录用户查看自己的所有朋友列表，同时返回与对应朋友的最新聊天

**接口类型**

GET

**接口定义**

参数

| 参数名  | 必选 | 请求类型 | 数据类型 | 说明          |
| ------- | ---- | -------- | -------- | ------------- |
| user_id | 否   | query    | int64    | 用户ID        |
| token   | 是   | query    | string   | 用户鉴权token |

**返回响应**

数据结构

```protobuf
message douyin_relation_friend_list_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  repeated FriendUser user_list = 3; // 用户列表
}

message User {
  required int64 id = 1; // 用户id
  required string name = 2; // 用户名称
  optional int64 follow_count = 3; // 关注总数
  optional int64 follower_count = 4; // 粉丝总数
  required bool is_follow = 5; // true-已关注，false-未关注
  optional string avatar = 6; //用户头像
  optional string background_image = 7; //用户个人页顶部大图
  optional string signature = 8; //个人简介
  optional int64 total_favorited = 9; //获赞数量
  optional int64 work_count = 10; //作品数量
  optional int64 favorite_count = 11; //点赞数量
}

message FriendUser extends User {
    optional string message = 1; // 和该好友的最新聊天消息
    required int64 msgType = 2; // message消息的类型，0 => 当前请求用户接收的消息， 1 => 当前请求用户发送的消息
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success",
    "user_list": [
        {
            "id": 1680688653,
            "name": "duryun1",
            "follow_count": 5,
            "follower_count": 1,
            "is_follow": true,
            "total_favorited": 0,
            "work_count": 0,
            "favorite_count": 0,
            "message": "你好",
            "msgType": 1
        }
    ]
}
```



##### /douyin/message/chat/ - 聊天记录

当前登录用户和其他指定用户的聊天消息记录

**接口类型**

GET

**接口定义**

参数

| 参数名     | 必选 | 请求类型 | 数据类型 | 说明          |
| ---------- | ---- | -------- | -------- | ------------- |
| to_user_id | 是   | query    | int64    | 用户ID        |
| token      | 是   | query    | string   | 用户鉴权token |

**返回响应**

数据结构

```protobuf
message douyin_message_chat_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  repeated Message message_list = 3; // 消息列表
}

message Message {
  required int64 id = 1; // 消息id
  required int64 to_user_id = 2; // 该消息接收者的id
  required int64 from_user_id =3; // 该消息发送者的id
  required string content = 4; // 消息内容
  optional string create_time = 5; // 消息创建时间
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success",
    "message_list": [
        {
            "id": 3,
            "to_user_id": 1680688653,
            "from_user_id": 1680688646,
            "content": "你好",
            "create_time": "2023-04-05 19:50:39"
        }
    ]
}
```



##### /douyin/message/action/ - 消息操作

登录用户对消息的相关操作，目前只支持消息发送

**接口类型**

POST

**接口定义**

参数

| 参数名      | 必选 | 请求类型 | 数据类型 | 说明          |
| ----------- | ---- | -------- | -------- | ------------- |
| to_user_id  | 是   | query    | int64    | 对方用户ID    |
| token       | 是   | query    | string   | 用户鉴权token |
| action_type | 是   | query    | int32    | 1-发送消息    |
| content     | 是   | query    | string   | 消息内容      |

**返回响应**

数据结构

```protobuf
message douyin_relation_action_response {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
}
```



返回示例

```
{
    "status_code": 0,
    "status_msg": "success"
}
```

