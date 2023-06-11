package cmdty

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/file/cmd/api/internal/logic/cmdty"
	"go-go-zero/service/file/cmd/api/internal/svc"
	"go-go-zero/service/file/cmd/api/internal/types"
)

func CheckHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := cmdty.NewCheckLogic(r.Context(), svcCtx)
		err := l.Check(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
