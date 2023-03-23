package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/common/tool"
	"tik-tok-brief/service/user/model"
	"tik-tok-brief/service/user/rpc/proto/pb"
	"time"

	"tik-tok-brief/service/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	logx.Debugf("req%#v",in)

	//验证是否存在该用户名
	u,err:=l.svcCtx.UserModel.FindOneByUsername(l.ctx,in.Username)
	if err!=nil&&err!=sqlx.ErrNotFound{
		logx.Error("user_register FindOneByUsername err:",err)
		return nil,errorx.NewStatusDBErr()
	}

	//查到该用户
	if u!=nil{
		return nil,errorx.NewStatusParamErr(errorx.ERRUSERNAME)

	}

	//没查到记录，注册新用户
	//密码加盐
	pwdHash,err:=tool.AddSalt(in.Password)
	if err!=nil{
		logx.Error("user_register AddSalt err:",err)
		return nil, errorx.NewStatusParamErr(errorx.ERRUSERNAME)
	}
	//生成userID(雪花算法)

	//创建用户
	user:=&model.User{
		UserId:time.Now().Unix(),
		Username: in.Username,
		Password: pwdHash,
	}
	_,err=l.svcCtx.UserModel.Insert(l.ctx,user)
	if err!=nil{
		logx.Error("user_register Insert err:",err)
		return nil, errorx.NewStatusDBErr()
	}

	//生成token
	gt:=NewGenerateTokenLogic(l.ctx,l.svcCtx)
	token,err:=gt.GenerateToken(&pb.GenerateTokenReq{UserID: user.UserId})
	if err!=nil{
		return nil, errorx.NewInternalErr()
	}
	return &pb.RegisterResp{UserID: user.UserId,Token: token.Token}, nil
}
