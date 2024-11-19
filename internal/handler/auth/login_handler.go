package auth

import (
	"github.com/wlevene/loginservice/internal/logic/auth"
	"github.com/wlevene/loginservice/internal/svc"
	"github.com/wlevene/loginservice/internal/types"
	"net/http"

	"github.com/wlevene/loginservice/internal/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := auth.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)

		response.Response(w, resp, err)
	}
}
