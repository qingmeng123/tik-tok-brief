package logic

import (
	"context"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"tik-tok-brief/common/errorx"
	fpb "tik-tok-brief/service/follow/rpc/proto/pb"
	"tik-tok-brief/service/user/api/internal/svc"
	"tik-tok-brief/service/user/api/internal/types"
	upb "tik-tok-brief/service/user/rpc/proto/pb"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 关注操作，分布式事务处理
func (l *FollowLogic) Follow(req *types.FollowReq) (resp *types.FollowResp, err error) {
	userId := l.ctx.Value("user_id").(int64)
	//获取BuildTarget
	userRPCBusiServer, err := l.svcCtx.Config.UserRPC.BuildTarget()
	if err != nil {
		logx.Error("svcCtx.Config.UserRPC.BuildTarget err:", err)
		return nil, errorx.NewInternalErr()
	}

	followRPCBusiServer, err := l.svcCtx.Config.FollowRPC.BuildTarget()
	if err != nil {
		logx.Error("svcCtx.Config.FollowRPC.BuildTarget err:", err)
		return nil, errorx.NewInternalErr()
	}
	// dtm 服务的 etcd 注册地址
	var dtmServer = "etcd://127.0.0.1:2379/dtmservice"

	//创建一个gid
	gid := dtmgrpc.MustGenGid(dtmServer)
	//创建一个saga协议的事务
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(userRPCBusiServer+"/user.user/FollowTxn", userRPCBusiServer+"/user.user/FollowRevertTxn", &upb.FollowTxnReq{
			UserId:     userId,
			ToUserId:   req.ToUserId,
			ActionType: int32(req.ActionType),
		}).
		Add(followRPCBusiServer+"/follow.follow/FollowTxn", followRPCBusiServer+"/follow.follow/FollowTxnRevert", &fpb.FollowTxnReq{
			UserId:     userId,
			ToUserId:   req.ToUserId,
			ActionType: int32(req.ActionType),
		})

	//事务提交
	err = saga.Submit()

	if err != nil {
		logx.Error("saga.Submit err:", err)
		return nil, err
	}

	return &types.FollowResp{StatusResponse: types.StatusResponse{
		StatusCode: int32(errorx.OK),
		StatusMsg:  errorx.SUCCESS,
	}}, nil
}

//
//func (l *FollowLogic) Follow(req *types.FollowReq) (resp *types.FollowResp, err error) {
//	userId := l.ctx.Value("user_id").(int64)
//
//
//	//检查是否存在toUserId
//	toUser, err := l.svcCtx.UserRPC.GetUser(l.ctx, &upb.GetUserReq{UserID: req.ToUserId})
//	if err != nil {
//		logx.Error("UserRPC.GetUser err:", err)
//		return nil, err
//	}
//	if toUser == nil {
//		return nil, errorx.NewParamErr(errorx.ERRUSERID)
//	}
//
//	//关注
//	if req.ActionType == 1 {
//		_, err = l.svcCtx.FollowRPC.Follow(l.ctx, &fpb.FollowReq{
//			UserId:   userId,
//			ToUserId: req.ToUserId,
//		})
//		if err != nil {
//			logx.Error("FollowRPC.Follow err:", err)
//			return nil, err
//		}
//	}
//
//	//取消关注
//	if req.ActionType == 2 {
//		_, err = l.svcCtx.FollowRPC.UnFollow(l.ctx, &fpb.UnFollowReq{
//			UserId:   userId,
//			ToUserId: req.ToUserId,
//		})
//		if err != nil {
//			logx.Error("FollowRPC.UnFollow err:", err)
//			return nil, err
//		}
//	}
//
//	//变更登录用户关注数
//	var number int64
//	number = 1
//	if req.ActionType == 2 {
//		number = -1
//	}
//
//	_, err = l.svcCtx.UserRPC.UpdateUserFollowCount(l.ctx, &upb.UpdateUserFollowCountReq{
//		UserId: userId,
//		Number: number,
//	})
//	if err != nil {
//		logx.Error("UserRPC.UpdateUserFollowCount err:", err)
//		return nil, err
//	}
//
//	//变更关注用户粉丝数
//	_, err = l.svcCtx.UserRPC.UpdateUserFollowerCount(l.ctx, &upb.UpdateUserFollowerCountReq{
//		UserId: req.ToUserId,
//		Number: number,
//	})
//	if err != nil {
//		logx.Error("UserRPC.UpdateUserFollowerCount err:", err)
//		return nil, err
//	}
//
//	return &types.FollowResp{StatusResponse: types.StatusResponse{
//		StatusCode: int32(errorx.OK),
//		StatusMsg:  errorx.SUCCESS,
//	}}, nil
//
//}
