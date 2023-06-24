package cmdty

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
	"go-go-zero/common/utils"
	"go-go-zero/service/file/cmd/api/internal/logic/cmdty"
	"go-go-zero/service/file/cmd/api/internal/svc"
	"go-go-zero/service/file/cmd/api/internal/types"
	"net/http"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req *types.CmdtyPicsReq
		if err := httpx.Parse(r, req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误！🤡"))
			return
		}

		userId := utils.GetUserId(r)

		l := cmdty.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(req, userId)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
