package cmdty

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/file/cmd/api/internal/logic/cmdty"
	"go-go-zero/service/file/cmd/api/internal/svc"
)

func RemoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := cmdty.NewRemoveLogic(r.Context(), svcCtx)
		err := l.Remove()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
