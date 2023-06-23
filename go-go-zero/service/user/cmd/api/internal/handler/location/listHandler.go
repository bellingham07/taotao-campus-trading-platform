package location

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"go-go-zero/service/user/cmd/api/internal/logic/location"
	"go-go-zero/service/user/cmd/api/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := location.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
