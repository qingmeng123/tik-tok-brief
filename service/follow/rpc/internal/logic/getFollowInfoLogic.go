package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

	"tik-tok-brief/service/follow/rpc/internal/svc"
	"tik-tok-brief/service/follow/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowInfoLogic {
	return &GetFollowInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取关注信息
func (l *GetFollowInfoLogic) GetFollowInfo(in *pb.GetFollowInfoReq) (*pb.GetFollowInfoResp, error) {
	follow, err := l.svcCtx.FollowModel.FindFollowByUsersId(l.ctx, in.UserId, in.ToUserId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("FollowModel.FindFollowByUsersId err:", err)
		return nil, errorx.NewStatusDBErr()
	}

	if err == sqlx.ErrNotFound {
		return &pb.GetFollowInfoResp{}, nil
	}

	flag := false
	if follow.IsFriend {
		flag = true
	}
	return &pb.GetFollowInfoResp{Follow: &pb.Follow{
		Id:       follow.Id,
		UserId:   follow.UserId,
		ToUserId: follow.ToUserId,
		IsFriend: flag,
	}}, nil
}
