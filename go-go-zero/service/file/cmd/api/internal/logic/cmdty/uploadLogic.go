package cmdty

import (
	"context"
	"errors"
	"github.com/yitter/idgenerator-go/idgen"
	"go-go-zero/service/cmdty/cmd/rpc/cmdtyservice"
	"go-go-zero/service/file/cmd/api/internal/logic"
	"go-go-zero/service/file/model"
	"mime/multipart"
	"strconv"
	"sync"
	"time"

	"go-go-zero/service/file/cmd/api/internal/svc"
	ctypes "go-go-zero/service/file/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadLogic {
	return &UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(req *ctypes.CmdtyPicsReq, userId int64) ([]ctypes.PicResp, error) {
	var (
		cmdtyId   = req.CmdtyId
		userIdStr = strconv.FormatInt(userId, 10)
		files     = make([]multipart.FileHeader, 0)
		orders    = make([]int64, 0)
		resp      = make([]ctypes.PicResp, 0)
		code      *cmdtyservice.CodeResp
	)

	commonLogic := logic.NewCommonLogic(l.ctx, l.svcCtx)
	for _, file := range req.Pics {
		files = append(files, file.Pic)
		orders = append(orders, file.Order)
	}
	urls, objectnames, err := commonLogic.MultiUpload(files, userIdStr)
	if err != nil {
		return nil, errors.New("上传图片失败！😥")
	}

	var wg sync.WaitGroup
	wg.Add(1)
	fas := make([]model.FileCmdty, 0)
	t := time.Now().Local()
	for idx, url := range urls {
		id := idgen.NextId()
		fa := model.FileCmdty{
			Id:         id,
			CmdtyId:    cmdtyId,
			UserId:     userId,
			Url:        url,
			Objectname: objectnames[idx],
			UploadAt:   t,
			Order:      orders[idx],
		}
		pr := ctypes.PicResp{
			Id:    id,
			Url:   url,
			Order: orders[idx],
		}
		resp = append(resp, pr)
		if orders[idx] == 1 {
			fa.IsCover = 1
			cr := &cmdtyservice.CoverReq{
				Id:    cmdtyId,
				Cover: url,
			}
			go func() {
				code, _ = l.svcCtx.CmdtyRpc.UpdateCover(l.ctx, cr)
				wg.Done()
			}()
		}
		fas = append(fas, fa)
	}
	wg.Wait()
	if code.GetCode() == -1 {
		commonLogic.MultiDelete(objectnames)
		return nil, errors.New("上传图片失败！😥")
	}
	_, err = l.svcCtx.FileCmdty.Insert(fas)
	if err != nil {
		commonLogic.MultiDelete(objectnames)
		return nil, errors.New("运气不好，上传失败！😥")
	}
	return resp, nil
}
