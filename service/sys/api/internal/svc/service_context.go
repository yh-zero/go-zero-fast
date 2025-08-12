package svc

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-fast/service/sys/api/internal/config"
	"go-zero-fast/service/sys/api/internal/middleware"
	"go-zero-fast/service/sys/rpc/client/menu"
	"go-zero-fast/service/sys/rpc/client/role"
	"go-zero-fast/service/sys/rpc/client/token"
	"go-zero-fast/service/sys/rpc/client/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware

	Casbin *casbin.SyncedCachedEnforcer

	UserRPC  user.User
	TokenRPC token.Token
	MenuRPC  menu.Menu
	RoleRpc  role.Role
}

func NewServiceContext(c config.Config) *ServiceContext {

	rds := redis.MustNewRedis(c.BizRedis, redis.WithPass(c.BizRedis.Pass))

	//casB := c.CasbinConf.MustNewCasbinWithOriginalRedisWatcher("mysql", c.DB.DataSource, c.BizRedis)
	casB := c.CasbinConf.MustNewCasbinWithRedisWatcher(c.DB.DataSource, c.BizRedis)

	svc := &ServiceContext{
		Config:   c,
		UserRPC:  user.NewUser(zrpc.MustNewClient(c.SysRPC)),
		TokenRPC: token.NewToken(zrpc.MustNewClient(c.SysRPC)),
		MenuRPC:  menu.NewMenu(zrpc.MustNewClient(c.SysRPC)),
		RoleRpc:  role.NewRole(zrpc.MustNewClient(c.SysRPC)),
		Casbin:   casB,
	}

	svc.Authority = middleware.NewAuthorityMiddleware(casB, rds).Handle

	return svc
}
