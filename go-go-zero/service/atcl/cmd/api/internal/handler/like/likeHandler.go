package like

import (
	"errors"
	xhttp "github.com/zeromicro/x/http"
	"go-go-zero/service/atcl/cmd/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/atcl/cmd/api/internal/logic/like"
	"go-go-zero/service/atcl/cmd/api/internal/svc"
)

func LikeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		err := httpx.Parse(r, &req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误！"))
		}

		l := like.NewLikeLogic(r.Context(), svcCtx)
		err = l.Like(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
