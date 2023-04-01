package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

	"tik-tok-brief/service/follow/rpc/internal/svc"
	"tik-tok-brief/service/follow/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFriendsListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFriendsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFriendsListLogic {
	return &GetFriendsListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取好友列表
func (l *GetFriendsListLogic) GetFriendsList(in *pb.GetFriendsListReq) (*pb.GetFriendsListResp, error) {
	follows, err := l.svcCtx.FollowModel.FindFriendsByUserId(l.ctx, in.UserId)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("FollowModel.FindFriendsByUserId err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err == sqlx.ErrNotFound {
		return nil, nil
	}
	res := make([]*pb.Follow, len(follows))
	err = copier.Copy(&res, follows)
	if err != nil {
		logx.Error("copier_copy err:", err)
		return nil, errorx.NewInternalErr()
	}

	return &pb.GetFriendsListResp{Follows: res}, nil
}
