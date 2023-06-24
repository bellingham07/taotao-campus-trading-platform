package avatar

import (
	"errors"
	xhttp "github.com/zeromicro/x/http"
	"go-go-zero/common/utils"
	"net/http"

	"go-go-zero/service/file/cmd/api/internal/logic/avatar"
	"go-go-zero/service/file/cmd/api/internal/svc"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, header, err := r.FormFile("avatar")
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误！😥"))
			return
		}

		userId := utils.GetUserId(r)

		l := avatar.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(header, userId)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
