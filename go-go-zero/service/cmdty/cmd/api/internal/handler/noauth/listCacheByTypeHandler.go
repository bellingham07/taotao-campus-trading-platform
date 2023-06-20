package noauth

import (
	"errors"
	xhttp "github.com/zeromicro/x/http"
	"go-go-zero/service/cmdty/cmd/api/internal/logic/cinfo"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"
)

func ListCacheByTypeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TypeReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误！🤡"))
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cinfo.NewListCacheByTypeLogic(r.Context(), svcCtx)
		resp, err := l.ListCacheByType(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
