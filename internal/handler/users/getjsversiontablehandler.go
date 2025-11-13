// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package users

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"wfw_user/internal/logic/users"
	"wfw_user/internal/svc"
)

func GetJsVersionTableHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := users.NewGetJsVersionTableLogic(r.Context(), svcCtx)
		resp, err := l.GetJsVersionTable()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
