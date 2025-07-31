package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserRolesModel = (*customUserRolesModel)(nil)

type (
	// UserRolesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRolesModel.
	UserRolesModel interface {
		userRolesModel
		FindRoleIdsByUserId(ctx context.Context, userId int64) ([]uint64, error) // 新增方法
	}

	customUserRolesModel struct {
		*defaultUserRolesModel
	}
)

// NewUserRolesModel returns a model for the database table.
func NewUserRolesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserRolesModel {
	return &customUserRolesModel{
		defaultUserRolesModel: newUserRolesModel(conn, c, opts...),
	}
}

// 在 defaultUserRolesModel 结构体中实现新方法
func (m *customUserRolesModel) FindRoleIdsByUserId(ctx context.Context, userId int64) ([]uint64, error) {
	var roleIds []uint64
	query := fmt.Sprintf("select `role_id` from %s where `user_id` = ?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &roleIds, query, userId)
	if err != nil {
		return nil, err
	}
	return roleIds, nil
}
