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
	videoFieldNames          = builder.RawFieldNames(&Video{})
	videoRows                = strings.Join(videoFieldNames, ",")
	videoRowsExpectAutoSet   = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	videoRowsWithPlaceHolder = strings.Join(stringx.Remove(videoFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheTikTokVideoVideoIdPrefix      = "cache:tikTokVideo:video:id:"
	cacheTikTokVideoVideoUserIdPrefix  = "cache:tikTokVideo:video:userId:"
	cacheTikTokVideoVideoVideoIdPrefix = "cache:tikTokVideo:video:videoId:"
)

type (
	videoModel interface {
		Insert(ctx context.Context, data *Video) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Video, error)
		FindOneByUserId(ctx context.Context, userId int64) (*Video, error)
		FindOneByVideoId(ctx context.Context, videoId int64) (*Video, error)
		Update(ctx context.Context, data *Video) error
		Delete(ctx context.Context, id int64) error
	}

	defaultVideoModel struct {
		sqlc.CachedConn
		table string
	}

	Video struct {
		Id            int64     `db:"id"`
		VideoId       int64     `db:"video_id"`       // 视频id
		UserId        int64     `db:"user_id"`        // 发布作者id
		Title         string    `db:"title"`          // 视频标题
		PlayUrl       string    `db:"play_url"`       // 视频播放地址
		CoverUrl      string    `db:"cover_url"`      // 封面地址
		FavoriteCount int64     `db:"favorite_count"` // 点赞数量
		CommentCount  int64     `db:"comment_count"`  // 评论数量
		CreateTime    time.Time `db:"create_time"`    // 创建时间
		UpdateTime    time.Time `db:"update_time"`    // 更新时间
	}
)

func newVideoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultVideoModel {
	return &defaultVideoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`video`",
	}
}

func (m *defaultVideoModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	tikTokVideoVideoIdKey := fmt.Sprintf("%s%v", cacheTikTokVideoVideoIdPrefix, id)
	tikTokVideoVideoUserIdKey := fmt.Sprintf("%s%v", cacheTikTokVideoVideoUserIdPrefix, data.UserId)
	tikTokVideoVideoVideoIdKey := fmt.Sprintf("%s%v", cacheTikTokVideoVideoVideoIdPrefix, data.VideoId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, tikTokVideoVideoIdKey, tikTokVideoVideoUserIdKey, tikTokVideoVideoVideoIdKey)
	return err
}

func (m *defaultVideoModel) FindOne(ctx context.Context, id int64) (*Video, error) {
	tikTokVideoVideoIdKey := fmt.Sprintf("%s%v", cacheTikTokVideoVideoIdPrefix, id)
	var resp Video
	err := m.QueryRowCtx(ctx, &resp, tikTokVideoVideoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
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

func (m *defaultVideoModel) FindOneByUserId(ctx context.Context, userId int64) (*Video, error) {
	tikTokVideoVideoUserIdKey := fmt.Sprintf("%s%v", cacheTikTokVideoVideoUserIdPrefix, userId)
	var resp Video
	err := m.QueryRowIndexCtx(ctx, &resp, tikTokVideoVideoUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", videoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoModel) FindOneByVideoId(ctx context.Context, videoId int64) (*Video, error) {
	tikTokVideoVideoVideoIdKey := fmt.Sprintf("%s%v", cacheTikTokVideoVideoVideoIdPrefix, videoId)
	var resp Video
	err := m.QueryRowIndexCtx(ctx, &resp, tikTokVideoVideoVideoIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `video_id` = ? limit 1", videoRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, videoId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultVideoModel) Insert(ctx context.Context, data *Video) (sql.Result, error) {
	tikTokVideoVideoIdKey := fmt.Sprintf("%s%v", cacheTikTokVideoVideoIdPrefix, data.Id)
	tikTokVideoVideoUserIdKey := fmt.Sprintf("%s%v", cacheTikTokVideoVideoUserIdPrefix, data.UserId)
	tikTokVideoVideoVideoIdKey := fmt.Sprintf("%s%v", cacheTikTokVideoVideoVideoIdPrefix, data.VideoId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, videoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.VideoId, data.UserId, data.Title, data.PlayUrl, data.CoverUrl, data.FavoriteCount, data.CommentCount)
	}, tikTokVideoVideoIdKey, tikTokVideoVideoUserIdKey, tikTokVideoVideoVideoIdKey)
	return ret, err
}

func (m *defaultVideoModel) Update(ctx context.Context, newData *Video) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	tikTokVideoVideoIdKey := fmt.Sprintf("%s%v", cacheTikTokVideoVideoIdPrefix, data.Id)
	tikTokVideoVideoUserIdKey := fmt.Sprintf("%s%v", cacheTikTokVideoVideoUserIdPrefix, data.UserId)
	tikTokVideoVideoVideoIdKey := fmt.Sprintf("%s%v", cacheTikTokVideoVideoVideoIdPrefix, data.VideoId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, videoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.VideoId, newData.UserId, newData.Title, newData.PlayUrl, newData.CoverUrl, newData.FavoriteCount, newData.CommentCount, newData.Id)
	}, tikTokVideoVideoIdKey, tikTokVideoVideoUserIdKey, tikTokVideoVideoVideoIdKey)
	return err
}

func (m *defaultVideoModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheTikTokVideoVideoIdPrefix, primary)
}

func (m *defaultVideoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", videoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultVideoModel) tableName() string {
	return m.table
}
