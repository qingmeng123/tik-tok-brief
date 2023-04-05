package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"time"
)

var _ VideoModel = (*customVideoModel)(nil)

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
		FindVideosByIds(ctx context.Context, ids []int64) ([]*Video, error)
		FindListByUserId(ctx context.Context, userId int64) ([]*Video, error)
		UpdateCommentCountByVideoId(ctx context.Context, videoId, number int64) error
		UpdateFavoriteCountByVideoId(ctx context.Context, videoId, number int64) error
		FindListByCTimeLimit(ctx context.Context, time time.Time, maxNum int) ([]*Video, error)
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

func (m *defaultVideoModel) FindVideosByIds(ctx context.Context, ids []int64) ([]*Video, error) {
	videos := make([]*Video, 0)
	query := fmt.Sprintf("select * from %s where video_id in (%s)", m.table, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]"))
	err := m.QueryRowsNoCacheCtx(ctx, &videos, query)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

func (m *defaultVideoModel) UpdateCommentCountByVideoId(ctx context.Context, videoId, number int64) error {
	query := fmt.Sprintf("update %s set comment_count=comment_count+? where video_id=?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, number, videoId)
	if err != nil {
		return err
	}

	return nil
}

func (m *defaultVideoModel) UpdateFavoriteCountByVideoId(ctx context.Context, videoId, number int64) error {
	query := fmt.Sprintf("update %s set favorite_count=favorite_count+? where video_id=?", m.table)
	_, err := m.ExecNoCacheCtx(ctx, query, number, videoId)
	if err != nil {
		return err
	}

	return nil
}

func (m *defaultVideoModel) FindListByCTimeLimit(ctx context.Context, time time.Time, maxNum int) ([]*Video, error) {
	videos := make([]*Video, 0)
	query := fmt.Sprintf("select * from %s where create_time <? order by create_time desc limit ?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &videos, query, time.Format("2006-01-02 15:04:05"), maxNum)
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (m *defaultVideoModel) FindListByUserId(ctx context.Context, userId int64) ([]*Video, error) {
	videos := make([]*Video, 0)
	query := fmt.Sprintf("select * from %s where `user_id`=? order by create_time desc", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &videos, query, userId)
	if err != nil {
		return nil, err
	}
	return videos, nil
}

// NewVideoModel returns a model for the database table.
func NewVideoModel(conn sqlx.SqlConn, c cache.CacheConf) VideoModel {
	return &customVideoModel{
		defaultVideoModel: newVideoModel(conn, c),
	}
}
