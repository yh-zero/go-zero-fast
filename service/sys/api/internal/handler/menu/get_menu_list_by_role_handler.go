package menu

import (
	"net/http"

	"go-zero-fast/common/result"
	"go-zero-fast/service/sys/api/internal/logic/menu"
	"go-zero-fast/service/sys/api/internal/svc"
)

// 获取菜单列表 -- 对应的用户角色的权限
func GetMenuListByRoleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewGetMenuListByRoleLogic(r.Context(), svcCtx)
		resp, err := l.GetMenuListByRole()
		result.HttpResult(r, w, resp, err)

	}
}
