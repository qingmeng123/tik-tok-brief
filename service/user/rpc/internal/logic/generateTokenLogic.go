package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"tik-tok-brief/service/user/rpc/proto/pb"
	"time"

	"tik-tok-brief/service/user/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyClaims struct {
	UserId int64 `json:"user_id"`
	jwt.RegisteredClaims
}

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取token
func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	expire := l.svcCtx.Config.JwtAuth.AccessExpire
	token, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret,expire, in.UserID)
	if err!=nil{
		logx.Error("generate token err:",err)
		return nil, err
	}
	return &pb.GenerateTokenResp{Token: token}, nil
}


func (l *GenerateTokenLogic) getJwtToken(secretKey string, expire, userId int64) (string, error) {
	iat := time.Now()
	claim := MyClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(iat.Add(time.Second * time.Duration(expire))),
			IssuedAt:  jwt.NewNumericDate(iat),
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	return token.SignedString([]byte(secretKey))
}