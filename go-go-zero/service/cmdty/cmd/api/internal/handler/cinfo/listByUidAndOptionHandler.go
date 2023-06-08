package cinfo

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/cmdty/cmd/api/internal/logic/cinfo"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
)

func ListByUidAndOptionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cinfo.NewListByUidAndOptionLogic(r.Context(), svcCtx)
		resp, err := l.ListByUidAndOption()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
