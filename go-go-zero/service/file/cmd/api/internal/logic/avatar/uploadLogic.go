package avatar

import (
	"context"
	"errors"
	"go-go-zero/service/file/cmd/api/internal/logic"
	"go-go-zero/service/file/cmd/api/internal/svc"
	utypes "go-go-zero/service/file/cmd/api/internal/types"
	"go-go-zero/service/file/model"
	"go-go-zero/service/user/cmd/rpc/userservice"
	"mime/multipart"
	"strconv"
	"sync"
	"xorm.io/xorm"

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

func (l *UploadLogic) Upload(header *multipart.FileHeader, userId int64) (*utypes.AvatarResp, error) {
	var userIdStr = strconv.FormatInt(userId, 10)
	// 1 先存到OSS
	commonLogic := logic.NewCommonLogic(l.ctx, l.svcCtx)
	url, objectName, err := commonLogic.Upload(header, userIdStr)
	if err != nil {
		return nil, errors.New("图片上传失败！😥")
	}

	// 2 存到数据库
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
	resp := &utypes.AvatarResp{Url: url}
	return resp, nil
}

// SaveOrUpdateByUserId TODO
func (l *UploadLogic) SaveOrUpdateByUserId(url string, objectName string, userId int64) error {
	var fa = new(model.FileAvatar)
	_, err := l.svcCtx.Xorm.Transaction(func(session *xorm.Session) (interface{}, error) {
		s := session.Table("file_avatar")
		has, err := s.Where("`user_id` = ?", userId).Get(fa)
		if !has {
			fa.Id = userId
			fa.Url = url
			fa.Objectname = objectName
			_, err = s.Insert(fa)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}

		fa.Url = url
		fa.Objectname = objectName
		_, err = s.Where("`user_id` = ?", userId).Update(fa)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	return err
}
