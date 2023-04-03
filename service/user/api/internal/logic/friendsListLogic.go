package logic

import (
	"context"

	"tik-tok-brief/service/user/api/internal/svc"
	"tik-tok-brief/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendsListLogic {
	return &FriendsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendsListLogic) FriendsList(req *types.FriendsListReq) (resp *types.FriendsListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
