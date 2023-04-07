## 功能介绍

- 视频：视频推送、视频投稿、发布列表
- 用户：用户注册、用户登录、用户信息
- 点赞：点赞操作、点赞列表
- 评论：评论操作、评论列表
- 关注：关注操作、关注列表、粉丝列表、好友列表、好友最新消息
- 聊天：聊天操作、聊天记录

## 项目部署

安装`docker`和`docker-compose`

- git clone
- 配置各个etc下的yaml文件
- 使用`docker-compose build`将六个服务构建，再使用`docker-compose up -d`启动服务

## 接口文档
见doc文件

## 技术选型
本项目采用基于`go-zero`的RPC框架，包含了`go-zero`以及相关`go-zero`作者开发的一些中间件，所用到的技术栈基本是`go-zero`
项目组的自研组件。

- Go-zero
- Mysql
- Redis
- COS
- grpc

### 小工具
- "github.com/sony/sonyflake"
- "github.com/u2takey/ffmpeg-go v0.4.1"
- "github.com/jinzhu/copier v0.3.5"