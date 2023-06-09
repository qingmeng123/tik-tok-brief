// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package user

import (
	"context"

	"tik-tok-brief/service/user/rpc/proto/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	FollowTxnReq                = pb.FollowTxnReq
	FollowTxnResp               = pb.FollowTxnResp
	GenerateTokenReq            = pb.GenerateTokenReq
	GenerateTokenResp           = pb.GenerateTokenResp
	GetUserListByIdsReq         = pb.GetUserListByIdsReq
	GetUserListByIdsResp        = pb.GetUserListByIdsResp
	GetUserReq                  = pb.GetUserReq
	GetUserResp                 = pb.GetUserResp
	LoginReq                    = pb.LoginReq
	LoginResp                   = pb.LoginResp
	RegisterReq                 = pb.RegisterReq
	RegisterResp                = pb.RegisterResp
	UpdateUserFollowCountReq    = pb.UpdateUserFollowCountReq
	UpdateUserFollowCountResp   = pb.UpdateUserFollowCountResp
	UpdateUserFollowerCountReq  = pb.UpdateUserFollowerCountReq
	UpdateUserFollowerCountResp = pb.UpdateUserFollowerCountResp
	UpdateUserWorkCountReq      = pb.UpdateUserWorkCountReq
	UpdateUserWorkCountResp     = pb.UpdateUserWorkCountResp
	User                        = pb.User

	UserZrpcClient interface {
		// 用户注册
		Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
		// 用户登录
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		// 获取用户
		GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
		// 获取token
		GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error)
		// 批量获取用户信息
		GetUserListByIds(ctx context.Context, in *GetUserListByIdsReq, opts ...grpc.CallOption) (*GetUserListByIdsResp, error)
		// 更新用户粉丝数
		UpdateUserFollowerCount(ctx context.Context, in *UpdateUserFollowerCountReq, opts ...grpc.CallOption) (*UpdateUserFollowerCountResp, error)
		// 更新用户关注数
		UpdateUserFollowCount(ctx context.Context, in *UpdateUserFollowCountReq, opts ...grpc.CallOption) (*UpdateUserFollowCountResp, error)
		// 更新用户作品数
		UpdateUserWorkCount(ctx context.Context, in *UpdateUserWorkCountReq, opts ...grpc.CallOption) (*UpdateUserWorkCountResp, error)
		// 关注操作（事务处理）
		FollowTxn(ctx context.Context, in *FollowTxnReq, opts ...grpc.CallOption) (*FollowTxnResp, error)
		// 关注操作回滚（事务处理）
		FollowRevertTxn(ctx context.Context, in *FollowTxnReq, opts ...grpc.CallOption) (*FollowTxnResp, error)
	}

	defaultUserZrpcClient struct {
		cli zrpc.Client
	}
)

func NewUserZrpcClient(cli zrpc.Client) UserZrpcClient {
	return &defaultUserZrpcClient{
		cli: cli,
	}
}

// 用户注册
func (m *defaultUserZrpcClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

// 用户登录
func (m *defaultUserZrpcClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

// 获取用户
func (m *defaultUserZrpcClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetUser(ctx, in, opts...)
}

// 获取token
func (m *defaultUserZrpcClient) GenerateToken(ctx context.Context, in *GenerateTokenReq, opts ...grpc.CallOption) (*GenerateTokenResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

// 批量获取用户信息
func (m *defaultUserZrpcClient) GetUserListByIds(ctx context.Context, in *GetUserListByIdsReq, opts ...grpc.CallOption) (*GetUserListByIdsResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.GetUserListByIds(ctx, in, opts...)
}

// 更新用户粉丝数
func (m *defaultUserZrpcClient) UpdateUserFollowerCount(ctx context.Context, in *UpdateUserFollowerCountReq, opts ...grpc.CallOption) (*UpdateUserFollowerCountResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UpdateUserFollowerCount(ctx, in, opts...)
}

// 更新用户关注数
func (m *defaultUserZrpcClient) UpdateUserFollowCount(ctx context.Context, in *UpdateUserFollowCountReq, opts ...grpc.CallOption) (*UpdateUserFollowCountResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UpdateUserFollowCount(ctx, in, opts...)
}

// 更新用户作品数
func (m *defaultUserZrpcClient) UpdateUserWorkCount(ctx context.Context, in *UpdateUserWorkCountReq, opts ...grpc.CallOption) (*UpdateUserWorkCountResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.UpdateUserWorkCount(ctx, in, opts...)
}

// 关注操作（事务处理）
func (m *defaultUserZrpcClient) FollowTxn(ctx context.Context, in *FollowTxnReq, opts ...grpc.CallOption) (*FollowTxnResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.FollowTxn(ctx, in, opts...)
}

// 关注操作回滚（事务处理）
func (m *defaultUserZrpcClient) FollowRevertTxn(ctx context.Context, in *FollowTxnReq, opts ...grpc.CallOption) (*FollowTxnResp, error) {
	client := pb.NewUserClient(m.cli.Conn())
	return client.FollowRevertTxn(ctx, in, opts...)
}
