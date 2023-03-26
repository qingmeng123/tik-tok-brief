package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/common/response"
	"tik-tok-brief/common/tool"
)

type ParseFormMiddleware struct {
	accessSecret string
	MaxVideoSize int64
}

func NewParseFormMiddleware(accessSecret string, MaxVideoSize int64) *ParseFormMiddleware {
	return &ParseFormMiddleware{accessSecret: accessSecret, MaxVideoSize: MaxVideoSize}
}

func (m *ParseFormMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var token string
		//解析请求体
		err := r.ParseMultipartForm(m.MaxVideoSize)
		if err != nil {
			response.Response(w, nil, errorx.NewParamErr(errorx.ERRFILEPARAM))
			return
		}

		//验证视频文件
		if r.MultipartForm.File["data"] == nil {
			response.Response(w, nil, errorx.NewParamErr(errorx.ERRFILEPARAM))
			return
		}

		ctx := r.Context()
		//获取token
		if r.MultipartForm.Value["token"] != nil {
			token = r.MultipartForm.Value["token"][0]
			claims, err := tool.ParseToken(token, m.accessSecret)
			if err != nil {
				logx.Error("jwt auth parse token err:", err)
				response.Response(w, nil, errorx.NewParamErr(errorx.ERRTOKEN))
				return
			}
			//写入上下文,执行之后的api
			ctx = context.WithValue(ctx, "user_id", claims.UserId)
		} else {
			response.Response(w, nil, errorx.NewParamErr(errorx.ERRTOKEN))
			return
		}
		next(w, r.WithContext(ctx))
	}
}
