package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/follow/model"
	"tik-tok-brief/service/follow/rpc/internal/svc"
	"tik-tok-brief/service/follow/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关注
func (l *FollowLogic) Follow(in *pb.FollowReq) (*pb.FollowResp, error) {
	//检查是否已关注对方
	_, err := l.svcCtx.FollowModel.FindFollowByUsersId(l.ctx, in.UserId, in.ToUserId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("FollowModel.FindFollowByUsersId err:", err)
		return nil, errorx.NewStatusDBErr()
	}

	if err == nil {
		//已关注
		return nil, errorx.NewStatusParamErr(errorx.ERRFOLLOWUSER)
	}

	//查看对方是否关注自己
	follower, err := l.svcCtx.FollowModel.FindFollowByUsersId(l.ctx, in.ToUserId, in.UserId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("FollowModel.FindFollowByUsersId err:", err)
		return nil, errorx.NewStatusDBErr()
	}

	if err != nil {
		//对方未关注自己
		_, err = l.svcCtx.FollowModel.Insert(l.ctx, &model.Follow{
			UserId:   in.UserId,
			ToUserId: in.ToUserId,
			IsFriend: false,
		})
		return &pb.FollowResp{}, nil
	}

	//对方已关注自己，成为朋友
	_, err = l.svcCtx.FollowModel.Insert(l.ctx, &model.Follow{
		UserId:   in.UserId,
		ToUserId: in.ToUserId,
		IsFriend: true,
	})
	if err != nil {
		logx.Error("FollowModel.Insert err:", err)
		return nil, errorx.NewStatusDBErr()
	}

	//对方也修改为朋友
	follower.IsFriend = true
	err = l.svcCtx.FollowModel.Update(l.ctx, follower)

	if err != nil {
		logx.Error("FollowModel.Update err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	return &pb.FollowResp{}, nil
}
