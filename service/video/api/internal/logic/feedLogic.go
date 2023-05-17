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
	userId := l.ctx.Value("user_id")

	if req.LastTime == 0 {
		req.LastTime = time.Now().Unix()
	}
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
	ids := make([]int64, len(feedResp.VideoList))
	for i := 0; i < len(feedResp.VideoList); i++ {
		ids[i] = feedResp.VideoList[i].UserId
	}

	//若登录还会获取关注和点赞信息
	if userId != nil {
		uid := userId.(int64)
		userListByIdsResp, err := l.svcCtx.UserRPC.GetUserListByIds(l.ctx, &userpb.GetUserListByIdsReq{
			UserId: &uid,
			Ids:    ids,
		})

		if err != nil {
			logx.Error("UserRPC.GetUserListByIds err:", err)
			return nil, err
		}

		//点赞信息
		vids := make([]int64, len(feedResp.VideoList))
		for i := 0; i < len(vids); i++ {
			vids[i] = feedResp.VideoList[i].VideoId
		}
		videoListByIdsResp, err := l.svcCtx.VideoRPC.GetVideoListByIds(l.ctx, &pb.GetVideoListByIdsReq{
			UserId:   &uid,
			VideoIds: vids,
		})
		if err != nil {
			logx.Error("VideoRPC.GetVideoListByIds err:", err)
			return nil, err
		}
		for i := 0; i < len(userListByIdsResp.Users); i++ {
			//点赞信息
			resp.VideoList[i].IsFavorite = videoListByIdsResp.VideoList[i].IsFavorite
			//作者信息
			err = copier.Copy(&(resp.VideoList[i].Author), userListByIdsResp.Users[i])
			if err != nil {
				logx.Error("copier.copy err:", err)
				return nil, errorx.NewInternalErr()
			}
		}

	} else {
		//未登录
		userListByIdsResp, err := l.svcCtx.UserRPC.GetUserListByIds(l.ctx, &userpb.GetUserListByIdsReq{
			Ids: ids,
		})

		if err != nil {
			logx.Error("UserRPC.GetUserListByIds err:", err)
			return nil, err
		}

		for i := 0; i < len(userListByIdsResp.Users); i++ {
			//作者信息
			err = copier.Copy(&(resp.VideoList[i].Author), userListByIdsResp.Users[i])
			if err != nil {
				logx.Error("copier.copy err:", err)
				return nil, errorx.NewInternalErr()
			}
		}
	}
	//获取视频最后时间
	resp.NextTime = feedResp.GetNextTime()
	resp.StatusCode = errorx.OK
	resp.StatusMsg = errorx.SUCCESS

	return resp, nil
}
