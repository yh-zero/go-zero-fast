package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"sync"
	"time"
)

var _ SysRolesModel = (*customSysRolesModel)(nil)

type (
	// SysRolesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRolesModel.
	SysRolesModel interface {
		sysRolesModel
		FindCodesByIds(ctx context.Context, ids []uint64) ([]string, error)
		FindRoleNamesByIds(ctx context.Context, ids []uint64) ([]string, error)
		FindPageByCursor(ctx context.Context, cursor uint64, pageSize uint64) ([]*SysRoles, uint64, uint64, error)
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

// 分页查询角色列表
/*
	cursor 游标（0表示第一页）
	pageSize 条数

*/
func (m *defaultSysRolesModel) FindPageByCursor(ctx context.Context, cursor uint64, pageSize uint64) ([]*SysRoles, uint64, uint64, error) {
	// 参数校验
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	// 并行获取数据与总数
	var list []*SysRoles
	var nextCursor, total uint64
	var listErr, totalErr error

	wg := sync.WaitGroup{}
	wg.Add(2)

	// A. 分页数据查询
	go func() {
		defer wg.Done()
		query := fmt.Sprintf("SELECT %s FROM %s WHERE id < ? ORDER BY id DESC LIMIT ?", sysRolesRows, m.table)
		if cursor == 0 {
			query = fmt.Sprintf("SELECT %s FROM %s ORDER BY id DESC LIMIT ?", sysRolesRows, m.table)
		}

		// 多查1条判断是否有下一页
		limit := pageSize + 1
		if cursor > 0 {
			listErr = m.QueryRowsNoCacheCtx(ctx, &list, query, cursor, limit)
		} else {
			listErr = m.QueryRowsNoCacheCtx(ctx, &list, query, limit)
		}

		if uint64(len(list)) > pageSize {
			list = list[:pageSize]
			nextCursor = list[len(list)-1].Id // 设置下一页游标
		}
	}()

	// B. 总数查询（带缓存）
	go func() {
		defer wg.Done()
		cacheKey := "cache:sysRoles:total"
		if err := m.GetCacheCtx(ctx, cacheKey, &total); err == nil {
			fmt.Println("------------- redis total", total)
			return
		}

		query := fmt.Sprintf("SELECT COUNT(id) FROM %s", m.table)
		totalErr = m.QueryRowNoCacheCtx(ctx, &total, query)
		fmt.Println("------------- total", total)
		// 异步更新缓存（30秒过期
		if totalErr == nil {
			go func() {
				err := m.SetCacheWithExpire(cacheKey, total, 30*time.Second)
				if err != nil {

				}
			}()
		}
	}()

	wg.Wait()

	// 错误合并返回
	if listErr != nil {
		return nil, 0, 0, listErr
	}
	if totalErr != nil {
		return nil, 0, 0, totalErr
	}
	return list, nextCursor, total, nil
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
func (m *customSysRolesModel) FindRoleNamesByIds(ctx context.Context, ids []uint64) ([]string, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	result := make([]string, 0, len(ids))
	for _, id := range ids {
		var roleName string
		query := fmt.Sprintf("select name from %s where id = ?", m.table)
		err := m.QueryRowNoCacheCtx(ctx, &roleName, query, id)
		if err != nil {
			if err == sqlc.ErrNotFound {
				continue // 跳过不存在的记录
			}
			return nil, err
		}
		result = append(result, roleName)
	}

	return result, nil
}
