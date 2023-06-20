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

func GetByIdDoneTypeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req    types.IdDoneTypeReq
			userId int64
		)
		if err := httpx.Parse(r, &req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误！"))
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		token := r.Header.Get("Authorization")
		claim, err := utils.ParseToken(token)
		if err != nil {
			userId = 0
		} else {
			userId = claim.Id
		}

		l := noauth.NewGetByIdDoneTypeLogic(r.Context(), svcCtx)
		resp, err := l.GetByIdDoneTypeLogic(&req, userId)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
