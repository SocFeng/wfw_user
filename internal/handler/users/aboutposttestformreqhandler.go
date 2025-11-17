// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package users

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wfw_user/internal/logic/users"
	"wfw_user/internal/svc"
	"wfw_user/internal/types"
)

func AboutPostTestFormReqHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AboutTestFormReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := users.NewAboutPostTestFormReqLogic(r.Context(), svcCtx)
		resp, err := l.AboutPostTestFormReq(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
