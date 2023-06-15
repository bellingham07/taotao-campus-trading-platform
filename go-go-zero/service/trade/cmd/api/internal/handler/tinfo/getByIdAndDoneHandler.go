package tinfo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/trade/cmd/api/internal/logic/tinfo"
	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"
)

func GetByIdAndDoneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdDoneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := tinfo.NewGetByIdAndDoneLogic(r.Context(), svcCtx)
		err := l.GetByIdAndDone(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
