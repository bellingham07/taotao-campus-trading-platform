package tag

import (
	"context"
	"errors"

	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveTagLogic {
	return &RemoveTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveTagLogic) RemoveTag(req *types.IdsReq) error {
	err := l.svcCtx.CmdtyTag.DeleteByIds(req.Ids)
	if err != nil {
		return errors.New("操作失败！😢")
	}
	return nil
}
