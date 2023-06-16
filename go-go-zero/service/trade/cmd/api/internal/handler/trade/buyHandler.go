package trade

import (
	xhttp "github.com/zeromicro/x/http"
	"go-go-zero/service/trade/cmd/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/trade/cmd/api/internal/logic/trade"
	"go-go-zero/service/trade/cmd/api/internal/svc"
)

func BuyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BuyReq
		err := httpx.Parse(r, &req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, "参数错误！")
		}

		l := trade.NewBuyLogic(r.Context(), svcCtx)
		id, err := l.Buy(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, id)
		}
	}
}
