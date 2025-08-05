package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysPositionsModel = (*customSysPositionsModel)(nil)

type (
	// SysPositionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysPositionsModel.
	SysPositionsModel interface {
		sysPositionsModel
	}

	customSysPositionsModel struct {
		*defaultSysPositionsModel
	}
)

// NewSysPositionsModel returns a model for the database table.
func NewSysPositionsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysPositionsModel {
	return &customSysPositionsModel{
		defaultSysPositionsModel: newSysPositionsModel(conn, c, opts...),
	}
}
