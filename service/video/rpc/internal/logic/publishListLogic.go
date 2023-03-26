package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

	"tik-tok-brief/service/video/rpc/internal/svc"
	"tik-tok-brief/service/video/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/jinzhu/copier"
)

type PublishListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查看某用户id的视频列表
func (l *PublishListLogic) PublishList(in *pb.PublishListReq) (*pb.PublishListResp, error) {
	videos, err := l.svcCtx.VideoModel.FindListByUserId(l.ctx, in.UserId)
	//数据库错误
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("videoModel_findListByUserId err:", err)
		return nil, errorx.NewStatusDBErr()
	}

	//封装video
	res := make([]*pb.Video, len(videos))
	err = copier.Copy(&res, videos)

	if err != nil {
		logx.Error("copier_copy err:", err)
		return nil, errorx.NewInternalErr()
	}

	return &pb.PublishListResp{VideoList: res}, nil
}
