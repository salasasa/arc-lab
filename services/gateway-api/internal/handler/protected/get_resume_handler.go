// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package protected

import (
	"net/http"

	"gateway-api/internal/logic/protected"
	"gateway-api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// 获取完整简历 (仅限登录用户)
func GetResumeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := protected.NewGetResumeLogic(r.Context(), svcCtx)
		resp, err := l.GetResume()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
