package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/chat/rpc/chat"
	cpb "tik-tok-brief/service/chat/rpc/proto/pb"
	fpb "tik-tok-brief/service/follow/rpc/proto/pb"
	"tik-tok-brief/service/user/rpc/proto/pb"

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
	userId := l.ctx.Value("user_id").(int64)
	list, err := l.svcCtx.FollowRPC.GetFriendsList(l.ctx, &fpb.GetFriendsListReq{UserId: userId})
	if err != nil {
		logx.Error("FollowRPC.GetFriendsList err:", err)
		return nil, err
	}

	ids := make([]int64, len(list.Follows))
	for i, follow := range list.Follows {
		ids[i] = follow.ToUserId
	}

	//获取users
	getUserListResp, err := l.svcCtx.UserRPC.GetUserListByIds(l.ctx, &pb.GetUserListByIdsReq{Ids: ids})
	if err != nil {
		logx.Error("UserRPC.GetUserListByIds err:", err)
		return nil, err
	}

	res := make([]types.FriendUser, len(getUserListResp.Users))
	err = copier.Copy(&res, getUserListResp.Users)
	if err != nil {
		logx.Error("copier Copy err:", err)
		return nil, errorx.NewInternalErr()
	}

	//获取最新消息
	var message *chat.GetLatestMessageResp
	for i := 0; i < len(res); i++ {
		message, err = l.svcCtx.ChatRPC.GetLatestMessage(l.ctx, &cpb.GetLatestMessageReq{
			FromUserId: userId,
			ToUserId:   res[i].UserId,
		})
		if err != nil {
			logx.Error("ChatRPC.GetLatestMessage err:", err)
			return nil, err
		}
		res[i].IsFollow = true
		//有消息
		if message.MsgType == 0 || message.MsgType == 1 {
			res[i].MsgType = message.MsgType
			res[i].Message = message.Message.Content
		}
	}

	return &types.FriendsListResp{
		StatusResponse: types.StatusResponse{
			StatusCode: int32(errorx.OK),
			StatusMsg:  errorx.SUCCESS,
		},
		UserList: res,
	}, nil
}
