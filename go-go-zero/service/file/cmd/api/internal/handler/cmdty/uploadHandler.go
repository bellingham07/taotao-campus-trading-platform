package cmdty

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	xhttp "github.com/zeromicro/x/http"
	"go-go-zero/service/file/cmd/api/internal/logic/cmdty"
	"go-go-zero/service/file/cmd/api/internal/svc"
	"mime/multipart"
	"net/http"
)

type PicReq struct {
	Pic   multipart.FileHeader `json:"pic"`
	Order int64                `json:"order"`
}

type CmdtyPicsReq struct {
	CmdtyId int64    `json:"cmdtyId"`
	Pics    []PicReq `json:"pics"`
}

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req *CmdtyPicsReq
		if err := httpx.Parse(r, req); err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, "参数错误！🤡")
			return
		}

		l := cmdty.NewUploadLogic(r.Context(), svcCtx)
		resp, err := l.Upload(req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
