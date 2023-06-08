package collect

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"go-go-zero/service/cmdty/cmd/api/internal/logic/collect"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := collect.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()

		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
