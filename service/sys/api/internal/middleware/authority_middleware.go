package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
)

// all  *casbin.Enforcer --> *casbin.SyncedCachedEnforce
type AuthorityMiddleware struct {
	CasB *casbin.SyncedCachedEnforcer
	Rds  *redis.Redis
}

func NewAuthorityMiddleware(CasB *casbin.SyncedCachedEnforcer, rds *redis.Redis) *AuthorityMiddleware {
	return &AuthorityMiddleware{
		CasB: CasB,
		Rds:  rds,
	}
}

func (m *AuthorityMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("=========== 中间件 ==============")

		// Passthrough to next handler if need
		next(w, r)
	}
}
