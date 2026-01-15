// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package protected

import (
	"net/http"

	"gateway-api/internal/logic/protected"
	"gateway-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 验证 Token 是否有效 (前端保活)
func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := protected.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
