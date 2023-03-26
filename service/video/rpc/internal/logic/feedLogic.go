package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/common/tool"
	"time"

	"tik-tok-brief/service/video/rpc/internal/svc"
	"tik-tok-brief/service/video/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 视频流
func (l *FeedLogic) Feed(in *pb.FeedReq) (*pb.FeedResp, error) {
	//时间戳设置
	var err error
	ctime := time.Now()

	if in.LastTime != nil {
		ctime, err = tool.UnixToTime(in.GetLastTime(), "2006-01-02 15:04:05")
		if err != nil {
			logx.Error("tool_unixToTime err:", err)
			return nil, errorx.NewStatusParamErr(errorx.ERRTIMEPARAM)
		}
	}

	videos, err := l.svcCtx.VideoModel.FindListByCTimeLimit(l.ctx, ctime, l.svcCtx.Config.MaxVideoNum)
	//数据库错误
	if err != nil && err != sqlx.ErrNotFound {
		return nil, errorx.NewStatusDBErr()
	}

	//封装video
	res := make([]*pb.Video, len(videos))
	err = copier.Copy(&res, videos)
	if err != nil {
		logx.Error("copier_copy err:", err)
		return nil, errorx.NewInternalErr()
	}

	now := time.Now().Unix()
	return &pb.FeedResp{NextTime: &now, VideoList: res}, nil
}
