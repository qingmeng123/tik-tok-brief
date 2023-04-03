package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"tik-tok-brief/common/response"
	"tik-tok-brief/service/user/api/internal/logic"
	"tik-tok-brief/service/user/api/internal/svc"
	"tik-tok-brief/service/user/api/internal/types"
)

func friendsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendsListReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParseErr(r.Context(), w, err)
			return
		}

		l := logic.NewFriendsListLogic(r.Context(), svcCtx)
		resp, err := l.FriendsList(&req)
		response.Response(w, resp, err)
	}
}
