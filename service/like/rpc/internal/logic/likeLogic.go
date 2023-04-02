package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/like/model"
	"tik-tok-brief/service/like/rpc/internal/svc"
	"tik-tok-brief/service/like/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 赞操作
func (l *LikeLogic) Like(in *pb.LikeReq) (*pb.LikeListResp, error) {
	like, err := l.svcCtx.LikeModel.FindLikeByUserIdVideoId(l.ctx, in.UserId, in.VideoId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("LikeModel FindLikeByUserIdVideoId err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	//点赞
	if err == sqlx.ErrNotFound && in.ActionType == 1 {
		_, err = l.svcCtx.LikeModel.Insert(l.ctx, &model.Like{
			UserId:  0,
			VideoId: 0,
		})
		if err != nil {
			logx.Error(".LikeModel.Insert err:", err)
			return nil, errorx.NewStatusDBErr()
		}
		return &pb.LikeListResp{}, nil
	}

	//取消点赞
	if err != sqlx.ErrNotFound && in.ActionType == 2 {
		err = l.svcCtx.LikeModel.Delete(l.ctx, like.Id)
		if err != nil {
			logx.Error("LikeModel.Delete err:", err)
			return nil, errorx.NewStatusDBErr()
		}
	}

	//操作错误
	return &pb.LikeListResp{}, errorx.NewParamErr(errorx.ERRLIKE)
}
