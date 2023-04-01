package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

	"tik-tok-brief/service/follow/rpc/internal/svc"
	"tik-tok-brief/service/follow/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowListByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFollowListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowListByUserIdLogic {
	return &GetFollowListByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取关注列表
func (l *GetFollowListByUserIdLogic) GetFollowListByUserId(in *pb.GetFollowListByUserIdReq) (*pb.GetFollowListByUserIdResp, error) {
	follows, err := l.svcCtx.FollowModel.FindFollowsByUserId(l.ctx, in.UserId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("FollowModel.FindFollowsByUserId err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err == sqlx.ErrNotFound {
		return nil, nil
	}
	res := make([]*pb.Follow, len(follows))
	err = copier.Copy(&res, follows)
	if err != nil {
		logx.Error("copier_copy err:", err)
		return nil, errorx.NewInternalErr()
	}

	return &pb.GetFollowListByUserIdResp{Follows: res}, nil
}
