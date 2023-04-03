package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindUserByIds(ctx context.Context, ids []int64) ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func (m *defaultUserModel) FindUserByIds(ctx context.Context, ids []int64) ([]*User, error) {
	users := make([]*User, 0)
	query := fmt.Sprintf("select * from %s where user_id in (%s)", m.table, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(ids)), ","), "[]"))
	err := m.QueryRowsNoCacheCtx(ctx, &users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}
