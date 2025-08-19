package model

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
	"sync"
	"time"
)

var _ SysRolesModel = (*customSysRolesModel)(nil)

var (
	cacheGoZeroFastSysRolesTotalPrefix = "cache:goZeroFast:sysRoles:total"
)

type (
	// SysRolesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRolesModel.
	SysRolesModel interface {
		sysRolesModel
		FindCodesByIds(ctx context.Context, ids []uint64) ([]string, error)
		FindRoleNamesByIds(ctx context.Context, ids []uint64) ([]string, error)
		FindPageByCursor(ctx context.Context, cursor uint64, pageSize uint64) ([]*SysRoles, uint64, uint64, error)
		DeleteByIds(ctx context.Context, ids []uint64) error
		FindPageByName(ctx context.Context, name string, pageNo uint64, pageSize uint64) ([]*SysRoles, uint64, error)
		UpdateWithMap(ctx context.Context, id uint64, data map[string]interface{}) error
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

// 支持条件的分页查询方法
func (m *defaultSysRolesModel) FindPageByName(ctx context.Context, name string, pageNo uint64, pageSize uint64) ([]*SysRoles, uint64, error) {
	// 参数校验与默认值处理
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}
	if pageNo <= 0 {
		pageNo = 1
	}

	// 并发执行数据查询和总数统计
	var (
		list     []*SysRoles
		total    uint64
		listErr  error
		totalErr error
	)

	wg := sync.WaitGroup{}
	wg.Add(2)

	// 1. 分页数据查询
	go func() {
		defer wg.Done()

		baseQuery := fmt.Sprintf("SELECT %s FROM %s", sysRolesRows, m.table)
		condition := ""
		params := []interface{}{}

		// 添加名称过滤条件
		if name != "" {
			condition = " WHERE name LIKE ?"
			params = append(params, "%"+name+"%")
		}

		// 计算偏移量
		offset := (pageNo - 1) * pageSize

		// 执行查询
		query := baseQuery + condition + " ORDER BY id DESC LIMIT ?, ?"
		params = append(params, offset, pageSize)

		listErr = m.QueryRowsNoCacheCtx(ctx, &list, query, params...)
	}()

	// 2. 总数查询
	go func() {
		defer wg.Done()

		// 根据查询条件生成唯一的缓存键
		cacheKey := cacheGoZeroFastSysRolesTotalPrefix
		if name != "" {
			cacheKey += ":name:" + name
		} else {
			cacheKey += ":all"
		}

		// 先尝试从缓存获取
		if err := m.GetCacheCtx(ctx, cacheKey, &total); err == nil {
			return
		}

		// 构建总数查询
		baseQuery := fmt.Sprintf("SELECT COUNT(id) FROM %s", m.table)
		condition := ""
		params := []interface{}{}

		if name != "" {
			condition = " WHERE name LIKE ?"
			params = append(params, "%"+name+"%")
		}

		query := baseQuery + condition

		// 执行查询
		totalErr = m.QueryRowNoCacheCtx(ctx, &total, query, params...)

		// 异步缓存结果
		if totalErr == nil {
			exp := 60 * time.Second
			if name == "" {
				exp = 5 * time.Minute
			}

			go func() {
				_ = m.SetCacheWithExpire(cacheKey, total, exp)
			}()
		}
	}()

	wg.Wait()

	// 错误处理
	if listErr != nil {
		return nil, 0, listErr
	}
	if totalErr != nil {
		return nil, 0, totalErr
	}

	return list, total, nil
}

// 分页查询角色列表
/*
	cursor 游标（0表示第一页）  --
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
		cacheKey := cacheGoZeroFastSysRolesTotalPrefix
		if err := m.GetCacheCtx(ctx, cacheKey, &total); err == nil {
			fmt.Println("------------- redis total", total)
			return
		}

		query := fmt.Sprintf("SELECT COUNT(id) FROM %s", m.table)
		totalErr = m.QueryRowNoCacheCtx(ctx, &total, query)
		fmt.Println("------------- total", total)
		// 异步更新缓存（60秒过期
		if totalErr == nil {
			go func() {
				err := m.SetCacheWithExpire(cacheKey, total, 60*time.Second)
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

// DeleteByIds 根据角色ID批量删除记录
func (m *customSysRolesModel) DeleteByIds(ctx context.Context, ids []uint64) error {
	if len(ids) == 0 {
		return nil
	}

	// 构建IN子句占位符（如"?,?,?"）
	placeholders := strings.Repeat("?,", len(ids))
	placeholders = placeholders[:len(placeholders)-1]

	// 执行批量删除
	query := fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", m.table, placeholders)

	// 转换参数类型
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	_, err := m.ExecNoCacheCtx(ctx, query, args...)
	if err != nil {
		return err
	}

	// 异步清理相关缓存
	go m.clearRoleCache(ids)
	return nil
}

func (m *customSysRolesModel) clearRoleCache(ids []uint64) {
	// 1. 清理总条数缓存
	cacheKey := cacheGoZeroFastSysRolesTotalPrefix
	m.DelCache(cacheKey)

	// 2. 清理单条记录缓存（如果有）
	for _, id := range ids {
		m.DelCache(fmt.Sprintf("%s%v", m.table, id))
	}
}

func (m *customSysRolesModel) UpdateWithMap(ctx context.Context, id uint64, data map[string]interface{}) error {
	// rpc里面做判断
	//if len(data) == 0 {
	//	return errors.New("无有效更新字段")
	//}

	// 添加更新时间戳
	//data["updated_at"] = time.Now() // 数据库同步更新了

	query := fmt.Sprintf("UPDATE %s SET ", m.table)
	var sets []string
	var args []interface{}

	allowedFields := map[string]bool{
		"status":         true,
		"name":           true,
		"code":           false,
		"default_router": true,
		"remark":         true,
		"sort":           true,
	}

	for k, v := range data {
		if !allowedFields[k] { // 不在白名单则跳过
			continue
		}
		sets = append(sets, fmt.Sprintf("%s = ?", k))
		args = append(args, v)
	}

	query += strings.Join(sets, ", ")
	query += " WHERE id = ?"
	args = append(args, id)

	_, err := m.ExecNoCacheCtx(ctx, query, args...)
	return err
}
