package cinfo

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"go-go-zero/service/cmdty/cmd/api/internal/logic/cinfo"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
)

func ListByUidAndOptionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cinfo.NewListByUidAndOptionLogic(r.Context(), svcCtx)
		resp, err := l.ListByUidAndOption()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
