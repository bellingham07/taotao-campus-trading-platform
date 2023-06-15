package avatar

import (
	"context"
	"errors"
	"go-go-zero/service/file/cmd/api/internal/logic"
	"go-go-zero/service/file/cmd/api/internal/svc"
	"go-go-zero/service/file/cmd/api/internal/types"
	"go-go-zero/service/user/cmd/rpc/userservice"
	"mime/multipart"
	"strconv"
	"sync"

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

func (l *UploadLogic) Upload(header *multipart.FileHeader) (*types.AvatarResp, error) {
	userIdStr := "408301323265285"
	// 1 先存到OSS
	commonLogic := logic.NewCommonLogic(l.ctx, l.svcCtx)
	url, objectName, err := commonLogic.Upload(header, userIdStr)
	if err != nil {
		return nil, errors.New("图片上传失败！😥")
	}
	// 2 存到数据库
	userId, _ := strconv.ParseInt(userIdStr, 10, 64)
	// 2.1 开协程去更新user服务的头像
	var wg sync.WaitGroup
	wg.Add(1)
	var code *userservice.CodeResp
	go func() {
		ar := &userservice.AvatarReq{
			Id:     userId,
			Avatar: url,
		}
		code, _ = l.svcCtx.UserRpc.UpdateAvatar(l.ctx, ar)
		wg.Done()
	}()
	// 2.1 OSS上传成功，就先更新file中的头像
	err = l.SaveOrUpdateByUserId(url, objectName, userId)
	if err != nil {
		commonLogic.Delete(objectName)
		return nil, errors.New("图片上传失败！😥")
	}
	wg.Wait()
	if code.GetCode() != 0 {
		commonLogic.Delete(objectName)
		go l.svcCtx.FileAvatar.Where("user_id = ?", userId).Delete()
		return nil, errors.New("图片上传失败！😥")
	}
	resp := &types.AvatarResp{Url: url}
	return resp, nil
}

// SaveOrUpdateByUserId TODO
func (l *UploadLogic) SaveOrUpdateByUserId(url string, name string, id int64) error {
	return nil
}
