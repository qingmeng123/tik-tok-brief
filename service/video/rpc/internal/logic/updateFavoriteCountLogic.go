package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

	"tik-tok-brief/service/video/rpc/internal/svc"
	"tik-tok-brief/service/video/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateFavoriteCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateFavoriteCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateFavoriteCountLogic {
	return &UpdateFavoriteCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新视频点赞数
func (l *UpdateFavoriteCountLogic) UpdateFavoriteCount(in *pb.UpdateFavoriteCountReq) (*pb.UpdateFavoriteCountResp, error) {
	err := l.svcCtx.VideoModel.UpdateFavoriteCountByVideoId(l.ctx, in.VideoId, in.Number)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("VideoModel.UpdateFavoriteCountByVideoId err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err == sqlx.ErrNotFound {
		return nil, errorx.NewStatusParamErr(errorx.ERRVIDEOID)
	}
	return &pb.UpdateFavoriteCountResp{}, nil
}
