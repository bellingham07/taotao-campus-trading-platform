package like

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/atcl/cmd/api/internal/logic/like"
	"go-go-zero/service/atcl/cmd/api/internal/svc"
)

func UnlikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := like.NewUnlikeLogic(r.Context(), svcCtx)
		err := l.Unlike()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
