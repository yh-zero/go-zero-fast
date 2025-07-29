package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUsersModel = (*customSysUsersModel)(nil)

type (
	// SysUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUsersModel.
	SysUsersModel interface {
		sysUsersModel
	}

	customSysUsersModel struct {
		*defaultSysUsersModel
	}
)

// NewSysUsersModel returns a model for the database table.
func NewSysUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysUsersModel {
	return &customSysUsersModel{
		defaultSysUsersModel: newSysUsersModel(conn, c, opts...),
	}
}
