package collect

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"go-go-zero/service/cmdty/cmd/api/internal/svc"
	"go-go-zero/service/cmdty/model"
)

type ListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListLogic) List() ([]*model.CmdtyCollect, error) {
	var userId int64 = 408301323265285
	ccList := l.svcCtx.CmdtyCollect.ListByUserId(userId)
	if ccList == nil {
		return nil, errors.New("出错啦😫")
	}
	return ccList, nil
}
