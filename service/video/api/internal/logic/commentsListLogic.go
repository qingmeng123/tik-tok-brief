package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/comment/rpc/proto/pb"
	upb "tik-tok-brief/service/user/rpc/proto/pb"
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
	for i := 0; i < len(res); i++ {
		res[i].CreateDate = listResp.Comments[i].CreateTime
	}

	//获取评论用户信息
	ids := make([]int64, len(res))
	for i := 0; i < len(ids); i++ {
		ids[i] = listResp.Comments[i].UserId
	}
	userListByIdsResp, err := l.svcCtx.UserRPC.GetUserListByIds(l.ctx, &upb.GetUserListByIdsReq{
		Ids: ids,
	})
	if err != nil {
		logx.Error("UserRPC.GetUserListByIds err:", err)
		return nil, err
	}
	for i := 0; i < len(res); i++ {
		res[i].User.UserId = listResp.Comments[i].UserId
	}
	SetCommentAuthorInfo(res, userListByIdsResp.Users)

	return &types.CommentsListResp{
		Status: types.Status{
			StatusCode: uint32(codes.OK),
			StatusMsg:  errorx.SUCCESS,
		},
		CommentsList: res,
	}, nil
}

func SetCommentAuthorInfo(res []types.Comment, authors []*upb.User) {
	for i := 0; i < len(res); i++ {
		for j := 0; j < len(authors); j++ {
			if authors[j].UserId == res[i].User.UserId {
				_ = copier.Copy(&res[i].User, authors[j])
				break
			}
		}
	}
}
