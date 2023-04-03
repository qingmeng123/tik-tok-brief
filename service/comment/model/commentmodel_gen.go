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
	commentFieldNames          = builder.RawFieldNames(&Comment{})
	commentRows                = strings.Join(commentFieldNames, ",")
	commentRowsExpectAutoSet   = strings.Join(stringx.Remove(commentFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	commentRowsWithPlaceHolder = strings.Join(stringx.Remove(commentFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTikTokCommentCommentIdPrefix = "cache:tikTokComment:comment:id:"
)

type (
	commentModel interface {
		Insert(ctx context.Context, data *Comment) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Comment, error)
		Update(ctx context.Context, data *Comment) error
		Delete(ctx context.Context, id int64) error
	}

	defaultCommentModel struct {
		sqlc.CachedConn
		table string
	}

	Comment struct {
		Id         int64     `db:"id"`       // id
		UserId     int64     `db:"user_id"`  // 发送用户id
		VideoId    int64     `db:"video_id"` // 接收消息用户id
		Content    string    `db:"content"`
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newCommentModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultCommentModel {
	return &defaultCommentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`comment`",
	}
}

func (m *defaultCommentModel) Delete(ctx context.Context, id int64) error {
	tikTokCommentCommentIdKey := fmt.Sprintf("%s%v", cacheTikTokCommentCommentIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tikTokCommentCommentIdKey)
	return err
}

func (m *defaultCommentModel) FindOne(ctx context.Context, id int64) (*Comment, error) {
	tikTokCommentCommentIdKey := fmt.Sprintf("%s%v", cacheTikTokCommentCommentIdPrefix, id)
	var resp Comment
	err := m.QueryRowCtx(ctx, &resp, tikTokCommentCommentIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentRows, m.table)
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

func (m *defaultCommentModel) Insert(ctx context.Context, data *Comment) (sql.Result, error) {
	tikTokCommentCommentIdKey := fmt.Sprintf("%s%v", cacheTikTokCommentCommentIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, commentRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.UserId, data.VideoId, data.Content)
	}, tikTokCommentCommentIdKey)
	return ret, err
}

func (m *defaultCommentModel) Update(ctx context.Context, data *Comment) error {
	tikTokCommentCommentIdKey := fmt.Sprintf("%s%v", cacheTikTokCommentCommentIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, commentRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.UserId, data.VideoId, data.Content, data.Id)
	}, tikTokCommentCommentIdKey)
	return err
}

func (m *defaultCommentModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTikTokCommentCommentIdPrefix, primary)
}

func (m *defaultCommentModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", commentRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultCommentModel) tableName() string {
	return m.table
}