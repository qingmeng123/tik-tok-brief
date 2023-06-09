// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	chat "tik-tok-brief/service/user/api/internal/handler/chat"
	"tik-tok-brief/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: RegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/",
					Handler: getUserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: followHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/follow/list",
					Handler: followListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/follower/list",
					Handler: followerListHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/friend/list",
					Handler: friendsListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/relation"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/message/action",
					Handler: chat.SendMessageHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/message/chat",
					Handler: chat.HistoryMessageHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin"),
	)
}
