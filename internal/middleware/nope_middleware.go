package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type NopeMiddleware struct {
}

func NewNopeMiddleware() *NopeMiddleware {
	return &NopeMiddleware{}
}

func (m *NopeMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		logx.Info("NopeMiddleware")
		// Passthrough to next handler if need
		next(w, r)
	}
}
