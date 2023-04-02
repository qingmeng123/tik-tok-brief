package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/like/rpc/internal/svc"
	"tik-tok-brief/service/like/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeListLogic {
	return &LikeListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取点赞视频列表
func (l *LikeListLogic) LikeList(in *pb.LikeListReq) (*pb.LikeListResp, error) {
	likes, err := l.svcCtx.LikeModel.FindLikeListByUserId(l.ctx, in.UserId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("LikeModel.FindLikeListByUserId err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err != sqlx.ErrNotFound {
		res := make([]*pb.Like, len(likes))
		copier.Copy(&res, likes)
		return &pb.LikeListResp{Likes: res}, nil
	}

	return &pb.LikeListResp{}, nil
}
