package main

import (
	"flag"
	"fmt"

	"go-zero-fast/service/sys/rpc/internal/config"
	menuServer "go-zero-fast/service/sys/rpc/internal/server/menu"
	roelServer "go-zero-fast/service/sys/rpc/internal/server/role"
	tokenServer "go-zero-fast/service/sys/rpc/internal/server/token"
	userServer "go-zero-fast/service/sys/rpc/internal/server/user"
	"go-zero-fast/service/sys/rpc/internal/svc"
	"go-zero-fast/service/sys/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/sys.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserServer(grpcServer, userServer.NewUserServer(ctx))
		pb.RegisterTokenServer(grpcServer, tokenServer.NewTokenServer(ctx))
		pb.RegisterMenuServer(grpcServer, menuServer.NewMenuServer(ctx))
		pb.RegisterRoleServer(grpcServer, roelServer.NewRoleServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
