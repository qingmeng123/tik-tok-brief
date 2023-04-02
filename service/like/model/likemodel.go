package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LikeModel = (*customLikeModel)(nil)

type (
	// LikeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLikeModel.
	LikeModel interface {
		likeModel
		FindLikeListByUserId(ctx context.Context, userId int64) ([]*Like, error)
		FindLikeByUserIdVideoId(ctx context.Context, userId, videoId int64) (*Like, error)
	}

	customLikeModel struct {
		*defaultLikeModel
	}
)

func (m *defaultLikeModel) FindLikeByUserIdVideoId(ctx context.Context, userId, videoId int64) (*Like, error) {
	like := new(Like)
	query := fmt.Sprintf("select * from %s where user_id =? and video_id=?", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &like, query, userId, videoId)
	if err != nil {
		return nil, err
	}
	return like, nil

}

func (m *defaultLikeModel) FindLikeListByUserId(ctx context.Context, userId int64) ([]*Like, error) {
	likes := make([]*Like, 0)
	query := fmt.Sprintf("select * from %s where user_id =?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &likes, query, userId)
	if err != nil {
		return nil, err
	}
	return likes, nil
}

// NewLikeModel returns a model for the database table.
func NewLikeModel(conn sqlx.SqlConn, c cache.CacheConf) LikeModel {
	return &customLikeModel{
		defaultLikeModel: newLikeModel(conn, c),
	}
}
