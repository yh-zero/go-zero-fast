package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-zero-fast/service/sys/model"
	"go-zero-fast/service/sys/rpc/internal/config"
)

type ServiceContext struct {
	Config              config.Config
	SysUsersModel       model.SysUsersModel
	SysTokensModel      model.SysTokensModel
	UserRolesModel      model.UserRolesModel
	SysRolesModel       model.SysRolesModel
	SysDepartmentsModel model.SysDepartmentsModel
	SysPositionsModel   model.SysPositionsModel
	UserPositionsModel  model.UserPositionsModel
	SysMenusModel       model.SysMenusModel
	RoleMenusModel      model.RoleMenusModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:              c,
		SysUsersModel:       model.NewSysUsersModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
		SysTokensModel:      model.NewSysTokensModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
		UserRolesModel:      model.NewUserRolesModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
		SysRolesModel:       model.NewSysRolesModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
		SysDepartmentsModel: model.NewSysDepartmentsModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
		SysPositionsModel:   model.NewSysPositionsModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
		UserPositionsModel:  model.NewUserPositionsModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
		SysMenusModel:       model.NewSysMenusModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
		RoleMenusModel:      model.NewRoleMenusModel(sqlx.NewMysql(c.DataSource), c.CacheRedis),
	}
}
