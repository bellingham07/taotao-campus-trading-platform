// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	cmt "go-go-zero/service/trade/cmd/api/internal/handler/cmt"
	trade "go-go-zero/service/trade/cmd/api/internal/handler/trade"
	"go-go-zero/service/trade/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: trade.GetByIdAndStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/",
				Handler: trade.ListByRoleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: trade.BeginTradeHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/:id/:stage",
				Handler: trade.ConfirmHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/:id",
				Handler: trade.CancelHandler(serverCtx),
			},
		},
		rest.WithPrefix("/trade"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/",
				Handler: cmt.CmtHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:userId",
				Handler: cmt.ListByToUserIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/:tradeId",
				Handler: cmt.ListByTradeIdHandler(serverCtx),
			},
		},
		rest.WithPrefix("/trade/cmt"),
	)
}
