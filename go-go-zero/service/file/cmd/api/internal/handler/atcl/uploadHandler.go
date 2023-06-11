package atcl

import (
	xhttp "github.com/zeromicro/x/http"
	"mime/multipart"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/file/cmd/api/internal/logic/atcl"
	"go-go-zero/service/file/cmd/api/internal/svc"
)

type PicReq struct {
	Pic   multipart.FileHeader `json:"pic"`
	Order int64                `json:"order"`
}

type AtclPicsReq struct {
	AtclId int64    `json:"cmdtyId"`
	Pics   []PicReq `json:"pics"`
}

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req *AtclPicsReq
		if err := httpx.Parse(r, req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, "参数错误！")
			return
		}

		l := atcl.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
