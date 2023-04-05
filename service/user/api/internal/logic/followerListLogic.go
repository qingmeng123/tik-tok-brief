package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"tik-tok-brief/common/errorx"
	fpb "tik-tok-brief/service/follow/rpc/proto/pb"
	"tik-tok-brief/service/user/rpc/proto/pb"

	"tik-tok-brief/service/user/api/internal/svc"
	"tik-tok-brief/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowerListLogic {
	return &FollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FollowerListLogic) FollowerList(req *types.FollowerListReq) (resp *types.FollowerListResp, err error) {
	userId := l.ctx.Value("user_id").(int64)
	list, err := l.svcCtx.FollowRPC.GetFansListByUserId(l.ctx, &fpb.GetFansListByUserIdReq{UserId: req.UserId})
	if err != nil {
		logx.Error("FollowRPC.GetFansListByUserId err:", err)
		return nil, err
	}

	ids := make([]int64, len(list.Follows))
	for i, follow := range list.Follows {
		ids[i] = follow.UserId
	}

	//获取users
	getUserListResp, err := l.svcCtx.UserRPC.GetUserListByIds(l.ctx, &pb.GetUserListByIdsReq{UserId: &userId, Ids: ids})
	if err != nil {
		logx.Error("UserRPC.GetUserListByIds err:", err)
		return nil, err
	}

	res := make([]types.User, len(getUserListResp.Users))
	err = copier.Copy(&res, getUserListResp)
	if err != nil {
		logx.Error("copier Copy err:", err)
		return nil, errorx.NewInternalErr()
	}

	return &types.FollowerListResp{
		StatusResponse: types.StatusResponse{
			StatusCode: int32(errorx.OK),
			StatusMsg:  errorx.SUCCESS,
		},
		UserList: res,
	}, nil
}
