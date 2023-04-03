package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

	"tik-tok-brief/service/comment/rpc/internal/svc"
	"tik-tok-brief/service/comment/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentListLogic {
	return &CommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 某视频的评论列表
func (l *CommentListLogic) CommentList(in *pb.CommentListReq) (*pb.CommentListResp, error) {
	likes, err := l.svcCtx.CommentModel.FindCommentListByVideoId(l.ctx, in.VideoId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("CommentModel.FindCommentListByVideoId err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err != sqlx.ErrNotFound {
		res := make([]*pb.Comment, len(likes))
		copier.Copy(&res, likes)
		for i := 0; i < len(res); i++ {
			res[i].CreateTime=likes[i].CreateTime.Format("2006-01-02 15:04:05")
		}
		return &pb.CommentListResp{Comments: res}, nil
	}

	return &pb.CommentListResp{}, nil
}
