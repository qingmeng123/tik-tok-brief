package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CommentModel = (*customCommentModel)(nil)

type (
	// CommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCommentModel.
	CommentModel interface {
		commentModel
		FindCommentListByVideoId(ctx context.Context, videoId int64) ([]*Comment, error)
	}

	customCommentModel struct {
		*defaultCommentModel
	}
)

func (m *defaultCommentModel) FindCommentListByVideoId(ctx context.Context, videoId int64) ([]*Comment, error) {
	comments := make([]*Comment, 0)
	query := fmt.Sprintf("select * from %s where video_id =?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &comments, query, videoId)
	if err != nil {
		return nil, err
	}
	return comments, nil
}

// NewCommentModel returns a model for the database table.
func NewCommentModel(conn sqlx.SqlConn, c cache.CacheConf) CommentModel {
	return &customCommentModel{
		defaultCommentModel: newCommentModel(conn, c),
	}
}
