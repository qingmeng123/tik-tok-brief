package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/comment/rpc/proto/pb"

	"tik-tok-brief/service/video/api/internal/svc"
	"tik-tok-brief/service/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentsListLogic {
	return &CommentsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentsListLogic) CommentsList(req *types.CommentsListReq) (resp *types.CommentsListResp, err error) {
	listResp, err := l.svcCtx.CommentRPC.CommentList(l.ctx, &pb.CommentListReq{VideoId: req.VideoId})
	if err != nil {
		logx.Error("CommentRPC.CommentList err:", err)
		return nil, err
	}
	res := make([]types.Comment, len(listResp.Comments))
	_ = copier.Copy(&res, listResp.Comments)
	return &types.CommentsListResp{
		Status: types.Status{
			StatusCode: uint32(codes.OK),
			StatusMsg:  errorx.SUCCESS,
		},
		CommentsList: res,
	}, nil
}
