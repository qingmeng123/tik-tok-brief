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

type GetFansListByUserIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFansListByUserIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFansListByUserIdLogic {
	return &GetFansListByUserIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取粉丝列表
func (l *GetFansListByUserIdLogic) GetFansListByUserId(in *pb.GetFansListByUserIdReq) (*pb.GetFansListByUserIdResp, error) {
	followers, err := l.svcCtx.FollowModel.FindFollowersByToUserId(l.ctx, in.UserId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("FollowModel.FindFollowersByToUserId err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err == sqlx.ErrNotFound {
		return &pb.GetFansListByUserIdResp{}, nil
	}
	res := make([]*pb.Follow, len(followers))
	err = copier.Copy(&res, followers)
	if err != nil {
		logx.Error("copier_copy err:", err)
		return nil, errorx.NewInternalErr()
	}

	return &pb.GetFansListByUserIdResp{Follows: res}, nil
}
