package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserPositionsModel = (*customUserPositionsModel)(nil)

type (
	// UserPositionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPositionsModel.
	UserPositionsModel interface {
		userPositionsModel
	}

	customUserPositionsModel struct {
		*defaultUserPositionsModel
	}
)

// NewUserPositionsModel returns a model for the database table.
func NewUserPositionsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserPositionsModel {
	return &customUserPositionsModel{
		defaultUserPositionsModel: newUserPositionsModel(conn, c, opts...),
	}
}
