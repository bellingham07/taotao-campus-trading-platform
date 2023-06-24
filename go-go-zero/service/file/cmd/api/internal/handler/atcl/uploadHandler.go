package atcl

import (
	"errors"
	xhttp "github.com/zeromicro/x/http"
	"go-go-zero/common/utils"
	"go-go-zero/service/file/cmd/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/file/cmd/api/internal/logic/atcl"
	"go-go-zero/service/file/cmd/api/internal/svc"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req *types.AtclPicsReq
		if err := httpx.Parse(r, req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误！"))
			return
		}

		userId := utils.GetUserId(r)

		l := atcl.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(req, userId)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
