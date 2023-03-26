package middleware

import (
	"net/http"
	"tik-tok-brief/common/middleware"
)

type JWTAuthMiddleware struct {
	accessSecret string
}

func NewJWTAuthMiddleware(accessSecret string) *JWTAuthMiddleware {
	return &JWTAuthMiddleware{accessSecret}
}

func (m *JWTAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	jwtAuth := middleware.NewJwtAuth(m.accessSecret)
	return jwtAuth.Handle(next)
}
