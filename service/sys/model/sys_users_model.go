package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"sync"
)

var _ SysUsersModel = (*customSysUsersModel)(nil)

type (
	// SysUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUsersModel.
	SysUsersModel interface {
		sysUsersModel
		CountUsersByIds(ctx context.Context, userIds []uint64) (uint64, error)
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

// 根据user_ids查询用户数量（数据库直接统计）
func (m *customSysUsersModel) CountUsersByIds(ctx context.Context, userIds []uint64) (uint64, error) {
	var count uint64
	if len(userIds) == 0 {
		return 0, nil
	}

	// 构建IN子句占位符（如"?,?,?"）
	placeholders := strings.Repeat("?,", len(userIds))
	placeholders = placeholders[:len(placeholders)-1]

	query := fmt.Sprintf("SELECT COUNT(DISTINCT `id`) FROM %s WHERE `id` IN (%s)", m.table, placeholders)

	// 转换参数类型
	args := make([]interface{}, len(userIds))
	for i, id := range userIds {
		args[i] = id
	}

	err := m.QueryRowNoCacheCtx(ctx, &count, query, args...)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 并发统计用户数量（适合超大规模数据）
func (m *customSysUsersModel) CountUsersByIdsConcurrent(ctx context.Context, userIds []uint64, batchSize int) (uint64, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var total uint64

	// 分批处理
	for i := 0; i < len(userIds); i += batchSize {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()

			end := start + batchSize
			if end > len(userIds) {
				end = len(userIds)
			}
			batch := userIds[start:end]

			// 调用基础统计方法
			count, err := m.CountUsersByIds(ctx, batch)
			if err != nil {
				return
			}

			mu.Lock()
			total += count
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	return total, nil
}
