// Code generated by goctl. DO NOT EDIT.

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
	likeFieldNames          = builder.RawFieldNames(&Like{})
	likeRows                = strings.Join(likeFieldNames, ",")
	likeRowsExpectAutoSet   = strings.Join(stringx.Remove(likeFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	likeRowsWithPlaceHolder = strings.Join(stringx.Remove(likeFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTikTokLikeLikeIdPrefix = "cache:tikTokLike:like:id:"
)

type (
	likeModel interface {
		Insert(ctx context.Context, data *Like) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Like, error)
		Update(ctx context.Context, data *Like) error
		Delete(ctx context.Context, id int64) error
	}

	defaultLikeModel struct {
		sqlc.CachedConn
		table string
	}

	Like struct {
		Id         int64     `db:"id"`          // id
		UserId     int64     `db:"user_id"`     // 用户id
		VideoId    int64     `db:"video_id"`    // 点赞视频id
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 更新时间
	}
)

func newLikeModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultLikeModel {
	return &defaultLikeModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`like`",
	}
}

func (m *defaultLikeModel) Delete(ctx context.Context, id int64) error {
	tikTokLikeLikeIdKey := fmt.Sprintf("%s%v", cacheTikTokLikeLikeIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tikTokLikeLikeIdKey)
	return err
}

func (m *defaultLikeModel) FindOne(ctx context.Context, id int64) (*Like, error) {
	tikTokLikeLikeIdKey := fmt.Sprintf("%s%v", cacheTikTokLikeLikeIdPrefix, id)
	var resp Like
	err := m.QueryRowCtx(ctx, &resp, tikTokLikeLikeIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", likeRows, m.table)
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

func (m *defaultLikeModel) Insert(ctx context.Context, data *Like) (sql.Result, error) {
	tikTokLikeLikeIdKey := fmt.Sprintf("%s%v", cacheTikTokLikeLikeIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, likeRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.VideoId)
	}, tikTokLikeLikeIdKey)
	return ret, err
}

func (m *defaultLikeModel) Update(ctx context.Context, data *Like) error {
	tikTokLikeLikeIdKey := fmt.Sprintf("%s%v", cacheTikTokLikeLikeIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, likeRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.VideoId, data.Id)
	}, tikTokLikeLikeIdKey)
	return err
}

func (m *defaultLikeModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTikTokLikeLikeIdPrefix, primary)
}

func (m *defaultLikeModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", likeRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultLikeModel) tableName() string {
	return m.table
}
