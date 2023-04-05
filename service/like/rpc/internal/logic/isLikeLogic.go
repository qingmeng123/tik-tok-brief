package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

	"tik-tok-brief/service/like/rpc/internal/svc"
	"tik-tok-brief/service/like/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsLikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsLikeLogic {
	return &IsLikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取UID是否对VID点赞
func (l *IsLikeLogic) IsLike(in *pb.IsLikeReq) (*pb.IsLikeResp, error) {
	_, err := l.svcCtx.LikeModel.FindLikeByUserIdVideoId(l.ctx, in.UserId, in.VideoId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("LikeModel.FindLikeByUserIdVideoId err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err == sqlx.ErrNotFound {
		return &pb.IsLikeResp{IsLike: false}, nil
	}

	return &pb.IsLikeResp{IsLike: true}, nil
}
