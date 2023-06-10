package cmt

import (
	xhttp "github.com/zeromicro/x/http"
	"go-go-zero/service/cmdty/cmd/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/cmdty/cmd/api/internal/logic/cmt"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
)

func ListByInfoIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(types.CmdtyIdReq)
		err := httpx.ParsePath(r, req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, "参数错误！🤡")
		}
		l := cmt.NewListByInfoIdLogic(r.Context(), svcCtx)
		resp, err := l.ListByInfoId(req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
