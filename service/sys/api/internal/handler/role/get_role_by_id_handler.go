package role

import (
	"net/http"

	"go-zero-fast/common/result"
	"go-zero-fast/service/sys/api/internal/logic/role"
	"go-zero-fast/service/sys/api/internal/svc"
	"go-zero-fast/service/sys/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 通过ID获取角色
func GetRoleByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IDReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewGetRoleByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetRoleById(&req)
		result.HttpResult(r, w, resp, err)

	}
}
