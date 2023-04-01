package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"tik-tok-brief/common/errorx"
	userpb "tik-tok-brief/service/user/rpc/proto/pb"
	"tik-tok-brief/service/video/rpc/proto/pb"
	"time"

	"tik-tok-brief/service/video/api/internal/svc"
	"tik-tok-brief/service/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedReq) (resp *types.FeedResp, err error) {
	// 是否登录,空值不能类型断言
	vUserId := l.ctx.Value("user_id")
	if vUserId == nil {

	} else {
		//tUserId:=vUserId.(int64)
	}
	req.LastTime = time.Now().Unix()
	//获取视频流
	feedResp, err := l.svcCtx.VideoRPC.Feed(l.ctx, &pb.FeedReq{
		LastTime: &req.LastTime,
	})
	if err != nil {
		logx.Error("videoRPC_feed err:", err)
		return nil, err
	}

	//封装信息
	resp = new(types.FeedResp)
	resp.VideoList = make([]types.Video, len(feedResp.VideoList))
	err = copier.Copy(&resp.VideoList, feedResp.VideoList)
	if err != nil {
		logx.Error("copier_copy err:", err)
		return nil, errorx.NewInternalErr()
	}

	//获取各个视频的作者
	for i, video := range feedResp.VideoList {
		userResp, err := l.svcCtx.UserRPC.GetUser(l.ctx, &userpb.GetUserReq{UserID: video.UserId})
		if err != nil {
			logx.Error("userRPC_getUser err:", err)
			return nil, err
		}
		err = copier.Copy(&resp.VideoList[i].Author, userResp.User)
		if err != nil {
			logx.Error("copier_copy err:", err)
			return nil, errorx.NewInternalErr()
		}
		resp.VideoList[i].PlayUrl = "https://www.w3schools.com/html/movie.mp4"
		resp.VideoList[i].CoverUrl = "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"
	}

	resp.StatusCode = errorx.OK
	resp.StatusMsg = errorx.SUCCESS

	return resp, nil
}
