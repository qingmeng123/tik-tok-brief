package logic

import (
	"context"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/user/api/internal/svc"
	"tik-tok-brief/service/user/api/internal/types"
	"tik-tok-brief/service/user/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	logx.Debugf("req%#v",req)

	//调用用户注册微服务
	registerResp, err := l.svcCtx.UserRPC.Register(l.ctx, &pb.RegisterReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err!=nil{
		logx.Error("rpc register err:",err)
		return nil,err
	}

	//返回结果
	return &types.RegisterResponse{
		StatusResponse:types.StatusResponse{
			StatusCode: int32(errorx.OK),
			StatusMsg:  errorx.SUCCESS,
		},
		UserID:     registerResp.UserID,
		Token:      registerResp.Token,
	},nil
}
