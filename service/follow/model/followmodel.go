package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FollowModel = (*customFollowModel)(nil)

type (
	// FollowModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFollowModel.
	FollowModel interface {
		followModel
		FindIsFriendByUsersId(ctx context.Context, userId, toUserId int64) (*Follow, error)
		FindFollowsByUserId(ctx context.Context, userId int64) ([]*Follow, error)
		FindFollowersByToUserId(ctx context.Context, userId int64) ([]*Follow, error)
		FindFriendsByUserId(ctx context.Context, userId int64) ([]*Follow, error)
	}

	customFollowModel struct {
		*defaultFollowModel
	}
)

func (m *defaultFollowModel) FindFollowsByUserId(ctx context.Context, userId int64) ([]*Follow, error) {
	follows := make([]*Follow, 0)
	query := fmt.Sprintf("select * from %s where user_id =?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &follows, query, userId)
	if err != nil {
		return nil, err
	}
	return follows, nil
}

func (m *defaultFollowModel) FindFollowersByToUserId(ctx context.Context, toUserId int64) ([]*Follow, error) {
	follows := make([]*Follow, 0)
	query := fmt.Sprintf("select * from %s where to_user_id =?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &follows, query, toUserId)
	if err != nil {
		return nil, err
	}
	return follows, nil
}

func (m *defaultFollowModel) FindFriendsByUserId(ctx context.Context, userId int64) ([]*Follow, error) {
	follows := make([]*Follow, 0)
	query := fmt.Sprintf("select * from %s where user_id =? and is_friend= ?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &follows, query, userId, true)
	if err != nil {
		return nil, err
	}
	return follows, nil
}

func (m *defaultFollowModel) FindIsFriendByUsersId(ctx context.Context, userId, toUserId int64) (*Follow, error) {
	query := fmt.Sprintf("select * from %s where user_id=? and to_user_id=? limit 1", m.table)
	follow := new(Follow)
	err := m.QueryRowNoCacheCtx(ctx, follow, query, userId, toUserId)
	if err != nil {
		return nil, err
	}
	return follow, nil
}

// NewFollowModel returns a model for the database table.
func NewFollowModel(conn sqlx.SqlConn, c cache.CacheConf) FollowModel {
	return &customFollowModel{
		defaultFollowModel: newFollowModel(conn, c),
	}
}
