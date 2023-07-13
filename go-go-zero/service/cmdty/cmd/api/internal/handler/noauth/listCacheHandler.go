package noauth

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"go-go-zero/service/cmdty/cmd/api/internal/logic/noauth"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
)

func ListCacheHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := noauth.NewListCacheLogic(r.Context(), svcCtx)
		resp := l.ListCache()
		xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
	}
}
