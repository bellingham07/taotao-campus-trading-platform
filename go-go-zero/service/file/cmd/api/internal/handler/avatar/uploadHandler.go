package avatar

import (
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"go-go-zero/service/file/cmd/api/internal/logic/avatar"
	"go-go-zero/service/file/cmd/api/internal/svc"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		_, header, err := r.FormFile("avatar")
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, "参数错误！😥")
			return
		}
		l := avatar.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(header)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
