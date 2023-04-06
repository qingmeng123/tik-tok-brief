package logic

import (
	"context"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/common/snowflake"
	"tik-tok-brief/service/video/model"
	"tik-tok-brief/service/video/rpc/internal/svc"
	"tik-tok-brief/service/video/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveVideoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveVideoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveVideoLogic {
	return &SaveVideoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 保存视频信息到数据库
func (l *SaveVideoLogic) SaveVideo(in *pb.SaveVideoReq) (*pb.SaveVideoResp, error) {
	//生成videoId
	//创建video
	vid, err := snowflake.GetID()
	if err != nil {
		logx.Error("snowflake.GetID err:", err)
		return nil, errorx.NewInternalErr()
	}
	video := &model.Video{
		VideoId:  int64(vid),
		UserId:   in.UserId,
		Title:    in.Title,
		PlayUrl:  in.PlayUrl,
		CoverUrl: in.CoverUrl,
	}
	//保存到数据库
	_, err = l.svcCtx.VideoModel.Insert(l.ctx, video)
	if err != nil {
		logx.Error("videoModel_insert err:", err)
		return nil, errorx.NewStatusDBErr()
	}

	return &pb.SaveVideoResp{VideoId: video.VideoId}, nil
}
