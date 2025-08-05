package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysDepartmentsModel = (*customSysDepartmentsModel)(nil)

type (
	// SysDepartmentsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysDepartmentsModel.
	SysDepartmentsModel interface {
		sysDepartmentsModel
	}

	customSysDepartmentsModel struct {
		*defaultSysDepartmentsModel
	}
)

// NewSysDepartmentsModel returns a model for the database table.
func NewSysDepartmentsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysDepartmentsModel {
	return &customSysDepartmentsModel{
		defaultSysDepartmentsModel: newSysDepartmentsModel(conn, c, opts...),
	}
}
