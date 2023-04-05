package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"
	lpb "tik-tok-brief/service/like/rpc/proto/pb"

	"tik-tok-brief/service/video/rpc/internal/svc"
	"tik-tok-brief/service/video/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetVideoListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetVideoListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetVideoListByIdsLogic {
	return &GetVideoListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据videoIds获取视频列表
func (l *GetVideoListByIdsLogic) GetVideoListByIds(in *pb.GetVideoListByIdsReq) (*pb.GetVideoListByIdsResp, error) {
	if in.VideoIds == nil {
		return &pb.GetVideoListByIdsResp{}, errorx.NewStatusParamErr(errorx.ERRUSERID)
	}

	users, err := l.svcCtx.VideoModel.FindVideosByIds(l.ctx, in.VideoIds)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("UserModel.FindVideosByIds err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err == sqlx.ErrNotFound {
		return &pb.GetVideoListByIdsResp{}, nil
	}
	res := make([]*pb.Video, len(users))
	err = copier.Copy(&res, users)
	if err != nil {
		logx.Error("copier.Copy err:", err)
		return nil, errorx.NewInternalErr()
	}
	//填充关注信息
	if in.UserId != nil {
		for i := 0; i < len(res); i++ {
			isLikeResp, err := l.svcCtx.LikeRPC.IsLike(l.ctx, &lpb.IsLikeReq{
				UserId:  in.GetUserId(),
				VideoId: res[i].VideoId,
			})
			if err != nil {
				logx.Error("FollowerRPC.GetFollowInfo err:", err)
				return nil, errorx.NewStatusDBErr()
			}
			res[i].IsFavorite = isLikeResp.IsLike
		}
	}
	return &pb.GetVideoListByIdsResp{VideoList: res}, nil
}
