package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysTokensModel = (*customSysTokensModel)(nil)

type (
	// SysTokensModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysTokensModel.
	SysTokensModel interface {
		sysTokensModel
	}

	customSysTokensModel struct {
		*defaultSysTokensModel
	}
)

// NewSysTokensModel returns a model for the database table.
func NewSysTokensModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysTokensModel {
	return &customSysTokensModel{
		defaultSysTokensModel: newSysTokensModel(conn, c, opts...),
	}
}
