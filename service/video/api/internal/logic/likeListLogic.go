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
	//获取点赞视频的详细信息
	videoListByIdsResp, err := l.svcCtx.VideoRPC.GetVideoListByIds(l.ctx, &vpb.GetVideoListByIdsReq{
		UserId:   &userId,
		VideoIds: ids,
	})
	if err != nil {
		logx.Error("VideoRPC.GetVideoListByIds err:", err)
		return nil, err
	}

	res := make([]types.Video, len(list.Likes))
	_ = copier.Copy(&res, videoListByIdsResp.VideoList)

	//获取视频作者
	uids := make([]int64, len(list.Likes))
	for i := 0; i < len(uids); i++ {
		uids[i] = videoListByIdsResp.VideoList[i].UserId
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
		res[i].Author.UserId = videoListByIdsResp.VideoList[i].UserId
	}

	SetAuthorInfo(res, userListByIdsResp.Users)
	return &types.LikeListResp{
		Status: types.Status{
			StatusCode: uint32(codes.OK),
			StatusMsg:  errorx.SUCCESS,
		},
		VideoList: res,
	}, nil
}

func SetAuthorInfo(res []types.Video, authors []*upb.User) {
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(authors); j++ {
			if authors[j].UserId == res[i].Author.UserId {
				_ = copier.Copy(&res[i].Author, authors[j])
				break
			}
		}
	}
}
