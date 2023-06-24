package collect

import (
	xhttp "github.com/zeromicro/x/http"
	"go-go-zero/common/utils"
	"go-go-zero/service/cmdty/cmd/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/cmdty/cmd/api/internal/logic/collect"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
)

func CollectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := new(types.IdReq)
		err := httpx.Parse(r, &req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, "参数错误！🤡")
		}

		userId := utils.GetUserId(r)

		l := collect.NewCollectLogic(r.Context(), svcCtx)
		err = l.Collect(req, userId)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, nil)
		}
	}
}
