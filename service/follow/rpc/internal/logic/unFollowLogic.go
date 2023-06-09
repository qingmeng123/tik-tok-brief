package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/follow/rpc/internal/svc"
	"tik-tok-brief/service/follow/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnFollowLogic {
	return &UnFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取消关注
func (l *UnFollowLogic) UnFollow(in *pb.UnFollowReq) (*pb.UnFollowResp, error) {
	//检查是否关注对方
	follow, err := l.svcCtx.FollowModel.FindFollowByUsersId(l.ctx, in.UserId, in.ToUserId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("FollowModel.FindFollowByUsersId err:", err)
		return nil, errorx.NewStatusDBErr()
	}

	if err != nil {
		//未关注
		return nil, errorx.NewStatusParamErr(errorx.ERRFOLLOWUSER)
	}

	//取消关注
	err = l.svcCtx.FollowModel.Delete(l.ctx, follow.Id)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("FollowModel.Delete err:", err)
		return nil, errorx.NewStatusDBErr()
	}

	//查看对方是否关注自己
	follower, err := l.svcCtx.FollowModel.FindFollowByUsersId(l.ctx, in.ToUserId, in.UserId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("FollowModel.FindFollowByUsersId err:", err)
		return nil, errorx.NewStatusDBErr()
	}

	//对方未关注自己，直接返回
	if err != nil {
		return &pb.UnFollowResp{}, nil
	}

	//对方已关注自己，取消朋友关系
	follower.IsFriend = false
	err = l.svcCtx.FollowModel.Update(l.ctx, follower)
	if err != nil {
		logx.Error("FollowModel.Update err:", err)
		return nil, errorx.NewStatusDBErr()
	}

	return &pb.UnFollowResp{}, nil
}
