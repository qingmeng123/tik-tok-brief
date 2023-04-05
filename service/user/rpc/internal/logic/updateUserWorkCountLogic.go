package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

	"tik-tok-brief/service/user/rpc/internal/svc"
	"tik-tok-brief/service/user/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserWorkCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserWorkCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserWorkCountLogic {
	return &UpdateUserWorkCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户作品数
func (l *UpdateUserWorkCountLogic) UpdateUserWorkCount(in *pb.UpdateUserWorkCountReq) (*pb.UpdateUserWorkCountResp, error) {
	err := l.svcCtx.UserModel.UpdateWorkCountByUserId(l.ctx, in.UserId, in.Number)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("UserModel.UpdateWorkCountByUserId err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err == sqlx.ErrNotFound {
		return nil, errorx.NewStatusParamErr(errorx.ERRUSERID)
	}

	return &pb.UpdateUserWorkCountResp{}, nil
}
