package middleware

import (
	"net/http"
	"tik-tok-brief/common/middleware"
)

type JWTOptionalAuthMiddleware struct {
	accessSecret string
}

func NewJWTOptionalAuthMiddleware(accessSecret string) *JWTOptionalAuthMiddleware {
	return &JWTOptionalAuthMiddleware{accessSecret: accessSecret}
}

func (m *JWTOptionalAuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	jwtAuth := middleware.NewJwtAuth(m.accessSecret)
	return jwtAuth.OptionalJWTHandle(next)
}
