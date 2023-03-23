/*******
* @Author:qingmeng
* @Description:
* @File:jwtauth
* @Date:2023/3/22
 */

package middleware

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"tik-tok-brief/common/errorx"
	"tik-tok-brief/common/response"
	"tik-tok-brief/common/tool"
)

type JwtAuth struct {
	accessSecret string
}

func NewJwtAuth(accessSecret string) *JwtAuth {
	return &JwtAuth{accessSecret: accessSecret}
}

func (m *JwtAuth) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var token string
		//解析请求体
		_= r.ParseForm()
		ctx := r.Context()
		//获取token
		if r.Form.Has("token"){
			token=r.Form.Get("token")
		}

		if token !=""{
			//解析token
			claims, err := tool.ParseToken(token, m.accessSecret)
			if err != nil {
				logx.Error("jwt auth parse token err:",err)
				response.Response(w,nil,errorx.NewParamErr(errorx.ERRTOKEN))
			}
			//写入上下文,执行之后的api
			ctx = context.WithValue(ctx, "user_id", claims.UserId)
			next(w,r.WithContext(ctx))
		}else {
			response.Response(w,nil,errorx.NewParamErr(errorx.ERRTOKEN))
		}
	}
}