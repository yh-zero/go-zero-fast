package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ SysMenusModel = (*customSysMenusModel)(nil)

type (
	// SysMenusModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMenusModel.
	SysMenusModel interface {
		sysMenusModel
		FindMenusByIds(ctx context.Context, ids []int64) ([]*SysMenus, error)
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

func (m *customSysMenusModel) FindMenusByIds(ctx context.Context, ids []int64) ([]*SysMenus, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("empty ids list")
	}

	// 构建占位符字符串 (?, ?, ?...)
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := fmt.Sprintf("SELECT %s FROM %s WHERE `id` IN (%s) AND `disabled` = 0 ORDER BY `sort` ASC;",
		sysMenusRows, m.table, strings.Join(placeholders, ","))

	var menus []*SysMenus
	err := m.QueryRowsNoCacheCtx(ctx, &menus, query, args...)

	// 检查结果完整性
	if err == nil {
		foundIDs := make(map[int64]bool, len(menus))
		for _, menu := range menus {
			foundIDs[menu.Id] = true
		}

		var missingIDs []int64
		for _, id := range ids {
			if !foundIDs[id] {
				missingIDs = append(missingIDs, id)
			}
		}

		if len(missingIDs) > 0 {
			return menus, fmt.Errorf("some records not found or disabled, missing IDs: %v", missingIDs)
		}
	}
	return menus, err
}
