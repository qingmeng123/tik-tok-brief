package logic

import (
	"context"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/user/rpc/proto/pb"

	"tik-tok-brief/service/user/api/internal/svc"
	"tik-tok-brief/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	loginResp,err:=l.svcCtx.UserRPC.Login(l.ctx,&pb.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err!=nil{
		logx.Error("login userRpc.login err:",err)
		return nil, err
	}

	return &types.LoginResp{
		StatusResponse: types.StatusResponse{
			StatusCode: int32(errorx.OK),
			StatusMsg:  errorx.SUCCESS,
		},
		UserId:         loginResp.UserId,
		Token:          loginResp.Token,
	},nil
}
