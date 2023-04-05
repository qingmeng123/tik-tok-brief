package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/like/rpc/proto/pb"
	upb "tik-tok-brief/service/user/rpc/proto/pb"
	vpb "tik-tok-brief/service/video/rpc/proto/pb"

	"tik-tok-brief/service/video/api/internal/svc"
	"tik-tok-brief/service/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeListLogic {
	return &LikeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeListLogic) LikeList(req *types.LikeListReq) (resp *types.LikeListResp, err error) {
	userId := l.ctx.Value("user_id").(int64)
	list, err := l.svcCtx.LikeRPC.LikeList(l.ctx, &pb.LikeListReq{
		UserId: req.UserId,
	})
	if err != nil {
		logx.Error("LikeRPC.LikeList err:", err)
		return nil, err
	}

	ids := make([]int64, len(list.Likes))
	for i, follow := range list.Likes {
		ids[i] = follow.VideoId
	}
	//获取请求用户Id的点赞视频
	videoListByIdsResp, err := l.svcCtx.VideoRPC.GetVideoListByIds(l.ctx, &vpb.GetVideoListByIdsReq{
		UserId:   &req.UserId,
		VideoIds: ids,
	})
	if err != nil {
		logx.Error("VideoRPC.GetVideoListByIds err:", err)
		return nil, err
	}

	res := make([]types.Video, len(videoListByIdsResp.VideoList))
	_ = copier.Copy(&res, videoListByIdsResp.VideoList)

	//获取视频作者
	uids := make([]int64, len(res))
	for i, video := range videoListByIdsResp.VideoList {
		uids[i] = video.UserId
	}

	userListByIdsResp, err := l.svcCtx.UserRPC.GetUserListByIds(l.ctx, &upb.GetUserListByIdsReq{
		UserId: &userId,
		Ids:    uids,
	})
	if err != nil {
		logx.Error("UserRPC.GetUserListByIds err:", err)
		return nil, err
	}

	for i := 0; i < len(res); i++ {
		_ = copier.Copy(&(res[i].Author), userListByIdsResp.Users[i])
	}
	return &types.LikeListResp{
		Status: types.Status{
			StatusCode: uint32(codes.OK),
			StatusMsg:  errorx.SUCCESS,
		},
		VideoList: res,
	}, nil
}
