package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysRolesModel = (*customSysRolesModel)(nil)

type (
	// SysRolesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRolesModel.
	SysRolesModel interface {
		sysRolesModel
		FindCodesByIds(ctx context.Context, ids []uint64) ([]string, error)
	}

	customSysRolesModel struct {
		*defaultSysRolesModel
	}
)

// NewSysRolesModel returns a model for the database table.
func NewSysRolesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysRolesModel {
	return &customSysRolesModel{
		defaultSysRolesModel: newSysRolesModel(conn, c, opts...),
	}
}

func (m *customSysRolesModel) FindCodesByIds(ctx context.Context, ids []uint64) ([]string, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	result := make([]string, 0, len(ids))
	for _, id := range ids {
		var code string
		query := fmt.Sprintf("select code from %s where id = ?", m.table)
		err := m.QueryRowNoCacheCtx(ctx, &code, query, id)
		if err != nil {
			if err == sqlc.ErrNotFound {
				continue // 跳过不存在的记录
			}
			return nil, err
		}
		result = append(result, code)
	}

	return result, nil
}

// 带缓存
//func (m *customSysRolesModel) FindCodesByIds(ctx context.Context, ids []uint64) ([]string, error) {
//	if len(ids) == 0 {
//		return nil, nil
//	}
//
//	// 创建结果切片，初始化为空字符串
//	result := make([]string, len(ids))
//
//	// 记录需要从数据库查询的ID及其在结果切片中的位置
//	missingIds := make([]uint64, 0)
//	idToIndex := make(map[uint64]int)
//
//	// 先尝试从缓存中获取
//	for i, id := range ids {
//		cacheKey := fmt.Sprintf("%s%v", cacheGoZeroFastSysRolesIdPrefix, id)
//		var role SysRoles
//		err := m.QueryRowCtx(ctx, &role, cacheKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
//			query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysRolesRows, m.table)
//			return conn.QueryRowCtx(ctx, v, query, id)
//		})
//		if err == nil {
//			result[i] = role.Code
//		} else if err == sqlc.ErrNotFound {
//			missingIds = append(missingIds, id)
//			idToIndex[id] = i
//		} else {
//			return nil, err
//		}
//	}
//
//	// 如果所有数据都在缓存中找到，直接返回
//	if len(missingIds) == 0 {
//		return result, nil
//	}
//
//	// 只查询一次数据库获取缺失的数据
//	query := fmt.Sprintf("select id, code from %s where id in (%s)", m.table, strings.Repeat("?,", len(missingIds)-1)+"?")
//	var roles []struct {
//		Id   uint64 `db:"id"`
//		Code string `db:"code"`
//	}
//
//	err := m.QueryRowsNoCacheCtx(ctx, &roles, query, missingIds)
//	if err != nil {
//		return nil, err
//	}
//
//	// 处理查询结果
//	for _, role := range roles {
//		if index, ok := idToIndex[role.Id]; ok {
//			result[index] = role.Code
//		}
//	}
//
//	return result, nil
//}
