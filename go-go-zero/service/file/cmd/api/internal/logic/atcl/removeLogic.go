package atcl

import (
	"context"
	"errors"
	"go-go-zero/service/file/cmd/api/internal/types"
	"go-go-zero/service/file/model"

	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/file/cmd/api/internal/svc"
)

type RemoveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveLogic) Remove(req *types.IdReq) error {
	_, err := l.svcCtx.FileAtcl.Delete(&model.FileAtcl{Id: req.Id})
	if err != nil {
		logx.Debugf("[DB ERROR] Remove 删除文章图片失败 %v\n", err)
		return errors.New("操作失败！😥")
	}
	return nil
}
