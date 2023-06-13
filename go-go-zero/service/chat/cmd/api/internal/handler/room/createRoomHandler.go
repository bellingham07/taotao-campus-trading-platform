package room

import (
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	xhttp "github.com/zeromicro/x/http"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-go-zero/service/chat/cmd/api/internal/logic/room"
	"go-go-zero/service/chat/cmd/api/internal/svc"
	"go-go-zero/service/chat/cmd/api/internal/types"
)

func CreateRoomHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateRoomReq
		if err := httpx.Parse(r, &req); err != nil {
			logx.Infof("%v \n", err)
			xhttp.JsonBaseResponseCtx(r.Context(), w, errors.New("参数错误！🤡"))
			return
		}

		l := room.NewCreateRoomLogic(r.Context(), svcCtx)
		resp, err := l.CreateRoom(&req)
		if err != nil {
			xhttp.JsonBaseResponseCtx(r.Context(), w, err)
		} else {
			xhttp.JsonBaseResponseCtx(r.Context(), w, resp)
		}
	}
}
