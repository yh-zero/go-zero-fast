package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/sync/errgroup"
	"strings"
	"time"
)

var _ SysMenusModel = (*customSysMenusModel)(nil)

var (
	cacheGoZeroFastSysMenusTotalPrefix = "cache:goZeroFast:sysMenus:total"
)

type (
	// SysMenusModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMenusModel.
	SysMenusModel interface {
		sysMenusModel
		FindMenusByIds(ctx context.Context, ids []int64) ([]*SysMenus, error)
		FindMenusList(ctx context.Context, pageNo uint64, pageSize uint64) ([]*SysMenus, uint64, error)
	}

	customSysMenusModel struct {
		*defaultSysMenusModel
	}
)

// NewSysMenusModel returns a model for the database table.
func NewSysMenusModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysMenusModel {
	return &customSysMenusModel{
		defaultSysMenusModel: newSysMenusModel(conn, c, opts...),
	}
}

// 分页查询
func (m defaultSysMenusModel) FindMenusList(ctx context.Context, pageNo uint64, pageSize uint64) ([]*SysMenus, uint64, error) {
	// 参数默认值设置
	if pageSize == 0 || pageSize > 200 {
		pageSize = 20
	}
	if pageNo == 0 {
		pageNo = 1
	}

	// 构建查询条件（这里保留了原有的条件构建方式，可根据实际业务简化）
	condition := ""
	params := []interface{}{}
	baseQuery := fmt.Sprintf("SELECT %s FROM %s%s", sysMenusRows, m.table, condition)
	countQuery := fmt.Sprintf("SELECT COUNT(id) FROM %s%s", m.table, condition)

	var (
		list  []*SysMenus
		total uint64
	)

	// 使用 errgroup 进行并发查询
	g, ctx := errgroup.WithContext(ctx)

	// 并发获取分页数据
	g.Go(func() error {
		offset := (pageNo - 1) * pageSize
		query := baseQuery + " ORDER BY sort ASC LIMIT ?, ?"
		queryParams := append(params, offset, pageSize)
		return m.QueryRowsNoCacheCtx(ctx, &list, query, queryParams...)
	})

	// 并发获取总数（带缓存）
	g.Go(func() error {
		cacheKey := fmt.Sprintf("%s:all", cacheGoZeroFastSysMenusTotalPrefix)

		// 尝试从缓存获取
		if err := m.GetCacheCtx(ctx, cacheKey, &total); err == nil {
			return nil
		}

		// 缓存未命中，查询数据库
		if err := m.QueryRowNoCacheCtx(ctx, &total, countQuery, params...); err != nil {
			return err
		}

		// 异步设置缓存（防雪崩）
		if total > 0 {
			go func() {
				_ = m.SetCacheWithExpire(cacheKey, total, time.Minute)
			}()
		}
		return nil
	})

	// 等待所有并发操作完成
	if err := g.Wait(); err != nil {
		return nil, 0, fmt.Errorf("query failed: %w", err)
	}

	return list, total, nil
}

//func (m *customSysMenusModel) FindByIds(ctx context.Context, ids []int64) ([]*SysMenus, error) {
//	if len(ids) == 0 {
//		return nil, fmt.Errorf("empty ids list")
//	}
//	query := fmt.Sprintf("SELECT %s FROM %s WHERE `id` IN (?) AND `disabled` != 0 ORDER BY `sort` ASC", sysMenusRows, m.table)
//
//	var menus []*SysMenus
//	err := m.QueryRowsNoCacheCtx(ctx, &menus, query, ids)
//
//	// 检查结果完整性
//	if err == nil {
//		foundIDs := make(map[int64]bool, len(menus))
//		for _, menu := range menus {
//			foundIDs[menu.Id] = true
//		}
//
//		var missingIDs []int64
//		for _, id := range ids {
//			if !foundIDs[id] {
//				missingIDs = append(missingIDs, id)
//			}
//		}
//
//		if len(missingIDs) > 0 {
//			return menus, fmt.Errorf("some records not found or disabled, missing IDs: %v", missingIDs)
//		}
//	}
//	return menus, err
//}

//func (m *customSysMenusModel) FindMenusByIds(ctx context.Context, ids []int64) ([]*SysMenus, error) {
//	if len(ids) == 0 {
//		return nil, fmt.Errorf("empty ids list")
//	}
//
//	// 构建占位符字符串 (?, ?, ?...)
//	placeholders := make([]string, len(ids))
//	args := make([]interface{}, len(ids))
//	for i, id := range ids {
//		placeholders[i] = "?"
//		args[i] = id
//	}
//
//	query := fmt.Sprintf("SELECT %s FROM %s WHERE `id` IN (%s) AND `disabled` = 0 ORDER BY `sort` ASC;",
//		sysMenusRows, m.table, strings.Join(placeholders, ","))
//
//	var menus []*SysMenus
//	err := m.QueryRowsNoCacheCtx(ctx, &menus, query, args...)
//
//	// 检查结果完整性
//	if err == nil {
//		foundIDs := make(map[int64]bool, len(menus))
//		for _, menu := range menus {
//			foundIDs[menu.Id] = true
//		}
//
//		var missingIDs []int64
//		for _, id := range ids {
//			if !foundIDs[id] {
//				missingIDs = append(missingIDs, id)
//			}
//		}
//
//		if len(missingIDs) > 0 {
//			return menus, fmt.Errorf("some records not found or disabled, missing IDs: %v", missingIDs)
//		}
//	}
//	return menus, err
//}

func (m *customSysMenusModel) FindMenusByIds(ctx context.Context, ids []int64) ([]*SysMenus, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("empty ids list")
	}

	// 构建占位符和参数
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	// 查询有效菜单（忽略不存在或禁用的记录）
	query := fmt.Sprintf("SELECT %s FROM %s WHERE `id` IN (%s) AND `disabled` = 0 ORDER BY `sort` ASC;",
		sysMenusRows, m.table, strings.Join(placeholders, ","))

	var menus []*SysMenus
	if err := m.QueryRowsNoCacheCtx(ctx, &menus, query, args...); err != nil {
		return nil, err // 数据库错误仍需返回
	}
	return menus, nil // 仅返回有效数据，跳过缺失项
}
