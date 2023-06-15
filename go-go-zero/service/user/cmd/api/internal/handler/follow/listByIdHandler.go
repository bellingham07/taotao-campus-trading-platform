package follow

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
	"go-go-zero/service/user/cmd/api/internal/types"
	"net/http"

	"go-go-zero/service/user/cmd/api/internal/logic/follow"
	"go-go-zero/service/user/cmd/api/internal/svc"
)

func ListByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PageReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, "参数错误！🤡")
			return
		}

		l := follow.NewListByIdLogic(r.Context(), svcCtx)
		resp, err := l.ListById(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
