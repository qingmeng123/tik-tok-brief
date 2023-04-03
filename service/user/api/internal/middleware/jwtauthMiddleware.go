package middleware

import (
	"net/http"
	"tik-tok-brief/common/middleware"
)

type JwtAuthMiddleware struct {
	accessSecret string
}

func NewJwtAuthMiddleware(accessToken string) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{accessSecret: accessToken}
}

func (m *JwtAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	jwtAuth := middleware.NewJwtAuth(m.accessSecret)
	return jwtAuth.Handle(next)
}
