package model

import (
	"context"
	"database/sql"
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
		FindFollowByUsersId(ctx context.Context, userId, toUserId int64) (*Follow, error)
		FindFollowsByUserId(ctx context.Context, userId int64) ([]*Follow, error)
		FindFollowersByToUserId(ctx context.Context, userId int64) ([]*Follow, error)
		FindFriendsByUserId(ctx context.Context, userId int64) ([]*Follow, error)
		TxInsert(tx *sql.Tx, data *Follow) (sql.Result, error)
		TxUpdate(tx *sql.Tx, data *Follow) error
		TxDelete(tx *sql.Tx, id int64) error
	}

	customFollowModel struct {
		*defaultFollowModel
	}
)

func (m *defaultFollowModel) TxDelete(tx *sql.Tx, id int64) error {
	tikTokFollowFollowIdKey := fmt.Sprintf("%s%v", cacheTikTokFollowFollowIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return tx.Exec(query, id)
	}, tikTokFollowFollowIdKey)
	return err
}

func (m *defaultFollowModel) TxUpdate(tx *sql.Tx, data *Follow) error {
	followIdKey := fmt.Sprintf("%s%v", cacheTikTokFollowFollowIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (sql.Result, error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, followRowsWithPlaceHolder)
		return tx.Exec(query, data.UserId, data.ToUserId, data.IsFriend, data.Id)
	}, followIdKey)
	return err
}

func (m *defaultFollowModel) TxInsert(tx *sql.Tx, data *Follow) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values(?,?,?)", m.table, followRowsExpectAutoSet)
	ret, err := tx.Exec(query, data.UserId, data.ToUserId, data.IsFriend)
	return ret, err
}

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

func (m *defaultFollowModel) FindFollowByUsersId(ctx context.Context, userId, toUserId int64) (*Follow, error) {
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
