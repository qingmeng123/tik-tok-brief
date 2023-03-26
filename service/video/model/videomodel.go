package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ VideoModel = (*customVideoModel)(nil)

type (
	// VideoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customVideoModel.
	VideoModel interface {
		videoModel
		FindListByUserId(ctx context.Context, userId int64) ([]*Video, error)
		FindListByCTimeLimit(ctx context.Context, time time.Time, maxNum int) ([]*Video, error)
	}

	customVideoModel struct {
		*defaultVideoModel
	}
)

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
