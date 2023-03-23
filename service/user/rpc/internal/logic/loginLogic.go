package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/common/tool"

	"tik-tok-brief/service/user/rpc/internal/svc"
	"tik-tok-brief/service/user/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录
func (l *LoginLogic) Login(in *pb.LoginReq) (*pb.LoginResp, error) {
	//验证用户名
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err!=nil&&err!=sqlx.ErrNotFound{
		logx.Error("user login FindOneByUsername err:",err)
		return nil, errorx.NewStatusDBErr()
	}

	if err!=nil{
		return nil, errorx.NewStatusParamErr(errorx.ERRUSERNAME)
	}

	//验证密码
	ok := tool.CheckPassword(user.Password, in.Password)
	if !ok{
		return nil, errorx.NewStatusParamErr(errorx.ERRUSERPASSWORD)
	}

	//生成token
	gt:=NewGenerateTokenLogic(l.ctx,l.svcCtx)
	token, err := gt.GenerateToken(&pb.GenerateTokenReq{UserID: user.UserId})
	if err!=nil{
		return nil, errorx.NewInternalErr()
	}
	return &pb.LoginResp{UserId: user.UserId,Token: token.Token}, nil
}
