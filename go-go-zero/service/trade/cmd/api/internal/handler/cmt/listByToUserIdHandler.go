package cmt

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/trade/cmd/api/internal/logic/cmt"
	"go-go-zero/service/trade/cmd/api/internal/svc"
	"go-go-zero/service/trade/cmd/api/internal/types"
)

func ListByToUserIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserIdReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cmt.NewListByToUserIdLogic(r.Context(), svcCtx)
		err := l.ListByToUserId(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
