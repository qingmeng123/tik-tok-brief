package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ChatModel = (*customChatModel)(nil)

type (
	// ChatModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatModel.
	ChatModel interface {
		chatModel
		FindChatList(ctx context.Context, fromUserId, toUserId int64) ([]*Chat, error)
	}

	customChatModel struct {
		*defaultChatModel
	}
)

func (m *defaultChatModel) FindChatList(ctx context.Context, fromUserId, toUserId int64) ([]*Chat, error) {
	chats := make([]*Chat, 0)
	query := fmt.Sprintf("select * from %s where from_user_id =? and to_user_id =? ", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &chats, query, fromUserId, toUserId)
	if err != nil {
		return nil, err
	}
	return chats, nil
}

// NewChatModel returns a model for the database table.
func NewChatModel(conn sqlx.SqlConn, c cache.CacheConf) ChatModel {
	return &customChatModel{
		defaultChatModel: newChatModel(conn, c),
	}
}
