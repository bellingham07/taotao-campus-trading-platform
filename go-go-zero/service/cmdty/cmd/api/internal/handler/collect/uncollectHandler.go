package collect

import (
	"go-go-zero/service/cmdty/cmd/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/cmdty/cmd/api/internal/logic/collect"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
)

func UncollectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(types.IdReq)
		err := httpx.Parse(r, &req)
		l := collect.NewUncollectLogic(r.Context(), svcCtx)
		resp, err := l.Uncollect(req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
