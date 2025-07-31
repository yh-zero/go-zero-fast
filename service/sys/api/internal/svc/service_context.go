package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-fast/service/sys/api/internal/config"
	"go-zero-fast/service/sys/rpc/client/token"
	"go-zero-fast/service/sys/rpc/client/user"
)

type ServiceContext struct {
	Config   config.Config
	UserRPC  user.User
	TokenRPC token.Token
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		UserRPC:  user.NewUser(zrpc.MustNewClient(c.SysRPC)),
		TokenRPC: token.NewToken(zrpc.MustNewClient(c.SysRPC)),
	}
}
