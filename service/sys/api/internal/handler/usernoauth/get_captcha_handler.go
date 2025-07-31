package usernoauth

import (
	"net/http"

	"go-zero-fast/common/result"
	"go-zero-fast/service/sys/api/internal/logic/usernoauth"
	"go-zero-fast/service/sys/api/internal/svc"
)

// 获取验证码
func GetCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := usernoauth.NewGetCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.GetCaptcha()
		result.HttpResult(r, w, resp, err)

	}
}
