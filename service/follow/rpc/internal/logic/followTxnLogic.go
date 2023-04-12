package logic

import (
	"context"
	"database/sql"

	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/service/follow/model"

	"tik-tok-brief/service/follow/rpc/internal/svc"
	"tik-tok-brief/service/follow/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowTxnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowTxnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowTxnLogic {
	return &FollowTxnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关注操作（事务处理）
func (l *FollowTxnLogic) FollowTxn(in *pb.FollowTxnReq) (*pb.FollowTxnResp, error) {

	// 获取RawDB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		logx.Error("NewMysql RawDB err:", err)
		return nil, errorx.NewStatusDBErr()
	}

	//获取子事务屏障对象
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		logx.Error("dtmgrpc.BarrierFromGrpc err:", err)
		return nil, errorx.NewStatusDBErr()
	}

	//关注操作
	if in.ActionType == 1 {
		//开启子事务屏障
		err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
			_, err = l.svcCtx.FollowModel.FindFollowByUsersId(l.ctx, in.UserId, in.ToUserId)
			if err != nil && err != sqlx.ErrNotFound {
				logx.Error("FollowModel.FindFollowByUsersId err:", err)
				return errorx.NewStatusDBErr()
			}

			//已关注过
			if err == nil {
				return errorx.NewStatusParamTxErr(errorx.ERRUSERID)
			}

			//查看对方是否关注自己
			follower, err := l.svcCtx.FollowModel.FindFollowByUsersId(l.ctx, in.ToUserId, in.UserId)
			if err != nil && err != sqlx.ErrNotFound {
				logx.Error("FollowModel.FindFollowByUsersId err:", err)
				return errorx.NewStatusDBErr()
			}

			if err != nil {
				//对方未关注自己
				_, err = l.svcCtx.FollowModel.TxInsert(tx, &model.Follow{
					UserId:   in.UserId,
					ToUserId: in.ToUserId,
					IsFriend: false,
				})
				if err != nil {
					logx.Error("FollowModel.TxInsert err:", err)

					return errorx.NewStatusDBErr()
				}
				//子事务返回nil
				return nil
			}

			//对方已关注自己，成为朋友
			_, err = l.svcCtx.FollowModel.TxInsert(tx, &model.Follow{
				UserId:   in.UserId,
				ToUserId: in.ToUserId,
				IsFriend: true,
			})
			if err != nil {
				logx.Error("FollowModel.Insert err:", err)
				return errorx.NewStatusDBErr()
			}

			//对方也修改为朋友
			follower.IsFriend = true
			err = l.svcCtx.FollowModel.TxUpdate(tx, follower)

			if err != nil {
				logx.Error("FollowModel.Update err:", err)
				return errorx.NewStatusDBErr()
			}
			return nil
		})

		if err != nil {
			return &pb.FollowTxnResp{}, errorx.NewStatusParamTxErr(err.Error())
		}

		return &pb.FollowTxnResp{}, nil
	}

	//取消关注操作
	if in.ActionType == 2 {
		//开启子事务屏障
		err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
			follow, err := l.svcCtx.FollowModel.FindFollowByUsersId(l.ctx, in.UserId, in.ToUserId)
			if err != nil && err != sqlx.ErrNotFound {
				logx.Error("FollowModel.FindFollowByUsersId err:", err)
				return errorx.NewStatusDBErr()
			}
			//未关注过
			if err != nil {
				return errorx.NewStatusParamErr(errorx.ERRFOLLOWUSER)
			}
			//删除关注
			err = l.svcCtx.FollowModel.TxDelete(tx, follow.Id)
			if err != nil {
				logx.Error("FollowModel.TxInsert err:", err)
				return errorx.NewStatusDBErr()
			}
			//查看对方是否关注自己
			follower, err := l.svcCtx.FollowModel.FindFollowByUsersId(l.ctx, in.ToUserId, in.UserId)
			if err != nil && err != sqlx.ErrNotFound {
				logx.Error("FollowModel.FindFollowByUsersId err:", err)
				return errorx.NewStatusDBErr()
			}

			//对方未关注自己，直接返回
			if err != nil {
				return nil
			}

			//对方已关注自己,取消朋友关系
			follower.IsFriend = false
			err = l.svcCtx.FollowModel.TxUpdate(tx, follower)

			if err != nil {
				logx.Error("FollowModel.Update err:", err)
				return errorx.NewStatusDBErr()
			}
			return nil
		})
		if err != nil {
			return &pb.FollowTxnResp{}, errorx.NewStatusParamTxErr(err.Error())
		}
		return &pb.FollowTxnResp{}, nil
	}

	//参数错误
	return &pb.FollowTxnResp{}, errorx.NewStatusParamTxErr(err.Error())
}
