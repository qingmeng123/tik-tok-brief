package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

	"github.com/zeromicro/go-zero/core/logx"
	fpb "tik-tok-brief/service/follow/rpc/proto/pb"
	"tik-tok-brief/service/user/rpc/internal/svc"
	"tik-tok-brief/service/user/rpc/proto/pb"
)

type GetUserListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListByIdsLogic {
	return &GetUserListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量获取用户信息
func (l *GetUserListByIdsLogic) GetUserListByIds(in *pb.GetUserListByIdsReq) (*pb.GetUserListByIdsResp, error) {
	if in.Ids == nil {
		return &pb.GetUserListByIdsResp{}, errorx.NewStatusParamErr(errorx.ERRUSERID)
	}
	users, err := l.svcCtx.UserModel.FindUserByIds(l.ctx, in.Ids)
	if err != nil && err != sqlx.ErrNotFound {
		logx.Error("UserModel.FindUserByIds err:", err)
		return nil, errorx.NewStatusDBErr()
	}
	if err == sqlx.ErrNotFound {
		return &pb.GetUserListByIdsResp{}, nil
	}

	res := make([]*pb.User, len(users))

	err = copier.Copy(&res, users)
	if err != nil {
		logx.Error("copier.Copy err:", err)
		return nil, errorx.NewInternalErr()
	}
	//填充关注信息
	if in.UserId != nil {
		for i := 0; i < len(res); i++ {
			infoResp, err := l.svcCtx.FollowerRPC.GetFollowInfo(l.ctx, &fpb.GetFollowInfoReq{
				UserId:   in.GetUserId(),
				ToUserId: res[i].UserId,
			})
			if err != nil {
				logx.Error("FollowerRPC.GetFollowInfo err:", err)
				return nil, errorx.NewStatusDBErr()
			}
			if infoResp != nil {
				res[i].IsFollow = infoResp.IsFollow
			}
		}
	}

	return &pb.GetUserListByIdsResp{Users: res}, nil
}
