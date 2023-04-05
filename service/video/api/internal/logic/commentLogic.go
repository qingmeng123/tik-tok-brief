package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"tik-tok-brief/common/errorx"
	cpb "tik-tok-brief/service/comment/rpc/proto/pb"
	vpb "tik-tok-brief/service/video/rpc/proto/pb"

	"tik-tok-brief/service/video/api/internal/svc"
	"tik-tok-brief/service/video/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommentLogic) Comment(req *types.CommentReq) (resp *types.CommentResp, err error) {
	userId := l.ctx.Value("user_id").(int64)
	//评论操作
	CommentResp, err := l.svcCtx.CommentRPC.Comment(l.ctx, &cpb.CommentReq{
		UserId:     userId,
		VideoId:    req.VideoId,
		ActionType: int32(req.ActionType),
		Content:    &req.CommentText,
	})
	if err != nil {
		logx.Error("CommentRPC.Comment err:", err)
		return nil, err
	}
	comment := new(types.Comment)

	//修改评论数
	var number int64
	number = -1
	if req.ActionType == 1 {
		number = 1
		_ = copier.Copy(comment, CommentResp.Comment)
	}

	_, err = l.svcCtx.VideoRPC.UpdateCommentCount(l.ctx, &vpb.UpdateCommentCountReq{
		VideoId: req.VideoId,
		Number:  number,
	})
	if err != nil {
		logx.Error("VideoRPC.UpdateCommentCount err:", err)
		return nil, err
	}

	return &types.CommentResp{
		Status: types.Status{
			StatusCode: uint32(codes.OK),
			StatusMsg:  errorx.SUCCESS,
		},
		Comment: *comment,
	}, nil

}
