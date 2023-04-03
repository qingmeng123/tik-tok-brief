package logic

import (
	"context"
	"tik-tok-brief/common/errorx"
	fpb "tik-tok-brief/service/follow/rpc/proto/pb"
	"tik-tok-brief/service/user/api/internal/svc"
	"tik-tok-brief/service/user/api/internal/types"
	upb "tik-tok-brief/service/user/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *FollowLogic) Follow(req *types.FollowReq) (resp *types.FollowResp, err error) {
	userId := l.ctx.Value("user_id").(int64)

	//检查是否存在toUserId
	toUser, err := l.svcCtx.UserRPC.GetUser(l.ctx, &upb.GetUserReq{UserID: req.ToUserId})
	if err != nil {
		logx.Error("UserRPC.GetUser err:", err)
		return nil, err
	}
	if toUser == nil {
		return nil, errorx.NewParamErr(errorx.ERRUSERID)
	}

	//关注
	if req.ActionType == 1 {
		_, err = l.svcCtx.FollowerRPC.Follow(l.ctx, &fpb.FollowReq{
			UserId:   userId,
			ToUserId: req.ToUserId,
		})
		if err != nil {
			logx.Error("FollowerRPC.Follow err:", err)
			return nil, err
		}
		return &types.FollowResp{StatusResponse: types.StatusResponse{
			StatusCode: int32(errorx.OK),
			StatusMsg:  errorx.SUCCESS,
		}}, nil
	}

	//取消关注
	if req.ActionType == 2 {
		_, err = l.svcCtx.FollowerRPC.UnFollow(l.ctx, &fpb.UnFollowReq{
			UserId:   userId,
			ToUserId: req.ToUserId,
		})
		if err != nil {
			logx.Error("FollowerRPC.UnFollow err:", err)
			return nil, err
		}
		return &types.FollowResp{StatusResponse: types.StatusResponse{
			StatusCode: int32(errorx.OK),
			StatusMsg:  errorx.SUCCESS,
		}}, nil
	}

	return &types.FollowResp{}, errorx.NewParamErr(errorx.ERRFOLLOWUSER)
}
