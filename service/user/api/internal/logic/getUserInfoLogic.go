package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"tik-tok-brief/common/errorx"
	fpb "tik-tok-brief/service/follow/rpc/proto/pb"
	"tik-tok-brief/service/user/api/internal/svc"
	"tik-tok-brief/service/user/api/internal/types"
	"tik-tok-brief/service/user/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// 获取token里的user_id
	userID := l.ctx.Value("user_id")
	//获取请求的userID信息
	user, err := l.svcCtx.UserRPC.GetUser(l.ctx, &pb.GetUserReq{UserID: req.UserId})
	if err != nil {
		logx.Error("get_userinfo_logic get user err:", err)
		return nil, err
	}

	resUser := new(types.User)

	_ = copier.Copy(resUser, user.User)

	//查看关注信息
	infoResp, err := l.svcCtx.FollowRPC.GetFollowInfo(l.ctx, &fpb.GetFollowInfoReq{
		UserId:   userID.(int64),
		ToUserId: req.UserId,
	})
	if err != nil {
		logx.Error("FollowRPC.GetFollowInfo err:", err)
		return nil, err
	}
	if infoResp != nil && infoResp.IsFollow {
		resUser.IsFollow = true
	}

	return &types.UserInfoResp{
		StatusResponse: types.StatusResponse{
			StatusCode: int32(errorx.OK),
			StatusMsg:  errorx.SUCCESS,
		},
		User: *resUser,
	}, nil
}
