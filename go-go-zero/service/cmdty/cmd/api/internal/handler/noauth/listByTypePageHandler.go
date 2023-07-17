package noauth

import (
	"errors"
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/cmdty/cmd/api/internal/logic/noauth"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"
)

func ListByTypePageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TypePageReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误"))
			return
		}

		l := noauth.NewListByTypePageLogic(r.Context(), svcCtx)
		resp, err := l.ListByTypePage(&req)
		if resp != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
