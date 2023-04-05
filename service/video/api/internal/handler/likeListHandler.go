package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"tik-tok-brief/common/response"
	"tik-tok-brief/service/video/api/internal/logic"
	"tik-tok-brief/service/video/api/internal/svc"
	"tik-tok-brief/service/video/api/internal/types"
)

func likeListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LikeListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParseErr(r.Context(), w, err)
			return
		}

		l := logic.NewLikeListLogic(r.Context(), svcCtx)
		resp, err := l.LikeList(&req)
		response.Response(w, resp, err)
	}
}
