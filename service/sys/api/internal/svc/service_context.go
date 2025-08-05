package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"go-zero-fast/service/sys/api/internal/config"
	"go-zero-fast/service/sys/api/internal/middleware"
	"go-zero-fast/service/sys/rpc/client/menu"
	"go-zero-fast/service/sys/rpc/client/token"
	"go-zero-fast/service/sys/rpc/client/user"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Authority rest.Middleware

	UserRPC  user.User
	TokenRPC token.Token
	MenuRPC  menu.Menu
}

func NewServiceContext(c config.Config) *ServiceContext {

	casB := c.CasbinConf.MustNewCasbinWithRedisWatcher(c.DB.DataSource, c.BizRedis)
	rds := redis.MustNewRedis(c.BizRedis, redis.WithPass(c.BizRedis.Pass))

	svc := &ServiceContext{
		Config:   c,
		UserRPC:  user.NewUser(zrpc.MustNewClient(c.SysRPC)),
		TokenRPC: token.NewToken(zrpc.MustNewClient(c.SysRPC)),
		MenuRPC:  menu.NewMenu(zrpc.MustNewClient(c.SysRPC)),
	}

	svc.Authority = middleware.NewAuthorityMiddleware(casB, rds).Handle

	return svc
}
