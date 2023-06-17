package like

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"go-go-zero/service/cmdty/cmd/api/internal/logic/like"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
)

func LikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := like.NewLikeLogic(r.Context(), svcCtx)
		err := l.Like()
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
