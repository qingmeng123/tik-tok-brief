// Code generated by goctl. DO NOT EDIT

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	followFieldNames          = builder.RawFieldNames(&Follow{})
	followRows                = strings.Join(followFieldNames, ",")
	followRowsExpectAutoSet   = strings.Join(stringx.Remove(followFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	followRowsWithPlaceHolder = strings.Join(stringx.Remove(followFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTikTokFollowFollowIdPrefix = "cache:tikTokFollow:follow:id:"
)

type (
	followModel interface {
		Insert(ctx context.Context, data *Follow) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Follow, error)
		Update(ctx context.Context, data *Follow) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFollowModel struct {
		sqlc.CachedConn
		table string
	}

	Follow struct {
		Id         int64     `db:"id"`          // id
		UserId     int64     `db:"user_id"`     // 关注用户id
		ToUserId   int64     `db:"to_user_id"`  // 被关注用户id
		IsFriend   bool      `db:"is_friend"`   // 0代表没有互相关注，1代表互相关注
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newFollowModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFollowModel {
	return &defaultFollowModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`follow`",
	}
}

func (m *defaultFollowModel) Delete(ctx context.Context, id int64) error {
	tikTokFollowFollowIdKey := fmt.Sprintf("%s%v", cacheTikTokFollowFollowIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tikTokFollowFollowIdKey)
	return err
}

func (m *defaultFollowModel) FindOne(ctx context.Context, id int64) (*Follow, error) {
	tikTokFollowFollowIdKey := fmt.Sprintf("%s%v", cacheTikTokFollowFollowIdPrefix, id)
	var resp Follow
	err := m.QueryRowCtx(ctx, &resp, tikTokFollowFollowIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", followRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFollowModel) Insert(ctx context.Context, data *Follow) (sql.Result, error) {
	tikTokFollowFollowIdKey := fmt.Sprintf("%s%v", cacheTikTokFollowFollowIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, followRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.ToUserId, data.IsFriend)
	}, tikTokFollowFollowIdKey)
	return ret, err
}

func (m *defaultFollowModel) Update(ctx context.Context, data *Follow) error {
	tikTokFollowFollowIdKey := fmt.Sprintf("%s%v", cacheTikTokFollowFollowIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, followRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.ToUserId, data.IsFriend, data.Id)
	}, tikTokFollowFollowIdKey)
	return err
}

func (m *defaultFollowModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTikTokFollowFollowIdPrefix, primary)
}

func (m *defaultFollowModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", followRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFollowModel) tableName() string {
	return m.table
}
