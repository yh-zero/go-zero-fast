package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-fast/common/ctxJwt"
	"net/http"
	"strconv"
)

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
		logx.Errorf("===========  AuthorityMiddleware:Handle | 中间件 ==============")
		path := r.URL.Path
		method := r.Method
		authorityId := ctxJwt.GetJwtDataRoleId(r.Context())
		fmt.Println("==== path", path)
		fmt.Println("==== method", method)
		fmt.Println("==== authorityId", strconv.FormatUint(authorityId, 10))
		result := batchCheck(m.CasB, strconv.FormatUint(authorityId, 10), path, method)
		if !result {
			logx.Errorf("---------- batchCheck: 不通过 ------------")
			errorResp := struct {
				Code int    `json:"code"`
				Msg  string `json:"msg"`
			}{
				Code: 401,
				Msg:  "无访问权限 - 联系管理员",
			}

			// 返回200状态码但包含错误信息
			httpx.WriteJson(w, http.StatusOK, errorResp)
			return
		}
		next(w, r)

	}
}

func batchCheck(CasB *casbin.SyncedCachedEnforcer, authorityId, path, method string) bool {
	// 打印传入的值
	logx.Infof("Checking permission - RoleID: %s, Path: %s, Method: %s", authorityId, path, method)

	// 获取当前所有的策略规则并打印
	allPolicies, _ := CasB.GetPolicy()
	logx.Infof("Current policies in Casbin:")
	for _, policy := range allPolicies {
		logx.Infof("Policy: %v", policy)
	}

	// 执行权限检查
	ok, err := CasB.Enforce(authorityId, path, method)
	if err != nil {
		logx.Errorf("Casbin enforce error: %v", err)
		return false
	}

	if !ok {
		// 打印更详细的调试信息
		logx.Errorf("Permission denied for RoleID: %s, Path: %s, Method: %s", authorityId, path, method)
		logx.Errorf("Available policies for this role:")
		rolePolicies, _ := CasB.GetFilteredPolicy(0, authorityId)
		for _, policy := range rolePolicies {
			logx.Errorf("Role policy: %v", policy)
		}

		return false
	}
	return true
}

//func batchCheck(CasB *casbin.SyncedCachedEnforcer, authorityId, path, method string) bool {
//	ok, _ := CasB.Enforce(authorityId, path, method)
//
//	if !ok {
//		//_, _ = CasB.AddPolicy(authorityId, path, method) // 如果权限数据不小心清了 把这个开启  然后api请求两次就会有权限  最后重新设置权限即可
//		logx.Errorf("---------- 权限不足 ------------")
//		return false
//	}
//	return true
//}
