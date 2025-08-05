package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ RoleMenusModel = (*customRoleMenusModel)(nil)

type (
	// RoleMenusModel is an interface to be customized, add more methods here,
	// and implement the added methods in customRoleMenusModel.
	RoleMenusModel interface {
		roleMenusModel
		FindMenuIdsByRoleId(ctx context.Context, roleId uint64) ([]int64, error)
	}

	customRoleMenusModel struct {
		*defaultRoleMenusModel
	}
)

// NewRoleMenusModel returns a model for the database table.
func NewRoleMenusModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) RoleMenusModel {
	return &customRoleMenusModel{
		defaultRoleMenusModel: newRoleMenusModel(conn, c, opts...),
	}
}

func (m *customRoleMenusModel) FindMenuIdsByRoleId(ctx context.Context, roleId uint64) ([]int64, error) {
	var result []int64
	query := fmt.Sprintf("select menu_id from %s where role_id = ?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &result, query, roleId)
	if err != nil {
		return nil, err
	}
	return result, nil
}
