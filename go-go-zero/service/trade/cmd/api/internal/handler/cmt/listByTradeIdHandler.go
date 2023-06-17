package cmt

import (
	"errors"
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/trade/cmd/api/internal/logic/cmt"
	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"
)

func ListByTradeIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TradeIdReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误！"))
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cmt.NewListByTradeIdLogic(r.Context(), svcCtx)
		resp, err := l.ListByTradeId(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
