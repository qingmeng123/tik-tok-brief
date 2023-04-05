package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

	"tik-tok-brief/service/user/rpc/internal/svc"
	"tik-tok-brief/service/user/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserFollowerCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserFollowerCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserFollowerCountLogic {
	return &UpdateUserFollowerCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户粉丝数
func (l *UpdateUserFollowerCountLogic) UpdateUserFollowerCount(in *pb.UpdateUserFollowerCountReq) (*pb.UpdateUserFollowerCountResp, error) {
	err := l.svcCtx.UserModel.UpdateFollowerCountByUserId(l.ctx, in.UserId, in.Number)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("UserModel.UpdateFollowerCountByUserId err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err == sqlx.ErrNotFound {
		return nil, errorx.NewStatusParamErr(errorx.ERRUSERID)
	}

	return &pb.UpdateUserFollowerCountResp{}, nil
}
