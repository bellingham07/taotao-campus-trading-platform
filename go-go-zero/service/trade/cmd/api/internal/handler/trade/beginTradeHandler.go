package trade

import (
	"errors"
	xhttp "github.com/zeromicro/x/http"
	"go-go-zero/common/utils"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/trade/cmd/api/internal/logic/trade"
	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"
)

func BeginTradeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TradeReq
		err := httpx.Parse(r, &req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误！"))
		}

		userId, name := utils.GetUserIdAndName(r)

		l := trade.NewBeginTradeLogic(r.Context(), svcCtx)
		id, err := l.BeginTrade(&req, userId, name)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, id)
		}
	}
}
