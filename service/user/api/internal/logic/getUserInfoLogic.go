package logic

import (
	"context"
	"tik-tok-brief/common/errorx"
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
	userID:=l.ctx.Value("user_id")

	//获取请求的userID信息
	user, err := l.svcCtx.UserRPC.GetUser(l.ctx, &pb.GetUserReq{UserID: req.UserId})
	if err!=nil{
		logx.Error("get_userinfo_logic get user err:",err)
		return nil, err
	}

	resUser:=&types.User{
		Id:            user.User.UserID,
		Username:        user.User.UserName,
		FollowCount:     user.User.FollowCount,
		FollowerCount:   user.User.FollowerCount,
		IsFollow:        true,
		Avatar:          user.User.Avatar,
		BackgroundImage: user.User.BackgroundImage,
		Signature:       user.User.Signature,
		TotalFavorited:  user.User.TotalFavorited,
		WorkCount:       user.User.WorkCount,
		FavoriteCount:   user.User.FavoriteCount,
	}

	//查看自身的信息
	if userID==req.UserId{
		resUser.IsFollow=false
	}


	return &types.UserInfoResp{
		StatusResponse: types.StatusResponse{
			StatusCode: int32(errorx.OK),
			StatusMsg:  errorx.SUCCESS,
		},
		User:           *resUser,
	},nil
}
