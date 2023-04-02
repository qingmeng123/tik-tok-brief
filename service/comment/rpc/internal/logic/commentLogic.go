package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/comment/model"
	"tik-tok-brief/service/comment/rpc/internal/svc"
	"tik-tok-brief/service/comment/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommentLogic {
	return &CommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 评论操作
func (l *CommentLogic) Comment(in *pb.CommentReq) (*pb.CommentResp, error) {
	//发布评论
	if in.ActionType == 1 && in.Content != nil {
		result, err := l.svcCtx.CommentModel.Insert(l.ctx, &model.Comment{
			UserId:  in.UserId,
			VideoId: in.VideoId,
			Content: *in.Content,
		})
		if err != nil {
			logx.Error("CommentModel.Insert err:", err)
			return nil, errorx.NewStatusDBErr()
		}
		res := new(pb.Comment)
		copier.Copy(res, &result)
		return &pb.CommentResp{Comment: res}, nil
	}

	//删除评论
	if in.ActionType == 2 && in.CommentId != nil {
		err := l.svcCtx.CommentModel.Delete(l.ctx, *in.CommentId)
		if err != nil {
			logx.Error("CommentModel.Delete err:", err)
			return nil, errorx.NewStatusDBErr()
		}
		return &pb.CommentResp{}, nil
	}

	//参数错误
	return &pb.CommentResp{}, errorx.NewParamErr(errorx.ERRCOMMENT)
}
