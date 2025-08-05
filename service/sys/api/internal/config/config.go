package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-fast/common/middlecasbin"
)

type Config struct {
	rest.RestConf
	SysRPC zrpc.RpcClientConf

	CasbinConf middlecasbin.CasbinConf
	BizRedis   redis.RedisConf
	DB         struct {
		DataSource   string
		MaxOpenConns int `json:",default=10"`
		MaxIdleConns int `json:",default=100"`
		MaxLifetime  int `json:",default=3600"`
	}

	JwtAuth struct {
		AccessSecret string
		AccessExpire int64
	}
}
