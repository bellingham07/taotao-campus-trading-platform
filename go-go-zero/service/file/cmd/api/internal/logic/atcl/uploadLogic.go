package atcl

import (
	"context"
	"errors"
	"go-go-zero/service/atcl/cmd/rpc/atclservice"
	"go-go-zero/service/file/cmd/api/internal/logic"
	"go-go-zero/service/file/model"
	"mime/multipart"
	"strconv"
	"sync"
	"time"

	"go-go-zero/service/file/cmd/api/internal/svc"
	atypers "go-go-zero/service/file/cmd/api/internal/types"

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

func (l *UploadLogic) Upload(req *atypers.AtclPicsReq, userId int64) ([]atypers.PicResp, error) {
	var (
		atclId    = req.AtclId
		userIdStr = strconv.FormatInt(userId, 10)
		files     = make([]multipart.FileHeader, 0)
		orders    = make([]int64, 0)
		code      *atclservice.CodeResp
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
	fas := make([]model.FileAtcl, 0)
	resp := make([]atypers.PicResp, 0)
	t := time.Now()
	for idx, url := range urls {
		fa := model.FileAtcl{
			AtclId:     atclId,
			UserId:     userId,
			Url:        url,
			Objectname: objectnames[idx],
			UploadAt:   t,
			Order:      orders[idx],
		}
		pr := atypers.PicResp{
			Url:   url,
			Order: orders[idx],
		}
		resp = append(resp, pr)
		if orders[idx] == 1 {
			fa.IsCover = 1
			cr := &atclservice.CoverReq{
				Id:    atclId,
				Cover: url,
			}
			go func() {
				defer wg.Done()
				code, _ = l.svcCtx.AtclRpc.UpdateCover(l.ctx, cr)
			}()
		}
		fas = append(fas, fa)
	}
	wg.Wait()
	if code.GetCode() == -1 {
		commonLogic.MultiDelete(objectnames)
		return nil, errors.New("上传图片失败！😥")
	}
	_, err = l.svcCtx.Xorm.Insert(fas)
	if err != nil {
		commonLogic.MultiDelete(objectnames)
		return nil, errors.New("运气不好，上传失败！😥")
	}
	return resp, nil
}
