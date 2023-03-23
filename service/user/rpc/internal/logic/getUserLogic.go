package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

	"tik-tok-brief/service/user/rpc/internal/svc"
	"tik-tok-brief/service/user/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *GetUserLogic) GetUser(in *pb.GetUserReq) (*pb.GetUserResp, error) {
	//验证请求的用户ID
	user, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, in.UserID)
	if err!=nil&&err!=sqlx.ErrNotFound{
		logx.Error("get_user FindOneByUserId err:",err)
		return nil, errorx.NewStatusDBErr()
	}

	if err!=nil{
		return nil, errorx.NewStatusParamErr(errorx.ERRUSERID)
	}

	res:=&pb.User{
		UserID:user.UserId,
		UserName: user.Username,

	}

	return &pb.GetUserResp{User: res}, nil
}
