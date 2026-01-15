// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package public

import (
	"net/http"

	"gateway-api/internal/logic/public"
	"gateway-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 简单的自我介绍 (公开)
func GetIntroHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := public.NewGetIntroLogic(r.Context(), svcCtx)
		resp, err := l.GetIntro()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
