package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ UserRolesModel = (*customUserRolesModel)(nil)

type (
	// UserRolesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserRolesModel.
	UserRolesModel interface {
		userRolesModel
		FindRoleIdsByUserId(ctx context.Context, userId uint64) ([]uint64, error) // 新增方法
		FindUserIdsByRoleId(ctx context.Context, roleId uint64) ([]uint64, error)
		FindUserIdsByRoleIds(ctx context.Context, roleIds []uint64) ([]uint64, error)
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

// 根据用户id 查找到全部的角色id
func (m *customUserRolesModel) FindRoleIdsByUserId(ctx context.Context, userId uint64) ([]uint64, error) {
	var roleIds []uint64
	query := fmt.Sprintf("select `role_id` from %s where `user_id` = ?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &roleIds, query, userId)
	if err != nil {
		return nil, err
	}
	return roleIds, nil
}

// 根据角色id  查找到全部的 用户id
func (m *customUserRolesModel) FindUserIdsByRoleId(ctx context.Context, roleId uint64) ([]uint64, error) {
	var userIds []uint64
	query := fmt.Sprintf("select `user_id` from %s where `role_id` = ?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &userIds, query, roleId)
	if err != nil {
		return nil, err
	}
	return userIds, nil
}

// 根据角色ids  查找到全部的 用户id
func (m *customUserRolesModel) FindUserIdsByRoleIds(ctx context.Context, roleIds []uint64) ([]uint64, error) {
	var userIds []uint64
	if len(roleIds) == 0 {
		return userIds, nil
	}

	// 构建临时表查询（MySQL语法）
	query := fmt.Sprintf(`
        SELECT DISTINCT ur.user_id FROM %s ur JOIN ( SELECT ? AS role_id %s ) AS roles ON ur.role_id = roles.role_id`, m.table, strings.Repeat("UNION ALL SELECT ? ", len(roleIds)-1))

	args := make([]interface{}, len(roleIds))
	for i, id := range roleIds {
		args[i] = id
	}

	err := m.QueryRowsNoCacheCtx(ctx, &userIds, query, args...)
	if err != nil {
		return nil, err
	}
	return userIds, nil
}
