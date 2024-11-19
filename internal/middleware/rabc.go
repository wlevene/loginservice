package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/casbin/casbin/v2"
)

type RABCMiddleware struct {
	rabc *casbin.Enforcer
}

func NewRABCMiddleware() *RABCMiddleware {
	rabc, _ := casbin.NewEnforcer("./assets/model.conf", "./assets/policy.csv")
	return &RABCMiddleware{
		rabc: rabc,
	}
}

func (m *RABCMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		if res, err := m.rabc.Enforce("alice", "data1", "read"); err == nil && res {
			next(w, r)
		} else {
			res, _ := json.Marshal(map[string]interface{}{
				"code": http.StatusForbidden,
				"msg":  "request has been denied",
			})
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write(res)
			return
		}

	}
}
