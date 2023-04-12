package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"tik-tok-brief/common/errorx"

	"tik-tok-brief/service/user/rpc/internal/svc"
	"tik-tok-brief/service/user/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowRevertTxnLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFollowRevertTxnLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowRevertTxnLogic {
	return &FollowRevertTxnLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 关注操作回滚（事务处理）
func (l *FollowRevertTxnLogic) FollowRevertTxn(in *pb.FollowTxnReq) (*pb.FollowTxnResp, error) {
	//验证参数
	ids := []int64{in.UserId, in.ToUserId}
	users, err := l.svcCtx.UserModel.FindUserByIds(l.ctx, ids)
	if err != nil {
		logx.Error("UserModel.FindUserByIds err:", err)
		return nil, errorx.NewStatusTxErr()
	}
	if len(users) != 2 || (in.ActionType != 1 && in.ActionType != 2) {
		return nil, errorx.NewStatusParamTxErr(errorx.ERRUSERID)
	}

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

	//取消关注操作回滚（关注回来）
	if in.ActionType == 2 {
		//开启子事务屏障
		err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
			//增加用户关注数
			_, err = l.svcCtx.UserModel.TxUpdateFollowCount(tx, in.UserId, 1)
			if err != nil {
				logx.Error("UserModel.TxUpdateFollowCount err:", err)
				return errorx.NewStatusDBErr()
			}

			//增加关注用户粉丝数
			_, err = l.svcCtx.UserModel.TxUpdateFollowerCount(tx, in.ToUserId, 1)
			if err != nil {
				logx.Error("UserModel.TxUpdateFollowerCount err:", err)
				return errorx.NewStatusDBErr()
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
		return &pb.FollowTxnResp{}, nil
	}

	//关注操作回滚（取消关注）
	//开启子事务屏障
	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		//减少用户关注数
		_, err = l.svcCtx.UserModel.TxUpdateFollowCount(tx, in.UserId, -1)
		if err != nil {
			logx.Error("UserModel.TxUpdateFollowCount err:", err)
			return errorx.NewStatusDBErr()
		}

		//减少关注用户粉丝数
		_, err = l.svcCtx.UserModel.TxUpdateFollowerCount(tx, in.ToUserId, -1)
		if err != nil {
			logx.Error("UserModel.TxUpdateFollowerCount err:", err)
			return errorx.NewStatusDBErr()
		}
		return nil
	})
	if err != nil {
		return nil, errorx.NewStatusParamTxErr(err.Error())
	}

	return &pb.FollowTxnResp{}, nil
}
