package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

	"tik-tok-brief/service/user/rpc/internal/svc"
	"tik-tok-brief/service/user/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserFollowCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserFollowCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserFollowCountLogic {
	return &UpdateUserFollowCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户关注数
func (l *UpdateUserFollowCountLogic) UpdateUserFollowCount(in *pb.UpdateUserFollowCountReq) (*pb.UpdateUserFollowCountResp, error) {
	err := l.svcCtx.UserModel.UpdateFollowCountByUserId(l.ctx, in.UserId, in.Number)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("UserModel.UpdateFollowCountByUserId err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err == sqlx.ErrNotFound {
		return nil, errorx.NewStatusParamErr(errorx.ERRUSERID)
	}

	return &pb.UpdateUserFollowCountResp{}, nil
}
