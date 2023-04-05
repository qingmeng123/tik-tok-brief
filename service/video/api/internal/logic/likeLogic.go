package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"tik-tok-brief/common/errorx"
	lpb "tik-tok-brief/service/like/rpc/proto/pb"
	"tik-tok-brief/service/video/api/internal/svc"
	"tik-tok-brief/service/video/api/internal/types"
	vpb "tik-tok-brief/service/video/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeLogic {
	return &LikeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeLogic) Like(req *types.LikeReq) (resp *types.LikeResp, err error) {
	userId := l.ctx.Value("user_id").(int64)
	//赞操作
	_, err = l.svcCtx.LikeRPC.Like(l.ctx, &lpb.LikeReq{
		UserId:     userId,
		VideoId:    req.VideoId,
		ActionType: int32(req.ActionType),
	})
	if err != nil {
		logx.Error("LikeRPC.Like err:", err)
		return nil, err
	}

	//修改点赞数
	var number int64
	if req.ActionType == 1 {
		number = 1
	}
	number = -1

	_, err = l.svcCtx.VideoRPC.UpdateFavoriteCount(l.ctx, &vpb.UpdateFavoriteCountReq{
		VideoId: req.VideoId,
		Number:  number,
	})
	if err != nil {
		logx.Error("VideoRPC.UpdateFavoriteCount err:", err)
		return nil, err
	}
	return &types.LikeResp{Status: types.Status{
		StatusCode: uint32(codes.OK),
		StatusMsg:  errorx.SUCCESS,
	}}, nil

}
