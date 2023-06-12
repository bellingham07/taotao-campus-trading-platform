package msg

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/chat/cmd/api/internal/logic/msg"
	"go-go-zero/service/chat/cmd/api/internal/svc"
	"go-go-zero/service/chat/cmd/api/internal/types"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListMessageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := msg.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
