package response

import (
	"net/http"

	"github.com/wlevene/loginservice/internal/errorx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {

	var body Body

	if err != nil {

		switch e := err.(type) {

		case *errorx.CodeError:
			body.Code = e.Code

		default:
			body.Code = -1

		}

		body.Msg = err.Error()

	} else {

		body.Msg = "OK"

		if resp != nil {
			body.Data = resp
		}
	}

	httpx.OkJson(w, body)
}
