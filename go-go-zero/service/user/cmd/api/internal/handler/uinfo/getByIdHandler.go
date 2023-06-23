package uinfo

import (
	"errors"
	"fmt"
	xhttp "github.com/zeromicro/x/http"
	"go-go-zero/common/utils"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/user/cmd/api/internal/logic/uinfo"
	"go-go-zero/service/user/cmd/api/internal/svc"
	"go-go-zero/service/user/cmd/api/internal/types"
)

func GetByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userIdAny := r.Context().Value(utils.JwtId("userId"))

		userId := userIdAny.(int64)

		var req types.IdReq
		if err := httpx.Parse(r, &req); err != nil {
			fmt.Println(req, err, userId)
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误！🤡"))
			return
		}

		l := uinfo.NewGetByIdLogic(r.Context(), svcCtx)
		resp, err := l.GetById(&req, userId)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
