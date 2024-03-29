package noauth

import (
	"errors"
	xhttp "github.com/zeromicro/x/http"
	"go-go-zero/common/utils"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/cmdty/cmd/api/internal/logic/noauth"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"
)

func GetByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IdReq
		if err := httpx.ParsePath(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误！"))
			return
		}

		userId := utils.GetUserIdWithNoAuth(r)

		l := noauth.NewGetByIdTypeLogic(r.Context(), svcCtx)
		resp, err := l.GetByIdTypeLogic(&req, userId)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
