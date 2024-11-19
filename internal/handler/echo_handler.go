package handler

import (
	"net/http"

	"github.com/wlevene/loginservice/internal/logic"
	"github.com/wlevene/loginservice/internal/svc"
	"github.com/wlevene/loginservice/internal/types"

	"github.com/wlevene/loginservice/internal/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func EchoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EchoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewEchoLogic(r.Context(), svcCtx)
		resp, err := l.Echo(&req)

		response.Response(w, resp, err)
	}
}
