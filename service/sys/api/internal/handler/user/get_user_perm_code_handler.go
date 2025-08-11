package user

import (
	"fmt"
	"net/http"

	"go-zero-fast/common/result"
	"go-zero-fast/service/sys/api/internal/logic/user"
	"go-zero-fast/service/sys/api/internal/svc"
)

// 获取用户权限码
func GetUserPermCodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewGetUserPermCodeLogic(r.Context(), svcCtx)
		resp, err := l.GetUserPermCode()
		fmt.Println("------------ GetUserPermCodeHandler", err)
		result.HttpResult(r, w, resp, err)

	}
}
