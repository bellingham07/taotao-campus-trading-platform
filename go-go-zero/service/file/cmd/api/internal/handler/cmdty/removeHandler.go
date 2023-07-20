package cmdty

import (
	"errors"
	xhttp "github.com/zeromicro/x/http"
	"go-go-zero/service/file/cmd/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/file/cmd/api/internal/logic/cmdty"
	"go-go-zero/service/file/cmd/api/internal/svc"
)

func RemoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdCmdtyIdReq
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误！🤡"))
			return
		}

		l := cmdty.NewRemoveLogic(r.Context(), svcCtx)
		err := l.Remove(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
