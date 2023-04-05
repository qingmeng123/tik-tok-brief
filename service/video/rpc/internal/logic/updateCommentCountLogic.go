package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

	"tik-tok-brief/service/video/rpc/internal/svc"
	"tik-tok-brief/service/video/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentCountLogic {
	return &UpdateCommentCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新视频评论数
func (l *UpdateCommentCountLogic) UpdateCommentCount(in *pb.UpdateCommentCountReq) (*pb.UpdateFavoriteCountResp, error) {
	err := l.svcCtx.VideoModel.UpdateCommentCountByVideoId(l.ctx, in.VideoId, in.Number)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("VideoModel.UpdateCommentCountByVideoId err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err == sqlx.ErrNotFound {
		return nil, errorx.NewStatusParamErr(errorx.ERRVIDEOID)
	}

	return &pb.UpdateFavoriteCountResp{}, nil
}
